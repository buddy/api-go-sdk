package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testRulesetCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.RulesetOps, out *buddy.Ruleset) func(t *testing.T) {
	return func(t *testing.T) {
		ruleset, _, err := client.RulesetService.Create(workspace.Domain, project.Name, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("RulesetService.Create", err))
		}
		err = CheckRuleset(ruleset, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *ruleset
	}
}

func testRulesetUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.RulesetOps, out *buddy.Ruleset) func(t *testing.T) {
	return func(t *testing.T) {
		ruleset, _, err := client.RulesetService.Update(workspace.Domain, project.Name, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("RulesetService.Update", err))
		}
		err = CheckRuleset(ruleset, out, ops)
		if err != nil {
			t.Fatal(err)
		}
		*out = *ruleset
	}
}

func testRulesetGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Ruleset) func(t *testing.T) {
	return func(t *testing.T) {
		ruleset, _, err := client.RulesetService.Get(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("RulesetService.Get", err))
		}
		err = CheckRuleset(ruleset, out, nil)
		if err != nil {
			t.Fatal(err)
		}
		*out = *ruleset
	}
}

func testRulesetGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		rulesets, _, err := client.RulesetService.GetList(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("RulesetService.GetList", err))
		}
		err = CheckRulesets(rulesets, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testRulesetDelete(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Ruleset) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.RulesetService.Delete(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("RulesetService.Delete", err))
		}
		_, resp, err := client.RulesetService.Get(workspace.Domain, project.Name, out.Id)
		if err == nil {
			t.Fatal("RulesetService.Get should fail after RulesetService.Delete")
		}
		if resp == nil || resp.StatusCode != 404 {
			t.Fatal(ErrorFormatted("RulesetService.Get after RulesetService.Delete: expected 404", err))
		}
	}
}

func TestRuleset(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace:     true,
		project:       true,
		pipeline:      true,
		group:         true,
		member:        true,
		projectMember: true,
		permission:    true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	pg := buddy.ProjectGroupOps{
		Id:            &seed.Group.Id,
		PermissionSet: &buddy.ProjectGroupOps{Id: &seed.Permission.Id},
	}
	_, _, err = seed.Client.ProjectGroupService.CreateProjectGroup(seed.Workspace.Domain, seed.Project.Name, &pg)
	if err != nil {
		t.Fatal(ErrorFormatted("ProjectGroupService.CreateProjectGroup", err))
	}
	identifier := UniqueString()
	newIdentifier := UniqueString()
	name := RandString(10)
	note := RandString(10)
	agentNote := RandString(10)
	empty := ""
	refs := []string{buddy.RulesetRefDefault, "refs/heads/release/*"}
	newRefs := []string{"refs/tags/v*"}
	admins := true
	disabled := true
	ops := buddy.RulesetOps{
		Identifier: &identifier,
		Name:       &name,
		Refs:       &refs,
		Note:       &note,
		AgentNote:  &agentNote,
		Bypass: &buddy.RulesetBypass{
			Admins: &admins,
			Others: buddy.RulesetBypassAccessDenied,
			Users: []*buddy.RulesetBypassEntry{
				{Id: seed.Member.Id, Access: buddy.RulesetBypassAccessAllowed},
			},
			Groups: []*buddy.RulesetBypassEntry{
				{Id: seed.Group.Id, Access: buddy.RulesetBypassAccessDenied},
			},
		},
		Rules: &buddy.RulesetRules{
			RequiredChecks: &buddy.RulesetRequiredChecks{
				Pipelines:       []string{seed.Pipeline.Identifier},
				RequireUpToDate: true,
			},
			RequirePullRequest: true,
			BlockForcePush:     true,
			RestrictDeletion:   true,
		},
	}
	// only the fields set change; Bypass and Rules are replaced as a whole
	updOps := buddy.RulesetOps{
		Identifier: &newIdentifier,
		Disabled:   &disabled,
		Refs:       &newRefs,
		Name:       &empty,
		Note:       &empty,
		AgentNote:  &empty,
		Bypass:     &buddy.RulesetBypass{},
		Rules: &buddy.RulesetRules{
			RestrictCreation: true,
			RestrictUpdate:   true,
		},
	}
	var ruleset buddy.Ruleset
	t.Run("Create", testRulesetCreate(seed.Client, seed.Workspace, seed.Project, &ops, &ruleset))
	t.Run("Update", testRulesetUpdate(seed.Client, seed.Workspace, seed.Project, &updOps, &ruleset))
	t.Run("Get", testRulesetGet(seed.Client, seed.Workspace, seed.Project, &ruleset))
	t.Run("GetList", testRulesetGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testRulesetDelete(seed.Client, seed.Workspace, seed.Project, &ruleset))
}

func TestRulesetRequiredChecks(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
		pipeline:  true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	domain := seed.Workspace.Domain
	projectName := seed.Project.Name
	refs := []string{buddy.RulesetRefDefault}
	t.Run("RejectPipelineNotMarkingCommits", func(t *testing.T) {
		name := UniqueString()
		doNotCreateCommitStatus := true
		pipeline, _, err := seed.Client.PipelineService.Create(domain, projectName, &buddy.PipelineOps{
			Name:                    &name,
			DoNotCreateCommitStatus: &doNotCreateCommitStatus,
		})
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Create", err))
		}
		identifier := UniqueString()
		_, resp, err := seed.Client.RulesetService.Create(domain, projectName, &buddy.RulesetOps{
			Identifier: &identifier,
			Refs:       &refs,
			Rules: &buddy.RulesetRules{
				RequiredChecks: &buddy.RulesetRequiredChecks{Pipelines: []string{pipeline.Identifier}},
			},
		})
		if err == nil {
			t.Fatal("RulesetService.Create should reject a pipeline that does not mark commits")
		}
		if resp == nil || resp.StatusCode != 400 {
			t.Fatal(ErrorFormatted("RulesetService.Create with a pipeline that does not mark commits: expected 400", err))
		}
	})
	identifier := UniqueString()
	ops := buddy.RulesetOps{
		Identifier: &identifier,
		Refs:       &refs,
		Rules: &buddy.RulesetRules{
			RequiredChecks: &buddy.RulesetRequiredChecks{Pipelines: []string{seed.Pipeline.Identifier}},
		},
	}
	var ruleset buddy.Ruleset
	t.Run("Create", testRulesetCreate(seed.Client, seed.Workspace, seed.Project, &ops, &ruleset))
	t.Run("BlockStopMarkingCommits", func(t *testing.T) {
		doNotCreateCommitStatus := true
		_, resp, err := seed.Client.PipelineService.Update(domain, projectName, seed.Pipeline.Id, &buddy.PipelineOps{
			DoNotCreateCommitStatus: &doNotCreateCommitStatus,
		})
		if err == nil {
			t.Fatal("PipelineService.Update should not turn off commit statuses of a required pipeline")
		}
		if resp == nil || resp.StatusCode != 400 {
			t.Fatal(ErrorFormatted("PipelineService.Update of a required pipeline: expected 400", err))
		}
	})
	t.Run("MissingSources", func(t *testing.T) {
		_, err := seed.Client.PipelineService.Delete(domain, projectName, seed.Pipeline.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("PipelineService.Delete", err))
		}
		ruleset.MissingSources = []*buddy.RulesetSource{
			{Type: buddy.RulesetCheckSourceTypePipeline, Identifier: seed.Pipeline.Identifier},
		}
		testRulesetGet(seed.Client, seed.Workspace, seed.Project, &ruleset)(t)
	})
	t.Run("Delete", testRulesetDelete(seed.Client, seed.Workspace, seed.Project, &ruleset))
}
