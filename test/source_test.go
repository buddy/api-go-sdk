package test

import (
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testSourceCreateFile(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.SourceFile) func(t *testing.T) {
	return func(t *testing.T) {
		message := RandString(10)
		name := RandString(6)
		path := "/" + name
		content := RandString(10)
		ops := buddy.SourceFileOps{
			Message:    &message,
			Path:       &path,
			ContentRaw: &content,
		}
		sf, _, err := client.SourceService.CreateFile(workspace.Domain, project.Name, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("SourceService.CreateFile", err))
		}
		err = CheckSourceFile(sf, name, name, message)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sf
	}
}

func testSourceUpdateFile(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, out *buddy.SourceFile) func(t *testing.T) {
	return func(t *testing.T) {
		message := RandString(10)
		content := RandString(10)
		ops := buddy.SourceFileOps{
			Message:    &message,
			ContentRaw: &content,
		}
		sf, _, err := client.SourceService.UpdateFile(workspace.Domain, project.Name, out.Content.Path, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("SourceService.CreateFile", err))
		}
		err = CheckSourceFile(sf, out.Content.Name, out.Content.Name, message)
		if err != nil {
			t.Fatal(err)
		}
		*out = *sf
	}
}

func testSourceGet(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, path string, revision *string, count int) func(t *testing.T) {
	return func(t *testing.T) {
		var query *buddy.SourceContentsGetQuery
		if revision != nil {
			query = &buddy.SourceContentsGetQuery{
				Revision: *revision,
			}
		}
		sc, _, err := client.SourceService.Get(workspace.Domain, project.Name, path, query)
		if err != nil {
			t.Fatal(ErrorFormatted("SourceService.Get", err))
		}
		err = CheckSourceContentsDir(sc, count)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testSourceGetFile(client *buddy.Client, workspace *buddy.Workspace, project *buddy.Project, path string) func(t *testing.T) {
	return func(t *testing.T) {
		sc, _, err := client.SourceService.Get(workspace.Domain, project.Name, path, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("SourceService.Get", err))
		}
		err = CheckSourceContentsFile(sc, path, path)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSource(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
		project:   true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var file1 buddy.SourceFile
	var file2 buddy.SourceFile
	t.Run("CreateFile1", testSourceCreateFile(seed.Client, seed.Workspace, seed.Project, &file1))
	t.Run("UpdateFile1", testSourceUpdateFile(seed.Client, seed.Workspace, seed.Project, &file1))
	t.Run("Get1", testSourceGet(seed.Client, seed.Workspace, seed.Project, "", nil, 1))
	t.Run("CreateFile2", testSourceCreateFile(seed.Client, seed.Workspace, seed.Project, &file2))
	t.Run("Get2", testSourceGet(seed.Client, seed.Workspace, seed.Project, "", &file1.Commit.Revision, 1))
	branch := "master"
	t.Run("Get3", testSourceGet(seed.Client, seed.Workspace, seed.Project, "", &branch, 2))
	t.Run("GetFile", testSourceGetFile(seed.Client, seed.Workspace, seed.Project, file2.Content.Path))
}
