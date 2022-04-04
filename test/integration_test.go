package test

import (
	"api-go-sdk/buddy"
	"testing"
)

func TestIntegration_amazon(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(ErrorFormatted("GetClient", err))
	}
	workspace, _, _, err := SeedInitialData(client)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	// CREATE INTEGRATION
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
	c := buddy.IntegrationOperationOptions{
		Name:            &name,
		Type:            &typ,
		Scope:           &scope,
		AccessKey:       &accessKey,
		SecretKey:       &secretKey,
		RoleAssumptions: &roleAssumptions,
	}
	integration, _, err := client.IntegrationService.Create(workspace.Domain, &c)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Create", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Integration.HashId", integration.HashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, name)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, scope)
	if err != nil {
		t.Fatal(err)
	}
	// UPDATE INTEGRATION
	newName := RandString(10)
	newScope := buddy.IntegrationScopeWorkspace
	u := buddy.IntegrationOperationOptions{
		Scope: &newScope,
		Name:  &newName,
	}
	hashId := integration.HashId
	integration, _, err = client.IntegrationService.Update(workspace.Domain, hashId, &u)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Update", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.HashId", integration.HashId, hashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, newName)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, newScope)
	if err != nil {
		t.Fatal(err)
	}
	// DELETE INTEGRATION
	_, err = client.IntegrationService.Delete(workspace.Domain, hashId)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Delete", err))
	}
}

func TestIntegration_digitalocean(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(ErrorFormatted("GetClient", err))
	}
	workspace, project, _, err := SeedInitialData(client)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	// CREATE INTEGRATION
	name := RandString(10)
	scope := buddy.IntegrationScopeProject
	typ := buddy.IntegrationTypeDigitalOcean
	token := RandString(10)
	c := buddy.IntegrationOperationOptions{
		Name:        &name,
		Scope:       &scope,
		Type:        &typ,
		Token:       &token,
		ProjectName: &project.Name,
	}
	integration, _, err := client.IntegrationService.Create(workspace.Domain, &c)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Create", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Integration.HashId", integration.HashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, name)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, scope)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.ProjectName", integration.ProjectName, project.Name)
	if err != nil {
		t.Fatal(err)
	}
	// UPDATE INTEGRATION
	newName := RandString(10)
	newScope := buddy.IntegrationScopePrivate
	u := buddy.IntegrationOperationOptions{
		Scope: &newScope,
		Name:  &newName,
	}
	hashId := integration.HashId
	integration, _, err = client.IntegrationService.Update(workspace.Domain, hashId, &u)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Update", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.HashId", integration.HashId, hashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, newName)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, newScope)
	if err != nil {
		t.Fatal(err)
	}
	// DELETE INTEGRATION
	_, err = client.IntegrationService.Delete(workspace.Domain, hashId)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Delete", err))
	}
}

func TestIntegration_shopify(t *testing.T) {
	client, err := GetClient()
	if err != nil {
		t.Fatal(ErrorFormatted("GetClient", err))
	}
	workspace, project, group, err := SeedInitialData(client)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	// CREATE INTEGRATION
	name := RandString(10)
	scope := buddy.IntegrationScopeGroup
	typ := buddy.IntegrationTypeShopify
	token := RandString(10)
	shop := RandString(10)
	c := buddy.IntegrationOperationOptions{
		Name:    &name,
		Scope:   &scope,
		Type:    &typ,
		Token:   &token,
		Shop:    &shop,
		GroupId: &group.Id,
	}
	integration, _, err := client.IntegrationService.Create(workspace.Domain, &c)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Create", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldSet("Integration.HashId", integration.HashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, name)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, scope)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckIntFieldEqualAndSet("Integration.GroupId", integration.GroupId, group.Id)
	if err != nil {
		t.Fatal(err)
	}
	// UPDATE INTEGRATION
	newName := RandString(10)
	newScope := buddy.IntegrationScopeAdminInProject
	u := buddy.IntegrationOperationOptions{
		Scope:       &newScope,
		ProjectName: &project.Name,
		Name:        &newName,
	}
	hashId := integration.HashId
	integration, _, err = client.IntegrationService.Update(workspace.Domain, hashId, &u)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Update", err))
	}
	err = CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.HashId", integration.HashId, hashId)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Name", integration.Name, newName)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Type", integration.Type, typ)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.Scope", integration.Scope, newScope)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckFieldEqualAndSet("Integration.ProjectName", integration.ProjectName, project.Name)
	if err != nil {
		t.Fatal(err)
	}
	// DELETE INTEGRATION
	_, err = client.IntegrationService.Delete(workspace.Domain, hashId)
	if err != nil {
		t.Fatal(ErrorFormatted("IntegrationService.Delete", err))
	}
}
