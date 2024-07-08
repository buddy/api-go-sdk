package buddy

import (
	"net/http"
)

const (
	IntegrationTypeDigitalOcean         = "DIGITAL_OCEAN"
	IntegrationTypeAmazon               = "AMAZON"
	IntegrationTypeStackHawk            = "STACK_HAWK"
	IntegrationTypeShopify              = "SHOPIFY"
	IntegrationTypePushover             = "PUSHOVER"
	IntegrationTypeRackspace            = "RACKSPACE"
	IntegrationTypeCloudflare           = "CLOUDFLARE"
	IntegrationTypeNewRelic             = "NEW_RELIC"
	IntegrationTypeSentry               = "SENTRY"
	IntegrationTypeRollbar              = "ROLLBAR"
	IntegrationTypeDatadog              = "DATADOG"
	IntegrationTypeDigitalOceanSpaces   = "DO_SPACES"
	IntegrationTypeHoneybadger          = "HONEYBADGER"
	IntegrationTypeVultr                = "VULTR"
	IntegrationTypeSentryEnterprise     = "SENTRY_ENTERPRISE"
	IntegrationTypeLoggly               = "LOGGLY"
	IntegrationTypeFirebase             = "FIREBASE"
	IntegrationTypeUpcloud              = "UPCLOUD"
	IntegrationTypeGhostInspector       = "GHOST_INSPECTOR"
	IntegrationTypeAzureCloud           = "AZURE_CLOUD"
	IntegrationTypeDockerHub            = "DOCKER_HUB"
	IntegrationTypeGoogleServiceAccount = "GOOGLE_SERVICE_ACCOUNT"
	IntegrationTypeGitHub               = "GIT_HUB"
	IntegrationTypeGitLab               = "GIT_LAB"

	IntegrationScopeWorkspace = "WORKSPACE"
	IntegrationScopeProject   = "PROJECT"

	IntegrationAuthTypeToken             = "TOKEN"
	IntegrationAuthTypeTokenAppExtension = "TOKEN_APP_EXTENSION"
	IntegrationAuthTypeDefault           = "DEFAULT"
	IntegrationAuthTypeTrusted           = "TRUSTED"
	IntegrationAuthTypeOidc              = "OIDC"

	IntegrationPermissionManage  = "MANAGE"
	IntegrationPermissionUseOnly = "USE_ONLY"
	IntegrationPermissionDenied  = "DENIED"
)

type Integration struct {
	Url                 string                  `json:"url"`
	HtmlUrl             string                  `json:"html_url"`
	HashId              string                  `json:"hash_id"`
	Name                string                  `json:"name"`
	Type                string                  `json:"type"`
	AuthType            string                  `json:"auth_type"`
	Scope               string                  `json:"scope"`
	ProjectName         string                  `json:"project_name"`
	Identifier          string                  `json:"identifier"`
	AllPipelinesAllowed bool                    `json:"all_pipelines_allowed"`
	AllowedPipelines    []*AllowedPipeline      `json:"allowed_pipelines"`
	Permissions         *IntegrationPermissions `json:"permissions"`
}

type Integrations struct {
	Url          string         `json:"url"`
	Integrations []*Integration `json:"integrations"`
}

type IntegrationOps struct {
	Name                *string                 `json:"name,omitempty"`
	Type                *string                 `json:"type,omitempty"`
	Scope               *string                 `json:"scope,omitempty"`
	ProjectName         *string                 `json:"project_name,omitempty"`
	Username            *string                 `json:"username,omitempty"`
	Shop                *string                 `json:"shop,omitempty"`
	Token               *string                 `json:"token,omitempty"`
	AccessKey           *string                 `json:"access_key,omitempty"`
	SecretKey           *string                 `json:"secret_key,omitempty"`
	Audience            *string                 `json:"audience,omitempty"`
	AppId               *string                 `json:"app_id,omitempty"`
	TenantId            *string                 `json:"tenant_id,omitempty"`
	Password            *string                 `json:"password,omitempty"`
	ApiKey              *string                 `json:"api_key,omitempty"`
	Email               *string                 `json:"email,omitempty"`
	AuthType            *string                 `json:"auth_type,omitempty"`
	PartnerToken        *string                 `json:"partner_token,omitempty"`
	GoogleProject       *string                 `json:"google_project,omitempty"`
	Config              *string                 `json:"config,omitempty"`
	Identifier          *string                 `json:"identifier,omitempty"`
	RoleAssumptions     *[]*RoleAssumption      `json:"role_assumptions,omitempty"`
	AllPipelinesAllowed *bool                   `json:"all_pipelines_allowed,omitempty"`
	AllowedPipelines    *[]*AllowedPipeline     `json:"allowed_pipelines,omitempty"`
	Permissions         *IntegrationPermissions `json:"permissions,omitempty"`
}

type IntegrationPermissions struct {
	Admins string                           `json:"admins"`
	Others string                           `json:"others"`
	Users  []*IntegrationResourcePermission `json:"users"`
	Groups []*IntegrationResourcePermission `json:"groups"`
}

type IntegrationResourcePermission struct {
	Id          int    `json:"id"`
	AccessLevel string `json:"access_level"`
}

type AllowedPipeline struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Url     string `json:"url,omitempty"`
	HtmlUrl string `json:"html_url,omitempty"`
}

type RoleAssumption struct {
	Arn        string `json:"arn"`
	ExternalId string `json:"external_id,omitempty"`
	Duration   int    `json:"duration,omitempty"`
}

type IntegrationService struct {
	client *Client
}

func (s *IntegrationService) Create(domain string, ops *IntegrationOps) (*Integration, *http.Response, error) {
	var i *Integration
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/integrations", domain), &ops, nil, &i)
	return i, resp, err
}

func (s *IntegrationService) Delete(domain string, hashId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/integrations/%s", domain, hashId), nil, nil)
}

func (s *IntegrationService) Update(domain string, hashId string, ops *IntegrationOps) (*Integration, *http.Response, error) {
	var i *Integration
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/integrations/%s", domain, hashId), &ops, nil, &i)
	return i, resp, err
}

func (s *IntegrationService) Get(domain string, hashId string) (*Integration, *http.Response, error) {
	var i *Integration
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/integrations/%s", domain, hashId), &i, nil)
	return i, resp, err
}

func (s *IntegrationService) GetList(domain string) (*Integrations, *http.Response, error) {
	var l *Integrations
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/integrations", domain), &l, nil)
	return l, resp, err
}
