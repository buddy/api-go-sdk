package test

import (
	"errors"
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyz"
)

func GetClient() (*buddy.Client, error) {
	return buddy.NewClient(os.Getenv("BUDDY_TOKEN"), os.Getenv("BUDDY_BASE_URL"), os.Getenv("BUDDY_INSECURE") == "true")
}

func RandStringFromCharSet(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}

func RandInt() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func RandString(strlen int) string {
	return RandStringFromCharSet(strlen, CharSetAlpha)
}

func RandEmail() string {
	return fmt.Sprintf("%s@0zxc.com", UniqueString())
}

func UniqueString() string {
	return fmt.Sprintf("%s%d", RandString(5), time.Now().UnixNano())
}

func ErrorFormatted(msg string, err error) error {
	return fmt.Errorf("%s\n%s", msg, err.Error())
}

func CheckFieldEqual(field string, got string, want string) error {
	if got != want {
		return ErrorFieldFormatted(field, got, want)
	}
	return nil
}

func CheckFieldEqualAndSet(field string, got string, want string) error {
	if err := CheckFieldEqual(field, got, want); err != nil {
		return err
	}
	return CheckFieldSet(field, got)
}

func CheckFieldSet(field string, got string) error {
	if got == "" {
		return ErrorFieldEmpty(field)
	}
	return nil
}

func CheckBoolFieldEqual(field string, got bool, want bool) error {
	if got != want {
		return ErrorFieldFormatted(field, strconv.FormatBool(got), strconv.FormatBool(want))
	}
	return nil
}

func CheckIntFieldEqual(field string, got int, want int) error {
	if got != want {
		return ErrorFieldFormatted(field, strconv.Itoa(got), strconv.Itoa(want))
	}
	return nil
}

func CheckIntFieldEqualAndSet(field string, got int, want int) error {
	if err := CheckIntFieldEqual(field, got, want); err != nil {
		return err
	}
	return CheckIntFieldSet(field, got)
}

func CheckIntFieldSet(field string, got int) error {
	if got == 0 {
		return ErrorFieldEmpty(field)
	}
	return nil
}

func ErrorFieldFormatted(field string, got string, want string) error {
	return fmt.Errorf("got %q %q; want %q", field, got, want)
}

func ErrorFieldEmpty(field string) error {
	return fmt.Errorf("expected %q not to be empty", field)
}

type SeedOps struct {
	workspace  bool
	project    bool
	group      bool
	member     bool
	permission bool
}

type Seed struct {
	Client      *buddy.Client
	Workspace   *buddy.Workspace
	Project     *buddy.Project
	Group       *buddy.Group
	Member      *buddy.Member
	Permission  *buddy.Permission
	Permission2 *buddy.Permission
}

func SeedInitialData(ops *SeedOps) (*Seed, error) {
	var seed Seed
	client, err := GetClient()
	if err != nil {
		return nil, err
	}
	seed.Client = client
	if ops != nil && ops.workspace {
		domain := UniqueString()
		w := buddy.WorkspaceCreateOps{
			Domain: &domain,
		}
		workspace, _, err := client.WorkspaceService.Create(&w)
		if err != nil {
			return nil, err
		}
		seed.Workspace = workspace
		if ops.project {
			projectDisplayName := UniqueString()
			p := buddy.ProjectCreateOps{
				DisplayName: &projectDisplayName,
			}
			project, _, err := client.ProjectService.Create(domain, &p)
			if err != nil {
				return nil, err
			}
			seed.Project = project
		}
		if ops.group {
			groupName := UniqueString()
			g := buddy.GroupOps{
				Name: &groupName,
			}
			group, _, err := client.GroupService.Create(domain, &g)
			if err != nil {
				return nil, err
			}
			seed.Group = group
		}
		if ops.member {
			email := RandEmail()
			m := buddy.MemberOps{
				Email: &email,
			}
			member, _, err := client.MemberService.Create(domain, &m)
			if err != nil {
				return nil, err
			}
			seed.Member = member
		}
		if ops.permission {
			// 1
			name := UniqueString()
			sandboxAccessLevel := buddy.PermissionAccessLevelReadWrite
			repositoryAccessLevel := buddy.PermissionAccessLevelReadWrite
			pipelineAccessLevel := buddy.PermissionAccessLevelReadWrite
			p := buddy.PermissionOps{
				Name:                  &name,
				SandboxAccessLevel:    &sandboxAccessLevel,
				RepositoryAccessLevel: &repositoryAccessLevel,
				PipelineAccessLevel:   &pipelineAccessLevel,
			}
			permission, _, err := client.PermissionService.Create(domain, &p)
			if err != nil {
				return nil, err
			}
			seed.Permission = permission
			// 2
			name = UniqueString()
			sandboxAccessLevel = buddy.PermissionAccessLevelReadOnly
			repositoryAccessLevel = buddy.PermissionAccessLevelManage
			pipelineAccessLevel = buddy.PermissionAccessLevelRunOnly
			p = buddy.PermissionOps{
				Name:                  &name,
				SandboxAccessLevel:    &sandboxAccessLevel,
				RepositoryAccessLevel: &repositoryAccessLevel,
				PipelineAccessLevel:   &pipelineAccessLevel,
			}
			permission, _, err = client.PermissionService.Create(domain, &p)
			if err != nil {
				return nil, err
			}
			seed.Permission2 = permission
		}
	}
	return &seed, nil
}

func CheckProject(project *buddy.Project, name string, displayName string) error {
	if err := CheckFieldSet("Project.Url", project.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.HtmlUrl", project.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Project.Name", project.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Project.DisplayName", project.DisplayName, displayName); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.Status", project.Status); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.CreateDate", project.CreateDate); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.HttpRepository", project.HttpRepository); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.SshRepository", project.SshRepository); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.SshPublicKey", project.SshPublicKey); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.KeyFingerprint", project.KeyFingerprint); err != nil {
		return err
	}
	if err := CheckFieldSet("Project.DefaultBranch", project.DefaultBranch); err != nil {
		return err
	}
	if err := CheckMember(project.CreatedBy, "", "", true, true, 0); err != nil {
		return err
	}
	return nil
}

func CheckProjectGroup(projectGroup *buddy.ProjectGroup, group *buddy.Group, permission *buddy.Permission) error {
	if err := CheckGroup(&projectGroup.Group, group.Name, group.Description, group.Id); err != nil {
		return err
	}
	if err := CheckPermission(projectGroup.PermissionSet, permission.Name, permission.Description, permission.Id, permission.PipelineAccessLevel, permission.RepositoryAccessLevel, permission.SandboxAccessLevel); err != nil {
		return err
	}
	return nil
}

func CheckProjectMember(projectMember *buddy.ProjectMember, member *buddy.Member, permission *buddy.Permission) error {
	if err := CheckMember(&projectMember.Member, member.Email, member.Name, member.Admin, member.WorkspaceOwner, member.Id); err != nil {
		return err
	}
	if err := CheckPermission(projectMember.PermissionSet, permission.Name, permission.Description, permission.Id, permission.PipelineAccessLevel, permission.RepositoryAccessLevel, permission.SandboxAccessLevel); err != nil {
		return err
	}
	return nil
}

func CheckMember(member *buddy.Member, email string, name string, admin bool, owner bool, id int) error {
	if err := CheckFieldSet("Member.Url", member.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Member.HtmlUrl", member.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Member.Id", member.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Member.Id", member.Id); err != nil {
			return err
		}
	}
	if name != "" {
		if err := CheckFieldEqualAndSet("Member.Name", member.Name, name); err != nil {
			return err
		}
	}
	if email != "" {
		if err := CheckFieldEqualAndSet("Member.Email", member.Email, email); err != nil {
			return err
		}
	} else {
		if err := CheckFieldSet("Member.Email", member.Email); err != nil {
			return err
		}
	}
	if err := CheckFieldSet("Member.AvatarUrl", member.AvatarUrl); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Member.Admin", member.Admin, admin); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Member.WorkspaceOwner", member.WorkspaceOwner, owner); err != nil {
		return err
	}
	return nil
}

func CheckMembers(members *buddy.Members, count int) error {
	if err := CheckFieldSet("Members.HtmlUrl", members.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Members.Url", members.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Members)", len(members.Members), count); err != nil {
		return err
	}
	return nil
}

func CheckProfile(profile *buddy.Profile, name string) error {
	if err := CheckFieldSet("Profile.HtmlUrl", profile.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Profile.Url", profile.Url); err != nil {
		return err
	}
	if err := CheckIntFieldSet("Profile.Id", profile.Id); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Profile.Name", profile.Name, name); err != nil {
		return err
	}
	if err := CheckFieldSet("Profile.AvatarUrl", profile.AvatarUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Profile.WorkspacesUrl", profile.WorkspacesUrl); err != nil {
		return err
	}
	return nil
}

func CheckProjects(projects *buddy.Projects, count int) error {
	if err := CheckFieldSet("Projects.HtmlUrl", projects.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Projects.Url", projects.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Projects)", len(projects.Projects), count); err != nil {
		return err
	}
	return nil
}

func CheckGroups(groups *buddy.Groups, count int) error {
	if err := CheckFieldSet("Groups.HtmlUrl", groups.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Groups.Url", groups.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Groups)", len(groups.Groups), count); err != nil {
		return err
	}
	return nil
}

func CheckPermission(permission *buddy.Permission, name string, desc string, id int, pipelineAccessLevel string, repoAccessLevel string, sandboxAccessLevel string) error {
	if err := CheckFieldSet("Permission.Url", permission.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Permission.HtmlUrl", permission.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Permission.Id", permission.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Permission.Id", permission.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqual("Permission.Name", permission.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqual("Permission.Description", permission.Description, desc); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Permission.PipelineAccessLevel", permission.PipelineAccessLevel, pipelineAccessLevel); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Permission.RepositoryAccessLevel", permission.RepositoryAccessLevel, repoAccessLevel); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Permission.SandboxAccessLevel", permission.SandboxAccessLevel, sandboxAccessLevel); err != nil {
		return err
	}
	return nil
}

func CheckWorkspaces(workspaces *buddy.Workspaces, atLeast int) error {
	if err := CheckFieldSet("Workspaces.HtmlUrl", workspaces.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Workspaces.Url", workspaces.Url); err != nil {
		return err
	}
	if len(workspaces.Workspaces) < atLeast {
		return errors.New("len(Workspaces) should be at least " + strconv.Itoa(atLeast))
	}
	return nil
}

func CheckWebhook(webhook *buddy.Webhook, targetUrl string, secretKey string, projectName string, event string, id int) error {
	if err := CheckFieldSet("Webhook.Url", webhook.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Webhook.HtmlUrl", webhook.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Webhook.Id", webhook.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Webhook.Id", webhook.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqualAndSet("Webhook.TargetUrl", webhook.TargetUrl, targetUrl); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Webhook.SecretKey", webhook.SecretKey, secretKey); err != nil {
		return err
	}
	if err := CheckIntFieldEqualAndSet("len(Webhook.Projects)", len(webhook.Projects), 1); err != nil {
		return err
	}
	if err := CheckIntFieldEqualAndSet("len(Webhook.Events)", len(webhook.Events), 1); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Webhook.Projects[0]", webhook.Projects[0], projectName); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Webhook.Events[0]", webhook.Events[0], event); err != nil {
		return err
	}
	return nil
}

func CheckWorkspace(workspace *buddy.Workspace, name string, domain string, id int) error {
	if err := CheckFieldSet("Workspace.Url", workspace.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Workspace.HtmlUrl", workspace.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Workspace.Id", workspace.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Workspace.Id", workspace.Id); err != nil {
			return err
		}
	}
	if err := CheckIntFieldSet("Workspace.OwnerId", workspace.OwnerId); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Workspace.Name", workspace.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Workspace.Domain", workspace.Domain, domain); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Workspace.Frozen", workspace.Frozen, false); err != nil {
		return err
	}
	if err := CheckFieldSet("Workspace.CreateDate", workspace.CreateDate); err != nil {
		return err
	}
	return nil
}

func CheckVariable(variable *buddy.Variable, key string, val string, typ string, desc string, set bool, enc bool, fileName string, filePath string, fileChmod string, filePlace string, id int) error {
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Variable.Id", variable.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Variable.Id", variable.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqualAndSet("Variable.Key", variable.Key, key); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Variable.Type", variable.Type, typ); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Variable.Encrypted", variable.Encrypted, enc); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Variable.Settable", variable.Settable, set); err != nil {
		return err
	}
	if err := CheckFieldEqual("Variable.Description", variable.Description, desc); err != nil {
		return err
	}
	if typ == buddy.VariableTypeSshKey {
		if err := CheckFieldEqualAndSet("Variable.FileName", variable.FileName, fileName); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Variable.FilePath", variable.FilePath, filePath); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Variable.FileChmod", variable.FileChmod, fileChmod); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Variable.FilePlace", variable.FilePlace, filePlace); err != nil {
			return err
		}
		if err := CheckFieldSet("Variable.PublicValue", variable.PublicValue); err != nil {
			return err
		}
		if err := CheckFieldSet("Variable.KeyFingerprint", variable.KeyFingerprint); err != nil {
			return err
		}
		if err := CheckFieldSet("Variable.Checksum", variable.Checksum); err != nil {
			return err
		}
	}
	if enc {
		if err := CheckFieldSet("Variable.Value", variable.Value); err != nil {
			return err
		}
	} else {
		if err := CheckFieldEqual("Variable.Value", variable.Value, val); err != nil {
			return err
		}
	}
	return nil
}

func CheckSourceFile(sf *buddy.SourceFile, name string, path string, message string) error {
	if sf.Content == nil {
		return errors.New("SourceFile.Content can not be nil")
	}
	if err := CheckFieldSet("SourceFile.Content.Url", sf.Content.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Content.HtmlUrl", sf.Content.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceFile.Content.ContentType", sf.Content.ContentType, buddy.SourceContentTypeFile); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceFile.Content.Encoding", sf.Content.Encoding, "base64"); err != nil {
		return err
	}
	if err := CheckIntFieldSet("SourceFile.Content.Size", sf.Content.Size); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceFile.Content.Name", sf.Content.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceFile.Content.Path", sf.Content.Path, path); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Content.Content", sf.Content.Content); err != nil {
		return err
	}
	if sf.Commit == nil {
		return errors.New("SourceFile.Commit can not be nil")
	}
	if sf.Commit.Committer == nil {
		return errors.New("SourceFile.Commit.Committer can not be nil")
	}
	if sf.Commit.Author == nil {
		return errors.New("SourceFile.Commit.Author can not be nil")
	}
	if err := CheckFieldSet("SourceFile.Commit.Url", sf.Commit.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Commit.HtmlUrl", sf.Commit.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Commit.Revision", sf.Commit.Revision); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Commit.AuthorDate", sf.Commit.AuthorDate); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceFile.Commit.CommitDate", sf.Commit.CommitDate); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceFile.Commit.Message", sf.Commit.Message, message); err != nil {
		return err
	}
	if err := CheckMember(sf.Commit.Committer, "", "", true, true, 0); err != nil {
		return err
	}
	if err := CheckMember(sf.Commit.Author, "", "", true, true, 0); err != nil {
		return err
	}
	return nil
}

func CheckGroup(group *buddy.Group, name string, desc string, id int) error {
	if err := CheckFieldSet("Group.Url", group.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Group.HtmlUrl", group.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Group.Id", group.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Group.Id", group.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqual("Group.Name", group.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqual("Group.Description", group.Description, desc); err != nil {
		return err
	}
	return nil
}

func CheckPublicKey(key *buddy.PublicKey, title string, content string, id int) error {
	if err := CheckFieldSet("PublicKey.HtmlUrl", key.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("PublicKey.Url", key.Url); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("PublicKey.Title", key.Title, title); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("PublicKey.Content", key.Content, content); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("PublicKey.Id", key.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("PublicKey.Id", key.Id); err != nil {
			return err
		}
	}
	return nil
}

func CheckSourceContentsDir(sc *buddy.SourceContents, count int) error {
	if err := CheckFieldSet("SourceContents.HtmlUrl", sc.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceContents.Url", sc.Url); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceContents.ContentType", sc.ContentType, buddy.SourceContentTypeDir); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(SourceContents)", len(sc.Contents), count); err != nil {
		return err
	}
	return nil
}

func CheckSourceContentsFile(sc *buddy.SourceContents, name string, path string) error {
	if err := CheckFieldSet("SourceContent.Url", sc.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceContent.HtmlUrl", sc.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceContent.ContentType", sc.ContentType, buddy.SourceContentTypeFile); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceContent.Encoding", sc.Encoding, "base64"); err != nil {
		return err
	}
	if err := CheckIntFieldSet("SourceContent.Size", sc.Size); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceContent.Name", sc.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("SourceContent.Path", sc.Path, path); err != nil {
		return err
	}
	if err := CheckFieldSet("SourceContent.Content", sc.Content); err != nil {
		return err
	}
	return nil
}

func CheckWebhooks(webhooks *buddy.Webhooks, count int) error {
	if err := CheckFieldSet("Webhooks.HtmlUrl", webhooks.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Webhooks.Url", webhooks.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Webhooks)", len(webhooks.Webhooks), count); err != nil {
		return err
	}
	return nil
}

func CheckVariables(variables *buddy.Variables, count int) error {
	if err := CheckFieldSet("Variables.HtmlUrl", variables.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Variables.Url", variables.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Variables)", len(variables.Variables), count); err != nil {
		return err
	}
	return nil
}

func CheckPermissions(permissions *buddy.Permissions, count int) error {
	if err := CheckFieldSet("Permissions.HtmlUrl", permissions.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Permissions.Url", permissions.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Permissions)", len(permissions.PermissionSets), count); err != nil {
		return err
	}
	return nil
}

func CheckIntegrations(integrations *buddy.Integrations, count int) error {
	if err := CheckFieldSet("Integrations.Url", integrations.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Integrations)", len(integrations.Integrations), count); err != nil {
		return err
	}
	return nil
}

func CheckIntegration(integration *buddy.Integration, expected *buddy.Integration, ops *buddy.IntegrationOps) error {
	name := expected.Name
	typ := expected.Type
	scope := expected.Scope
	projectName := expected.ProjectName
	groupId := expected.GroupId
	hashId := expected.HashId
	if ops != nil {
		if ops.Name != nil {
			name = *ops.Name
		}
		if ops.Type != nil {
			typ = *ops.Type
		}
		if ops.Scope != nil {
			scope = *ops.Scope
		}
		if ops.ProjectName != nil {
			projectName = *ops.ProjectName
		}
		if ops.GroupId != nil {
			groupId = *ops.GroupId
		}
	}
	if scope != buddy.IntegrationScopeProject && scope != buddy.IntegrationScopeGroupInProject && scope != buddy.IntegrationScopeAdminInProject {
		projectName = ""
	}
	if scope != buddy.IntegrationScopeGroup && scope != buddy.IntegrationScopeGroupInProject {
		groupId = 0
	}
	if err := CheckFieldSet("Integration.Url", integration.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Integration.HtmlUrl", integration.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Integration.Name", integration.Name, name); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Integration.Type", integration.Type, typ); err != nil {
		return err
	}
	if err := CheckFieldEqualAndSet("Integration.Scope", integration.Scope, scope); err != nil {
		return err
	}
	if err := CheckFieldEqual("Integration.ProjectName", integration.ProjectName, projectName); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("Integration.GroupId", integration.GroupId, groupId); err != nil {
		return err
	}
	if hashId != "" {
		if err := CheckFieldEqualAndSet("Integration.HashId", integration.HashId, hashId); err != nil {
			return err
		}
	} else {
		if err := CheckFieldSet("Integration.HashId", integration.HashId); err != nil {
			return err
		}
	}
	return nil
}
