package buddy

import "net/http"

type TokenService struct {
	client *Client
}

type Token struct {
	Url                   string   `json:"url"`
	HtmlUrl               string   `json:"html_url"`
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	ExpiresAt             string   `json:"expires_at"`
	Scopes                []string `json:"scopes"`
	IpRestrictions        []string `json:"ip_restrictions"`
	WorkspaceRestrictions []string `json:"workspace_restrictions"`
}

type TokenOps struct {
	Name                  *string   `json:"name,omitempty"`
	ExpiresIn             *int      `json:"expires_in,omitempty"`
	ExpiresAt             *string   `json:"expires_at,omitempty"`
	Scopes                *[]string `json:"scopes,omitempty"`
	IpRestrictions        *[]string `json:"ip_restrictions,omitempty"`
	WorkspaceRestrictions *[]string `json:"workspace_restrictions,omitempty"`
}

type Tokens struct {
	Url          string   `json:"url"`
	HtmlUrl      string   `json:"html_url"`
	AccessTokens []*Token `json:"access_tokens"`
}

func (s *TokenService) Create(ops *TokenOps) (*Token, *http.Response, error) {
	var t *Token
	resp, err := s.client.Create(s.client.NewUrlPath("/user/tokens"), &ops, nil, &t)
	return t, resp, err
}

func (s *TokenService) Get(tokenId string) (*Token, *http.Response, error) {
	var t *Token
	resp, err := s.client.Get(s.client.NewUrlPath("/user/tokens/%s", tokenId), &t, nil)
	return t, resp, err
}

func (s *TokenService) GetList() (*Tokens, *http.Response, error) {
	var t *Tokens
	resp, err := s.client.Get(s.client.NewUrlPath("/user/tokens"), &t, nil)
	return t, resp, err
}

func (s *TokenService) Delete(tokenId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/user/tokens/%s", tokenId), nil, nil)
}
