package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testPermissionCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		desc := RandString(10)
		pipelineAccessLevel := buddy.PermissionAccessLevelRunOnly
		repositoryAccessLevel := buddy.PermissionAccessLevelReadWrite
		sandboxAccessLevel := buddy.PermissionAccessLevelDenied
		ops := buddy.PermissionOps{
			Name:                  &name,
			Description:           &desc,
			PipelineAccessLevel:   &pipelineAccessLevel,
			RepositoryAccessLevel: &repositoryAccessLevel,
			SandboxAccessLevel:    &sandboxAccessLevel,
		}
		permission, _, err := client.PermissionService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("PermissionService.Create", err))
		}
		err = CheckPermission(permission, name, desc, 0, pipelineAccessLevel, repositoryAccessLevel, sandboxAccessLevel)
		if err != nil {
			t.Fatal(err)
		}
		*out = *permission
	}
}

func testPermissionUpdate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandString(10)
		desc := ""
		pipelineAccessLevel := buddy.PermissionAccessLevelReadWrite
		repositoryAccessLevel := buddy.PermissionAccessLevelManage
		sandboxAccessLevel := buddy.PermissionTypeReadOnly
		ops := buddy.PermissionOps{
			Name:                  &name,
			Description:           &desc,
			PipelineAccessLevel:   &pipelineAccessLevel,
			RepositoryAccessLevel: &repositoryAccessLevel,
			SandboxAccessLevel:    &sandboxAccessLevel,
		}
		permission, _, err := client.PermissionService.Update(workspace.Domain, out.Id, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("PermissionService.Patch", err))
		}
		err = CheckPermission(permission, name, desc, out.Id, pipelineAccessLevel, repositoryAccessLevel, sandboxAccessLevel)
		if err != nil {
			t.Fatal(err)
		}
		*out = *permission
	}
}

func testPermissionGet(client *buddy.Client, workspace *buddy.Workspace, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		permissionGet, _, err := client.PermissionService.Get(workspace.Domain, permission.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PermissionService.Get", err))
		}
		err = CheckPermission(permissionGet, permission.Name, permission.Description, permission.Id, permission.PipelineAccessLevel, permission.RepositoryAccessLevel, permission.SandboxAccessLevel)
		if err != nil {
			t.Fatal(err)
		}
		*permission = *permissionGet
	}
}

func testPermissionGetList(client *buddy.Client, workspace *buddy.Workspace, count int) func(t *testing.T) {
	return func(t *testing.T) {
		permissions, _, err := client.PermissionService.GetList(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("PermissionService.GetList", err))
		}
		err = CheckPermissions(permissions, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testPermissionDelete(client *buddy.Client, workspace *buddy.Workspace, permission *buddy.Permission) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.PermissionService.Delete(workspace.Domain, permission.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PermissionService.Delete", err))
		}
	}
}

func TestPermission(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var permission buddy.Permission
	t.Run("Create", testPermissionCreate(seed.Client, seed.Workspace, &permission))
	t.Run("Update", testPermissionUpdate(seed.Client, seed.Workspace, &permission))
	t.Run("Get", testPermissionGet(seed.Client, seed.Workspace, &permission))
	t.Run("GetList", testPermissionGetList(seed.Client, seed.Workspace, 3))
	t.Run("Delete", testPermissionDelete(seed.Client, seed.Workspace, &permission))
}
