package buddy

import (
	"net/http"
)

const (
	ProjectStatusActive = "ACTIVE"
	ProjectStatusClosed = "CLOSED"
)

type ProjectService struct {
	client *Client
}

type Project struct {
	Url            string  `json:"url"`
	HtmlUrl        string  `json:"html_url"`
	Name           string  `json:"name"`
	DisplayName    string  `json:"display_name"`
	Status         string  `json:"status"`
	CreateDate     string  `json:"create_date"`
	CreatedBy      *Member `json:"created_by"`
	HttpRepository string  `json:"http_repository"`
	SshRepository  string  `json:"ssh_repository"`
	SshPublicKey   string  `json:"ssh_public_key"`
	KeyFingerprint string  `json:"key_fingerprint"`
	DefaultBranch  string  `json:"default_branch"`
}

type Projects struct {
	Url      string     `json:"url"`
	HtmlUrl  string     `json:"html_url"`
	Projects []*Project `json:"projects"`
}

type ProjectIntegration struct {
	HashId string `json:"hash_id"`
}

type ProjectCreateOps struct {
	Name              *string             `json:"name,omitempty"`
	DisplayName       *string             `json:"display_name,omitempty"`
	ExternalProjectId *string             `json:"external_project_id,omitempty"`
	GitLabProjectId   *string             `json:"git_lab_project_id,omitempty"`
	CustomRepoUrl     *string             `json:"custom_repo_url,omitempty"`
	CustomRepoUser    *string             `json:"custom_repo_user,omitempty"`
	CustomRepoPass    *string             `json:"custom_repo_pass,omitempty"`
	Integration       *ProjectIntegration `json:"integration,omitempty"`
}

type ProjectUpdateOps struct {
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type ProjectListQuery struct {
	PageQuery
	Membership bool   `url:"membership,omitempty"`
	Status     string `url:"status,omitempty"`
}

func (s *ProjectService) Create(domain string, ops *ProjectCreateOps) (*Project, *http.Response, error) {
	var p *Project
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects", domain), &ops, nil, &p)
	return p, resp, err
}

func (s *ProjectService) Delete(domain string, projectName string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s", domain, projectName), nil, nil)
}

func (s *ProjectService) Update(domain string, projectName string, ops *ProjectUpdateOps) (*Project, *http.Response, error) {
	var p *Project
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s", domain, projectName), &ops, nil, &p)
	return p, resp, err
}

func (s *ProjectService) Get(domain string, projectName string) (*Project, *http.Response, error) {
	var p *Project
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s", domain, projectName), &p, nil)
	return p, resp, err
}

func (s *ProjectService) GetList(domain string, query *ProjectListQuery) (*Projects, *http.Response, error) {
	var l *Projects
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects", domain), &l, query)
	return l, resp, err
}

func (s *ProjectService) GetListAll(domain string, query *ProjectListQuery) (*Projects, *http.Response, error) {
	if query == nil {
		query = &ProjectListQuery{}
	}
	query.Page = 1
	query.PerPage = 30
	var all Projects
	for {
		l, resp, err := s.GetList(domain, query)
		if err != nil {
			return nil, resp, err
		}
		if len(l.Projects) == 0 {
			break
		}
		all.Url = l.Url
		all.HtmlUrl = l.HtmlUrl
		all.Projects = append(all.Projects, l.Projects...)
		query.Page += 1
	}
	return &all, nil, nil
}
