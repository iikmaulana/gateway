package models

const (
	BodyTypeJSON = "application/json"
	BodyTypeXML  = "application/xml"
	BodyTypeRaw  = "raw"
)

type AuthorizationInfo struct {
	UserID         string `json:"userid"`
	Username       string `json:"username"`
	IsOrgAdmin     int    `json:"isOrgAdmin"`
	IsActive       int    `json:"isActive"`
	OrganizationId string `json:"organization_id"`
	AppId          string `json:"app_id"`
}

type ClientInfo struct {
	ClientID string `json:"client_id"`
	Agent    string `json:"agent"`
}

type RequestContext struct {
	path       string
	method     string
	body       []byte
	params     map[string]string
	query      map[string]string
	header     map[string]string
	authInfo   AuthorizationInfo
	clientInfo ClientInfo
}
