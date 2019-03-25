package mono

import (
	"flag"

	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// persist contains the active API Client connection
var persist *mono_client.APIClient

// New will create a new API client using default settings (for an authenticated version use the NewWithAuth version)
func New() *mono_client.APIClient {
	return Init(api.GetSettings(api.ServiceMono), nil)
}

// NewWithAuth creates a new API client using default settings and the provided authentication info
func NewWithAuth(auth *runtime.ClientAuthInfoWriter) *mono_client.APIClient {
	return Init(api.GetSettings(api.ServiceMono), auth)
}

// Init initializes a new api client
func Init(apiSetting api.Settings, auth *runtime.ClientAuthInfoWriter) *mono_client.APIClient {
	transportRuntime := httptransport.New(apiSetting.Host, apiSetting.BasePath, []string{apiSetting.Schema})
	transportRuntime.Transport = api.NewUserAgentTripper()

	if flag.Lookup("test.v") != nil {
		transportRuntime.SetDebug(true)
	}

	if auth != nil {
		transportRuntime.DefaultAuthentication = *auth
	}
	return mono_client.New(transportRuntime, strfmt.Default)
}

// Get returns a cached version of the default api client
func Get() *mono_client.APIClient {
	if persist == nil {
		persist = New()
	}
	return persist
}
