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

func (s *PermissionService) Create(workspaceDomain string, ops *PermissionOps) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/permissions", workspaceDomain), &ops, nil, &p)
	return p, resp, err
}

func (s *PermissionService) Delete(workspaceDomain string, permissionId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/permissions/%d", workspaceDomain, permissionId), nil, nil)
}

func (s *PermissionService) Update(workspaceDomain string, permissionId int, ops *PermissionOps) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/permissions/%d", workspaceDomain, permissionId), &ops, nil, &p)
	return p, resp, err
}

func (s *PermissionService) Get(workspaceDomain string, permissionId int) (*Permission, *http.Response, error) {
	var p *Permission
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/permissions/%d", workspaceDomain, permissionId), &p, nil)
	return p, resp, err
}

func (s *PermissionService) GetList(workspaceDomain string) (*Permissions, *http.Response, error) {
	var l *Permissions
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/permissions", workspaceDomain), &l, nil)
	return l, resp, err
}
