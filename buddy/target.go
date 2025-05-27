package buddy

import "net/http"

const (
	TargetTypeFtp          = "FTP"
	TargetTypeSsh          = "SSH"
	TargetTypeMatch        = "MATCH"
	TargetTypeUpcloud      = "UPCLOUD"
	TargetTypeVultr        = "VULTR"
	TargetTypeDigitalOcean = "DIGITAL_OCEAN"
	TargetTypeGit          = "GIT"

	TargetScopeWorkspace   = "WORKSPACE"
	TargetScopeProject     = "PROJECT"
	TargetScopeEnvironment = "ENVIRONMENT"
	TargetScopePipeline    = "PIPELINE"
	TargetScopeAction      = "ACTION"
	TargetScopeAny         = "ANY"

	TargetAuthMethodPassword         = "PASSWORD"
	TargetAuthMethodSshKey           = "SSH_KEY"
	TargetAuthMethodAssetsKey        = "ASSETS_KEY"
	TargetAuthMethodProxyCredentials = "PROXY_CREDENTIALS"
	TargetAuthMethodProxyKey         = "PROXY_KEY"
	TargetAuthMethodHttp             = "HTTP"

	TargetPermissionManage  = "MANAGE"
	TargetPermissionUseOnly = "USE_ONLY"
)

type TargetService struct {
	client *Client
}

type TargetAuth struct {
	Method     string `json:"method"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Asset      string `json:"asset"`
	Passphrase string `json:"passphrase"`
	Key        string `json:"key"`
	KeyPath    string `json:"key_path"`
}

type TargetProject struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

type TargetPipeline struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TargetEnvironment struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TargetProxy struct {
	Name string      `json:"name"`
	Host string      `json:"host"`
	Port string      `json:"port"`
	Auth *TargetAuth `json:"auth"`
}

type TargetResourcePermission struct {
	Id          int    `json:"id"`
	AccessLevel string `json:"access_level"`
}

type TargetPermissions struct {
	Others string                      `json:"others"`
	Users  []*TargetResourcePermission `json:"users"`
	Groups []*TargetResourcePermission `json:"groups"`
}

type Target struct {
	Url         string             `json:"url"`
	HtmlUrl     string             `json:"html_url"`
	Id          string             `json:"id"`
	Identifier  string             `json:"identifier"`
	Tags        []string           `json:"tags"`
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	Host        string             `json:"host"`
	Scope       string             `json:"scope"`
	Repository  string             `json:"repository"`
	Port        string             `json:"port"`
	Path        string             `json:"path"`
	Secure      bool               `json:"secure"`
	Integration string             `json:"integration"`
	Disabled    bool               `json:"disabled"`
	Auth        *TargetAuth        `json:"auth"`
	Proxy       *TargetProxy       `json:"proxy"`
	Permissions *TargetPermissions `json:"permissions"`
}

type TargetOps struct {
	Identifier  *string            `json:"identifier,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Tags        *[]string          `json:"tags,omitempty"`
	Type        *string            `json:"type,omitempty"`
	Host        *string            `json:"host,omitempty"`
	Scope       *string            `json:"scope,omitempty"`
	Repository  *string            `json:"repository,omitempty"`
	Port        *string            `json:"port,omitempty"`
	Path        *string            `json:"path,omitempty"`
	Secure      *bool              `json:"secure,omitempty"`
	Integration *string            `json:"integration,omitempty"`
	Disabled    *bool              `json:"disabled,omitempty"`
	Auth        *TargetAuth        `json:"auth,omitempty"`
	Project     *TargetProject     `json:"project,omitempty"`
	Pipeline    *TargetPipeline    `json:"pipeline,omitempty"`
	Environment *TargetEnvironment `json:"environment,omitempty"`
	Proxy       *TargetProxy       `json:"proxy,omitempty"`
	Permissions *TargetPermissions `json:"permissions,omitempty"`
}

type TargetGetListQuery struct {
	ProjectName   string `url:"projectName,omitempty"`
	PipelineId    int    `url:"pipelineId,omitempty"`
	ActionId      int    `url:"actionId,omitempty"`
	EnvironmentId string `url:"environmentId,omitempty"`
}

type Targets struct {
	Targets []*Target `json:"targets"`
}

func (s *TargetService) Create(workspaceDomain string, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/targets", workspaceDomain), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Delete(workspaceDomain string, targetId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/targets/%s", workspaceDomain, targetId), nil, nil)
}

func (s *TargetService) Update(workspaceDomain string, targetId string, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/targets/%s", workspaceDomain, targetId), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Get(workspaceDomain string, targetId string) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/targets/%s", workspaceDomain, targetId), &t, nil)
	return t, resp, err
}

func (s *TargetService) GetList(workspaceDomain string, query *TargetGetListQuery) (*Targets, *http.Response, error) {
	var t *Targets
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/targets", workspaceDomain), &t, &query)
	return t, resp, err
}
