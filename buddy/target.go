package buddy

import "net/http"

// TargetService handles communication with target related methods of the Buddy API.
type TargetService struct {
	client *Client
}

type TargetAuth struct {
	Method     string `json:"method"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Asset      string `json:"asset,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Key        string `json:"key,omitempty"`
	KeyPath    string `json:"key_path,omitempty"`
}

type TargetProxy struct {
	Name string      `json:"name"`
	Host string      `json:"host"`
	Port string      `json:"port"`
	Auth *TargetAuth `json:"auth"`
}

type TargetProject struct {
	Name string `json:"name"`
}

type TargetPipeline struct {
	Id int `json:"id"`
}

type TargetEnvironment struct {
	Id string `json:"id"`
}

type TargetMatch struct {
	Identifier string   `json:"identifier"`
	Scope      string   `json:"scope"`
	Tags       []string `json:"tags"`
}

type Target struct {
	Url         string               `json:"url"`
	HtmlUrl     string               `json:"html_url"`
	Id          string               `json:"id"`
	Identifier  string               `json:"identifier"`
	Permissions *PipelinePermissions `json:"permissions"`
	Tags        []string             `json:"tags"`
	Name        string               `json:"name"`
	Host        string               `json:"host"`
	Port        string               `json:"port"`
	Path        string               `json:"path"`
	Auth        *TargetAuth          `json:"auth"`
	Type        string               `json:"type"`
	Repository  string               `json:"repository"`
	Project     *TargetProject       `json:"project"`
	Pipeline    *TargetPipeline      `json:"pipeline"`
	Environment *TargetEnvironment   `json:"environment"`
	Integration string               `json:"integration"`
	Secure      bool                 `json:"secure"`
	Proxy       *TargetProxy         `json:"proxy"`
	Match       *TargetMatch         `json:"match"`
}

type Targets struct {
	Url     string    `json:"url"`
	HtmlUrl string    `json:"html_url"`
	Targets []*Target `json:"targets"`
}

type TargetOps struct {
	Name        *string              `json:"name,omitempty"`
	Identifier  *string              `json:"identifier,omitempty"`
	Type        *string              `json:"type,omitempty"`
	Host        *string              `json:"host,omitempty"`
	Auth        *TargetAuth          `json:"auth,omitempty"`
	Repository  *string              `json:"repository,omitempty"`
	Project     *TargetProject       `json:"project,omitempty"`
	Pipeline    *TargetPipeline      `json:"pipeline,omitempty"`
	Environment *TargetEnvironment   `json:"environment,omitempty"`
	Permissions *PipelinePermissions `json:"permissions,omitempty"`
	Port        *string              `json:"port,omitempty"`
	Path        *string              `json:"path,omitempty"`
	Integration *string              `json:"integration,omitempty"`
	Secure      *bool                `json:"secure,omitempty"`
	Proxy       *TargetProxy         `json:"proxy,omitempty"`
	Tags        *[]string            `json:"tags,omitempty"`
	Match       *TargetMatch         `json:"match,omitempty"`
}

func (s *TargetService) Create(workspaceDomain, projectName string, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets", workspaceDomain, projectName), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Update(workspaceDomain, projectName string, targetId string, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%s", workspaceDomain, projectName, targetId), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Get(workspaceDomain, projectName string, targetId string) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%s", workspaceDomain, projectName, targetId), &t, nil)
	return t, resp, err
}

func (s *TargetService) GetList(workspaceDomain, projectName string) (*Targets, *http.Response, error) {
	var l *Targets
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets", workspaceDomain, projectName), &l, nil)
	return l, resp, err
}

func (s *TargetService) Delete(workspaceDomain, projectName string, targetId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%s", workspaceDomain, projectName, targetId), nil, nil)
}
