package buddy

import (
	"net/http"
)

const (
	WebhookEventPush                = "PUSH"
	WebhookEventExecutionStarted    = "EXECUTION_STARTED"
	WebhookEventExecutionSuccessful = "EXECUTION_SUCCESSFUL"
	WebhookEventExecutionFailed     = "EXECUTION_FAILED"
	WebhookEventExecutionFinished   = "EXECUTION_FINISHED"
)

type WebhookService struct {
	client *Client
}

type Webhook struct {
	Url       string            `json:"url"`
	HtmlUrl   string            `json:"html_url"`
	Id        int               `json:"id"`
	TargetUrl string            `json:"target_url"`
	SecretKey string            `json:"secret_key"`
	Projects  []string          `json:"projects"`
	Events    []string          `json:"events"`
	Requests  []*WebhookRequest `json:"requests"`
}

type Webhooks struct {
	Url      string     `json:"url"`
	HtmlUrl  string     `json:"html_url"`
	Webhooks []*Webhook `json:"webhooks"`
}

type WebhookRequest struct {
	PostDate       string `json:"post_date"`
	ResponseStatus string `json:"response_status"`
	Body           string `json:"body"`
}

type WebhookOps struct {
	Events    *[]string `json:"events,omitempty"`
	Projects  *[]string `json:"projects,omitempty"`
	TargetUrl *string   `json:"target_url,omitempty"`
	SecretKey *string   `json:"secret_key,omitempty"`
}

func (s *WebhookService) Create(workspaceDomain string, ops *WebhookOps) (*Webhook, *http.Response, error) {
	var w *Webhook
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/webhooks", workspaceDomain), &ops, nil, &w)
	return w, resp, err
}

func (s *WebhookService) Delete(workspaceDomain string, webhookId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/webhooks/%d", workspaceDomain, webhookId), nil, nil)
}

func (s *WebhookService) Update(workspaceDomain string, webhookId int, ops *WebhookOps) (*Webhook, *http.Response, error) {
	var w *Webhook
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/webhooks/%d", workspaceDomain, webhookId), &ops, nil, &w)
	return w, resp, err
}

func (s *WebhookService) Get(workspaceDomain string, webhookId int) (*Webhook, *http.Response, error) {
	var w *Webhook
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/webhooks/%d", workspaceDomain, webhookId), &w, nil)
	return w, resp, err
}

func (s *WebhookService) GetList(workspaceDomain string) (*Webhooks, *http.Response, error) {
	var all *Webhooks
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/webhooks", workspaceDomain), &all, nil)
	return all, resp, err
}
