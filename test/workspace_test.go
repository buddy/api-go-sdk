package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testWorkspaceCreate(client *buddy.Client, out *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		domain := UniqueString()
		name := UniqueString()
		salt := UniqueString()
		ops := buddy.WorkspaceCreateOps{
			Domain:         &domain,
			Name:           &name,
			EncryptionSalt: &salt,
		}
		workspace, _, err := client.WorkspaceService.Create(&ops)
		if err != nil {
			t.Fatal(ErrorFormatted("WorkspaceService.Create", err))
		}
		err = CheckWorkspace(workspace, name, domain, 0)
		if err != nil {
			t.Fatal(err)
		}
		*out = *workspace
	}
}

func testWorkspaceUpdate(client *buddy.Client, out *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		salt := UniqueString()
		name := UniqueString()
		ops := buddy.WorkspaceUpdateOps{
			Name:           &name,
			EncryptionSalt: &salt,
		}
		workspace, _, err := client.WorkspaceService.Update(out.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("WorkspaceService.Patch", err))
		}
		err = CheckWorkspace(workspace, name, out.Domain, out.Id)
		if err != nil {
			t.Fatal(err)
		}
		*out = *workspace
	}
}

func testWorkspaceGet(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		workspaceGet, _, err := client.WorkspaceService.Get(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("WorkspaceService.Get", err))
		}
		err = CheckWorkspace(workspaceGet, workspace.Name, workspace.Domain, workspace.Id)
		if err != nil {
			t.Fatal(err)
		}
		*workspace = *workspaceGet
	}
}

func testWorkspaceGetList(client *buddy.Client, atLeast int) func(t *testing.T) {
	return func(t *testing.T) {
		workspaces, _, err := client.WorkspaceService.GetList()
		if err != nil {
			t.Fatal(ErrorFormatted("WorkspaceService.GetList", err))
		}
		err = CheckWorkspaces(workspaces, atLeast)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testWorkspaceDelete(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.WorkspaceService.Delete(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("WorkspaceService.Delete", err))
		}
	}
}

func TestWorkspace(t *testing.T) {
	seed, err := SeedInitialData(nil)
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var workspace buddy.Workspace
	t.Run("Create", testWorkspaceCreate(seed.Client, &workspace))
	t.Run("Update", testWorkspaceUpdate(seed.Client, &workspace))
	t.Run("Get", testWorkspaceGet(seed.Client, &workspace))
	t.Run("GetList", testWorkspaceGetList(seed.Client, 2))
	t.Run("Delete", testWorkspaceDelete(seed.Client, &workspace))
}
