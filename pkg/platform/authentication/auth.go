package authentication

import (
	"os"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/mono"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/authentication"
	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/viper"

	httptransport "github.com/go-openapi/runtime/client"
)

var (
	// FailNoCredentials is the failure type used when trying to authenticate without credentials
	FailNoCredentials = failures.Type("authentication.fail.nocredentials")

	// FailAuthAPI identifies a failure due to the auth api call
	FailAuthAPI = failures.Type("authentication.fail.api", api.FailUnknown)

	// FailTokenList identifies a failure in listing tokens from the api
	FailTokenList = failures.Type("authentication.fail.tokenlist", api.FailUnknown)

	// FailTokenDelete identifies a failure in deleting tokens from the api
	FailTokenDelete = failures.Type("authentication.fail.tokendelete", api.FailUnknown)

	// FailTokenCreate identifies a failure in creating tokens through the api
	FailTokenCreate = failures.Type("authentication.fail.tokencreate", api.FailUnknown)

	// FailNotAuthenticated is a helper failure that can be used by other libraries when they require authentication but
	// we aren't authenticated
	FailNotAuthenticated = failures.Type("authentication.fail.notauthed")
)

var exit = os.Exit

var persist *Auth

// Auth is the base structure used to record the authenticated state
type Auth struct {
	client      *mono_client.Mono
	clientAuth  *runtime.ClientAuthInfoWriter
	bearerToken string
	user        *mono_models.User
}

// Get returns a cached version of Auth
func Get() *Auth {
	if persist == nil {
		persist = New()
	}
	return persist
}

// Client is a shortcut for calling Client() on the persisted auth
func Client() *mono_client.Mono {
	return Get().Client()
}

// ClientAuth is a shortcut for calling ClientAuth() on the persisted auth
func ClientAuth() runtime.ClientAuthInfoWriter {
	return Get().ClientAuth()
}

// Reset clears the cache
func Reset() {
	persist = nil
}

// Logout will remove the stored apiToken
func Logout() {
	Get().Logout()
	Reset()
}

// Init creates a new version of Auth with default settings
func Init() *Auth {
	return New()
}

// New creates a new version of Auth
func New() *Auth {
	auth := &Auth{}

	if availableAPIToken() != "" {
		auth.Authenticate()
	}

	return auth
}

// Authenticated checks whether we are currently authenticated
func (s *Auth) Authenticated() bool {
	return s.clientAuth != nil
}

// ClientAuth returns the auth type required by swagger api calls
func (s *Auth) ClientAuth() runtime.ClientAuthInfoWriter {
	if s.clientAuth == nil {
		return nil
	}
	return *s.clientAuth
}

// BearerToken returns the current bearerToken
func (s *Auth) BearerToken() string {
	return s.bearerToken
}

// Authenticate will try to authenticate using stored credentials
func (s *Auth) Authenticate() *failures.Failure {
	if s.Authenticated() {
		return nil
	}

	apiToken := availableAPIToken()
	if apiToken == "" {
		return FailNoCredentials.New("err_no_credentials")
	}

	return s.AuthenticateWithToken(apiToken)
}

// AuthenticateWithModel will try to authenticate using the given swagger model
func (s *Auth) AuthenticateWithModel(credentials *mono_models.Credentials) *failures.Failure {
	params := authentication.NewPostLoginParams()
	params.SetCredentials(credentials)
	loginOK, err := mono.Get().Authentication.PostLogin(params)

	if err != nil {
		s.Logout()
		return FailAuthAPI.Wrap(err)
	}

	payload := loginOK.Payload
	s.user = payload.User
	s.bearerToken = payload.Token
	clientAuth := httptransport.BearerToken(s.bearerToken)
	s.clientAuth = &clientAuth

	if credentials.Token != "" {
		viper.Set("apiToken", credentials.Token)
	} else {
		s.CreateToken()
	}

	return nil
}

// AuthenticateWithUser will try to authenticate using the given credentials
func (s *Auth) AuthenticateWithUser(username, password, totp string) *failures.Failure {
	return s.AuthenticateWithModel(&mono_models.Credentials{
		Username: username,
		Password: password,
		Totp:     totp,
	})
}

// AuthenticateWithToken will try to authenticate using the given token
func (s *Auth) AuthenticateWithToken(token string) *failures.Failure {
	return s.AuthenticateWithModel(&mono_models.Credentials{
		Token: token,
	})
}

// WhoAmI returns the username of the currently authenticated user, or an empty string if not authenticated
func (s *Auth) WhoAmI() string {
	if s.user != nil {
		return s.user.Username
	}
	return ""
}

// UserID returns the user ID for the currently authenticated user, or nil if not authenticated
func (s *Auth) UserID() *strfmt.UUID {
	if s.user != nil {
		return &s.user.UserID
	}
	return nil
}

// Logout will destroy any session tokens and reset the current Auth instance
func (s *Auth) Logout() {
	viper.Set("apiToken", "")
	s.client = nil
	s.clientAuth = nil
	s.bearerToken = ""
	s.user = nil
}

// Client will return an API client that has authentication set up
func (s *Auth) Client() *mono_client.Mono {
	if s.client == nil {
		s.client = mono.NewWithAuth(s.clientAuth)
	}
	if !s.Authenticated() {
		if fail := s.Authenticate(); fail != nil {
			logging.Error("Trying to get the Client while not authenticated")
			print.Error(locale.T("err_api_not_authenticated"))
			exit(1)
			return nil
		}
	}
	return s.client
}

// CreateToken will create an API token for the current authenticated user
func (s *Auth) CreateToken() *failures.Failure {
	client := s.Client()
	tokensOK, err := client.Authentication.ListTokens(nil, s.ClientAuth())
	if err != nil {
		return FailTokenList.New(locale.Tr("err_token_list", err.Error()))
	}

	for _, token := range tokensOK.Payload {
		if token.Name == constants.APITokenName {
			params := authentication.NewDeleteTokenParams()
			params.SetTokenID(token.TokenID)
			_, err := client.Authentication.DeleteToken(params, s.ClientAuth())
			if err != nil {
				return FailTokenDelete.New(locale.Tr("err_token_delete", err.Error()))
			}
			break
		}
	}

	params := authentication.NewAddTokenParams()
	params.SetTokenOptions(&mono_models.TokenEditable{Name: constants.APITokenName})
	tokenOK, err := client.Authentication.AddToken(params, s.ClientAuth())
	if err != nil {
		return FailTokenCreate.New(locale.Tr("err_token_create", err.Error()))
	}

	token := tokenOK.Payload.Token
	viper.Set("apiToken", token)

	return nil
}

func availableAPIToken() string {
	if tkn := os.Getenv(constants.APIKeyEnvVarName); tkn != "" {
		return tkn
	}
	return viper.GetString("apiToken")
}
