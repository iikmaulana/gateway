package handler

import (
	"encoding/json"
	"github.com/iikmaulana/gateway/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	L "github.com/iikmaulana/gateway/libs/helper/logger"
	"github.com/iikmaulana/gateway/service"
)

func logger(err error) {
	if err != nil {
		L.Warn(err)
	}
}

func (fwd chiForwarder) hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)

	logger(json.NewEncoder(w).Encode(models.Response{
		Response:   http.StatusOK,
		Controller: r.RequestURI,
		Action:     r.Method,
		Result: map[string]interface{}{
			"id": "Selamat datang di API",
		},
	}))
}

func (fwd chiForwarder) forward(w http.ResponseWriter, r *http.Request) {
	var header metadata.MD

	composite := r.Context().Value(models.ServiceContextValueKey).(*service.Composite)

	if composite.Connection == nil && composite.ServiceClient == nil {
		basePath := composite.Endpoints()
		path := r.RequestURI[strings.Index(r.RequestURI, basePath)+len(basePath):]
		baseHttp := "http://" + composite.Url + composite.Endpoint + path
		if baseHttp != "" {
			httpReq, err := http.NewRequest(r.Method, baseHttp, r.Body)
			if err != nil {
				fwd.notFound(composite.Key, w, r)
				return
			}

			client := &http.Client{}
			httpReq.Header = r.Header

			resp, err := client.Do(httpReq)
			if err != nil {
				fwd.notFound(composite.Key, w, r)
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			fwd.responseFromHttp(composite.Key, w, body)
			defer resp.Body.Close()

			return
		}
	}

	ctx, req, err := transformRequestFromHttp(r)
	if err != nil {
		w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
		w.WriteHeader(http.StatusInternalServerError)
		logger(json.NewEncoder(w).Encode(models.Response{
			Response:   http.StatusInternalServerError,
			Error:      err.Error(),
			Controller: r.RequestURI,
			Action:     r.Method,
		}))

		return
	}

	resp, err := composite.Dispatch(ctx, req, grpc.Header(&header))
	if err != nil {
		var code int

		message := make(map[string]string)
		stat, _ := status.FromError(err)
		switch stat.Code() {
		default:
			code = http.StatusInternalServerError
			message["id"] = "Kesalahan pada server"

		case codes.Unavailable:
			code = http.StatusServiceUnavailable
			message["id"] = "Layanan tidak dapat diakses"
		}

		w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
		w.WriteHeader(code)
		logger(json.NewEncoder(w).Encode(models.Response{
			Response:   code,
			Error:      err.Error(),
			Controller: r.RequestURI,
			Action:     r.Method,
		}))

		return
	}

	L.Infof("request %s has been served by %s", r.RequestURI, resp.Server)
	logger(transformResponseToHTTP(resp, header, w))
}

func (fwd chiForwarder) notFound(serviceName string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
	w.WriteHeader(http.StatusNotImplemented)
	logger(json.NewEncoder(w).Encode(models.Response{
		Response:   http.StatusNotImplemented,
		Error:      "Layanan tidak terdaftar",
		Appid:      "",
		Svcid:      serviceName,
		Controller: r.RequestURI,
		Action:     r.Method,
		Result:     "",
	}))
}

func (fwd chiForwarder) responseFromHttp(serviceName string, w http.ResponseWriter, r []byte) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)

	asd := models.Response{}
	if err := json.Unmarshal(r, &asd); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(r)
	return

}

func (fwd chiForwarder) unauthorized(w http.ResponseWriter, message map[string]string, errs []models.Error) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
	w.WriteHeader(http.StatusUnauthorized)
	logger(json.NewEncoder(w).Encode(models.Response{
		Response: http.StatusUnauthorized,
		Error:    message["id"],
	}))
}
