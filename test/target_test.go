package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testTargetCreate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.TargetOps, out *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		target, _, err := client.TargetService.Create(workspace.Domain, project.Name, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Create", err))
		}
		if err := CheckTarget(target, out, ops); err != nil {
			t.Fatal(err)
		}
		*out = *target
	}
}

func testTargetUpdate(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, ops *buddy.TargetOps, out *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		target, _, err := client.TargetService.Update(workspace.Domain, project.Name, out.Id, ops)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Update", err))
		}
		if err := CheckTarget(target, out, ops); err != nil {
			t.Fatal(err)
		}
		*out = *target
	}
}

func testTargetGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		target, _, err := client.TargetService.Get(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Get", err))
		}
		if err := CheckTarget(target, out, nil); err != nil {
			t.Fatal(err)
		}
		*out = *target
	}
}

func testTargetGetList(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, count int) func(t *testing.T) {
	return func(t *testing.T) {
		targets, _, err := client.TargetService.GetList(workspace.Domain, project.Name)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.GetList", err))
		}
		if err := CheckTargets(targets, count); err != nil {
			t.Fatal(err)
		}
	}
}

func testTargetDelete(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.Target) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.TargetService.Delete(workspace.Domain, project.Name, out.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("TargetService.Delete", err))
		}
	}
}

func TestTarget(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	name := RandString(10)
	typ := "LOCAL"
	ops := buddy.TargetOps{
		Name: &name,
		Type: &typ,
	}
	newName := RandString(10)
	newType := "REMOTE"
	updOps := buddy.TargetOps{
		Name: &newName,
		Type: &newType,
	}
	var target buddy.Target
	t.Run("Create", testTargetCreate(seed.Client, seed.Workspace, seed.Project, &ops, &target))
	t.Run("Update", testTargetUpdate(seed.Client, seed.Workspace, seed.Project, &updOps, &target))
	t.Run("Get", testTargetGet(seed.Client, seed.Workspace, seed.Project, &target))
	t.Run("GetList", testTargetGetList(seed.Client, seed.Workspace, seed.Project, 1))
	t.Run("Delete", testTargetDelete(seed.Client, seed.Workspace, seed.Project, &target))
}
