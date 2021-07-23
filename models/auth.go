package models

import (
	"github.com/dgrijalva/jwt-go"
	"net/url"
	"time"
)

const (
	TimestampFormat string = "2006-01-02T15:04:05-0700"
	TokenTypeBearer string = "bearer"
)

type Request struct {
	Method      string
	URL         *url.URL
	Body        []byte
	CompositeID CompositeID
	Timestamp   time.Time
	Signature   string
	Token       string
}

type CompositeID string

type Composite struct {
	ID     CompositeID `json:"id"`
	Secret string      `json:"secret"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID         string `json:"userid"`
	Username       string `json:"username"`
	IsOrgAdmin     int    `json:"isOrgAdmin"`
	IsActive       int    `json:"isActive"`
	OrganizationId string `json:"organization_id"`
	AppId          string `json:"app_id"`
}

type Token struct {
	Type  string
	Value string
}
