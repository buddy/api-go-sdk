package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testSandboxWaitForRunning(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, err := client.SandboxService.WaitForStatuses(workspace.Domain, sandbox.Id, 60, []string{buddy.SandboxStatusRunning})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.WaitForRunning", err))
		}
		err = CheckSandbox(sb, sandbox, nil, buddy.SandboxStatusRunning, "", "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxWaitForSetupDone(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, err := client.SandboxService.WaitForSetupStatuses(workspace.Domain, sandbox.Id, 60, []string{buddy.SandboxSetupStatusDone})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.WaitForRunning", err))
		}
		err = CheckSandbox(sb, sandbox, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusDone, "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxWaitForAppRunning(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, err := client.SandboxService.WaitForAppStatuses(workspace.Domain, sandbox.Id, 60, []string{buddy.SandboxAppStatusRunning})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.WaitForRunning", err))
		}
		err = CheckSandbox(sb, sandbox, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusDone, buddy.SandboxAppStatusRunning)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.SandboxOps, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Create(workspace.Domain, project.Name, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Create", err))
		}
		err = CheckSandbox(sb, out, ops, buddy.SandboxStatusStarting, buddy.SandboxSetupStatusConfiguring, buddy.SandboxAppStatusNone)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxEdit(client *buddy.Client, workspace *buddy.Workspace, ops *buddy.SandboxOps, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Update(workspace.Domain, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Update", err))
		}
		err = CheckSandbox(sb, out, ops, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusDone, buddy.SandboxAppStatusRunning)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxGet(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Get(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Get", err))
		}
		err = CheckSandbox(sb, out, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusDone, buddy.SandboxAppStatusRunning)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxes(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.SandboxService.GetList(workspace.Domain, buddy.Query{
			ProjectName: &project.Name,
		})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.GetList", err))
		}
		err = CheckSandboxes(list, 1)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSandboxDelete(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.SandboxService.Delete(workspace.Domain, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.Delete", err))
		}
	}
}

func TestSandbox(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	newName := RandString(10)
	identifier := UniqueString()
	newIdentifier := UniqueString()
	os := buddy.SandboxOsUbuntu2204
	resources := buddy.SandboxResource2X4
	installCommands := "pwd"
	runCommands := "while :; do foo; sleep 2; done"
	appDir := "/etc"
	appType := buddy.SandboxAppTypeCmd
	tags := []string{"a"}
	newTags := []string{"b"}
	endpointName := RandString(10)
	endpointPort := "22"
	endpointType := buddy.SandboxEndpointTypeTcp
	endpoints := []buddy.SandboxEndpoint{{
		Name:     &endpointName,
		Endpoint: &endpointPort,
		Type:     &endpointType,
	}}
	createOps := buddy.SandboxOps{
		Name:            &name,
		Identifier:      &identifier,
		Os:              &os,
		Resources:       &resources,
		InstallCommands: &installCommands,
		RunCommand:      &runCommands,
		AppDir:          &appDir,
		AppType:         &appType,
		Tags:            &tags,
		Endpoints:       &endpoints,
	}
	updateOps := buddy.SandboxOps{
		Name:       &newName,
		Identifier: &newIdentifier,
		Tags:       &newTags,
	}
	var sandbox buddy.Sandbox
	t.Run("Create", testSandboxCreate(seed.Client, seed.Workspace, seed.Project, &createOps, &sandbox))
	t.Run("Wait For Running", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for Setup", testSandboxWaitForSetupDone(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for App", testSandboxWaitForAppRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Update", testSandboxEdit(seed.Client, seed.Workspace, &updateOps, &sandbox))
	t.Run("Get", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList", testSandboxes(seed.Client, seed.Workspace, seed.Project))
	t.Run("Delete", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
}
