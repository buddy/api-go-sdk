package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func CheckTarget(target *buddy.Target, want *buddy.TargetOps) error {
	if err := CheckFieldSet("HtmlUrl", target.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Url", target.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Id", target.Id); err != nil {
		return err
	}
	if want.Name != nil {
		if err := CheckFieldEqualAndSet("Name", target.Name, *want.Name); err != nil {
			return err
		}
	}
	if want.Identifier != nil {
		if err := CheckFieldEqualAndSet("Identifier", target.Identifier, *want.Identifier); err != nil {
			return err
		}
	}
	if want.Type != nil {
		if err := CheckFieldEqualAndSet("Type", target.Type, *want.Type); err != nil {
			return err
		}
	}
	if want.Disabled != nil {
		if err := CheckBoolFieldEqual("Disabled", target.Disabled, *want.Disabled); err != nil {
			return err
		}
	}
	if want.Tags != nil {
		if err := CheckIntFieldEqualAndSet("len(Tags)", len(target.Tags), len(*want.Tags)); err != nil {
			return err
		}
	}
	if want.Host != nil {
		if err := CheckFieldEqualAndSet("Host", target.Host, *want.Host); err != nil {
			return err
		}
	}
	if want.Auth != nil {
		if err := CheckFieldEqual("Auth.Method", target.Auth.Method, want.Auth.Method); err != nil {
			return err
		}
		isPass := target.Auth.Method == buddy.TargetAuthMethodPassword || target.Auth.Method == ""
		isHttp := target.Auth.Method == buddy.TargetAuthMethodHttp
		isAsset := target.Auth.Method == buddy.TargetAuthMethodAssetsKey
		isSshKey := target.Auth.Method == buddy.TargetAuthMethodSshKey
		if isPass || isHttp {
			if err := CheckFieldSet("Auth.Password", target.Auth.Password); err != nil {
				return err
			}
			if err := CheckFieldEqualAndSet("Auth.Username", target.Auth.Username, want.Auth.Username); err != nil {
				return err
			}
		}
		if isAsset {
			if want.Auth.Username != "" {
				if err := CheckFieldEqualAndSet("Auth.Username", target.Auth.Username, want.Auth.Username); err != nil {
					return err
				}
			}
			if err := CheckFieldEqualAndSet("Auth.Asset", target.Auth.Asset, want.Auth.Asset); err != nil {
				return err
			}
		}
		if isSshKey {
			if want.Auth.Username != "" {
				if err := CheckFieldEqualAndSet("Auth.Username", target.Auth.Username, want.Auth.Username); err != nil {
					return err
				}
			}
			if err := CheckFieldSet("Auth.Key", target.Auth.Key); err != nil {
				return err
			}
			if want.Auth.Passphrase != "" {
				if err := CheckFieldSet("Auth.Passphrase", target.Auth.Passphrase); err != nil {
					return err
				}
			}
		}
	}
	if want.Integration != nil {
		if err := CheckFieldEqualAndSet("Integration", target.Integration, *want.Integration); err != nil {
			return err
		}
	}
	if want.Proxy != nil {
		if err := CheckFieldEqualAndSet("Proxy.Name", target.Proxy.Name, want.Proxy.Name); err != nil {
			return err
		}
	}
	if want.Repository != nil {
		if err := CheckFieldEqualAndSet("Repository", target.Repository, *want.Repository); err != nil {
			return err
		}
	}
	if want.Port != nil {
		if err := CheckFieldEqualAndSet("Port", target.Port, *want.Port); err != nil {
			return err
		}
	}
	if want.Path != nil {
		if err := CheckFieldEqualAndSet("Path", target.Path, *want.Path); err != nil {
			return err
		}
	}
	if want.Secure != nil {
		if err := CheckBoolFieldEqual("Secure", target.Secure, *want.Secure); err != nil {
			return err
		}
	}
	if want.Permissions == nil && target.Permissions != nil {
		if err := CheckFieldEqualAndSet("Permissions.Others", target.Permissions.Others, buddy.TargetPermissionUseOnly); err != nil {
			return err
		}
	}
	if want.Permissions != nil {
		if err := CheckFieldEqualAndSet("Permissions.Others", target.Permissions.Others, want.Permissions.Others); err != nil {
			return err
		}
		if err := CheckIntFieldEqualAndSet("Permissions.Users[0].Id", target.Permissions.Users[0].Id, want.Permissions.Users[0].Id); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Permissions.Users[0].AccessLevel", target.Permissions.Users[0].AccessLevel, want.Permissions.Users[0].AccessLevel); err != nil {
			return err
		}
		if err := CheckIntFieldEqualAndSet("Permissions.Groups[0].Id", target.Permissions.Groups[0].Id, want.Permissions.Groups[0].Id); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Permissions.Groups[0].AccessLevel", target.Permissions.Groups[0].AccessLevel, want.Permissions.Groups[0].AccessLevel); err != nil {
			return err
		}
	}
	return nil
}

func testTargetFtps(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeFtp
		host := "1.1.1.1"
		port := "33"
		secure := true
		disabled := true
		username := RandString(10)
		password := RandString(10)
		auth := buddy.TargetAuth{
			Username: username,
			Password: password,
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Type:       &typ,
			Host:       &host,
			Port:       &port,
			Secure:     &secure,
			Auth:       &auth,
			Disabled:   &disabled,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func CheckTargets(targets *buddy.Targets, count int) error {
	if err := CheckIntFieldEqualAndSet("len(targets.Targets)", len(targets.Targets), count); err != nil {
		return err
	}
	return nil
}

func testTargetGetList(client *buddy.Client, workspaceDomain string, count int) func(t *testing.T) {
	return func(t *testing.T) {
		targets, _, err := client.TargetService.GetList(workspaceDomain, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.GetList", err))
		}
		err = CheckTargets(targets, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetDelete(client *buddy.Client, workspaceDomain string, target *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.TargetService.Delete(workspaceDomain, target.Id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetGet(client *buddy.Client, workspaceDomain string, target *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		targ, _, err := client.TargetService.Get(workspaceDomain, target.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Get", err))
		}
		ops := buddy.TargetOps{
			Name:        &target.Name,
			Identifier:  &target.Identifier,
			Permissions: target.Permissions,
			Port:        &target.Port,
			Type:        &target.Type,
			Scope:       &target.Scope,
			Path:        &target.Path,
			Host:        &target.Host,
		}
		err = CheckTarget(targ, &ops)
		if err != nil {
			t.Fatal(err)
		}
		*target = *targ
	}
}

func testTargetPermissionsUpdate(client *buddy.Client, workspaceDomain string, memberId int, groupId int, target *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		perms := buddy.TargetPermissions{
			Others: buddy.TargetPermissionUseOnly,
			Users: []*buddy.TargetResourcePermission{
				{
					Id:          memberId,
					AccessLevel: buddy.TargetPermissionManage,
				},
			},
			Groups: []*buddy.TargetResourcePermission{
				{
					Id:          groupId,
					AccessLevel: buddy.TargetPermissionManage,
				},
			},
		}
		newName := UniqueString()
		newIdentifier := UniqueString()
		ops := buddy.TargetOps{
			Name:        &newName,
			Identifier:  &newIdentifier,
			Type:        &target.Type,
			Permissions: &perms,
		}
		targ, _, err := client.TargetService.Update(workspaceDomain, target.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Update", err))
		}
		err = CheckTarget(targ, &ops)
		if err != nil {
			t.Fatal(err)
		}
		*target = *targ
	}
}

func testTargetPermissions(client *buddy.Client, workspaceDomain string, memberId int, groupId int, target *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeSsh
		host := "1.1.1.1"
		port := "44"
		username := RandString(10)
		key := RandString(10)
		path := RandString(10)
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodSshKey,
			Username: username,
			Key:      key,
		}
		perms := buddy.TargetPermissions{
			Others: buddy.TargetPermissionManage,
			Users: []*buddy.TargetResourcePermission{
				{
					Id:          memberId,
					AccessLevel: buddy.TargetPermissionUseOnly,
				},
			},
			Groups: []*buddy.TargetResourcePermission{
				{
					Id:          groupId,
					AccessLevel: buddy.TargetPermissionManage,
				},
			},
		}
		ops := buddy.TargetOps{
			Name:        &name,
			Identifier:  &identifier,
			Type:        &typ,
			Host:        &host,
			Port:        &port,
			Auth:        &auth,
			Path:        &path,
			Permissions: &perms,
		}
		targ, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(targ, &ops)
		if err != nil {
			t.Fatal(err)
		}
		*target = *targ
	}
}

func testTargetSshKey(client *buddy.Client, workspaceDomain string, projectName string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeSsh
		host := "1.1.1.1"
		port := "44"
		username := RandString(10)
		key := RandString(10)
		passphrase := RandString(10)
		path := RandString(10)
		auth := buddy.TargetAuth{
			Method:     buddy.TargetAuthMethodSshKey,
			Username:   username,
			Key:        key,
			Passphrase: passphrase,
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Type:       &typ,
			Host:       &host,
			Port:       &port,
			Auth:       &auth,
			Path:       &path,
			Project: &buddy.TargetProject{
				Name: projectName,
			},
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetSshProxyCredentials(client *buddy.Client, workspaceDomain string, pipelineId int) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeSsh
		host := "1.1.1.1"
		port := "44"
		path := RandString(10)
		auth := buddy.TargetAuth{
			Method: buddy.TargetAuthMethodProxyCredentials,
		}
		proxyAuth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodPassword,
			Username: RandString(10),
			Password: RandString(10),
		}
		proxy := buddy.TargetProxy{
			Name: UniqueString(),
			Host: "2.2.2.2",
			Port: "33",
			Auth: &proxyAuth,
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Type:       &typ,
			Host:       &host,
			Port:       &port,
			Auth:       &auth,
			Path:       &path,
			Proxy:      &proxy,
			Pipeline: &buddy.TargetPipeline{
				Id: pipelineId,
			},
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetDigitalOcean(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		integrationName := UniqueString()
		integrationType := buddy.IntegrationTypeDigitalOcean
		integrationToken := UniqueString()
		integrationScope := buddy.IntegrationScopeWorkspace
		integrationOps := buddy.IntegrationOps{
			Name:  &integrationName,
			Type:  &integrationType,
			Token: &integrationToken,
			Scope: &integrationScope,
		}
		integration, _, err := client.IntegrationService.Create(workspaceDomain, &integrationOps)
		if err != nil {
			t.Fatal(err)
		}
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeDigitalOcean
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodPassword,
			Username: UniqueString(),
			Password: UniqueString(),
		}
		host := "1.1.1.1"
		port := "44"
		ops := buddy.TargetOps{
			Name:        &name,
			Identifier:  &identifier,
			Type:        &typ,
			Host:        &host,
			Port:        &port,
			Auth:        &auth,
			Integration: &integration.Identifier,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetVultr(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		integrationName := UniqueString()
		integrationType := buddy.IntegrationTypeVultr
		integrationToken := UniqueString()
		integrationScope := buddy.IntegrationScopeWorkspace
		integrationOps := buddy.IntegrationOps{
			Name:  &integrationName,
			Type:  &integrationType,
			Token: &integrationToken,
			Scope: &integrationScope,
		}
		integration, _, err := client.IntegrationService.Create(workspaceDomain, &integrationOps)
		if err != nil {
			t.Fatal(err)
		}
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeVultr
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodPassword,
			Username: UniqueString(),
			Password: UniqueString(),
		}
		host := "1.1.1.1"
		port := "44"
		ops := buddy.TargetOps{
			Name:        &name,
			Identifier:  &identifier,
			Type:        &typ,
			Host:        &host,
			Port:        &port,
			Auth:        &auth,
			Integration: &integration.Identifier,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetUpcloud(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		integrationName := UniqueString()
		integrationType := buddy.IntegrationTypeUpcloud
		integrationUser := UniqueString()
		integrationPass := UniqueString()
		integrationScope := buddy.IntegrationScopeWorkspace
		integrationOps := buddy.IntegrationOps{
			Name:     &integrationName,
			Type:     &integrationType,
			Username: &integrationUser,
			Password: &integrationPass,
			Scope:    &integrationScope,
		}
		integration, _, err := client.IntegrationService.Create(workspaceDomain, &integrationOps)
		if err != nil {
			t.Fatal(err)
		}
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeUpcloud
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodPassword,
			Username: UniqueString(),
			Password: UniqueString(),
		}
		host := "1.1.1.1"
		port := "44"
		ops := buddy.TargetOps{
			Name:        &name,
			Identifier:  &identifier,
			Type:        &typ,
			Host:        &host,
			Port:        &port,
			Auth:        &auth,
			Integration: &integration.Identifier,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetSshAsset(client *buddy.Client, workspaceDomain string, projectName string) func(t *testing.T) {
	return func(t *testing.T) {
		envName := UniqueString()
		envId := UniqueString()
		envTyp := buddy.EnvironmentTypeDev
		envOps := buddy.EnvironmentOps{
			Name:       &envName,
			Identifier: &envId,
			Type:       &envTyp,
		}
		env, _, err := client.EnvironmentService.Create(workspaceDomain, projectName, &envOps)
		if err != nil {
			t.Fatal(err)
		}
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeSsh
		host := "1.1.1.1"
		port := "44"
		path := RandString(10)
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodAssetsKey,
			Username: RandString(10),
			Asset:    "id_workspace",
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Type:       &typ,
			Host:       &host,
			Port:       &port,
			Auth:       &auth,
			Path:       &path,
			Environment: &buddy.TargetEnvironment{
				Id: env.Id,
			},
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetSshPassword(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		typ := buddy.TargetTypeSsh
		host := "1.1.1.1"
		port := "44"
		username := RandString(10)
		password := RandString(10)
		path := RandString(10)
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodPassword,
			Username: username,
			Password: password,
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Type:       &typ,
			Host:       &host,
			Port:       &port,
			Auth:       &auth,
			Path:       &path,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetGitHttp(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		repository := "https://a" + UniqueString() + ".com"
		typ := buddy.TargetTypeGit
		auth := buddy.TargetAuth{
			Method:   buddy.TargetAuthMethodHttp,
			Username: RandString(10),
			Password: RandString(10),
		}
		tags := []string{"a", "b"}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Repository: &repository,
			Tags:       &tags,
			Type:       &typ,
			Auth:       &auth,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetGitAsset(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		repository := "https://a" + UniqueString() + ".com"
		typ := buddy.TargetTypeGit
		auth := buddy.TargetAuth{
			Method: buddy.TargetAuthMethodAssetsKey,
			Asset:  "id_workspace",
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Repository: &repository,
			Type:       &typ,
			Auth:       &auth,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetGitSshKey(client *buddy.Client, workspaceDomain string) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		identifier := UniqueString()
		repository := "https://a" + UniqueString() + ".com"
		typ := buddy.TargetTypeGit
		auth := buddy.TargetAuth{
			Method: buddy.TargetAuthMethodSshKey,
			Key:    RandString(10),
		}
		ops := buddy.TargetOps{
			Name:       &name,
			Identifier: &identifier,
			Repository: &repository,
			Type:       &typ,
			Auth:       &auth,
		}
		target, _, err := client.TargetService.Create(workspaceDomain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		err = CheckTarget(target, &ops)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestTarget(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
		pipeline:  true,
		member:    true,
		group:     true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var target buddy.Target
	t.Run("CreateFtps", testTargetFtps(seed.Client, seed.Workspace.Domain))
	t.Run("CreateGitHttp", testTargetGitHttp(seed.Client, seed.Workspace.Domain))
	t.Run("CreateGitSshKey", testTargetGitSshKey(seed.Client, seed.Workspace.Domain))
	t.Run("CreateGitAsset", testTargetGitAsset(seed.Client, seed.Workspace.Domain))
	t.Run("CreateSshPassword", testTargetSshPassword(seed.Client, seed.Workspace.Domain))
	t.Run("CreateSshKey", testTargetSshKey(seed.Client, seed.Workspace.Domain, seed.Project.Name))
	t.Run("CreateSshAsset", testTargetSshAsset(seed.Client, seed.Workspace.Domain, seed.Project.Name))
	t.Run("CreateSshProxyCredentials", testTargetSshProxyCredentials(seed.Client, seed.Workspace.Domain, seed.Pipeline.Id))
	t.Run("CreateUpcloud", testTargetUpcloud(seed.Client, seed.Workspace.Domain))
	t.Run("CreateVultr", testTargetVultr(seed.Client, seed.Workspace.Domain))
	t.Run("CreateDigitalOcean", testTargetDigitalOcean(seed.Client, seed.Workspace.Domain))
	t.Run("CreateWithPerms", testTargetPermissions(seed.Client, seed.Workspace.Domain, seed.Member.Id, seed.Group.Id, &target))
	t.Run("UpdateWithPerms", testTargetPermissionsUpdate(seed.Client, seed.Workspace.Domain, seed.Member.Id, seed.Group.Id, &target))
	t.Run("Get", testTargetGet(seed.Client, seed.Workspace.Domain, &target))
	t.Run("GetList", testTargetGetList(seed.Client, seed.Workspace.Domain, 9))
	t.Run("Delete", testTargetDelete(seed.Client, seed.Workspace.Domain, &target))
}
