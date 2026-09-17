package buddy

import (
	"net/http"
)

const (
	VariableTypeVar    = "VAR"
	VariableTypeSshKey = "SSH_KEY"

	VariableSshKeyFilePlaceContainer = "CONTAINER"
	VariableSshKeyFilePlaceNone      = "NONE"

	VariableAccessLevelUseOnly = "USE_ONLY"
	VariableAccessLevelDenied  = "DENIED"
)

type VariableService struct {
	client *Client
}

type Variable struct {
	Url                  string                     `json:"url"`
	HtmlUrl              string                     `json:"html_url"`
	Id                   int                        `json:"id"`
	Key                  string                     `json:"key"`
	Value                string                     `json:"value"`
	Type                 string                     `json:"type"`
	Encrypted            bool                       `json:"encrypted"`
	Settable             bool                       `json:"settable"`
	RunOnlySettable      bool                       `json:"run_only_settable"`
	Disabled             bool                       `json:"disabled"`
	PipelinesAccessLevel string                     `json:"pipelines_access_level"`
	SandboxesAccessLevel string                     `json:"sandboxes_access_level"`
	AllowedPipelines     []*VariableAllowedPipeline `json:"allowed_pipelines"`
	AllowedSandboxes     []*VariableAllowedSandbox  `json:"allowed_sandboxes"`
	Note                 string                     `json:"note"`
	AgentNote            string                     `json:"agent_note"`
	FilePath             string                     `json:"file_path"`
	FileChmod            string                     `json:"file_chmod"`
	FilePlace            string                     `json:"file_place"`
	PublicValue          string                     `json:"public_value"`
	KeyFingerprint       string                     `json:"key_fingerprint"`
	Checksum             string                     `json:"checksum"`
	Project              *VariableProject           `json:"project"`
	Pipeline             *VariablePipeline          `json:"pipeline"`
	Action               *VariableAction            `json:"action"`
	Environment          *VariableEnvironment       `json:"environment"`
	Sandbox              *VariableSandbox           `json:"sandbox"`
}

type Variables struct {
	Url       string      `json:"url"`
	HtmlUrl   string      `json:"html_url"`
	Variables []*Variable `json:"variables"`
}

// VariableOps is used both to create and to update a variable. The scope fields
// (Project, Pipeline, Action, Environment, Sandbox) and Type are immutable once the
// variable exists: passing any of them to VariableService.Update with a value other
// than the current one fails with 400. Leave them nil on update, and delete and
// recreate the variable to move it between scopes.
//
// RunOnlySettable requires Settable to be true, otherwise both Create and Update fail
// with 400.
//
// PipelinesAccessLevel, SandboxesAccessLevel, AllowedPipelines and AllowedSandboxes can
// only be set on workspace- and project-scoped variables, setting them on any other scope
// fails with 400. PipelinesAccessLevel defaults to VariableAccessLevelUseOnly and
// SandboxesAccessLevel to VariableAccessLevelDenied.
//
// AllowedPipelines and AllowedSandboxes are exceptions to those defaults, so every rule
// needs the opposite AccessLevel of the default it overrides. A rule either covers a whole
// pipeline or a single Action, and rules for one pipeline cannot mix the two forms nor
// carry different access levels. Pointing at an empty slice removes every rule of that
// kind; the default level itself cannot be changed while rules still exist.
type VariableOps struct {
	Key                  *string                     `json:"key,omitempty"`
	Value                *string                     `json:"value,omitempty"`
	Type                 *string                     `json:"type,omitempty"`
	Note                 *string                     `json:"note,omitempty"`
	AgentNote            *string                     `json:"agent_note,omitempty"`
	Settable             *bool                       `json:"settable,omitempty"`
	RunOnlySettable      *bool                       `json:"run_only_settable,omitempty"`
	Disabled             *bool                       `json:"disabled,omitempty"`
	PipelinesAccessLevel *string                     `json:"pipelines_access_level,omitempty"`
	SandboxesAccessLevel *string                     `json:"sandboxes_access_level,omitempty"`
	AllowedPipelines     *[]*VariableAllowedPipeline `json:"allowed_pipelines,omitempty"`
	AllowedSandboxes     *[]*VariableAllowedSandbox  `json:"allowed_sandboxes,omitempty"`
	Encrypted            *bool                       `json:"encrypted,omitempty"`
	Project              *VariableProject            `json:"project,omitempty"`
	Pipeline             *VariablePipeline           `json:"pipeline,omitempty"`
	Action               *VariableAction             `json:"action,omitempty"`
	Environment          *VariableEnvironment        `json:"environment,omitempty"`
	Sandbox              *VariableSandbox            `json:"sandbox,omitempty"`
	FilePlace            *string                     `json:"file_place,omitempty"`
	FilePath             *string                     `json:"file_path,omitempty"`
	FileChmod            *string                     `json:"file_chmod,omitempty"`
}

type VariableGetListQuery struct {
	ProjectName   string `url:"projectName,omitempty"`
	PipelineId    int    `url:"pipelineId,omitempty"`
	ActionId      int    `url:"actionId,omitempty"`
	EnvironmentId string `url:"environment_id,omitempty"`
	SandboxId     string `url:"sandbox_id,omitempty"`
}

type VariableProject struct {
	Name string `json:"name"`
}

type VariablePipeline struct {
	Id int `json:"id"`
}

type VariableAction struct {
	Id int `json:"id"`
}

type VariableEnvironment struct {
	Id string `json:"id"`
}

type VariableSandbox struct {
	Id string `json:"id"`
}

type VariableAllowedPipeline struct {
	Project     string `json:"project"`
	Pipeline    string `json:"pipeline"`
	Action      string `json:"action,omitempty"`
	AccessLevel string `json:"access_level"`
}

type VariableAllowedSandbox struct {
	Project     string `json:"project"`
	Sandbox     string `json:"sandbox"`
	AccessLevel string `json:"access_level"`
}

func (s *VariableService) Create(workspaceDomain string, ops *VariableOps) (*Variable, *http.Response, error) {
	var v *Variable
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/variables", workspaceDomain), &ops, nil, &v)
	return v, resp, err
}

func (s *VariableService) Delete(workspaceDomain string, variableId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/variables/%d", workspaceDomain, variableId), nil, nil)
}

func (s *VariableService) Update(workspaceDomain string, variableId int, ops *VariableOps) (*Variable, *http.Response, error) {
	var v *Variable
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/variables/%d", workspaceDomain, variableId), &ops, nil, &v)
	return v, resp, err
}

func (s *VariableService) Get(workspaceDomain string, variableId int) (*Variable, *http.Response, error) {
	var v *Variable
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/variables/%d", workspaceDomain, variableId), &v, nil)
	return v, resp, err
}

func (s *VariableService) GetList(workspaceDomain string, query *VariableGetListQuery) (*Variables, *http.Response, error) {
	var all *Variables
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/variables", workspaceDomain), &all, &query)
	return all, resp, err
}
