package buddy

import (
	"net/http"
)

type ProfileEmail struct {
	Email     string `json:"email"`
	Confirmed bool   `json:"confirmed"`
}

type ProfileEmailSet struct {
	Emails []*ProfileEmail `json:"emails"`
}

type ProfileEmailOps struct {
	Email *string `json:"email"`
}

type ProfileEmailService struct {
	client *Client
}

func (s *ProfileEmailService) Create(ops *ProfileEmailOps) (*ProfileEmail, *http.Response, error) {
	var p *ProfileEmail
	resp, err := s.client.Create(s.client.NewUrlPath("/user/emails"), &ops, &p)
	return p, resp, err
}

func (s *ProfileEmailService) Delete(email string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/user/emails/%s", email))
}

func (s *ProfileEmailService) GetList() (*ProfileEmailSet, *http.Response, error) {
	var p *ProfileEmailSet
	resp, err := s.client.Get(s.client.NewUrlPath("/user/emails"), &p, nil)
	return p, resp, err
}
