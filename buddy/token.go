package buddy

import "net/http"

const (
	TokenScopeWorkspace         = "WORKSPACE"
	TokenScopeWorkspacesManage  = "WORKSPACES_MANAGE"
	TokenScopeProjectDelete     = "PROJECT_DELETE"
	TokenScopeRepositoryRead    = "REPOSITORY_READ"
	TokenScopeRepositoryWrite   = "REPOSITORY_WRITE"
	TokenScopeExecutionInfo     = "EXECUTION_INFO"
	TokenScopeExecutionRun      = "EXECUTION_RUN"
	TokenScopeExecutionManage   = "EXECUTION_MANAGE"
	TokenScopeEnvironmentInfo   = "ENVIRONMENT_INFO"
	TokenScopeEnvironmentAdd    = "ENVIRONMENT_ADD"
	TokenScopeEnvironmentManage = "ENVIRONMENT_MANAGE"
	TokenScopeZoneRead          = "ZONE_READ"
	TokenScopeZoneWrite         = "ZONE_WRITE"
	TokenScopeZoneManage        = "ZONE_MANAGE"
	TokenScopeTargetInfo        = "TARGET_INFO"
	TokenScopeTargetAdd         = "TARGET_ADD"
	TokenScopeTargetManage      = "TARGET_MANAGE"
	TokenScopeUserInfo          = "USER_INFO"
	TokenScopeUserKey           = "USER_KEY"
	TokenScopeUserEmail         = "USER_EMAIL"
	TokenScopeIntegrationInfo   = "INTEGRATION_INFO"
	TokenScopeIntegrationAdd    = "INTEGRATION_ADD"
	TokenScopeIntegrationManage = "INTEGRATION_MANAGE"
	TokenScopeMemberEmail       = "MEMBER_EMAIL"
	TokenScopeManageEmails      = "MANAGE_EMAILS"
	TokenScopeWebhookInfo       = "WEBHOOK_INFO"
	TokenScopeWebhookAdd        = "WEBHOOK_ADD"
	TokenScopeWebhookManage     = "WEBHOOK_MANAGE"
	TokenScopeVariableAdd       = "VARIABLE_ADD"
	TokenScopeVariableInfo      = "VARIABLE_INFO"
	TokenScopeVariableManage    = "VARIABLE_MANAGE"
	TokenScopeTokenInfo         = "TOKEN_INFO"
	TokenScopeTokenManage       = "TOKEN_MANAGE"
)

type TokenService struct {
	client *Client
}

type Token struct {
	Url                   string   `json:"url"`
	HtmlUrl               string   `json:"html_url"`
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	Token                 string   `json:"token"`
	ExpiresAt             string   `json:"expires_at"`
	Scopes                []string `json:"scopes"`
	IpRestrictions        []string `json:"ip_restrictions"`
	WorkspaceRestrictions []string `json:"workspace_restrictions"`
}

type TokenOps struct {
	Name                  *string   `json:"name,omitempty"`
	ExpiresIn             *int      `json:"expires_in,omitempty"`
	ExpiresAt             *string   `json:"expires_at,omitempty"`
	Scopes                *[]string `json:"scopes,omitempty"`
	IpRestrictions        *[]string `json:"ip_restrictions,omitempty"`
	WorkspaceRestrictions *[]string `json:"workspace_restrictions,omitempty"`
}

type Tokens struct {
	Url          string   `json:"url"`
	HtmlUrl      string   `json:"html_url"`
	AccessTokens []*Token `json:"access_tokens"`
}

func (s *TokenService) Create(ops *TokenOps) (*Token, *http.Response, error) {
	var t *Token
	resp, err := s.client.Create(s.client.NewUrlPath("/user/tokens"), &ops, nil, &t)
	return t, resp, err
}

func (s *TokenService) Get(tokenId string) (*Token, *http.Response, error) {
	var t *Token
	resp, err := s.client.Get(s.client.NewUrlPath("/user/tokens/%s", tokenId), &t, nil)
	return t, resp, err
}

func (s *TokenService) GetMe() (*Token, *http.Response, error) {
	var t *Token
	resp, err := s.client.Get(s.client.NewUrlPath("/user/token"), &t, nil)
	return t, resp, err
}

func (s *TokenService) GetList() (*Tokens, *http.Response, error) {
	var t *Tokens
	resp, err := s.client.Get(s.client.NewUrlPath("/user/tokens"), &t, nil)
	return t, resp, err
}

func (s *TokenService) Delete(tokenId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/user/tokens/%s", tokenId), nil, nil)
}
