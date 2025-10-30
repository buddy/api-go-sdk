package buddy

import (
	"net/http"
)

const (
	PipelineCpuX64 = "X64"
	PipelineCpuArm = "ARM"

	PipelineEventTypePush        = "PUSH"
	PipelineEventTypeSchedule    = "SCHEDULE"
	PipelineEventTypeCreateRef   = "CREATE_REF"
	PipelineEventTypeDeleteRef   = "DELETE_REF"
	PipelineEventTypePullRequest = "PULL_REQUEST"
	PipelineEventTypeWebhook     = "WEBHOOK"
	PipelineEventTypeEmail       = "EMAIL"

	PipelinePullRequestEventAssigned             = "assigned"
	PipelinePullRequestEventUnassigned           = "unassigned"
	PipelinePullRequestEventLabeled              = "labeled"
	PipelinePullRequestEventUnlabeled            = "unlabeled"
	PipelinePullRequestEventOpened               = "opened"
	PipelinePullRequestEventEdited               = "edited"
	PipelinePullRequestEventClosed               = "closed"
	PipelinePullRequestEventReopened             = "reopened"
	PipelinePullRequestEventSynchronize          = "synchronize"
	PipelinePullRequestEventConvertedToDraft     = "converted_to_draft"
	PipelinePullRequestEventLocked               = "locked"
	PipelinePullRequestEventUnlocked             = "unlocked"
	PipelinePullRequestEventEnqueued             = "enqueued"
	PipelinePullRequestEventDequeued             = "dequeued"
	PipelinePullRequestEventMilestoned           = "milestoned"
	PipelinePullRequestEventDemilestoned         = "demilestoned"
	PipelinePullRequestEventReadyForReview       = "ready_for_review"
	PipelinePullRequestEventReviewRequested      = "review_requested"
	PipelinePullRequestEventReviewRequestRemoved = "review_request_removed"
	PipelinePullRequestEventAutoMergeEnabled     = "auto_merge_enabled"
	PipelinePullRequestEventAutoMergeDisabled    = "auto_merge_disabled"

	PipelineTriggerConditionOnChange                   = "ON_CHANGE"
	PipelineTriggerConditionOnChangeAtPath             = "ON_CHANGE_AT_PATH"
	PipelineTriggerConditionVarIs                      = "VAR_IS"
	PipelineTriggerConditionVarIsNot                   = "VAR_IS_NOT"
	PipelineTriggerConditionVarContains                = "VAR_CONTAINS"
	PipelineTriggerConditionVarNotContains             = "VAR_NOT_CONTAINS"
	PipelineTriggerConditionDateTime                   = "DATETIME"
	PipelineTriggerConditionSuccessPipeline            = "SUCCESS_PIPELINE"
	PipelineTriggerConditionTriggeringUserIs           = "TRIGGERING_USER_IS"
	PipelineTriggerConditionTriggeringUserIsNot        = "TRIGGERING_USER_IS_NOT"
	PipelineTriggerConditionTriggeringUserIsInGroup    = "TRIGGERING_USER_IS_IN_GROUP"
	PipelineTriggerConditionTriggeringUserIsNotInGroup = "TRIGGERING_USER_IS_NOT_IN_GROUP"

	PipelinePriorityHigh   = "HIGH"
	PipelinePriorityNormal = "NORMAL"
	PipelinePriorityLow    = "LOW"

	PipelineDefinitionSourceLocal  = "LOCAL"
	PipelineDefinitionSourceRemote = "REMOTE"

	PipelinePermissionDefault   = "DEFAULT"
	PipelinePermissionDenied    = "DENIED"
	PipelinePermissionReadOnly  = "READ_ONLY"
	PipelinePermissionRunOnly   = "RUN_ONLY"
	PipelinePermissionReadWrite = "READ_WRITE"

	PipelineGitConfigRefNone    = "NONE"
	PipelineGitConfigRefDynamic = "DYNAMIC"
	PipelineGitConfigRefFixed   = "FIXED"

	PipelineGitChangeSetBaseLatestRun            = "LATEST_RUN"
	PipelineGitChangeSetBaseLatestRunMatchingRef = "LATEST_RUN_MATCHING_REF"
	PipelineGitChangeSetBasePullRequest          = "PULL_REQUEST"

	PipelineFilesystemChangeSetBaseDateModified = "DATE_MODIFIED"
	PipelineFilesystemChangeSetBaseContents     = "CONTENTS"
)

type Pipeline struct {
	Url                       string                      `json:"url"`
	HtmlUrl                   string                      `json:"html_url"`
	Id                        int                         `json:"id"`
	Identifier                string                      `json:"identifier"`
	Name                      string                      `json:"name"`
	Refs                      []string                    `json:"refs"`
	Events                    []*PipelineEvent            `json:"events"`
	TriggerConditions         []*PipelineTriggerCondition `json:"trigger_conditions"`
	ExecutionMessageTemplate  string                      `json:"execution_message_template"`
	LastExecutionStatus       string                      `json:"last_execution_status"`
	LastExecutionRevision     string                      `json:"last_execution_revision"`
	CreateDate                string                      `json:"create_date"`
	Priority                  string                      `json:"priority"`
	AlwaysFromScratch         bool                        `json:"always_from_scratch"`
	FailOnPrepareEnvWarning   bool                        `json:"fail_on_prepare_env_warning"`
	FetchAllRefs              bool                        `json:"fetch_all_refs"`
	AutoClearCache            bool                        `json:"auto_clear_cache"`
	NoSkipToMostRecent        bool                        `json:"no_skip_to_most_recent"`
	DoNotCreateCommitStatus   bool                        `json:"do_not_create_commit_status"`
	IgnoreFailOnProjectStatus bool                        `json:"ignore_fail_on_project_status"`
	CloneDepth                int                         `json:"clone_depth"`
	Cpu                       string                      `json:"cpu"`
	Paused                    bool                        `json:"paused"`
	Worker                    string                      `json:"worker"`
	TargetSiteUrl             string                      `json:"target_site_url"`
	Tags                      []string                    `json:"tags"`
	Project                   *Project                    `json:"project"`
	Creator                   *Member                     `json:"creator"`
	ConcurrentPipelineRuns    bool                        `json:"concurrent_pipeline_runs"`
	GitConfigRef              string                      `json:"git_config_ref"`
	GitConfig                 *PipelineGitConfig          `json:"git_config"`
	DefinitionSource          string                      `json:"definition_source"`
	RemotePath                string                      `json:"remote_path"`
	RemoteBranch              string                      `json:"remote_branch"`
	RemoteRef                 string                      `json:"remote_ref"`
	RemoteProjectName         string                      `json:"remote_project_name"`
	RemoteParameters          []*PipelineRemoteParameter  `json:"remote_parameters"`
	ManageVariablesByYaml     bool                        `json:"manage_variables_by_yaml"`
	ManagePermissionsByYaml   bool                        `json:"manage_permissions_by_yaml"`
	DescriptionRequired       bool                        `json:"description_required"`
	FilesystemChangesetBase   string                      `json:"filesystem_changeset_base"`
	GitChangesetBase          string                      `json:"git_changeset_base"`
	Disabled                  bool                        `json:"disabled"`
	DisabledReason            string                      `json:"disabled_reason"`
	Permissions               *PipelinePermissions        `json:"permissions"`
	PauseOnRepeatedFailures   int                         `json:"pause_on_repeated_failures"`
	Loop                      []string                    `json:"loop"`
}

type Pipelines struct {
	Url       string      `json:"url"`
	HtmlUrl   string      `json:"html_url"`
	Pipelines []*Pipeline `json:"pipelines"`
}

type PipelineRemoteParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PipelineEvent struct {
	Type      string   `json:"type"`
	Refs      []string `json:"refs"`
	Events    []string `json:"events"`
	Branches  []string `json:"branches"`
	StartDate string   `json:"start_date,omitempty"`
	Delay     int      `json:"delay,omitempty"`
	Cron      string   `json:"cron,omitempty"`
	Timezone  string   `json:"timezone,omitempty"`
	Prefix    string   `json:"prefix,omitempty"`
	Whitelist []string `json:"whitelist,omitempty"`
	Totp      bool     `json:"totp"`
}

type PipelineResourcePermission struct {
	Id          int    `json:"id"`
	AccessLevel string `json:"access_level"`
}

type PipelinePermissions struct {
	Others string                        `json:"others"`
	Users  []*PipelineResourcePermission `json:"users"`
	Groups []*PipelineResourcePermission `json:"groups"`
}

type PipelineTriggerCondition struct {
	TriggerCondition      string   `json:"trigger_condition"`
	TriggerConditionPaths []string `json:"trigger_condition_paths"`
	TriggerVariableKey    string   `json:"trigger_variable_key"`
	TriggerVariableValue  string   `json:"trigger_variable_value"`
	TriggerHours          []int    `json:"trigger_hours"`
	TriggerDays           []int    `json:"trigger_days"`
	Timezone              string   `json:"timezone"`
	TriggerProjectName    string   `json:"trigger_project_name"`
	TriggerPipelineName   string   `json:"trigger_pipeline_name"`
	TriggerUser           string   `json:"trigger_user"`
	TriggerGroup          string   `json:"trigger_group"`
}

type PipelineGitConfig struct {
	Path    string `json:"path"`
	Branch  string `json:"branch"`
	Project string `json:"project"`
}

type PipelineService struct {
	client *Client
}

type PipelineOps struct {
	Name                      *string                      `json:"name,omitempty"`
	Identifier                *string                      `json:"identifier,omitempty"`
	Refs                      *[]string                    `json:"refs,omitempty"`
	Tags                      *[]string                    `json:"tags,omitempty"`
	Events                    *[]*PipelineEvent            `json:"events,omitempty"`
	TriggerConditions         *[]*PipelineTriggerCondition `json:"trigger_conditions,omitempty"`
	AlwaysFromScratch         *bool                        `json:"always_from_scratch,omitempty"`
	Priority                  *string                      `json:"priority,omitempty"`
	FailOnPrepareEnvWarning   *bool                        `json:"fail_on_prepare_env_warning,omitempty"`
	FetchAllRefs              *bool                        `json:"fetch_all_refs,omitempty"`
	AutoClearCache            *bool                        `json:"auto_clear_cache,omitempty"`
	NoSkipToMostRecent        *bool                        `json:"no_skip_to_most_recent,omitempty"`
	DoNotCreateCommitStatus   *bool                        `json:"do_not_create_commit_status,omitempty"`
	CloneDepth                *int                         `json:"clone_depth,omitempty"`
	Paused                    *bool                        `json:"paused,omitempty"`
	IgnoreFailOnProjectStatus *bool                        `json:"ignore_fail_on_project_status,omitempty"`
	ExecutionMessageTemplate  *string                      `json:"execution_message_template,omitempty"`
	Worker                    *string                      `json:"worker,omitempty"`
	Cpu                       *string                      `json:"cpu,omitempty"`
	TargetSiteUrl             *string                      `json:"target_site_url,omitempty"`
	GitConfigRef              *string                      `json:"git_config_ref,omitempty"`
	GitConfig                 *PipelineGitConfig           `json:"git_config,omitempty"`
	ConcurrentPipelineRuns    *bool                        `json:"concurrent_pipeline_runs,omitempty"`
	DefinitionSource          *string                      `json:"definition_source,omitempty"`
	RemotePath                *string                      `json:"remote_path,omitempty"`
	RemoteBranch              *string                      `json:"remote_branch,omitempty"`
	RemoteRef                 *string                      `json:"remote_ref,omitempty"`
	RemoteProjectName         *string                      `json:"remote_project_name,omitempty"`
	RemoteParameters          *[]*PipelineRemoteParameter  `json:"remote_parameters,omitempty"`
	DescriptionRequired       *bool                        `json:"description_required,omitempty"`
	FilesystemChangesetBase   *string                      `json:"filesystem_changeset_base,omitempty"`
	GitChangesetBase          *string                      `json:"git_changeset_base,omitempty"`
	ManageVariablesByYaml     *bool                        `json:"manage_variables_by_yaml"`
	ManagePermissionsByYaml   *bool                        `json:"manage_permissions_by_yaml"`
	Disabled                  *bool                        `json:"disabled,omitempty"`
	DisabledReason            *string                      `json:"disabled_reason,omitempty"`
	PauseOnRepeatedFailures   *int                         `json:"pause_on_repeated_failures,omitempty"`
	Permissions               *PipelinePermissions         `json:"permissions,omitempty"`
	Loop                      *[]string                    `json:"loop,omitempty"`
}

func (s *PipelineService) Create(workspaceDomain string, projectName string, ops *PipelineOps) (*Pipeline, *http.Response, error) {
	var p *Pipeline
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/pipelines", workspaceDomain, projectName), &ops, nil, &p)
	return p, resp, err
}

func (s *PipelineService) Delete(workspaceDomain string, projectName string, pipelineId int) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/pipelines/%d", workspaceDomain, projectName, pipelineId), nil, nil)
}

func (s *PipelineService) Update(workspaceDomain string, projectName string, pipelineId int, ops *PipelineOps) (*Pipeline, *http.Response, error) {
	var p *Pipeline
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/pipelines/%d", workspaceDomain, projectName, pipelineId), &ops, nil, &p)
	return p, resp, err
}

func (s *PipelineService) Get(workspaceDomain string, projectName string, pipelineId int) (*Pipeline, *http.Response, error) {
	var p *Pipeline
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/pipelines/%d", workspaceDomain, projectName, pipelineId), &p, nil)
	return p, resp, err
}

func (s *PipelineService) GetList(workspaceDomain string, projectName string, query *PageQuery) (*Pipelines, *http.Response, error) {
	var l *Pipelines
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/pipelines", workspaceDomain, projectName), &l, query)
	return l, resp, err
}

func (s *PipelineService) GetListAll(workspaceDomain string, projectName string) (*Pipelines, *http.Response, error) {
	var all Pipelines
	page := 1
	perPage := 30
	for {
		l, resp, err := s.GetList(workspaceDomain, projectName, &PageQuery{
			Page:    page,
			PerPage: perPage,
		})
		if err != nil {
			return nil, resp, err
		}
		if len(l.Pipelines) == 0 {
			break
		}
		all.Url = l.Url
		all.HtmlUrl = l.HtmlUrl
		all.Pipelines = append(all.Pipelines, l.Pipelines...)
		page += 1
	}
	return &all, nil, nil
}
