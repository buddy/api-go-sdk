package buddy

import (
	"net/http"
)

type MemberService struct {
	client *Client
}

type Member struct {
	Url                       string `json:"url"`
	HtmlUrl                   string `json:"html_url"`
	Id                        int    `json:"id"`
	Name                      string `json:"name"`
	Email                     string `json:"email"`
	AvatarUrl                 string `json:"avatar_url"`
	AutoAssignToNewProjects   bool   `json:"auto_assign_to_new_projects"`
	AutoAssignPermissionSetId int    `json:"auto_assign_permission_set_id"`
	Status                    string `json:"status"`
	Admin                     bool   `json:"admin"`
	WorkspaceOwner            bool   `json:"workspace_owner"`
}

type Members struct {
	Url     string    `json:"url"`
	HtmlUrl string    `json:"html_url"`
	Members []*Member `json:"members"`
}

type MemberCreateOps struct {
	Email *string `json:"email"`
}

type MemberUpdateOps struct {
	Admin                     *bool `json:"admin,omitempty"`
	AutoAssignToNewProjects   *bool `json:"auto_assign_to_new_projects,omitempty"`
	AutoAssignPermissionSetId *int  `json:"auto_assign_permission_set_id,omitempty"`
}

func (s *MemberService) Get(domain string, memberId int) (*Member, *http.Response, error) {
	var m *Member
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/members/%d", domain, memberId), &m, nil)
	return m, resp, err
}

func (s *MemberService) GetList(domain string, query *PageQuery) (*Members, *http.Response, error) {
	var l *Members
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/members", domain), &l, query)
	return l, resp, err
}

func (s *MemberService) GetListAll(domain string) (*Members, *http.Response, error) {
	var all Members
	page := 1
	perPage := 30
	for {
		var l *Members
		l, resp, err := s.GetList(domain, &PageQuery{
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

func (s *MemberService) Create(domain string, ops *MemberCreateOps) (*Member, *http.Response, error) {
	var m *Member
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/members", domain), &ops, nil, &m)
	return m, resp, err
}

func (s *MemberService) Delete(domain string, memberId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/members/%d", domain, memberId), nil, nil)
}

func (s *MemberService) Update(domain string, memberId int, ops *MemberUpdateOps) (*Member, *http.Response, error) {
	var m *Member
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/members/%d", domain, memberId), &ops, nil, &m)
	return m, resp, err
}
