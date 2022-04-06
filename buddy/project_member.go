package buddy

import (
	"net/http"
)

type ProjectMemberService struct {
	client *Client
}

type ProjectMember struct {
	Member
	PermissionSet *Permission `json:"permission_set"`
}

type ProjectMemberOps struct {
	Id            *int              `json:"id,omitempty"`
	PermissionSet *ProjectMemberOps `json:"permission_set,omitempty"`
}

func (s *ProjectMemberService) CreateProjectMember(domain string, projectName string, ops *ProjectMemberOps) (*ProjectMember, *http.Response, error) {
	var pm *ProjectMember
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/members", domain, projectName), &ops, nil, &pm)
	return pm, resp, err
}

func (s *ProjectMemberService) DeleteProjectMember(domain string, projectName string, memberId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/members/%d", domain, projectName, memberId), nil, nil)
}

func (s *ProjectMemberService) GetProjectMember(domain string, projectName string, memberId int) (*ProjectMember, *http.Response, error) {
	var pm *ProjectMember
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/members/%d", domain, projectName, memberId), &pm, nil)
	return pm, resp, err
}

func (s *ProjectMemberService) GetProjectMembers(domain string, projectName string, query *PageQuery) (*Members, *http.Response, error) {
	var l *Members
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/members", domain, projectName), &l, query)
	return l, resp, err
}

func (s *ProjectMemberService) GetProjectMembersAll(domain string, projectName string) (*Members, *http.Response, error) {
	var all Members
	page := 1
	perPage := 30
	for {
		l, resp, err := s.GetProjectMembers(domain, projectName, &PageQuery{
			Page:    page,
			PerPage: perPage,
		})
		if err != nil {
			return nil, resp, err
		}
		if len(l.Members) == 0 {
			break
		}
		all.Url = l.Url
		all.HtmlUrl = l.HtmlUrl
		all.Members = append(all.Members, l.Members...)
		page += 1
	}
	return &all, nil, nil
}

func (s *ProjectMemberService) UpdateProjectMember(domain string, projectName string, memberId int, ops *ProjectMemberOps) (*ProjectMember, *http.Response, error) {
	var pm *ProjectMember
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/members/%d", domain, projectName, memberId), &ops, nil, &pm)
	return pm, resp, err
}
