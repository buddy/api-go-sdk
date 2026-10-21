package test

import (
	"encoding/base64"
	"fmt"
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

func testSandboxPatch(client *buddy.Client, workspace *buddy.Workspace, ops *buddy.SandboxOps, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Update(workspace.Domain, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Patch", err))
		}
		err = CheckSandbox(sb, out, ops, buddy.SandboxStatusRunning, "", "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxFetchWaitForFetched(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, expected *buddy.SandboxFetch, out *buddy.SandboxFetch) func(t *testing.T) {
	return func(t *testing.T) {
		f, err := client.SandboxService.WaitForFetchStatuses(workspace.Domain, sandbox.Id, expected.Id, 120, []string{buddy.SandboxFetchStatusFetched, buddy.SandboxFetchStatusError})
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.WaitForFetchStatuses", err))
		}
		err = CheckSandboxFetch(f, expected, buddy.SandboxFetchStatusFetched)
		if err != nil {
			t.Fatal(err)
		}
		*out = *f
	}
}

func testSandboxFetchCreate(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, ops *buddy.SandboxFetchOps, out *buddy.SandboxFetch) func(t *testing.T) {
	return func(t *testing.T) {
		f, _, err := client.SandboxService.CreateFetch(workspace.Domain, sandbox.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.CreateFetch", err))
		}
		err = CheckSandboxFetch(f, SandboxFetchFromOps(ops), "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *f
	}
}

func testSandboxFetchGet(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.SandboxFetch) func(t *testing.T) {
	return func(t *testing.T) {
		f, _, err := client.SandboxService.GetFetch(workspace.Domain, sandbox.Id, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.GetFetch", err))
		}
		err = CheckSandboxFetch(f, out, buddy.SandboxFetchStatusFetched)
		if err != nil {
			t.Fatal(err)
		}
		*out = *f
	}
}

func testSandboxFetchRefetch(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.SandboxFetch) func(t *testing.T) {
	return func(t *testing.T) {
		f, _, err := client.SandboxService.Refetch(workspace.Domain, sandbox.Id, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.Refetch", err))
		}
		err = CheckSandboxFetch(f, out, "")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSandboxFetches(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, count int) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.SandboxService.GetFetches(workspace.Domain, sandbox.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.GetFetches", err))
		}
		err = CheckSandboxFetches(list, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSandboxFetchDelete(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, out *buddy.SandboxFetch) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.SandboxService.DeleteFetch(workspace.Domain, sandbox.Id, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("SandboxService.DeleteFetch", err))
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

func testSandboxCreatePipelines(client *buddy.Client, workspace *buddy.Workspace, query *buddy.Query, ops *buddy.SandboxOps, out *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, _, err := client.SandboxService.Create(workspace.Domain, query, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Create", err))
		}
		err = CheckSandboxPipelines(sb, *ops.PipelinesAccessLevel, *ops.AllowedPipelines)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sb
	}
}

func testSandboxWaitForSetupPipelines(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		sb, err := client.SandboxService.WaitForSetupStatuses(workspace.Domain, sandbox.Id, 60, []string{buddy.SandboxSetupStatusSuccess})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.WaitForSetup", err))
		}
		err = CheckSandboxPipelines(sb, sandbox.PipelinesAccessLevel, sandbox.AllowedPipelines)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSandboxYamlRoundtrip(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox) func(t *testing.T) {
	return func(t *testing.T) {
		y, _, err := client.SandboxService.GetYaml(workspace.Domain, sandbox.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.GetYaml", err))
		}
		if err = CheckFieldSet("SandboxYaml.Yaml", y.Yaml); err != nil {
			t.Fatal(err)
		}
		sb, _, err := client.SandboxService.UpdateYaml(workspace.Domain, sandbox.Id, &buddy.SandboxYamlOps{
			Yaml: &y.Yaml,
		})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.UpdateYaml", err))
		}
		err = CheckSandboxPipelines(sb, sandbox.PipelinesAccessLevel, sandbox.AllowedPipelines)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSandboxUpdateYamlPipelines(client *buddy.Client, workspace *buddy.Workspace, sandbox *buddy.Sandbox, yaml string, pipelinesAccessLevel string, allowedPipelines []*buddy.SandboxAllowedPipeline) func(t *testing.T) {
	return func(t *testing.T) {
		encoded := base64.StdEncoding.EncodeToString([]byte(yaml))
		sb, _, err := client.SandboxService.UpdateYaml(workspace.Domain, sandbox.Id, &buddy.SandboxYamlOps{
			Yaml: &encoded,
		})
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.UpdateYaml", err))
		}
		err = CheckSandboxPipelines(sb, pipelinesAccessLevel, allowedPipelines)
		if err != nil {
			t.Fatal(err)
		}
		sb, _, err = client.SandboxService.Get(workspace.Domain, sandbox.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TestSandbox.Get", err))
		}
		err = CheckSandboxPipelines(sb, pipelinesAccessLevel, allowedPipelines)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSandboxYamlPipelines(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
		pipeline:  true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	identifier := UniqueString()
	os := buddy.SandboxOsUbuntu2404
	resources := buddy.SandboxResource1X2
	pipelinesAccessLevelDenied := buddy.SandboxPipelineAccessLevelDenied
	allowedPipelines := []*buddy.SandboxAllowedPipeline{{
		Project:     seed.Project.Name,
		Pipeline:    seed.Pipeline.Identifier,
		AccessLevel: buddy.SandboxPipelineAccessLevelUseOnly,
	}}
	createOps := buddy.SandboxOps{
		Name:                 &identifier,
		Identifier:           &identifier,
		Os:                   &os,
		Resources:            &resources,
		PipelinesAccessLevel: &pipelinesAccessLevelDenied,
		AllowedPipelines:     &allowedPipelines,
	}
	projectQuery := buddy.Query{
		ProjectName: &seed.Project.Name,
	}
	yamlWithoutPipelines := fmt.Sprintf("- sandbox: %s\n  os: %s\n  resources: %s\n", identifier, os, resources)
	yamlWithoutAll := yamlWithoutPipelines + fmt.Sprintf("  permissions:\n    pipelines:\n    - project: %s\n      pipeline: %s\n      access: %s\n", seed.Project.Name, seed.Pipeline.Identifier, buddy.SandboxPipelineAccessLevelDenied)
	yamlOnlyAll := yamlWithoutPipelines + fmt.Sprintf("  permissions:\n    pipelines:\n    - all: %s\n", buddy.SandboxPipelineAccessLevelDenied)
	deniedPipelines := []*buddy.SandboxAllowedPipeline{{
		Project:     seed.Project.Name,
		Pipeline:    seed.Pipeline.Identifier,
		AccessLevel: buddy.SandboxPipelineAccessLevelDenied,
	}}
	var sandbox buddy.Sandbox
	t.Run("Create", testSandboxCreatePipelines(seed.Client, seed.Workspace, &projectQuery, &createOps, &sandbox))
	t.Run("Wait for Setup", testSandboxWaitForSetupPipelines(seed.Client, seed.Workspace, &sandbox))
	t.Run("Yaml Roundtrip", testSandboxYamlRoundtrip(seed.Client, seed.Workspace, &sandbox))
	// YAML is the target state: no pipelines section resets the permissions to the defaults
	t.Run("Yaml Without Pipelines", testSandboxUpdateYamlPipelines(seed.Client, seed.Workspace, &sandbox, yamlWithoutPipelines, buddy.SandboxPipelineAccessLevelUseOnly, nil))
	// a section without all gets the default all
	t.Run("Yaml Without All", testSandboxUpdateYamlPipelines(seed.Client, seed.Workspace, &sandbox, yamlWithoutAll, buddy.SandboxPipelineAccessLevelUseOnly, deniedPipelines))
	t.Run("Yaml Only All", testSandboxUpdateYamlPipelines(seed.Client, seed.Workspace, &sandbox, yamlOnlyAll, buddy.SandboxPipelineAccessLevelDenied, nil))
	t.Run("Delete", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
}

func TestSandbox(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:     true,
		project:       true,
		member:        true,
		permission:    true,
		projectMember: true,
		pipeline:      true,
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
		Others: buddy.SandboxPermissionUseOnly,
		Users: []*buddy.SandboxResourcePermission{
			{
				Id:          seed.Member.Id,
				AccessLevel: buddy.SandboxPermissionManage,
			},
		},
	}
	appIdentifier := UniqueString()
	appPath := "/tmp"
	apps := []*buddy.SandboxAppOps{{
		Identifier: &appIdentifier,
		Command:    &app,
		Path:       &appPath,
	}}
	fetchType := buddy.SandboxFetchTypePublicRepo
	fetchRepository := "https://github.com/octocat/Hello-World"
	fetchRef := "master"
	fetchPath := "/tmp/hello-world"
	fetchBuildCommand := "ls"
	fetch := []*buddy.SandboxFetchOps{{
		Type:         &fetchType,
		Repository:   &fetchRepository,
		Ref:          &fetchRef,
		Path:         &fetchPath,
		BuildCommand: &fetchBuildCommand,
	}}
	newFetchRepository := "https://github.com/octocat/Spoon-Knife"
	newFetchPath := "/tmp/spoon-knife"
	newFetchOps := buddy.SandboxFetchOps{
		Type:       &fetchType,
		Repository: &newFetchRepository,
		Path:       &newFetchPath,
	}
	updateFetchPath := "/tmp/hello-world-2"
	updateFetchOps := buddy.SandboxOps{
		Fetch: &[]*buddy.SandboxFetchOps{{
			Type:       &fetchType,
			Repository: &fetchRepository,
			Ref:        &fetchRef,
			Path:       &updateFetchPath,
		}},
	}
	clearFetchOps := buddy.SandboxOps{
		Fetch: &[]*buddy.SandboxFetchOps{},
	}
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
		Fetch:             &fetch,
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
				AccessLevel: buddy.SandboxPermissionUseOnly,
			},
		},
	}
	pipelinesAccessLevelDenied := buddy.SandboxPipelineAccessLevelDenied
	allowedPipelines := []*buddy.SandboxAllowedPipeline{{
		Project:     seed.Project.Name,
		Pipeline:    seed.Pipeline.Identifier,
		AccessLevel: buddy.SandboxPipelineAccessLevelUseOnly,
	}}
	updateOps := buddy.SandboxOps{
		Name:                 &newName,
		Note:                 &newNote,
		AgentNote:            &newAgentNote,
		Identifier:           &newIdentifier,
		Tags:                 &newTags,
		Permissions:          &newPerms,
		PipelinesAccessLevel: &pipelinesAccessLevelDenied,
		AllowedPipelines:     &allowedPipelines,
	}
	pipelinesAccessLevelUseOnly := buddy.SandboxPipelineAccessLevelUseOnly
	clearAllowedPipelinesOps := buddy.SandboxOps{
		PipelinesAccessLevel: &pipelinesAccessLevelUseOnly,
		AllowedPipelines:     &[]*buddy.SandboxAllowedPipeline{},
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
	if len(sandbox.Fetch) == 0 {
		t.Fatal("Sandbox.Fetch is empty")
	}
	var sandboxFetch buddy.SandboxFetch
	t.Run("Wait For Fetch", testSandboxFetchWaitForFetched(seed.Client, seed.Workspace, &sandbox, sandbox.Fetch[0], &sandboxFetch))
	t.Run("GetFetches", testSandboxFetches(seed.Client, seed.Workspace, &sandbox, 1))
	var newSandboxFetch buddy.SandboxFetch
	t.Run("CreateFetch", testSandboxFetchCreate(seed.Client, seed.Workspace, &sandbox, &newFetchOps, &newSandboxFetch))
	t.Run("Wait For New Fetch", testSandboxFetchWaitForFetched(seed.Client, seed.Workspace, &sandbox, &newSandboxFetch, &newSandboxFetch))
	t.Run("GetFetch", testSandboxFetchGet(seed.Client, seed.Workspace, &sandbox, &newSandboxFetch))
	t.Run("Refetch", testSandboxFetchRefetch(seed.Client, seed.Workspace, &sandbox, &newSandboxFetch))
	t.Run("GetFetches After Create", testSandboxFetches(seed.Client, seed.Workspace, &sandbox, 2))
	t.Run("DeleteFetch", testSandboxFetchDelete(seed.Client, seed.Workspace, &sandbox, &newSandboxFetch))
	t.Run("GetFetches After Delete", testSandboxFetches(seed.Client, seed.Workspace, &sandbox, 1))
	t.Run("Update Fetch", testSandboxPatch(seed.Client, seed.Workspace, &updateFetchOps, &sandbox))
	t.Run("Clear Fetch", testSandboxPatch(seed.Client, seed.Workspace, &clearFetchOps, &sandbox))
	t.Run("Clear Allowed Pipelines", testSandboxPatch(seed.Client, seed.Workspace, &clearAllowedPipelinesOps, &sandbox))
	t.Run("Delete", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
	sandbox = buddy.Sandbox{}
	t.Run("Create In Workspace", testSandboxCreate(seed.Client, seed.Workspace, nil, &createOps, &sandbox))
	t.Run("Wait For Running In Workspace", testSandboxWaitForRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for Setup In Workspace", testSandboxWaitForSetupDone(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for App In Workspace", testSandboxWaitForAppRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
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
	t.Run("Wait for Setup In Workspace Env", testSandboxWaitForSetupDone(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for App In Workspace Env", testSandboxWaitForAppRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
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
	t.Run("Wait for Setup In Project Env", testSandboxWaitForSetupDone(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Wait for App In Project Env", testSandboxWaitForAppRunning(seed.Client, seed.Workspace, &sandbox, &sandbox))
	t.Run("Get In Project Env", testSandboxGet(seed.Client, seed.Workspace, &sandbox))
	t.Run("GetList In Project Env", testSandboxes(seed.Client, seed.Workspace, &envQuery))
	t.Run("Delete In Project Env", testSandboxDelete(seed.Client, seed.Workspace, &sandbox))
}
