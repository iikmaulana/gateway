package manual

import (
	"fmt"
	"github.com/iikmaulana/gateway/controller/resolver"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

type manualResolver struct{}

func NewManualResolver() (res resolver.Resolver, errx serror.SError) {
	res = manualResolver{}
	return res, errx
}

func (ox manualResolver) GenerateURL(service string, port string) (url string) {
	url = fmt.Sprintf("%s:%s", service, port)
	return url
}

func (ox manualResolver) Register() (errx serror.SError) {
	return errx
}
