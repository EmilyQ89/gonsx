package securitytags

import (
	"github.com/sky-uk/gonsx/api"
	"net/http"
)

// GetAllSecurityTagsAPI - struct
type GetAllSecurityTagsAPI struct {
	*api.BaseAPI
}

// NewGetAll - Generates a new GetAllSecurityTagsAPI object.
func NewGetAll() *GetAllSecurityTagsAPI {
	this := new(GetAllSecurityTagsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/2.0/services/securitytags/tags", nil, new(SecurityTags))
	return this
}

// GetResponse returns the ResponseObject from CreateSecurityTagAPI
func (getAPI GetAllSecurityTagsAPI) GetResponse() *SecurityTags {
	return getAPI.ResponseObject().(*SecurityTags)
}