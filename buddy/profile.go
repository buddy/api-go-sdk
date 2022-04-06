package buddy

import "net/http"

type Profile struct {
	Url           string `json:"url"`
	HtmlUrl       string `json:"html_url"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	AvatarUrl     string `json:"avatar_url"`
	WorkspacesUrl string `json:"workspaces_url"`
}

type ProfileOps struct {
	Name *string `json:"name"`
}

type ProfileService struct {
	client *Client
}

func (s *ProfileService) Get() (*Profile, *http.Response, error) {
	var p *Profile
	resp, err := s.client.Get(s.client.NewUrlPath("/user"), &p, nil)
	return p, resp, err
}

func (s *ProfileService) Update(ops *ProfileOps) (*Profile, *http.Response, error) {
	var p *Profile
	resp, err := s.client.Patch(s.client.NewUrlPath("/user"), &ops, nil, &p)
	return p, resp, err
}
