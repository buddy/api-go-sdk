package buddy

import (
	"net/http"
)

const (
	PermissionAccessLevelDenied    = "DENIED"
	PermissionAccessLevelReadOnly  = "READ_ONLY"
	PermissionAccessLevelReadWrite = "READ_WRITE"
	PermissionAccessLevelManage    = "MANAGE"
	PermissionAccessLevelRunOnly   = "RUN_ONLY"

	PermissionTypeDeveloper      = "DEVELOPER"
	PermissionTypeReadOnly       = "READ_ONLY"
	PermissionTypeProjectManager = "PROJECT_MANAGER"
	PermissionTypeCustom         = "CUSTOM"
)

type Permission struct {
	Url                    string `json:"url"`
	HtmlUrl                string `json:"html_url"`
	Id                     int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	Type                   string `json:"type"`
	RepositoryAccessLevel  string `json:"repository_access_level"`
	PipelineAccessLevel    string `json:"pipeline_access_level"`
	SandboxAccessLevel     string `json:"sandbox_access_level"`
	ProjectTeamAccessLevel string `json:"project_team_access_level"`
}

type Permissions struct {
	HtmlUrl        string        `json:"html_url"`
	Url            string        `json:"url"`
	PermissionSets []*Permission `json:"permission_sets"`
}

type PermissionOps struct {
	Description            *string `json:"description,omitempty"`
	Name                   *string `json:"name"`
	PipelineAccessLevel    *string `json:"pipeline_access_level"`
	ProjectTeamAccessLevel *string `json:"project_team_access_level"`
	RepositoryAccessLevel  *string `json:"repository_access_level"`
	SandboxAccessLevel     *string `json:"sandbox_access_level"`
}

type PermissionService struct {
	client *Client
}

func (s *PermissionService) Create(domain string, ops *PermissionOps) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/permissions", domain), &ops, nil, &p)
	return p, resp, err
}

func (s *PermissionService) Delete(domain string, permissionId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/permissions/%d", domain, permissionId), nil, nil)
}

func (s *PermissionService) Update(domain string, permissionId int, ops *PermissionOps) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/permissions/%d", domain, permissionId), &ops, nil, &p)
	return p, resp, err
}

func (s *PermissionService) Get(domain string, permissionId int) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/permissions/%d", domain, permissionId), &p, nil)
	return p, resp, err
}

func (s *PermissionService) GetList(domain string) (*Permissions, *http.Response, error) {
	var l *Permissions
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/permissions", domain), &l, nil)
	return l, resp, err
}
