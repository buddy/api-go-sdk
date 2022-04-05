package buddy

import "net/http"

type SourceService struct {
	client *Client
}

type SourceFileOps struct {
	Content *string `json:"content,omitempty"`
	Message *string `json:"message,omitempty"`
	Path    *string `json:"path,omitempty"`
	Branch  *string `json:"branch,omitempty"`
}

type SourceContent struct {
}

func (s *SourceService) CreateFile(domain string, projectName string, ops *SourceFileOps) (*http.Response, error) {
	var c *SourceContent
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/repository/contents", domain, projectName), &ops, &c)
	return resp, err
}
