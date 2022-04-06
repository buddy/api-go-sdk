package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testWebhookCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Webhook) func(t *testing.T) {
	return func(t *testing.T) {
		secretKey := RandString(10)
		targetUrl := "https://" + RandString(10) + ".0zxc.com"
		event := buddy.WebhookEventPush
		events := []string{event}
		projects := []string{project.Name}
		ops := buddy.WebhookOps{
			SecretKey: &secretKey,
			TargetUrl: &targetUrl,
			Events:    &events,
			Projects:  &projects,
		}
		webhook, _, err := client.WebhookService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("WebhookService.Create", err))
		}
		err = CheckWebhook(webhook, targetUrl, secretKey, project.Name, event, 0)
		if err != nil {
			t.Fatal(err)
		}
		*out = *webhook
	}
}

func testWebhookUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Webhook) func(t *testing.T) {
	return func(t *testing.T) {
		secretKey := RandString(10)
		targetUrl := "https://" + RandString(10) + ".0zxc.com"
		event := buddy.WebhookEventExecutionStarted
		events := []string{event}
		projects := []string{project.Name}
		ops := buddy.WebhookOps{
			SecretKey: &secretKey,
			TargetUrl: &targetUrl,
			Events:    &events,
			Projects:  &projects,
		}
		webhook, _, err := client.WebhookService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("WebhookService.Update", err))
		}
		err = CheckWebhook(webhook, targetUrl, secretKey, project.Name, event, out.Id)
		if err != nil {
			t.Fatal(err)
		}
		*out = *webhook
	}
}

func testWebhookGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Webhook) func(t *testing.T) {
	return func(t *testing.T) {
		webhook, _, err := client.WebhookService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("WebhookService.Get", err))
		}
		err = CheckWebhook(webhook, out.TargetUrl, out.SecretKey, project.Name, out.Events[0], out.Id)
		if err != nil {
			t.Fatal(err)
		}
		*out = *webhook
	}
}

func testWebhookGetList(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		webhooks, _, err := client.WebhookService.GetList(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("WebhookService.GetList", err))
		}
		err = CheckWebhooks(webhooks, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testWebhookDelete(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Webhook) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.WebhookService.Delete(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("WebhookService.Delete", err))
		}
	}
}

func TestWebhook(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var webhook buddy.Webhook
	t.Run("Create", testWebhookCreate(seed.Client, seed.Workspace, seed.Project, &webhook))
	t.Run("Update", testWebhookUpdate(seed.Client, seed.Workspace, seed.Project, &webhook))
	t.Run("Get", testWebhookGet(seed.Client, seed.Workspace, seed.Project, &webhook))
	t.Run("GetList", testWebhookGetList(seed.Client, seed.Workspace, 1))
	t.Run("Delete", testWebhookDelete(seed.Client, seed.Workspace, &webhook))
}
