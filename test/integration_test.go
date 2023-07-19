package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func TestIntegration(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
		group:     true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	t.Run("Google OIDC", testIntegrationGoogleOIDC(seed.Client, seed.Workspace))
	t.Run("Amazon OIDC", testIntegrationAmazonOidc(seed.Client, seed.Workspace))
	t.Run("Amazon", testIntegrationAmazon(seed.Client, seed.Workspace, seed.Project))
	t.Run("GitHub", testIntegrationGitHub(seed.Client, seed.Workspace))
	t.Run("GitLab", testIntegrationGitLab(seed.Client, seed.Workspace))
	t.Run("DigitalOcean", testIntegrationDigitalOcean(seed.Client, seed.Workspace, seed.Project))
	t.Run("Shopify", testIntegrationShopify(seed.Client, seed.Workspace, seed.Group, seed.Project))
	t.Run("Shopify Partner", testIntegrationShopifyPartner(seed.Client, seed.Workspace, seed.Group, seed.Project))
	t.Run("Stack Hawk", testIntegrationStackHawk(seed.Client, seed.Workspace))
}

func testIntegrationUpdate(client *buddy.Client, workspace *buddy.Workspace, hashId string, ops *buddy.IntegrationOps, out *buddy.Integration) func(t *testing.T) {
	return func(t *testing.T) {
		integrationUpdated, _, err := client.IntegrationService.Update(workspace.Domain, hashId, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("IntegrationService.Patch", err))
		}
		err = CheckIntegration(integrationUpdated, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *integrationUpdated
	}
}

func testIntegrationCreate(client *buddy.Client, workspace *buddy.Workspace, ops *buddy.IntegrationOps, out *buddy.Integration) func(t *testing.T) {
	return func(t *testing.T) {
		integrationAdded, _, err := client.IntegrationService.Create(workspace.Domain, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("IntegrationService.Create", err))
		}
		err = CheckIntegration(integrationAdded, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *integrationAdded
	}
}

func testIntegrationGet(client *buddy.Client, workspace *buddy.Workspace, hashId string, out *buddy.Integration) func(t *testing.T) {
	return func(t *testing.T) {
		integrationGet, _, err := client.IntegrationService.Get(workspace.Domain, hashId)
		if err != nil {
			t.Fatal(ErrorFormatted("IntegrationService.Get", err))
		}
		err = CheckIntegration(integrationGet, out, nil)
		if err != nil {
			t.Fatal(err)
		}
		*out = *integrationGet
	}
}

func testIntegrationGetList(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		integrations, _, err := client.IntegrationService.GetList(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("IntegrationService.GetList", err))
		}
		err = CheckIntegrations(integrations, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testIntegrationDelete(client *buddy.Client, workspace *buddy.Workspace, hashId string) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.IntegrationService.Delete(workspace.Domain, hashId)
		if err != nil {
			t.Fatal(ErrorFormatted("IntegrationService.Delete", err))
		}
	}
}

func testIntegrationStackHawk(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		typ := buddy.IntegrationTypeStackHawk
		scope := buddy.IntegrationScopeWorkspace
		apiKey := RandString(10)
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:   &name,
			Type:   &typ,
			Scope:  &scope,
			ApiKey: &apiKey,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdmin
		updateOps := buddy.IntegrationOps{
			Scope: &newScope,
			Name:  &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationGitLab(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		typ := buddy.IntegrationTypeGitLab
		scope := buddy.IntegrationScopeWorkspace
		token := RandString(10)
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:  &name,
			Type:  &typ,
			Scope: &scope,
			Token: &token,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdmin
		updateOps := buddy.IntegrationOps{
			Scope: &newScope,
			Name:  &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationGitHub(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		typ := buddy.IntegrationTypeGitHub
		scope := buddy.IntegrationScopeWorkspace
		token := RandString(10)
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:  &name,
			Type:  &typ,
			Scope: &scope,
			Token: &token,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdmin
		updateOps := buddy.IntegrationOps{
			Scope: &newScope,
			Name:  &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationAmazonOidc(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		scope := buddy.IntegrationScopeWorkspace
		typ := buddy.IntegrationTypeAmazon
		authType := buddy.IntegrationAuthTypeOidc
		roleAssumptions := []*buddy.RoleAssumption{
			{
				Arn: RandString(10),
			},
		}
		audience := RandString(10)
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:            &name,
			AuthType:        &authType,
			Type:            &typ,
			Scope:           &scope,
			RoleAssumptions: &roleAssumptions,
			Audience:        &audience,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdmin
		newAudience := RandString(10)
		updateOps := buddy.IntegrationOps{
			Name:     &newName,
			AuthType: &authType,
			Scope:    &newScope,
			Audience: &newAudience,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationGoogleOIDC(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		scope := buddy.IntegrationScopeAdmin
		typ := buddy.IntegrationTypeGoogleServiceAccount
		authType := buddy.IntegrationAuthTypeOidc
		googleProject := RandString(10)
		config := "{}"
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:          &name,
			Type:          &typ,
			Scope:         &scope,
			AuthType:      &authType,
			GoogleProject: &googleProject,
			Config:        &config,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeWorkspace
		updateOps := buddy.IntegrationOps{
			Name:          &newName,
			Scope:         &newScope,
			AuthType:      &authType,
			GoogleProject: &googleProject,
			Config:        &config,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationAmazon(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		scope := buddy.IntegrationScopeAdmin
		typ := buddy.IntegrationTypeAmazon
		accessKey := RandString(10)
		secretKey := RandString(10)
		roleAssumptions := []*buddy.RoleAssumption{
			{
				Arn: RandString(10),
			},
			{
				Arn:        RandString(10),
				Duration:   10,
				ExternalId: RandString(10),
			},
		}
		var integration buddy.Integration
		createOps := buddy.IntegrationOps{
			Name:            &name,
			Type:            &typ,
			Scope:           &scope,
			AccessKey:       &accessKey,
			SecretKey:       &secretKey,
			RoleAssumptions: &roleAssumptions,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopePrivateInProject
		updateOps := buddy.IntegrationOps{
			Scope:       &newScope,
			ProjectName: &project.Name,
			Name:        &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationDigitalOcean(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		var integration buddy.Integration
		name := RandString(10)
		scope := buddy.IntegrationScopeProject
		typ := buddy.IntegrationTypeDigitalOcean
		token := RandString(10)
		createOps := buddy.IntegrationOps{
			Name:        &name,
			Scope:       &scope,
			Type:        &typ,
			Token:       &token,
			ProjectName: &project.Name,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopePrivate
		updateOps := buddy.IntegrationOps{
			Scope: &newScope,
			Name:  &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationShopify(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, project *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		var integration buddy.Integration
		name := RandString(10)
		scope := buddy.IntegrationScopeGroup
		typ := buddy.IntegrationTypeShopify
		token := RandString(10)
		shop := RandString(10)
		createOps := buddy.IntegrationOps{
			Name:    &name,
			Scope:   &scope,
			Type:    &typ,
			Token:   &token,
			Shop:    &shop,
			GroupId: &group.Id,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdminInProject
		updateOps := buddy.IntegrationOps{
			Scope:       &newScope,
			ProjectName: &project.Name,
			Name:        &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}

func testIntegrationShopifyPartner(client *buddy.Client, workspace *buddy.Workspace, group *buddy.Group, project *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		var integration buddy.Integration
		name := RandString(10)
		scope := buddy.IntegrationScopeGroup
		typ := buddy.IntegrationTypeShopify
		token := RandString(10)
		partnerToken := RandString(10)
		authType := buddy.IntegrationAuthTypeTokenAppExtension
		createOps := buddy.IntegrationOps{
			Name:         &name,
			Scope:        &scope,
			Type:         &typ,
			Token:        &token,
			PartnerToken: &partnerToken,
			AuthType:     &authType,
			GroupId:      &group.Id,
		}
		newName := RandString(10)
		newScope := buddy.IntegrationScopeAdminInProject
		updateOps := buddy.IntegrationOps{
			Scope:       &newScope,
			ProjectName: &project.Name,
			AuthType:    &authType,
			Name:        &newName,
		}
		t.Run("Create", testIntegrationCreate(client, workspace, &createOps, &integration))
		t.Run("Update", testIntegrationUpdate(client, workspace, integration.HashId, &updateOps, &integration))
		t.Run("Get", testIntegrationGet(client, workspace, integration.HashId, &integration))
		t.Run("GetList", testIntegrationGetList(client, workspace, 1))
		t.Run("Delete", testIntegrationDelete(client, workspace, integration.HashId))
	}
}
