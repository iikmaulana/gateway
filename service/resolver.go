package service

import (
	"github.com/iikmaulana/gateway/controller/resolver"
	"github.com/iikmaulana/gateway/controller/resolver/manual"
	"github.com/iikmaulana/gateway/libs"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

func NewResolver() (resolv resolver.Resolver, errx serror.SError) {
	switch helper.Env(libs.AppCluster, libs.ClusterLocal) {

	default:
		resolv, errx = manual.NewManualResolver()
	}

	if errx != nil {
		return resolv, errx
	}

	errx = resolv.Register()
	if errx != nil {
		return resolv, errx
	}

	return resolv, errx
}
