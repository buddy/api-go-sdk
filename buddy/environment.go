package buddy

import "net/http"

const (
	EnvironmentPermissionAccessLevelManage  = "MANAGE"
	EnvironmentPermissionAccessLevelUseOnly = "USE_ONLY"
	EnvironmentPermissionAccessLevelDefault = "DEFAULT"
	EnvironmentPermissionAccessLevelDenied  = "DENIED"
)

type EnvironmentService struct {
	client *Client
}

type EnvironmentResourcePermissions struct {
	Id          int    `json:"id"`
	AccessLevel string `json:"access_level"`
}

type EnvironmentPermissions struct {
	Others string                            `json:"others"`
	Users  []*EnvironmentResourcePermissions `json:"users"`
	Groups []*EnvironmentResourcePermissions `json:"groups"`
}

type Environment struct {
	Url                 string                  `json:"url"`
	HtmlUrl             string                  `json:"html_url"`
	Id                  string                  `json:"id"`
	Name                string                  `json:"name"`
	Identifier          string                  `json:"identifier"`
	Tags                []string                `json:"tags"`
	PublicUrl           string                  `json:"public_url"`
	AllPipelinesAllowed bool                    `json:"all_pipelines_allowed"`
	AllowedPipelines    []*AllowedPipeline      `json:"allowed_pipelines"`
	Project             *Project                `json:"project"`
	Variables           []*Variable             `json:"variables"`
	Permissions         *EnvironmentPermissions `json:"permissions"`
}

type Environments struct {
	Url          string         `json:"url"`
	HtmlUrl      string         `json:"html_url"`
	Environments []*Environment `json:"environments"`
}

type EnvironmentOps struct {
	Name                *string                 `json:"name,omitempty"`
	Identifier          *string                 `json:"identifier,omitempty"`
	PublicUrl           *string                 `json:"public_url,omitempty"`
	Tags                *[]string               `json:"tags,omitempty"`
	Variables           *[]*Variable            `json:"variables,omitempty"`
	Permissions         *EnvironmentPermissions `json:"permissions,omitempty"`
	AllPipelinesAllowed *bool                   `json:"all_pipelines_allowed,omitempty"`
	AllowedPipelines    *[]*AllowedPipeline     `json:"allowed_pipelines,omitempty"`
}

func (s *EnvironmentService) Create(workspaceDomain string, projectName string, ops *EnvironmentOps) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/environments", workspaceDomain, projectName), &ops, nil, &e)
	return e, resp, err
}

func (s *EnvironmentService) Update(workspaceDomain string, projectName string, environmentId string, ops *EnvironmentOps) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/environments/%s", workspaceDomain, projectName, environmentId), &ops, nil, &e)
	return e, resp, err
}

func (s *EnvironmentService) Get(workspaceDomain string, projectName string, environmentId string) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/environments/%s", workspaceDomain, projectName, environmentId), &e, nil)
	return e, resp, err
}

func (s *EnvironmentService) GetList(workspaceDomain string, projectName string) (*Environments, *http.Response, error) {
	var e *Environments
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/environments", workspaceDomain, projectName), &e, nil)
	return e, resp, err
}

func (s *EnvironmentService) Delete(workspaceDomain string, projectName string, environmentId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/environments/%s", workspaceDomain, projectName, environmentId), nil, nil)
}
