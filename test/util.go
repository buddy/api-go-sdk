package test

import (
	"bytes"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"golang.org/x/crypto/ssh"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func CheckStringArrayEqual(field string, got []string, want []string) error {
	if len(got) != len(want) {
		return fmt.Errorf("expected %q to be equal length", field)
	}
	for _, s1 := range want {
		found := false
		for _, s2 := range got {
			if s1 == s2 {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("expected %q to have all elements in array", field)
		}
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
	workspace      bool
	project        bool
	group          bool
	member         bool
	permission     bool
	gitIntegration bool
}

type Seed struct {
	Client         *buddy.Client
	Workspace      *buddy.Workspace
	Project        *buddy.Project
	Group          *buddy.Group
	Member         *buddy.Member
	Permission     *buddy.Permission
	Permission2    *buddy.Permission
	GitIntegration *buddy.Integration
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
		if ops.gitIntegration {
			in := UniqueString()
			it := buddy.IntegrationTypeGitHub
			is := buddy.IntegrationScopeWorkspace
			io := os.Getenv("BUDDY_GH_TOKEN")
			i := buddy.IntegrationOps{
				Name:  &in,
				Type:  &it,
				Scope: &is,
				Token: &io,
			}
			integration, _, err := client.IntegrationService.Create(domain, &i)
			if err != nil {
				return nil, err
			}
			seed.GitIntegration = integration
		}
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
			m := buddy.MemberCreateOps{
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

func CheckProject(project *buddy.Project, name string, displayName string, short bool, updateDefaultBranch bool, allowPullRequests bool, fetchSubmodules bool, fetchSubmodulesKey string, access string, withoutRepository bool) error {
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
	if !short {
		if err := CheckBoolFieldEqual("Project.UpdateDefaultBranchFromExternal", project.UpdateDefaultBranchFromExternal, updateDefaultBranch); err != nil {
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
		if err := CheckFieldSet("Project.DefaultBranch", project.DefaultBranch); err != nil {
			return err
		}
		if err := CheckFieldEqual("Project.Access", project.Access, access); err != nil {
			return err
		}
		if err := CheckBoolFieldEqual("Project.AllowPullRequests", project.AllowPullRequests, allowPullRequests); err != nil {
			return err
		}
		if err := CheckBoolFieldEqual("Project.WithoutRepository", project.WithoutRepository, withoutRepository); err != nil {
			return err
		}
		if err := CheckBoolFieldEqual("Project.FetchSubmodules", project.FetchSubmodules, fetchSubmodules); err != nil {
			return err
		}
		if fetchSubmodules {
			if err := CheckFieldEqual("Project.FetchSubmodulesEnvKey", project.FetchSubmodulesEnvKey, fetchSubmodulesKey); err != nil {
				return err
			}
		}
		if err := CheckMember(project.CreatedBy, "", "", false, 0, true, true, 0, ""); err != nil {
			return err
		}
	}
	return nil
}

func CheckProjectGroup(projectGroup *buddy.ProjectGroup, group *buddy.Group, permission *buddy.Permission) error {
	if err := CheckGroup(&projectGroup.Group, group.Name, group.Description, group.AutoAssignToNewProjects, group.AutoAssignPermissionSetId, group.Id); err != nil {
		return err
	}
	if err := CheckPermission(projectGroup.PermissionSet, permission.Name, permission.Description, permission.Id, permission.PipelineAccessLevel, permission.RepositoryAccessLevel, permission.SandboxAccessLevel, permission.ProjectTeamAccessLevel); err != nil {
		return err
	}
	return nil
}

func CheckProjectMember(projectMember *buddy.ProjectMember, member *buddy.Member, permission *buddy.Permission) error {
	if err := CheckMember(&projectMember.Member, member.Email, member.Name, member.AutoAssignToNewProjects, member.AutoAssignPermissionSetId, member.Admin, member.WorkspaceOwner, member.Id, ""); err != nil {
		return err
	}
	if err := CheckPermission(projectMember.PermissionSet, permission.Name, permission.Description, permission.Id, permission.PipelineAccessLevel, permission.RepositoryAccessLevel, permission.SandboxAccessLevel, permission.ProjectTeamAccessLevel); err != nil {
		return err
	}
	return nil
}

func CheckMember(member *buddy.Member, email string, name string, assignToProject bool, assignToProjectPermId int, admin bool, owner bool, id int, status string) error {
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
	if err := CheckBoolFieldEqual("Member.AutoAssignToNewProjects", member.AutoAssignToNewProjects, assignToProject); err != nil {
		return err
	}
	if assignToProject {
		if err := CheckIntFieldEqual("Member.AutoAssignPermissionSetId", member.AutoAssignPermissionSetId, assignToProjectPermId); err != nil {
			return err
		}
	}
	if status != "" {
		if err := CheckFieldEqualAndSet("Member.Status", member.Status, status); err != nil {
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

func CheckTokens(tokens *buddy.Tokens, count int) error {
	if err := CheckFieldSet("Tokens.HtmlUrl", tokens.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Tokens.Url", tokens.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Tokens)", len(tokens.AccessTokens), count); err != nil {
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

func CheckPermission(permission *buddy.Permission, name string, desc string, id int, pipelineAccessLevel string, repoAccessLevel string, sandboxAccessLevel string, projectTeamAccessLevel string) error {
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
	if err := CheckFieldEqualAndSet("Permission.ProjectTeamAccessLevel", permission.ProjectTeamAccessLevel, projectTeamAccessLevel); err != nil {
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

func CheckSso(sso *buddy.Sso, typ string, ssoUrl string, issuer string, certificate string, signature string, digest string, requireSsoForAllMembers bool) error {
	if err := CheckFieldSet("Sso.Url", sso.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Sso.HtmlUrl", sso.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldEqual("Sso.Type", sso.Type, typ); err != nil {
		return err
	}
	if typ != buddy.SsoTypeOidc {
		if err := CheckFieldEqual("Sso.SsoUrl", sso.SsoUrl, ssoUrl); err != nil {
			return err
		}
		if err := CheckFieldEqual("Sso.Certificate", sso.Certificate, certificate); err != nil {
			return err
		}
		if err := CheckFieldEqual("Sso.SignatureMethod", sso.SignatureMethod, signature); err != nil {
			return err
		}
		if err := CheckFieldEqual("Sso.DigestMethod", sso.DigestMethod, digest); err != nil {
			return err
		}
	}
	if err := CheckFieldEqual("Sso.Issuer", sso.Issuer, issuer); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Sso.RequireSsoForAllMembers", sso.RequireSsoForAllMembers, requireSsoForAllMembers); err != nil {
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

func CheckVariable(variable *buddy.Variable, key string, val string, typ string, desc string, set bool, enc bool, filePath string, fileChmod string, filePlace string, id int) error {
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
	if err := CheckMember(sf.Commit.Committer, "", "", false, 0, true, true, 0, ""); err != nil {
		return err
	}
	if err := CheckMember(sf.Commit.Author, "", "", false, 0, true, true, 0, ""); err != nil {
		return err
	}
	return nil
}

func CheckToken(token *buddy.Token, name string, expiresIn int, expiresAt string, scopes []string, workspaceRestrictions []string, ipRestrictions []string, id string, hasToken bool) error {
	if err := CheckFieldSet("Token.Url", token.Url); err != nil {
		return err
	}
	if hasToken {
		if err := CheckFieldSet("Token.Token", token.Token); err != nil {
			return err
		}
	}
	if err := CheckFieldSet("Token.HtmlUrl", token.HtmlUrl); err != nil {
		return err
	}
	if id != "" {
		if err := CheckFieldEqualAndSet("Token.Id", token.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckFieldSet("Token.Id", token.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqualAndSet("Token.Name", token.Name, name); err != nil {
		return err
	}
	expiresGot, _ := time.Parse(time.RFC3339, token.ExpiresAt)
	var expiresWant time.Time
	if expiresIn != 0 {
		expiresWant = time.Now().AddDate(0, 0, expiresIn)
	} else if expiresAt != "" {
		expiresWant, _ = time.Parse(time.RFC3339, expiresAt)
	}
	if err := CheckIntFieldEqualAndSet("Token.Expires.Year", expiresGot.Year(), expiresWant.Year()); err != nil {
		return err
	}
	if err := CheckIntFieldEqualAndSet("Token.Expires.Month", int(expiresGot.Month()), int(expiresWant.Month())); err != nil {
		return err
	}
	if err := CheckIntFieldEqualAndSet("Token.Expires.Day", expiresGot.Day(), expiresWant.Day()); err != nil {
		return err
	}
	if err := CheckStringArrayEqual("Token.Scopes", token.Scopes, scopes); err != nil {
		return err
	}
	if err := CheckStringArrayEqual("Token.WorkspaceRestrictions", token.WorkspaceRestrictions, workspaceRestrictions); err != nil {
		return err
	}
	if err := CheckStringArrayEqual("Token.IpRestrictions", token.IpRestrictions, ipRestrictions); err != nil {
		return err
	}
	return nil
}

func CheckGroup(group *buddy.Group, name string, desc string, assignToProjects bool, assignToProjectsPermId int, id int) error {
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
	if err := CheckBoolFieldEqual("Group.AutoAssignToNewProjects", group.AutoAssignToNewProjects, assignToProjects); err != nil {
		return err
	}
	if assignToProjects {
		if err := CheckIntFieldEqual("Group.AutoAssignPermissionSetId", group.AutoAssignPermissionSetId, assignToProjectsPermId); err != nil {
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
	if err := CheckFieldEqualAndSet("PublicKey.Content", strings.TrimSpace(key.Content), strings.TrimSpace(content)); err != nil {
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

func CheckPipelines(pipelines *buddy.Pipelines, count int) error {
	if err := CheckFieldSet("Pipelines.HtmlUrl", pipelines.HtmlUrl); err != nil {
		return err
	}
	if err := CheckFieldSet("Pipelines.Url", pipelines.Url); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Pipelines)", len(pipelines.Pipelines), count); err != nil {
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

func CheckPipeline(project *buddy.Project, pipeline *buddy.Pipeline, expected *buddy.Pipeline, ops *buddy.PipelineOps) error {
	name := expected.Name
	on := expected.On
	refs := expected.Refs
	tags := expected.Tags
	events := expected.Events
	triggerConditions := expected.TriggerConditions
	alwaysFromScratch := expected.AlwaysFromScratch
	priority := expected.Priority
	failOnPrepareEnvWarning := expected.FailOnPrepareEnvWarning
	fetchAllRefs := expected.FetchAllRefs
	autoClearCache := expected.AutoClearCache
	noSkipToMostRecent := expected.NoSkipToMostRecent
	doNotCreateCommitStatus := expected.DoNotCreateCommitStatus
	startDate := expected.StartDate
	delay := expected.Delay
	cloneDepth := expected.CloneDepth
	cron := expected.Cron
	paused := expected.Paused
	ignoreFailOnProjectStatus := expected.IgnoreFailOnProjectStatus
	executionMessageTemplate := expected.ExecutionMessageTemplate
	worker := expected.Worker
	targetSiteUrl := expected.TargetSiteUrl
	definitionSource := expected.DefinitionSource
	remotePath := expected.RemotePath
	remoteBranch := expected.RemoteBranch
	remoteProjectName := expected.RemoteProjectName
	remoteParameters := expected.RemoteParameters
	disabled := expected.Disabled
	disabledReason := expected.DisabledReason
	permissions := expected.Permissions
	id := expected.Id
	if ops != nil {
		if ops.Permissions != nil {
			permissions = ops.Permissions
		}
		if ops.Name != nil {
			name = *ops.Name
		}
		if ops.On != nil {
			on = *ops.On
		}
		if ops.Refs != nil {
			refs = *ops.Refs
		}
		if ops.Tags != nil {
			tags = *ops.Tags
		}
		if ops.Events != nil {
			events = *ops.Events
		}
		if ops.TriggerConditions != nil {
			triggerConditions = *ops.TriggerConditions
		}
		if ops.AlwaysFromScratch != nil {
			alwaysFromScratch = *ops.AlwaysFromScratch
		}
		if ops.Priority != nil {
			priority = *ops.Priority
		}
		if ops.FailOnPrepareEnvWarning != nil {
			failOnPrepareEnvWarning = *ops.FailOnPrepareEnvWarning
		}
		if ops.FetchAllRefs != nil {
			fetchAllRefs = *ops.FetchAllRefs
		}
		if ops.AutoClearCache != nil {
			autoClearCache = *ops.AutoClearCache
		}
		if ops.NoSkipToMostRecent != nil {
			noSkipToMostRecent = *ops.NoSkipToMostRecent
		}
		if ops.DoNotCreateCommitStatus != nil {
			doNotCreateCommitStatus = *ops.DoNotCreateCommitStatus
		}
		if ops.StartDate != nil {
			startDate = *ops.StartDate
		}
		if ops.Delay != nil {
			delay = *ops.Delay
		}
		if ops.CloneDepth != nil {
			cloneDepth = *ops.CloneDepth
		}
		if ops.Cron != nil {
			cron = *ops.Cron
		}
		if ops.Paused != nil {
			paused = *ops.Paused
		}
		if ops.IgnoreFailOnProjectStatus != nil {
			ignoreFailOnProjectStatus = *ops.IgnoreFailOnProjectStatus
		}
		if ops.ExecutionMessageTemplate != nil {
			executionMessageTemplate = *ops.ExecutionMessageTemplate
		}
		if ops.Worker != nil {
			worker = *ops.Worker
		}
		if ops.TargetSiteUrl != nil {
			targetSiteUrl = *ops.TargetSiteUrl
		}
		if ops.DefinitionSource != nil {
			definitionSource = *ops.DefinitionSource
		}
		if ops.RemotePath != nil {
			remotePath = *ops.RemotePath
		}
		if ops.RemoteBranch != nil {
			remoteBranch = *ops.RemoteBranch
		}
		if ops.RemoteProjectName != nil {
			remoteProjectName = *ops.RemoteProjectName
		}
		if ops.RemoteParameters != nil {
			remoteParameters = *ops.RemoteParameters
		}
		if ops.Disabled != nil {
			disabled = *ops.Disabled
		}
		if ops.DisabledReason != nil {
			disabledReason = *ops.DisabledReason
		}
	}
	lenRefs := len(refs)
	lenEvents := len(events)
	lenTriggerConditions := len(triggerConditions)
	lenTags := len(tags)
	lenRemoteParameters := len(remoteParameters)
	if err := CheckFieldSet("Pipeline.Url", pipeline.Url); err != nil {
		return err
	}
	if err := CheckFieldSet("Pipeline.HtmlUrl", pipeline.HtmlUrl); err != nil {
		return err
	}
	if id != 0 {
		if err := CheckIntFieldEqualAndSet("Pipeline.Id", pipeline.Id, id); err != nil {
			return err
		}
	} else {
		if err := CheckIntFieldSet("Pipeline.Id", pipeline.Id); err != nil {
			return err
		}
	}
	if err := CheckFieldEqualAndSet("Pipeline.Name", pipeline.Name, name); err != nil {
		return err
	}
	if on != "" {
		if err := CheckFieldEqualAndSet("Pipeline.On", pipeline.On, on); err != nil {
			return err
		}
	}
	if err := CheckIntFieldEqual("len(Pipeline.Refs)", len(pipeline.Refs), lenRefs); err != nil {
		return err
	}
	if lenRefs > 0 {
		if err := CheckFieldEqualAndSet("Pipeline.Refs[0]", pipeline.Refs[0], refs[0]); err != nil {
			return err
		}
	}
	if err := CheckIntFieldEqual("len(Pipeline.Events)", len(pipeline.Events), lenEvents); err != nil {
		return err
	}
	if lenEvents > 0 {
		if err := CheckFieldEqualAndSet("Pipeline.Events[0].Type", pipeline.Events[0].Type, events[0].Type); err != nil {
			return err
		}
		if err := CheckIntFieldEqualAndSet("len(Pipeline.Events[0].Refs)", len(pipeline.Events[0].Refs), len(events[0].Refs)); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Pipeline.Events[0].Refs[0]", pipeline.Events[0].Refs[0], events[0].Refs[0]); err != nil {
			return err
		}
	}
	if permissions != nil {
		if err := CheckFieldEqualAndSet("Pipeline.Permissions.Others", pipeline.Permissions.Others, permissions.Others); err != nil {
			return err
		}
		usersLen := len(permissions.Users)
		groupsLen := len(permissions.Groups)
		if err := CheckIntFieldEqual("len(Pipeline.Permissions.Users)", len(pipeline.Permissions.Users), usersLen); err != nil {
			return err
		}
		if err := CheckIntFieldEqual("len(Pipeline.Permissions.Groups)", len(pipeline.Permissions.Groups), groupsLen); err != nil {
			return err
		}
		if usersLen > 0 {
			if err := CheckIntFieldEqual("Pipeline.Permissions.Users[0].Id", pipeline.Permissions.Users[0].Id, permissions.Users[0].Id); err != nil {
				return err
			}
			if err := CheckFieldEqual("Pipeline.Permissions.Users[0].AccessLevel", pipeline.Permissions.Users[0].AccessLevel, permissions.Users[0].AccessLevel); err != nil {
				return err
			}
		}
		if groupsLen > 0 {
			if err := CheckIntFieldEqual("Pipeline.Permissions.Groups[0].Id", pipeline.Permissions.Groups[0].Id, permissions.Groups[0].Id); err != nil {
				return err
			}
			if err := CheckFieldEqual("Pipeline.Permissions.Groups[0].AccessLevel", pipeline.Permissions.Groups[0].AccessLevel, permissions.Groups[0].AccessLevel); err != nil {
				return err
			}
		}
	}
	if err := CheckIntFieldEqual("len(Pipeline.TriggerConditions)", len(pipeline.TriggerConditions), lenTriggerConditions); err != nil {
		return err
	}
	if lenTriggerConditions > 0 {
		expectedTriggerCondition := triggerConditions[0]
		pipelineTriggerCondition := pipeline.TriggerConditions[0]
		if err := CheckFieldEqualAndSet("PipelineTriggerCondition.TriggerCondition", pipelineTriggerCondition.TriggerCondition, expectedTriggerCondition.TriggerCondition); err != nil {
			return err
		}
		lenPaths := len(expectedTriggerCondition.TriggerConditionPaths)
		if err := CheckIntFieldEqual("len(PipelineTriggerCondition.TriggerConditionPaths)", len(pipelineTriggerCondition.TriggerConditionPaths), lenPaths); err != nil {
			return err
		}
		if lenPaths > 0 {
			if err := CheckFieldEqualAndSet("PipelineTriggerCondition.TriggerConditionPaths[0]", pipelineTriggerCondition.TriggerConditionPaths[0], expectedTriggerCondition.TriggerConditionPaths[0]); err != nil {
				return err
			}
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerVariableKey", pipelineTriggerCondition.TriggerVariableKey, expectedTriggerCondition.TriggerVariableKey); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerVariableValue", pipelineTriggerCondition.TriggerVariableValue, expectedTriggerCondition.TriggerVariableValue); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.ZoneId", pipelineTriggerCondition.ZoneId, expectedTriggerCondition.ZoneId); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerProjectName", pipelineTriggerCondition.TriggerProjectName, expectedTriggerCondition.TriggerProjectName); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerPipelineName", pipelineTriggerCondition.TriggerPipelineName, expectedTriggerCondition.TriggerPipelineName); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerUser", pipelineTriggerCondition.TriggerUser, expectedTriggerCondition.TriggerUser); err != nil {
			return err
		}
		if err := CheckFieldEqual("PipelineTriggerCondition.TriggerGroup", pipelineTriggerCondition.TriggerGroup, expectedTriggerCondition.TriggerGroup); err != nil {
			return err
		}
		lenHours := len(expectedTriggerCondition.TriggerHours)
		lenDays := len(expectedTriggerCondition.TriggerDays)
		if err := CheckIntFieldEqual("len(PipelineTriggerCondition.TriggerHours)", len(pipelineTriggerCondition.TriggerHours), lenHours); err != nil {
			return err
		}
		if lenHours > 0 {
			if err := CheckIntFieldEqualAndSet("PipelineTriggerCondition.TriggerHours[0]", pipelineTriggerCondition.TriggerHours[0], expectedTriggerCondition.TriggerHours[0]); err != nil {
				return err
			}
		}
		if err := CheckIntFieldEqual("len(PipelineTriggerCondition.TriggerDays)", len(pipelineTriggerCondition.TriggerDays), lenDays); err != nil {
			return err
		}
		if lenDays > 0 {
			if err := CheckIntFieldEqualAndSet("PipelineTriggerCondition.TriggerDays[0]", pipelineTriggerCondition.TriggerDays[0], expectedTriggerCondition.TriggerDays[0]); err != nil {
				return err
			}
		}
	}
	if executionMessageTemplate != "" {
		if err := CheckFieldEqualAndSet("Pipeline.ExecutionMessageTemplate", pipeline.ExecutionMessageTemplate, executionMessageTemplate); err != nil {
			return err
		}
	}
	if err := CheckFieldSet("Pipeline.LastExecutionStatus", pipeline.LastExecutionStatus); err != nil {
		return err
	}
	if err := CheckFieldSet("Pipeline.CreateDate", pipeline.CreateDate); err != nil {
		return err
	}
	if priority != "" {
		if err := CheckFieldEqualAndSet("Pipeline.Priority", pipeline.Priority, priority); err != nil {
			return err
		}
	}
	if err := CheckBoolFieldEqual("Pipeline.AlwaysFromScratch", pipeline.AlwaysFromScratch, alwaysFromScratch); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.FailOnPrepareEnvWarning", pipeline.FailOnPrepareEnvWarning, failOnPrepareEnvWarning); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.FetchAllRefs", pipeline.FetchAllRefs, fetchAllRefs); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.AutoClearCache", pipeline.AutoClearCache, autoClearCache); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.NoSkipToMostRecent", pipeline.NoSkipToMostRecent, noSkipToMostRecent); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.DoNotCreateCommitStatus", pipeline.DoNotCreateCommitStatus, doNotCreateCommitStatus); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.IgnoreFailOnProjectStatus", pipeline.IgnoreFailOnProjectStatus, ignoreFailOnProjectStatus); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.StartDate", pipeline.StartDate, startDate); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("Pipeline.Delay", pipeline.Delay, delay); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("Pipeline.CloneDepth", pipeline.CloneDepth, cloneDepth); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.Cron", pipeline.Cron, cron); err != nil {
		return err
	}
	if err := CheckBoolFieldEqual("Pipeline.Paused", pipeline.Paused, paused); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.Worker", pipeline.Worker, worker); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.TargetSiteUrl", pipeline.TargetSiteUrl, targetSiteUrl); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Pipeline.Tags)", len(pipeline.Tags), lenTags); err != nil {
		return err
	}
	if lenTags > 0 {
		if err := CheckFieldEqualAndSet("Pipeline.Tags[0]", pipeline.Tags[0], tags[0]); err != nil {
			return err
		}
	}
	if pipeline.Project == nil {
		return errors.New("Pipeline.Project must be set")
	}
	if err := CheckProject(pipeline.Project, project.Name, project.DisplayName, true, false, false, false, "", buddy.ProjectAccessPrivate, false); err != nil {
		return err
	}
	if pipeline.Creator == nil {
		return errors.New("Pipeline.Creator must be set")
	}
	if err := CheckMember(pipeline.Creator, "", "", false, 0, true, true, 0, ""); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.DefinitionSource", pipeline.DefinitionSource, definitionSource); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.RemotePath", pipeline.RemotePath, remotePath); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.RemoteBranch", pipeline.RemoteBranch, remoteBranch); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.RemoteProjectName", pipeline.RemoteProjectName, remoteProjectName); err != nil {
		return err
	}
	if err := CheckIntFieldEqual("len(Pipeline.RemoteParameters)", len(pipeline.RemoteParameters), lenRemoteParameters); err != nil {
		return err
	}
	if lenRemoteParameters > 0 {
		if err := CheckFieldEqualAndSet("Pipeline.RemoteParameters[0].Key", pipeline.RemoteParameters[0].Key, remoteParameters[0].Key); err != nil {
			return err
		}
		if err := CheckFieldEqualAndSet("Pipeline.RemoteParameters[0].Value", pipeline.RemoteParameters[0].Value, remoteParameters[0].Value); err != nil {
			return err
		}
	}
	if err := CheckBoolFieldEqual("Pipeline.Disabled", pipeline.Disabled, disabled); err != nil {
		return err
	}
	if err := CheckFieldEqual("Pipeline.DisabledReason", pipeline.DisabledReason, disabledReason); err != nil {
		return err
	}
	return nil
}

func GenerateCertificate() (error, string) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Company, INC."},
			Country:      []string{"US"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certPrivKey, err := rsa.GenerateKey(crand.Reader, 4096)
	if err != nil {
		return err, ""
	}
	certBytes, err := x509.CreateCertificate(crand.Reader, cert, cert, &certPrivKey.PublicKey, certPrivKey)
	if err != nil {
		return err, ""
	}
	certPEM := new(bytes.Buffer)
	err = pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	if err != nil {
		return err, ""
	}
	return nil, certPEM.String()
}

func GenerateRsaKeyPair() (error, string, string) {
	privateKey, err := rsa.GenerateKey(crand.Reader, 4096)
	if err != nil {
		return err, "", ""
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyBytesEncoded := pem.EncodeToMemory(privateKeyBlock)
	sshPublicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err, "", ""
	}
	sshPublicKeyBytes := ssh.MarshalAuthorizedKey(sshPublicKey)
	return nil, string(sshPublicKeyBytes), string(privateKeyBytesEncoded)
}

func CheckIntegration(integration *buddy.Integration, expected *buddy.Integration, ops *buddy.IntegrationOps) error {
	name := expected.Name
	typ := expected.Type
	scope := expected.Scope
	projectName := expected.ProjectName
	groupId := expected.GroupId
	hashId := expected.HashId
	authType := expected.AuthType
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
		if ops.AuthType != nil {
			authType = *ops.AuthType
		}
	}
	if scope != buddy.IntegrationScopeProject && scope != buddy.IntegrationScopeGroupInProject && scope != buddy.IntegrationScopeAdminInProject && scope != buddy.IntegrationScopePrivateInProject {
		projectName = ""
	}
	if scope != buddy.IntegrationScopeGroup && scope != buddy.IntegrationScopeGroupInProject {
		groupId = 0
	}
	if authType != "" {
		if err := CheckFieldEqual("Integration.AuthType", integration.AuthType, authType); err != nil {
			return err
		}
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
