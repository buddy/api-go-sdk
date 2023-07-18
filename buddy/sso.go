package buddy

import (
	"net/http"
)

const (
	SignatureMethodSha1   = "sha1"
	SignatureMethodSha256 = "sha256"
	SignatureMethodSha512 = "sha512"

	DigestMethodSha1   = "sha1"
	DigestMethodSha256 = "sha256"
	DigestMethodSha512 = "sha512"

	SsoTypeSaml = "SAML"
	SsoTypeOidc = "OIDC"
)

type SsoService struct {
	client *Client
}

type SsoUpdateOps struct {
	SsoUrl                  *string `json:"sso_url,omitempty"`
	Issuer                  *string `json:"issuer,omitempty"`
	Type                    *string `json:"type,omitempty"`
	ClientId                *string `json:"client_id"`
	ClientSecret            *string `json:"client_secret"`
	Certificate             *string `json:"certificate,omitempty"`
	SignatureMethod         *string `json:"signature_method,omitempty"`
	DigestMethod            *string `json:"digest_method,omitempty"`
	RequireSsoForAllMembers *bool   `json:"require_sso_for_all_members,omitempty"`
}

type Sso struct {
	Url                     string `json:"url"`
	HtmlUrl                 string `json:"html_url"`
	Type                    string `json:"type"`
	SsoUrl                  string `json:"sso_url"`
	Issuer                  string `json:"issuer"`
	Certificate             string `json:"certificate"`
	SignatureMethod         string `json:"signature_method"`
	DigestMethod            string `json:"digest_method"`
	RequireSsoForAllMembers bool   `json:"require_sso_for_all_members"`
}

func (s *SsoService) Update(domain string, ops *SsoUpdateOps) (*Sso, *http.Response, error) {
	var r *Sso
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/sso", domain), &ops, nil, &r)
	return r, resp, err
}

func (s *SsoService) Get(domain string) (*Sso, *http.Response, error) {
	var r *Sso
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sso", domain), &r, nil)
	return r, resp, err
}

func (s *SsoService) Enable(domain string) (*http.Response, error) {
	return s.client.Create(s.client.NewUrlPath("/workspaces/%s/enable-sso", domain), nil, nil, nil)
}

func (s *SsoService) Disable(domain string) (*http.Response, error) {
	return s.client.Create(s.client.NewUrlPath("/workspaces/%s/disable-sso", domain), nil, nil, nil)
}
