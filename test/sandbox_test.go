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
		sb, err := client.SandboxService.WaitForSetupStatuses(workspace.Domain, sandbox.Id, 60, []string{buddy.SandboxSetupStatusSuccess})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.WaitForRunning", err))
		}
		err = CheckSandbox(sb, sandbox, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusSuccess, "")
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
		err = CheckSandbox(sb, sandbox, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusSuccess, buddy.SandboxAppStatusRunning)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxCreate(client *buddy.Client, workspace *buddy.Workspace, query *buddy.Query, ops *buddy.SandboxOps, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Create(workspace.Domain, query, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Create", err))
		}
		err = CheckSandbox(sb, out, ops, buddy.SandboxStatusStarting, buddy.SandboxSetupStatusInProgress, buddy.SandboxAppStatusNone)
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
		err = CheckSandbox(sb, out, ops, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusSuccess, buddy.SandboxAppStatusRunning)
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
		err = CheckSandbox(sb, out, nil, buddy.SandboxStatusRunning, buddy.SandboxSetupStatusSuccess, buddy.SandboxAppStatusRunning)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxes(client *buddy.Client, workspace *buddy.Workspace, query *buddy.Query) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.SandboxService.GetList(workspace.Domain, query)
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
		workspace:     true,
		project:       true,
		member:        true,
		permission:    true,
		projectMember: true,
	})

	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	note := RandString(10)
	agentNote := RandString(10)
	newName := RandString(10)
	newNote := ""
	newAgentNote := RandString(10)
	identifier := UniqueString()
	newIdentifier := UniqueString()
	os := buddy.SandboxOsUbuntu2204
	resources := buddy.SandboxResource2X4
	installCommands := "pwd"
	app := "while :; do foo; sleep 2; done"
	appDir := "/etc"
	tags := []string{"a"}
	newTags := []string{"b"}
	endpointName := RandString(10)
	endpointPort := "22"
	endpointType := buddy.SandboxEndpointTypeTcp
	timeout := 300
	variableKey := RandString(10)
	variableVal := RandString(10)
	endpoints := []*buddy.SandboxEndpoint{{
		Name:     &endpointName,
		Endpoint: &endpointPort,
		Type:     &endpointType,
	}}
	variables := []*buddy.VariableOps{{
		Key:   &variableKey,
		Value: &variableVal,
	}}
	perms := buddy.SandboxPermissions{
		Others: buddy.SandboxPermissionReadOnly,
		Users: []*buddy.SandboxResourcePermission{
			{
				Id:          seed.Member.Id,
				AccessLevel: buddy.SandboxPermissionManage,
			},
		},
	}
	apps := []string{app}
	createOps := buddy.SandboxOps{
		Name:              &name,
		Note:              &note,
		AgentNote:         &agentNote,
		Identifier:        &identifier,
		Os:                &os,
		Resources:         &resources,
		FirstBootCommands: &installCommands,
		AppDir:            &appDir,
		Apps:              &apps,
		Tags:              &tags,
		Timeout:           &timeout,
		Endpoints:         &endpoints,
		Variables:         &variables,
		Permissions:       &perms,
	}
	newPerms := buddy.SandboxPermissions{
		Others: buddy.SandboxPermissionDenied,
		Users: []*buddy.SandboxResourcePermission{
			{
				Id:          seed.Member.Id,
				AccessLevel: buddy.SandboxPermissionReadOnly,
			},
		},
	}
	updateOps := buddy.SandboxOps{
		Name:        &newName,
		Note:        &newNote,
		AgentNote:   &newAgentNote,
		Identifier:  &newIdentifier,
		Tags:        &newTags,
		Permissions: &newPerms,
	}
	projectQuery := buddy.Query{
		ProjectName: &seed.Project.Name,
	}
	wsEnvName := UniqueString()
	wsEnvId := UniqueString()
	pjEnvName := UniqueString()
	pjEnvId := UniqueString()
	wsEnv, _, err := seed.Client.EnvironmentService.Create(seed.Workspace.Domain, &buddy.EnvironmentOps{
		Name:       &wsEnvName,
		Identifier: &wsEnvId,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("Create workspace env", err))
	}
	pjEnv, _, err := seed.Client.EnvironmentService.Create(seed.Workspace.Domain, &buddy.EnvironmentOps{
		Name:       &pjEnvName,
		Identifier: &pjEnvId,
		Project: &buddy.ProjectSimple{
			Name: seed.Project.Name,
		},
	})
	if err != nil {
		t.Fatal(ErrorFormatted("Create project env", err))
	}
	var sandbox buddy.Sandbox
	t.Run("Create", testSandboxCreate(seed.Client, seed.Workspace, &projectQuery, &createOps, &sandbox))
	t.Run("Wait For Running", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for Setup", testSandboxWaitForSetupDone(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for App", testSandboxWaitForAppRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Update", testSandboxEdit(seed.Client, seed.Workspace, &updateOps, &sandbox))
	t.Run("Get", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList", testSandboxes(seed.Client, seed.Workspace, &projectQuery))
	t.Run("Delete", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
	sandbox = buddy.Sandbox{}
	t.Run("Create In Workspace", testSandboxCreate(seed.Client, seed.Workspace, nil, &createOps, &sandbox))
	t.Run("Wait For Running In Workspace", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Get In Workspace", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList In Workspace", testSandboxes(seed.Client, seed.Workspace, nil))
	t.Run("Delete In Workspace", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
	newIdentifier = UniqueString()
	createOps.Identifier = &newIdentifier
	createOps.Environment = &buddy.SandboxEnvironmentOps{
		Id: &wsEnv.Id,
	}
	envQuery := buddy.Query{
		EnvironmentId: &wsEnv.Id,
	}
	sandbox = buddy.Sandbox{}
	t.Run("Create In Workspace Env", testSandboxCreate(seed.Client, seed.Workspace, nil, &createOps, &sandbox))
	t.Run("Wait For Running In Workspace Env", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Get In Workspace Env", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList In Workspace Env", testSandboxes(seed.Client, seed.Workspace, &envQuery))
	t.Run("Delete In Workspace Env", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
	sandbox = buddy.Sandbox{}
	newIdentifier = UniqueString()
	createOps.Identifier = &newIdentifier
	createOps.Environment = &buddy.SandboxEnvironmentOps{
		Id: &pjEnv.Id,
	}
	envQuery.ProjectName = &seed.Project.Name
	envQuery.EnvironmentId = &pjEnv.Id
	t.Run("Create In Project Env", testSandboxCreate(seed.Client, seed.Workspace, &projectQuery, &createOps, &sandbox))
	t.Run("Wait For Running In Project Env", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Get In Project Env", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList In Project Env", testSandboxes(seed.Client, seed.Workspace, &envQuery))
	t.Run("Delete In Project Env", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
}
