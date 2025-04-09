package buddy

import (
	"encoding/base64"
	"net/http"
)

const (
	SourceContentTypeDir       = "DIR"
	SourceContentTypeFile      = "FILE"
	SourceContentTypeSubModule = "SUB_MODULE"
	SourceContentTypeSymlink   = "SYMLINK"
)

type SourceService struct {
	client *Client
}

type SourceContent struct {
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	ContentType string `json:"content_type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
}

type SourceCommit struct {
	Url        string  `json:"url"`
	HtmlUrl    string  `json:"html_url"`
	Revision   string  `json:"revision"`
	AuthorDate string  `json:"author_date"`
	CommitDate string  `json:"commit_date"`
	Message    string  `json:"message"`
	Committer  *Member `json:"committer"`
	Author     *Member `json:"author"`
}

type SourceContentsGetQuery struct {
	Revision string `json:"revision,omitempty"`
}

type SourceContents struct {
	SourceContent
	Contents []*SourceContent `json:"contents"`
}

type SourceFile struct {
	Content *SourceContent `json:"content"`
	Commit  *SourceCommit  `json:"commit"`
}

type SourceFileOps struct {
	Content    *string `json:"content,omitempty"`
	ContentRaw *string `json:"content_raw,omitempty"`
	Message    *string `json:"message,omitempty"`
	Path       *string `json:"path,omitempty"`
	Branch     *string `json:"branch,omitempty"`
}

func (s *SourceService) prepareContent(ops *SourceFileOps) {
	if ops != nil && ops.ContentRaw != nil {
		b := base64.StdEncoding.EncodeToString([]byte(*ops.ContentRaw))
		ops.Content = &b
		ops.ContentRaw = nil
	}
}

func (s *SourceService) CreateFile(workspaceDomain string, projectName string, ops *SourceFileOps) (*SourceFile, *http.Response, error) {
	var c *SourceFile
	s.prepareContent(ops)
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/repository/contents", workspaceDomain, projectName), &ops, nil, &c)
	return c, resp, err
}

func (s *SourceService) UpdateFile(workspaceDomain string, projectName string, path string, ops *SourceFileOps) (*SourceFile, *http.Response, error) {
	var c *SourceFile
	s.prepareContent(ops)
	resp, err := s.client.Put(s.client.NewUrlPath("/workspaces/%s/projects/%s/repository/contents/%s", workspaceDomain, projectName, path), &ops, nil, &c)
	return c, resp, err
}

func (s *SourceService) Get(workspaceDomain string, projectName string, path string, query *SourceContentsGetQuery) (*SourceContents, *http.Response, error) {
	var c *SourceContents
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/repository/contents/%s", workspaceDomain, projectName, path), &c, query)
	if c.ContentType == "" {
		c.ContentType = SourceContentTypeDir
	}
	return c, resp, err
}

func (s *SourceService) DeleteFile(workspaceDomain string, projectName string, path string, ops *SourceFileOps) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/repository/contents/%s", workspaceDomain, projectName, path), &ops, nil)
}
