package buddy

import "net/http"

type ProjectGroupService struct {
	client *Client
}

type ProjectGroup struct {
	Group
	PermissionSet *Permission `json:"permission_set"`
}

type ProjectGroupOps struct {
	Id            *int             `json:"id,omitempty"`
	PermissionSet *ProjectGroupOps `json:"permission_set,omitempty"`
}

func (s *ProjectGroupService) CreateProjectGroup(workspaceDomain string, projectName string, ops *ProjectGroupOps) (*ProjectGroup, *http.Response, error) {
	var pg *ProjectGroup
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/groups", workspaceDomain, projectName), &ops, nil, &pg)
	return pg, resp, err
}

func (s *ProjectGroupService) DeleteProjectGroup(workspaceDomain string, projectName string, groupId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/groups/%d", workspaceDomain, projectName, groupId), nil, nil)
}

func (s *ProjectGroupService) GetProjectGroup(workspaceDomain string, projectName string, groupId int) (*ProjectGroup, *http.Response, error) {
	var pg *ProjectGroup
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/groups/%d", workspaceDomain, projectName, groupId), &pg, nil)
	return pg, resp, err
}

func (s *ProjectGroupService) GetProjectGroups(workspaceDomain string, projectName string) (*Groups, *http.Response, error) {
	var all *Groups
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/groups", workspaceDomain, projectName), &all, nil)
	return all, resp, err
}

func (s *ProjectGroupService) UpdateProjectGroup(workspaceDomain string, projectName string, groupId int, ops *ProjectGroupOps) (*ProjectGroup, *http.Response, error) {
	var pg *ProjectGroup
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/groups/%d", workspaceDomain, projectName, groupId), &ops, nil, &pg)
	return pg, resp, err
}
