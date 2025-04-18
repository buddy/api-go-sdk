package buddy

import (
	"net/http"
)

type WorkspaceService struct {
	client *Client
}

type WorkspaceUpdateOps struct {
	EncryptionSalt *string `json:"encryption_salt,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type WorkspaceCreateOps struct {
	EncryptionSalt *string `json:"encryption_salt,omitempty"`
	Name           *string `json:"name,omitempty"`
	Domain         *string `json:"domain,omitempty"`
}

type Workspace struct {
	Url        string `json:"url"`
	HtmlUrl    string `json:"html_url"`
	Id         int    `json:"id"`
	OwnerId    int    `json:"owner_id"`
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Frozen     bool   `json:"frozen"`
	CreateDate string `json:"create_date"`
}

type Workspaces struct {
	Url        string       `json:"url"`
	HtmlUrl    string       `json:"html_url"`
	Workspaces []*Workspace `json:"workspaces"`
}

func (s *WorkspaceService) Create(ops *WorkspaceCreateOps) (*Workspace, *http.Response, error) {
	var w *Workspace
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces"), &ops, nil, &w)
	return w, resp, err
}

func (s *WorkspaceService) Delete(workspaceDomain string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("workspaces/%s", workspaceDomain), nil, nil)
}

func (s *WorkspaceService) Update(workspaceDomain string, ops *WorkspaceUpdateOps) (*Workspace, *http.Response, error) {
	var w *Workspace
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s", workspaceDomain), &ops, nil, &w)
	return w, resp, err
}

func (s *WorkspaceService) Get(workspaceDomain string) (*Workspace, *http.Response, error) {
	var w *Workspace
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s", workspaceDomain), &w, nil)
	return w, resp, err
}

func (s *WorkspaceService) GetList() (*Workspaces, *http.Response, error) {
	var all *Workspaces
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces"), &all, nil)
	return all, resp, err
}
