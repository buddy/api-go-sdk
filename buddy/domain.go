package buddy

import "net/http"

type DomainService struct {
	client *Client
}

type Record struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Ttl    int      `json:"ttl"`
	Values []string `json:"values"`
}

type Domain struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type DomainCreateOps struct {
	Name *string `json:"name"`
}

type RecordUpsertOps struct {
	Ttl    *int      `json:"ttl,omitempty"`
	Values *[]string `json:"values,omitempty"`
}

type Domains struct {
	Domains []*Domain `json:"domains"`
}

type Records struct {
	Records []*Record `json:"records"`
}

func (s *DomainService) GetList(workspaceDomain string) (*Domains, *http.Response, error) {
	var d *Domains
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains", workspaceDomain), &d, nil)
	return d, resp, err
}

func (s *DomainService) Create(workspaceDomain string, ops *DomainCreateOps) (*Domain, *http.Response, error) {
	var d *Domain
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/domains", workspaceDomain), &ops, nil, &d)
	return d, resp, err
}

func (s *DomainService) GetRecords(workspaceDomain string, domain string) (*Records, *http.Response, error) {
	var r *Records
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains/%s/records", workspaceDomain, domain), &r, nil)
	return r, resp, err
}

func (s *DomainService) GetRecord(workspaceDomain string, domain string, typ string) (*Record, *http.Response, error) {
	var r *Record
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), &r, nil)
	return r, resp, err
}

func (s *DomainService) UpsertRecord(workspaceDomain string, domain string, typ string, ops *RecordUpsertOps) (*Record, *http.Response, error) {
	var r *Record
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), &ops, nil, &r)
	return r, resp, err
}

func (s *DomainService) DeleteRecord(workspaceDomain string, domain string, typ string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), nil, nil)
}
