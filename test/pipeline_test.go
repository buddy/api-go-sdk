package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
	"time"
)

func testPipelineCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.PipelineOps, out *buddy.Pipeline) func(t *testing.T) {
	return func(t *testing.T) {
		pipeline, _, err := client.PipelineService.Create(workspace.Domain, project.Name, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Create", err))
		}
		err = CheckPipeline(project, pipeline, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *pipeline
	}
}

func testPipelineUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.PipelineOps, out *buddy.Pipeline) func(t *testing.T) {
	return func(t *testing.T) {
		pipeline, _, err := client.PipelineService.Update(workspace.Domain, project.Name, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Update", err))
		}
		err = CheckPipeline(project, pipeline, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *pipeline
	}
}

func testPipelineGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Pipeline) func(t *testing.T) {
	return func(t *testing.T) {
		pipeline, _, err := client.PipelineService.Get(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Get", err))
		}
		err = CheckPipeline(project, pipeline, out, nil)
		if err != nil {
			t.Fatal(err)
		}
		*out = *pipeline
	}
}

func testPipelineDelete(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Pipeline) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.PipelineService.Delete(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Delete", err))
		}
	}
}

func testPipelineGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		query := buddy.PageQuery{
			Page:    1,
			PerPage: 30,
		}
		pipelines, _, err := client.PipelineService.GetList(workspace.Domain, project.Name, &query)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.GetList", err))
		}
		err = CheckPipelines(pipelines, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testPipelineGetListAll(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		pipelines, _, err := client.PipelineService.GetListAll(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.GetList", err))
		}
		err = CheckPipelines(pipelines, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testCreateProjectWithYaml(client *buddy.Client, workspace *buddy.Workspace, remotePath string, out *buddy.Project) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		ops := buddy.ProjectCreateOps{
			DisplayName: &name,
		}
		project, _, err := client.ProjectService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("ProjectService.Create", err))
		}
		*out = *project
		message := RandString(10)
		yaml := `
- pipeline: "test"
  on: "CLICK"
  refs:
  - "refs/heads/master"
  actions:
  - action: "Execute: ls"
    type: "BUILD"
    working_directory: "/buddy/test"
    docker_image_name: "library/ubuntu"
    docker_image_tag: "18.04"
    execute_commands:
    - "!{cmd}"
    volume_mappings:
    - "/:/buddy/test"
    cache_base_image: true
    shell: "BASH"
`
		fileOps := buddy.SourceFileOps{
			ContentRaw: &yaml,
			Path:       &remotePath,
			Message:    &message,
		}
		_, _, err = client.SourceService.CreateFile(workspace.Domain, project.Name, &fileOps)
		if err != nil {
			t.Fatal(ErrorFormatted("SourceService.CreateFile", err))
		}
	}
}

func TestPipelineSchedule(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	newName := RandString(10)
	startDate := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	newStartDate := time.Now().UTC().Add(5 * time.Hour).Format(time.RFC3339)
	priority := buddy.PipelinePriorityLow
	newPriority := buddy.PipelinePriorityHigh
	delay := 5
	newDelay := 7
	paused := true
	newPaused := false
	failOnPrepareEnvWarning := true
	newFailOnPrepareEnvWarning := false
	fetchAllRefs := true
	newFetchAllRefs := false
	on := buddy.PipelineOnSchedule
	ops := buddy.PipelineOps{
		Name:                    &name,
		StartDate:               &startDate,
		Priority:                &priority,
		Delay:                   &delay,
		On:                      &on,
		Paused:                  &paused,
		FailOnPrepareEnvWarning: &failOnPrepareEnvWarning,
		FetchAllRefs:            &fetchAllRefs,
	}
	var pipeline buddy.Pipeline
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	updateOps := buddy.PipelineOps{
		Name:                    &newName,
		StartDate:               &newStartDate,
		Priority:                &newPriority,
		Delay:                   &newDelay,
		Paused:                  &newPaused,
		FetchAllRefs:            &newFetchAllRefs,
		FailOnPrepareEnvWarning: &newFailOnPrepareEnvWarning,
	}
	t.Run("Update", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	t.Run("Get", testPipelineGet(seed.Client, seed.Workspace, seed.Project, &pipeline))
	t.Run("GetList", testPipelineGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("GetListAll", testPipelineGetListAll(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testPipelineDelete(seed.Client, seed.Workspace, seed.Project, &pipeline))
}

func TestPipelineScheduleCron(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	cron := "15 14 1 * *"
	newCron := "0 22 * * 1-5"
	newDisabled := true
	newDisabledReason := RandString(10)
	on := buddy.PipelineOnSchedule
	ops := buddy.PipelineOps{
		Name: &name,
		On:   &on,
		Cron: &cron,
	}
	updateOps := buddy.PipelineOps{
		Cron:           &newCron,
		Disabled:       &newDisabled,
		DisabledReason: &newDisabledReason,
	}
	var pipeline buddy.Pipeline
	// by default its true
	pipeline.FailOnPrepareEnvWarning = true
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	t.Run("Update", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	t.Run("Get", testPipelineGet(seed.Client, seed.Workspace, seed.Project, &pipeline))
	t.Run("GetList", testPipelineGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("GetListAll", testPipelineGetListAll(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testPipelineDelete(seed.Client, seed.Workspace, seed.Project, &pipeline))
}

func TestPipelineEvent(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
		member:    true,
		group:     true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	on := buddy.PipelineOnEvent
	eventType := buddy.PipelineEventTypePush
	eventRef := RandString(10)
	eventRefs := []string{eventRef}
	event := buddy.PipelineEvent{
		Type: eventType,
		Refs: eventRefs,
	}
	events := []*buddy.PipelineEvent{&event}
	tc := buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionOnChange,
	}
	tcs := []*buddy.PipelineTriggerCondition{&tc}
	ops := buddy.PipelineOps{
		Name:              &name,
		On:                &on,
		Events:            &events,
		TriggerConditions: &tcs,
	}
	var pipeline buddy.Pipeline
	// by default true
	pipeline.FailOnPrepareEnvWarning = true
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	newEventType := buddy.PipelineEventTypeCreateRef
	newEventRef := RandString(10)
	newEventRefs := []string{newEventRef}
	newEvent := buddy.PipelineEvent{
		Type: newEventType,
		Refs: newEventRefs,
	}
	newEvents := []*buddy.PipelineEvent{&newEvent}
	updateEventOps := buddy.PipelineOps{
		Events: &newEvents,
	}
	t.Run("UpdateEvent", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateEventOps, &pipeline))
	newTcPaths := []string{RandString(10)}
	newTc := buddy.PipelineTriggerCondition{
		TriggerCondition:      buddy.PipelineTriggerConditionOnChangeAtPath,
		TriggerConditionPaths: newTcPaths,
	}
	newTcs := []*buddy.PipelineTriggerCondition{&newTc}
	updateOps := buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcChangeAtPath", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition:     buddy.PipelineTriggerConditionVarIs,
		TriggerVariableKey:   RandString(10),
		TriggerVariableValue: RandString(10),
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcVarIs", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition:     buddy.PipelineTriggerConditionVarIsNot,
		TriggerVariableKey:   RandString(10),
		TriggerVariableValue: RandString(10),
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcVarIsNot", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition:     buddy.PipelineTriggerConditionVarContains,
		TriggerVariableKey:   RandString(10),
		TriggerVariableValue: RandString(10),
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcVarContains", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition:     buddy.PipelineTriggerConditionVarNotContains,
		TriggerVariableKey:   RandString(10),
		TriggerVariableValue: RandString(10),
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcVarNotContains", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionDateTime,
		TriggerHours:     []int{10},
		TriggerDays:      []int{1},
		ZoneId:           "America/Monterrey",
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcDateTime", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionTriggeringUserIs,
		TriggerUser:      seed.Member.Email,
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcTriggeringUserIs", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionTriggeringUserIsNot,
		TriggerUser:      seed.Member.Email,
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcTriggeringUserIsNot", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionTriggeringUserIsInGroup,
		TriggerGroup:     seed.Group.Name,
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcTriggeringUserIsInGroup", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition: buddy.PipelineTriggerConditionTriggeringUserIsNotInGroup,
		TriggerGroup:     seed.Group.Name,
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcTriggeringUserIsNotInGroup", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	newTc = buddy.PipelineTriggerCondition{
		TriggerCondition:    buddy.PipelineTriggerConditionSuccessPipeline,
		TriggerProjectName:  seed.Project.Name,
		TriggerPipelineName: name,
	}
	updateOps = buddy.PipelineOps{
		TriggerConditions: &newTcs,
	}
	t.Run("UpdateTcSuccessPipeline", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
}

func TestPipelineRemote(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	newName := RandString(10)
	remotePath := RandString(10)
	remotePath2 := RandString(10)
	remoteBranch := "master"
	paramKey := "cmd"
	paramVal := "ls"
	source := buddy.PipelineDefinitionSourceRemote
	remoteParam := buddy.PipelineRemoteParameter{
		Key:   paramKey,
		Value: paramVal,
	}
	remoteParams := []*buddy.PipelineRemoteParameter{&remoteParam}
	var remoteProject buddy.Project
	var remoteProject2 buddy.Project
	t.Run("YamlFile", testCreateProjectWithYaml(seed.Client, seed.Workspace, remotePath, &remoteProject))
	t.Run("YamlFile2", testCreateProjectWithYaml(seed.Client, seed.Workspace, remotePath2, &remoteProject2))
	ops := buddy.PipelineOps{
		Name:              &name,
		RemoteBranch:      &remoteBranch,
		RemotePath:        &remotePath,
		RemoteProjectName: &remoteProject.Name,
		RemoteParameters:  &remoteParams,
		DefinitionSource:  &source,
	}
	var pipeline buddy.Pipeline
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	updateOps := buddy.PipelineOps{
		Name:              &newName,
		RemoteBranch:      &remoteBranch,
		RemotePath:        &remotePath2,
		RemoteProjectName: &remoteProject2.Name,
		RemoteParameters:  &remoteParams,
		DefinitionSource:  &source,
	}
	t.Run("Update", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	t.Run("Get", testPipelineGet(seed.Client, seed.Workspace, seed.Project, &pipeline))
	t.Run("GetList", testPipelineGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("GetListAll", testPipelineGetListAll(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testPipelineDelete(seed.Client, seed.Workspace, seed.Project, &pipeline))
}

func TestPipelineWithPermissions(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:  true,
		project:    true,
		member:     true,
		group:      true,
		permission: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	t.Run("CreateProjectMember", testProjectMemberCreate(seed.Client, seed.Workspace, seed.Project, seed.Member, seed.Permission))
	t.Run("CreateProjectGroup", testProjectGroupCreate(seed.Client, seed.Workspace, seed.Project, seed.Group, seed.Permission2))
	name := RandString(10)
	on := buddy.PipelineOnClick
	ref := RandString(10)
	refs := []string{ref}
	failOnPrepareEnvWarning := true
	userPerm := buddy.PipelineResourcePermission{
		Id:          seed.Member.Id,
		AccessLevel: buddy.PipelinePermissionReadWrite,
	}
	groupPerm := buddy.PipelineResourcePermission{
		Id:          seed.Group.Id,
		AccessLevel: buddy.PipelinePermissionDefault,
	}
	perms := buddy.PipelinePermissions{
		Others: buddy.PipelinePermissionDenied,
		Groups: []*buddy.PipelineResourcePermission{&groupPerm},
		Users:  []*buddy.PipelineResourcePermission{&userPerm},
	}
	ops := buddy.PipelineOps{
		Name:                    &name,
		On:                      &on,
		Refs:                    &refs,
		Permissions:             &perms,
		FailOnPrepareEnvWarning: &failOnPrepareEnvWarning,
	}
	updUserPem := buddy.PipelineResourcePermission{
		Id:          seed.Member.Id,
		AccessLevel: buddy.PipelinePermissionReadOnly,
	}
	updPerms := buddy.PipelinePermissions{
		Others: buddy.PipelinePermissionReadWrite,
		Users:  []*buddy.PipelineResourcePermission{&updUserPem},
	}
	updOps := buddy.PipelineOps{
		Permissions: &updPerms,
	}
	var pipeline buddy.Pipeline
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	t.Run("Update", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updOps, &pipeline))
	t.Run("Get", testPipelineGet(seed.Client, seed.Workspace, seed.Project, &pipeline))
	t.Run("GetList", testPipelineGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("GetListAll", testPipelineGetListAll(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testPipelineDelete(seed.Client, seed.Workspace, seed.Project, &pipeline))
}

func TestPipelineClick(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	on := buddy.PipelineOnClick
	ref := RandString(10)
	alwaysFromScratch := true
	failOnPrepareEnvWarning := true
	fetchAllRefs := true
	autoClearCache := true
	noSkipToMostRecent := true
	ignoreFailOnProjectStatus := true
	executionMessageTemplate := RandString(10)
	targetSiteUrl := RandString(10)
	refs := []string{ref}
	cloneDepth := 10
	ops := buddy.PipelineOps{
		Name:                      &name,
		On:                        &on,
		AlwaysFromScratch:         &alwaysFromScratch,
		FailOnPrepareEnvWarning:   &failOnPrepareEnvWarning,
		FetchAllRefs:              &fetchAllRefs,
		AutoClearCache:            &autoClearCache,
		NoSkipToMostRecent:        &noSkipToMostRecent,
		IgnoreFailOnProjectStatus: &ignoreFailOnProjectStatus,
		ExecutionMessageTemplate:  &executionMessageTemplate,
		TargetSiteUrl:             &targetSiteUrl,
		Refs:                      &refs,
		CloneDepth:                &cloneDepth,
	}
	newName := RandString(10)
	newRef := RandString(10)
	newRefs := []string{newRef}
	newMsgTemplate := RandString(10)
	newTargetSiteUrl := RandString(10)
	newCloneDepth := 0
	updateOps := buddy.PipelineOps{
		Name:                     &newName,
		Refs:                     &newRefs,
		ExecutionMessageTemplate: &newMsgTemplate,
		TargetSiteUrl:            &newTargetSiteUrl,
		CloneDepth:               &newCloneDepth,
	}
	var pipeline buddy.Pipeline
	t.Run("Create", testPipelineCreate(seed.Client, seed.Workspace, seed.Project, &ops, &pipeline))
	t.Run("Update", testPipelineUpdate(seed.Client, seed.Workspace, seed.Project, &updateOps, &pipeline))
	t.Run("Get", testPipelineGet(seed.Client, seed.Workspace, seed.Project, &pipeline))
	t.Run("GetList", testPipelineGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("GetListAll", testPipelineGetListAll(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testPipelineDelete(seed.Client, seed.Workspace, seed.Project, &pipeline))
}
