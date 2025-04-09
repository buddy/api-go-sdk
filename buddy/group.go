package buddy

import (
	"net/http"
)

const (
	GroupMemberStatusManager = "MANAGER"
	GroupMemberStatusMember  = "MEMBER"
)

type GroupService struct {
	client *Client
}

type Group struct {
	Url                       string `json:"url"`
	HtmlUrl                   string `json:"html_url"`
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	AutoAssignToNewProjects   bool   `json:"auto_assign_to_new_projects"`
	AutoAssignPermissionSetId int    `json:"auto_assign_permission_set_id"`
}

type Groups struct {
	Url     string   `json:"url"`
	HtmlUrl string   `json:"html_url"`
	Groups  []*Group `json:"groups"`
}

type GroupOps struct {
	Name                      *string `json:"name,omitempty"`
	Description               *string `json:"description,omitempty"`
	AutoAssignToNewProjects   *bool   `json:"auto_assign_to_new_projects,omitempty"`
	AutoAssignPermissionSetId *int    `json:"auto_assign_permission_set_id,omitempty"`
}

type GroupMemberOps struct {
	Id     *int    `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

func (s *GroupService) Get(workspaceDomain string, groupId int) (*Group, *http.Response, error) {
	var g *Group
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/groups/%d", workspaceDomain, groupId), &g, nil)
	return g, resp, err
}

func (s *GroupService) GetList(workspaceDomain string) (*Groups, *http.Response, error) {
	var l *Groups
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/groups", workspaceDomain), &l, nil)
	return l, resp, err
}

func (s *GroupService) Create(workspaceDomain string, ops *GroupOps) (*Group, *http.Response, error) {
	var g *Group
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/groups", workspaceDomain), &ops, nil, &g)
	return g, resp, err
}

func (s *GroupService) Update(workspaceDomain string, groupId int, ops *GroupOps) (*Group, *http.Response, error) {
	var g *Group
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/groups/%d", workspaceDomain, groupId), &ops, nil, &g)
	return g, resp, err
}

func (s *GroupService) Delete(workspaceDomain string, groupId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/groups/%d", workspaceDomain, groupId), nil, nil)
}

func (s *GroupService) AddGroupMember(workspaceDomain string, groupId int, ops *GroupMemberOps) (*Member, *http.Response, error) {
	var m *Member
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/groups/%d/members", workspaceDomain, groupId), &ops, nil, &m)
	return m, resp, err
}

func (s *GroupService) DeleteGroupMember(workspaceDomain string, groupId int, memberId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/groups/%d/members/%d", workspaceDomain, groupId, memberId), nil, nil)
}

func (s *GroupService) UpdateGroupMember(workspaceDomain string, groupId int, memberId int, status string) (*Member, *http.Response, error) {
	var m *Member
	ops := GroupMemberOps{
		Status: &status,
	}
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/groups/%d/members/%d", workspaceDomain, groupId, memberId), &ops, nil, &m)
	return m, resp, err
}

func (s *GroupService) GetGroupMember(workspaceDomain string, groupId int, memberId int) (*Member, *http.Response, error) {
	var m *Member
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/groups/%d/members/%d", workspaceDomain, groupId, memberId), &m, nil)
	return m, resp, err
}

func (s *GroupService) GetGroupMembers(workspaceDomain string, groupId int) (*Members, *http.Response, error) {
	var l *Members
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/groups/%d/members", workspaceDomain, groupId), &l, nil)
	return l, resp, err
}
