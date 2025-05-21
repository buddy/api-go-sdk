package buddy

import "net/http"

// TargetService handles communication with target related methods of the Buddy API.
type TargetService struct {
	client *Client
}

type Target struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type Targets struct {
	Url     string    `json:"url"`
	HtmlUrl string    `json:"html_url"`
	Targets []*Target `json:"targets"`
}

type TargetOps struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (s *TargetService) Create(workspaceDomain, projectName string, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets", workspaceDomain, projectName), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Update(workspaceDomain, projectName string, targetId int, ops *TargetOps) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%d", workspaceDomain, projectName, targetId), &ops, nil, &t)
	return t, resp, err
}

func (s *TargetService) Get(workspaceDomain, projectName string, targetId int) (*Target, *http.Response, error) {
	var t *Target
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%d", workspaceDomain, projectName, targetId), &t, nil)
	return t, resp, err
}

func (s *TargetService) GetList(workspaceDomain, projectName string) (*Targets, *http.Response, error) {
	var l *Targets
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets", workspaceDomain, projectName), &l, nil)
	return l, resp, err
}

func (s *TargetService) Delete(workspaceDomain, projectName string, targetId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/targets/%d", workspaceDomain, projectName, targetId), nil, nil)
}
