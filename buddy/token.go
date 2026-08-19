package buddy

import "net/http"

const (
	TokenScopeUserRead         = "USER_READ"
	TokenScopeUserWrite        = "USER_WRITE"
	TokenScopeUserSshKeyRead   = "USER_SSH_KEY_READ"
	TokenScopeUserSshKeyWrite  = "USER_SSH_KEY_WRITE"
	TokenScopeUserSshKeyManage = "USER_SSH_KEY_MANAGE"
	TokenScopeUserEmailRead    = "USER_EMAIL_READ"
	TokenScopeUserEmailWrite   = "USER_EMAIL_WRITE"
	TokenScopeUserEmailManage  = "USER_EMAIL_MANAGE"
	TokenScopeUserPatRead      = "USER_PAT_READ"
	TokenScopeUserPatWrite     = "USER_PAT_WRITE"
	TokenScopeUserPatManage    = "USER_PAT_MANAGE"

	TokenScopeWorkspaceProvisioning = "WORKSPACE_PROVISIONING"
	TokenScopeWorkspaceRead         = "WORKSPACE_READ"
	TokenScopeWorkspaceManage       = "WORKSPACE_MANAGE"
	TokenScopeProjectRead           = "PROJECT_READ"
	TokenScopeProjectWrite          = "PROJECT_WRITE"
	TokenScopeProjectManage         = "PROJECT_MANAGE"
	TokenScopeMemberRead            = "MEMBER_READ"
	TokenScopeMemberWrite           = "MEMBER_WRITE"
	TokenScopeMemberManage          = "MEMBER_MANAGE"
	TokenScopeMemberEmailRead       = "MEMBER_EMAIL_READ"
	TokenScopeGroupRead             = "GROUP_READ"
	TokenScopeGroupWrite            = "GROUP_WRITE"
	TokenScopeGroupManage           = "GROUP_MANAGE"
	TokenScopePermissionRead        = "PERMISSION_READ"
	TokenScopePermissionWrite       = "PERMISSION_WRITE"
	TokenScopePermissionManage      = "PERMISSION_MANAGE"

	TokenScopeRepositoryRead  = "REPOSITORY_READ"
	TokenScopeRepositoryWrite = "REPOSITORY_WRITE"

	TokenScopePipelineRead   = "PIPELINE_READ"
	TokenScopePipelineRun    = "PIPELINE_RUN"
	TokenScopePipelineWrite  = "PIPELINE_WRITE"
	TokenScopePipelineManage = "PIPELINE_MANAGE"

	TokenScopeTargetRead   = "TARGET_READ"
	TokenScopeTargetWrite  = "TARGET_WRITE"
	TokenScopeTargetManage = "TARGET_MANAGE"

	TokenScopeDomainRead   = "DOMAIN_READ"
	TokenScopeDomainWrite  = "DOMAIN_WRITE"
	TokenScopeDomainManage = "DOMAIN_MANAGE"

	TokenScopeIntegrationRead   = "INTEGRATION_READ"
	TokenScopeIntegrationWrite  = "INTEGRATION_WRITE"
	TokenScopeIntegrationManage = "INTEGRATION_MANAGE"

	TokenScopeWebhookRead   = "WEBHOOK_READ"
	TokenScopeWebhookWrite  = "WEBHOOK_WRITE"
	TokenScopeWebhookManage = "WEBHOOK_MANAGE"

	TokenScopeVariableRead   = "VARIABLE_READ"
	TokenScopeVariableWrite  = "VARIABLE_WRITE"
	TokenScopeVariableManage = "VARIABLE_MANAGE"

	TokenScopeEnvironmentRead   = "ENVIRONMENT_READ"
	TokenScopeEnvironmentWrite  = "ENVIRONMENT_WRITE"
	TokenScopeEnvironmentManage = "ENVIRONMENT_MANAGE"

	TokenScopeDistributionRead  = "DISTRIBUTION_READ"
	TokenScopeDistributionWrite = "DISTRIBUTION_WRITE"

	TokenScopeSandboxRead  = "SANDBOX_READ"
	TokenScopeSandboxWrite = "SANDBOX_WRITE"

	TokenScopeUnitTestRead   = "UNIT_TEST_READ"
	TokenScopeUnitTestWrite  = "UNIT_TEST_WRITE"
	TokenScopeUnitTestManage = "UNIT_TEST_MANAGE"

	TokenScopeVisualTestRead  = "VISUAL_TEST_READ"
	TokenScopeVisualTestWrite = "VISUAL_TEST_WRITE"

	TokenScopeCrawlRead   = "CRAWL_READ"
	TokenScopeCrawlWrite  = "CRAWL_WRITE"
	TokenScopeCrawlManage = "CRAWL_MANAGE"

	TokenScopeTunnelRead   = "TUNNEL_READ"
	TokenScopeTunnelManage = "TUNNEL_MANAGE"
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
