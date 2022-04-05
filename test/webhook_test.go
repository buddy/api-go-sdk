package test

import (
	"api-go-sdk/buddy"
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

func TestWebhook(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var webhook buddy.Webhook
	t.Run("Create", testWebhookCreate(seed.client, seed.workspace, seed.project, &webhook))
}
