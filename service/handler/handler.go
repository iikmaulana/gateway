package handler

import (
	"encoding/json"
	"github.com/iikmaulana/gateway/models"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

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

func (fwd chiForwarder) notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
	w.WriteHeader(http.StatusNotImplemented)
	logger(json.NewEncoder(w).Encode(models.Response{
		Response:   http.StatusNotImplemented,
		Error:      "Layanan tidak terdaftar",
		Controller: r.RequestURI,
		Action:     r.Method,
	}))
}

func (fwd chiForwarder) unauthorized(w http.ResponseWriter, message map[string]string, errs []models.Error) {
	w.Header().Set(models.ContentTypeHeaderKey, models.ContentTypeValueJSON)
	w.WriteHeader(http.StatusUnauthorized)
	logger(json.NewEncoder(w).Encode(models.Response{
		Response: http.StatusUnauthorized,
		Error:    message["id"],
	}))
}