package buddy

import (
	"net/http"
)

type PublicKeyService struct {
	client *Client
}

type PublicKey struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PublicKeyOps struct {
	Content *string `json:"content"`
	Title   *string `json:"title,omitempty"`
}

func (s *PublicKeyService) Create(ops *PublicKeyOps) (*PublicKey, *http.Response, error) {
	var k *PublicKey
	resp, err := s.client.Create(s.client.NewUrlPath("/user/keys"), &ops, nil, &k)
	return k, resp, err
}

func (s *PublicKeyService) Delete(keyId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/user/keys/%d", keyId), nil, nil)
}

func (s *PublicKeyService) Get(keyId int) (*PublicKey, *http.Response, error) {
	var k *PublicKey
	resp, err := s.client.Get(s.client.NewUrlPath("/user/keys/%d", keyId), &k, nil)
	return k, resp, err
}
