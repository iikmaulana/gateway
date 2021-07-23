package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"time"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"

	"github.com/iikmaulana/gateway/controller/resolver"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/logger"
	"github.com/iikmaulana/gateway/packets"
)

type Forwarder interface {
	Mount(*Composite)
}

type Composite struct {
	packets.ServiceClient
	key             string
	endpoint        string
	connection      *grpc.ClientConn
	protectedRoutes map[string][]protectedRoute
}

func NewComposite(resolv resolver.Resolver, cfg Config) (*Composite, error) {
	connURL := resolv.GenerateURL(cfg.Host, helper.IntToString(cfg.Port))

	logger.Infof("connecting gRPC composite %s(%s)", cfg.Key, connURL)

	var err error
	var conn *grpc.ClientConn
	var sc packets.ServiceClient
	var res *packets.Ack

	i := 0
	max := 10
	for {
		if i > 0 {
			cooldown := 1
			switch {
			case i > 8:
				cooldown = 5

			case i > 5 && i <= 8:
				cooldown = 3
			}

			logger.Infof("will retrying (%d/%d) in %d seconds...", i, max, cooldown)
			time.Sleep(time.Duration(cooldown) * time.Second)
		}

		conn, err = grpc.Dial(
			connURL,
			grpc.WithInsecure(),
			grpc.WithBalancerName(roundrobin.Name),
			grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()),
		)
		if err != nil {
			if i < max {
				logger.Warnf("failed to dial service %s, detail: %v", cfg.Key, err)

				i++
				continue
			}

			return nil, fmt.Errorf("while dialing server: %v", err)
		}

		host, err := os.Hostname()
		if err != nil {
			host = "?"
		}

		logger.Infof("doing handshake with %s", cfg.Key)
		sc = packets.NewServiceClient(conn)
		res, err = sc.Handshake(context.Background(), &packets.AckRequest{
			From: host,
		})
		if err != nil {
			if i < max {
				logger.Warnf("failed to handshaking with service %s, detail: %v", cfg.Key, err)

				i++
				continue
			}

			return nil, fmt.Errorf("while handshaking with service %s: %v", cfg.Key, err)
		}

		break
	}

	logger.Infof("handshake has been served by %s", res.Server)
	if !cfg.Check(res.Checksum) {
		return nil, errors.New("invalid service checksum")
	}

	protectedRoutes := make(map[string][]protectedRoute)
	for method, prs := range res.ProtectedRoutes {
		for _, route := range prs.Routes {
			pattern, err := regexp.Compile(route.Pattern)
			if err != nil {
				return nil, err
			}

			if route.Method == "" {
				route.Method = "protect"
				if route.IsStrict {
					route.Method = "strict"
				}
			}

			protectedRoutes[method] = append(protectedRoutes[method], protectedRoute{
				pattern: pattern,
				method:  route.Method,
			})
		}
	}

	return &Composite{
		key:             cfg.Key,
		endpoint:        cfg.gatewayEndpoint,
		connection:      conn,
		ServiceClient:   sc,
		protectedRoutes: protectedRoutes,
	}, nil
}

func (c Composite) Key() string {
	return c.key
}

func (c Composite) Endpoint() string {
	return c.endpoint
}

func (c Composite) Stop() error {
	err := c.connection.Close()
	if err != nil {
		return fmt.Errorf("while closing composite connection: %v", err)
	}

	return nil
}

func (c Composite) IsNeedProtection(method string, path string) (bool, bool, bool) {
	if routes, ok := c.protectedRoutes[method]; ok {
		for _, route := range routes {
			if route.pattern.MatchString(path) {
				switch route.method {
				case "strict":
					return true, true, false

				case "private":
					return true, false, true

				case "protect":
					return true, false, false

				default:
					return false, false, false
				}
			}
		}
	}

	return false, false, false
}
