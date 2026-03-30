package buddy

import "net/http"

const (
	EnvironmentPermissionAccessLevelManage  = "MANAGE"
	EnvironmentPermissionAccessLevelUseOnly = "USE_ONLY"
	EnvironmentPermissionAccessLevelDefault = "DEFAULT"
	EnvironmentPermissionAccessLevelDenied  = "DENIED"

	EnvironmentScopeProject   = "PROJECT"
	EnvironmentScopeWorkspace = "WORKSPACE"

	EnvironmentAccessLevelUseOnly = "USE_ONLY"
	EnvironmentAccessLevelDenied  = "DENIED"
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

type EnvironmentAllowedPipeline struct {
	Project     string `json:"project"`
	Pipeline    string `json:"pipeline"`
	AccessLevel string `json:"access_level"`
}

type EnvironmentAllowedEnvironment struct {
	Project     string `json:"project"`
	Environment string `json:"environment"`
	AccessLevel string `json:"access_level"`
}

type Environment struct {
	Url                     string                           `json:"url"`
	HtmlUrl                 string                           `json:"html_url"`
	Id                      string                           `json:"id"`
	Name                    string                           `json:"name"`
	Identifier              string                           `json:"identifier"`
	Tags                    []string                         `json:"tags"`
	Icon                    string                           `json:"icon"`
	PublicUrl               string                           `json:"public_url"`
	PipelinesAccessLevel    string                           `json:"pipelines_access_level"`
	EnvironmentsAccessLevel string                           `json:"environments_access_level"`
	AllowedPipelines        []*EnvironmentAllowedPipeline    `json:"allowed_pipelines"`
	AllowedEnvironments     []*EnvironmentAllowedEnvironment `json:"allowed_environments"`
	CreateDate              string                           `json:"create_date"`
	Project                 *Project                         `json:"project"`
	Variables               []*Variable                      `json:"variables"`
	Permissions             *EnvironmentPermissions          `json:"permissions"`
	BaseOnly                bool                             `json:"base_only,omitempty"`
	BaseEnvironments        []string                         `json:"base_environments"`
	Scope                   string                           `json:"scope,omitempty"`
	Targets                 []*Target                        `json:"targets"`
}

type Environments struct {
	Url          string         `json:"url"`
	HtmlUrl      string         `json:"html_url"`
	Environments []*Environment `json:"environments"`
}

type EnvironmentOps struct {
	Name                    *string                           `json:"name,omitempty"`
	Identifier              *string                           `json:"identifier,omitempty"`
	PublicUrl               *string                           `json:"public_url,omitempty"`
	Icon                    *string                           `json:"icon,omitempty"`
	Tags                    *[]string                         `json:"tags,omitempty"`
	Permissions             *EnvironmentPermissions           `json:"permissions,omitempty"`
	PipelinesAccessLevel    *string                           `json:"pipelines_access_level,omitempty"`
	EnvironmentsAccessLevel *string                           `json:"environments_access_level,omitempty"`
	AllowedPipelines        *[]*EnvironmentAllowedPipeline    `json:"allowed_pipelines,omitempty"`
	AllowedEnvironments     *[]*EnvironmentAllowedEnvironment `json:"allowed_environments,omitempty"`
	Project                 *ProjectSimple                    `json:"project,omitempty"`
	BaseEnvironments        *[]string                         `json:"base_environments,omitempty"`
	BaseOnly                *bool                             `json:"base_only,omitempty"`
	Scope                   *string                           `json:"scope,omitempty"`
}

func (s *EnvironmentService) Create(workspaceDomain string, ops *EnvironmentOps) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/environments", workspaceDomain), &ops, nil, &e)
	return e, resp, err
}

func (s *EnvironmentService) Get(workspaceDomain string, environmentId string) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/environments/%s", workspaceDomain, environmentId), &e, nil)
	return e, resp, err
}

func (s *EnvironmentService) GetList(workspaceDomain string, projectName string) (*Environments, *http.Response, error) {
	var e *Environments
	query := Query{}
	if projectName != "" {
		query.ProjectName = &projectName
	}
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/environments", workspaceDomain), &e, query)
	return e, resp, err
}

func (s *EnvironmentService) Update(workspaceDomain string, environmentId string, ops *EnvironmentOps) (*Environment, *http.Response, error) {
	var e *Environment
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/environments/%s", workspaceDomain, environmentId), &ops, nil, &e)
	return e, resp, err
}

func (s *EnvironmentService) Delete(workspaceDomain string, environmentId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/environments/%s", workspaceDomain, environmentId), nil, nil)
}
