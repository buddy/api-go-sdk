package buddy

import (
	"errors"
	"net/http"
	"time"
)

const (
	SandboxOsUbuntu2204 = "ubuntu:22.04"
	SandboxOsUbuntu2404 = "ubuntu:24.04"

	SandboxResource1X2    = "1x2"
	SandboxResource2X4    = "2x4"
	SandboxResource3X6    = "3x6"
	SandboxResource4X8    = "4x8"
	SandboxResource5X10   = "5x10"
	SandboxResource6X12   = "6x12"
	SandboxResource7X14   = "7x14"
	SandboxResource8X16   = "8x16"
	SandboxResource9X18   = "9x18"
	SandboxResource10X20  = "10x20"
	SandboxResource11X22  = "11x22"
	SandboxResource12X24  = "12x24"
	SandboxResourceCustom = "CUSTOM"

	SandboxEndpointTypeHttp = "HTTP"
	SandboxEndpointTypeTls  = "TLS"
	SandboxEndpointTypeTcp  = "TCP"
	SandboxEndpointTypeSsh  = "SSH"

	SandboxEndpointRegionEu = "EU"
	SandboxEndpointRegionUs = "US"
	SandboxEndpointRegionAs = "AS"

	SandboxScopeProject     = "PROJECT"
	SandboxScopeEnvironment = "ENVIRONMENT"
	SandboxScopeWorkspace   = "WORKSPACE"

	SandboxFetchTypeProjectRepo = "PROJECT_REPO"
	SandboxFetchTypePublicRepo  = "PUBLIC_REPO"
	SandboxFetchTypeArtifact    = "ARTIFACT"

	SandboxFetchStatusFetching = "FETCHING"
	SandboxFetchStatusFetched  = "FETCHED"
	SandboxFetchStatusError    = "ERROR"
	SandboxFetchStatusStale    = "STALE"

	SandboxPermissionDefault  = "DEFAULT"
	SandboxPermissionManage   = "READ_WRITE"
	SandboxPermissionReadOnly = "READ_ONLY"
	SandboxPermissionUseOnly  = "USE_ONLY"
	SandboxPermissionDenied   = "DENIED"

	SandboxPipelineAccessLevelUseOnly = "USE_ONLY"
	SandboxPipelineAccessLevelDenied  = "DENIED"

	SandboxStatusStarting  = "STARTING"
	SandboxStatusStopping  = "STOPPING"
	SandboxStatusFailed    = "FAILED"
	SandboxStatusRunning   = "RUNNING"
	SandboxStatusStopped   = "STOPPED"
	SandboxStatusRestoring = "RESTORING"

	SandboxSetupStatusInProgress = "INPROGRESS"
	SandboxSetupStatusSuccess    = "SUCCESS"
	SandboxSetupStatusFailed     = "FAILED"
	SandboxSetupStatusStale      = "STALE"

	SandboxAppStatusNone    = "NONE"
	SandboxAppStatusEnded   = "ENDED"
	SandboxAppStatusRunning = "RUNNING"
	SandboxAppStatusFailed  = "FAILED"

	SandboxEndpointTlsTerminateAtRegion = "REGION"
	SandboxEndpointTlsTerminateAtAgent  = "AGENT"
	SandboxEndpointTlsTerminateAtTarget = "TARGET"

	SandboxEndpointHttpAuthTypeNone  = "NONE"
	SandboxEndpointHttpAuthTypeBasic = "BASIC"
	SandboxEndpointHttpAuthTypeBuddy = "BUDDY"
)

type SandboxService struct {
	client *Client
}

type SandboxEnvironmentOps struct {
	Id *string `json:"id,omitempty"`
}

type SandboxEndpointHttp struct {
	VerifyCertificate   *bool              `json:"verify_certificate,omitempty"`
	Compression         *bool              `json:"compression,omitempty"`
	Http2               *bool              `json:"http2,omitempty"`
	LogRequests         *bool              `json:"log_requests,omitempty"`
	RewriteHostHeader   *string            `json:"rewrite_host_header,omitempty"`
	WhitelistUserAgents *[]string          `json:"whitelist_user_agents,omitempty"`
	RequestHeaders      *map[string]string `json:"request_headers,omitempty"`
	ResponseHeaders     *map[string]string `json:"response_headers,omitempty"`
	Login               *string            `json:"login,omitempty"`
	Password            *string            `json:"password,omitempty"`
	AuthType            *string            `json:"auth_type,omitempty"`
	CircuitBreaker      *int               `json:"circuit_breaker,omitempty"`
	TlsCa               *string            `json:"tls_ca,omitempty"`
	ServePath           *string            `json:"serve_path,omitempty"`
}

type SandboxEndpointTls struct {
	TerminateAt   *string `json:"terminate_at,omitempty"`
	PrivateKey    *string `json:"private_key,omitempty"`
	Certificate   *string `json:"certificate,omitempty"`
	CaCertificate *string `json:"ca_certificate,omitempty"`
}

type SandboxEndpoint struct {
	Name          *string              `json:"name,omitempty"`
	Endpoint      *string              `json:"endpoint,omitempty"`
	Type          *string              `json:"type,omitempty"`
	Region        *string              `json:"region,omitempty"`
	Whitelist     *[]string            `json:"whitelist,omitempty"`
	Timeout       *int                 `json:"timeout,omitempty"`
	EndpointUrl   *string              `json:"endpoint_url,omitempty"`
	Http          *SandboxEndpointHttp `json:"http,omitempty"`
	Tls           *SandboxEndpointTls  `json:"tls,omitempty"`
	Active        *bool                `json:"active,omitempty"`
	TargetLatency *int                 `json:"target_latency,omitempty"`
}

type SandboxApp struct {
	Id         string `json:"id"`
	Identifier string `json:"identifier"`
	Command    string `json:"command"`
	Path       string `json:"path"`
	AppStatus  string `json:"app_status"`
}

type SandboxAppOps struct {
	Identifier *string `json:"identifier,omitempty"`
	Command    *string `json:"command,omitempty"`
	Path       *string `json:"path,omitempty"`
}

type SandboxFetch struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	Repository      string `json:"repository"`
	Ref             string `json:"ref"`
	Path            string `json:"path"`
	SourcePath      string `json:"source_path"`
	Artifact        string `json:"artifact"`
	BuildCommand    string `json:"build_command"`
	CurrentRevision string `json:"current_revision"`
	Behind          *int   `json:"behind"`
	Outdated        *bool  `json:"outdated"`
	FetchedDate     string `json:"fetched_date"`
	Status          string `json:"status"`
	ErrorMessage    string `json:"error_message"`
}

type SandboxFetches struct {
	Url     string          `json:"url"`
	HtmlUrl string          `json:"html_url"`
	Fetches []*SandboxFetch `json:"fetches"`
}

type SandboxFetchOps struct {
	Type         *string `json:"type,omitempty"`
	Repository   *string `json:"repository,omitempty"`
	Ref          *string `json:"ref,omitempty"`
	Path         *string `json:"path,omitempty"`
	SourcePath   *string `json:"source_path,omitempty"`
	Artifact     *string `json:"artifact,omitempty"`
	BuildCommand *string `json:"build_command,omitempty"`
}

type Sandboxes struct {
	Url       string     `json:"url"`
	HtmlUrl   string     `json:"html_url"`
	Sandboxes []*Sandbox `json:"sandboxes"`
}

type SandboxResourcePermission struct {
	Id          int    `json:"id"`
	AccessLevel string `json:"access_level"`
}

type SandboxPermissions struct {
	Others string                       `json:"others"`
	Users  []*SandboxResourcePermission `json:"users"`
	Groups []*SandboxResourcePermission `json:"groups"`
}

type SandboxAllowedPipeline struct {
	Project     string `json:"project"`
	Pipeline    string `json:"pipeline"`
	AccessLevel string `json:"access_level"`
}

type Sandbox struct {
	Url                  string                    `json:"url"`
	HtmlUrl              string                    `json:"html_url"`
	Id                   string                    `json:"id"`
	Identifier           string                    `json:"identifier"`
	Name                 string                    `json:"name"`
	Note                 string                    `json:"note"`
	AgentNote            string                    `json:"agent_note"`
	Project              *Project                  `json:"project"`
	Status               string                    `json:"status"`
	Os                   string                    `json:"os"`
	Resources            string                    `json:"resources"`
	FirstBootCommands    string                    `json:"first_boot_commands"`
	BootLogs             []string                  `json:"boot_logs"`
	AppDir               string                    `json:"app_dir"`
	Timeout              int                       `json:"timeout"`
	Apps                 []*SandboxApp             `json:"apps"`
	Fetch                []*SandboxFetch           `json:"fetch"`
	SetupStatus          string                    `json:"setup_status"`
	Tags                 []string                  `json:"tags"`
	Endpoints            []*SandboxEndpoint        `json:"endpoints"`
	Variables            []*Variable               `json:"variables"`
	Environment          *Environment              `json:"environment"`
	Permissions          *SandboxPermissions       `json:"permissions"`
	PipelinesAccessLevel string                    `json:"pipelines_access_level"`
	AllowedPipelines     []*SandboxAllowedPipeline `json:"allowed_pipelines"`
	Scope                string                    `json:"scope"`
	SshHost              string                    `json:"ssh_host"`
	SshPort              int                       `json:"ssh_port"`
	CreateDate           string                    `json:"create_date"`
}

type SandboxStatusOps struct {
}

// SandboxYaml holds the sandbox configuration as base64-encoded YAML
type SandboxYaml struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
	Yaml    string `json:"yaml"`
}

type SandboxYamlOps struct {
	// Yaml is the base64-encoded YAML, in the format returned by SandboxService.GetYaml
	Yaml *string `json:"yaml"`
}

type SandboxOps struct {
	Name                 *string                    `json:"name,omitempty"`
	Note                 *string                    `json:"note,omitempty"`
	AgentNote            *string                    `json:"agent_note,omitempty"`
	Identifier           *string                    `json:"identifier,omitempty"`
	Os                   *string                    `json:"os,omitempty"`
	Resources            *string                    `json:"resources,omitempty"`
	FirstBootCommands    *string                    `json:"first_boot_commands,omitempty"`
	AppDir               *string                    `json:"app_dir,omitempty"`
	Apps                 *[]*SandboxAppOps          `json:"apps,omitempty"`
	Fetch                *[]*SandboxFetchOps        `json:"fetch,omitempty"`
	Tags                 *[]string                  `json:"tags,omitempty"`
	Timeout              *int                       `json:"timeout,omitempty"`
	Environment          *SandboxEnvironmentOps     `json:"environment,omitempty"`
	Variables            *[]*VariableOps            `json:"variables,omitempty"`
	Endpoints            *[]*SandboxEndpoint        `json:"endpoints,omitempty"`
	Permissions          *SandboxPermissions        `json:"permissions"`
	PipelinesAccessLevel *string                    `json:"pipelines_access_level,omitempty"`
	AllowedPipelines     *[]*SandboxAllowedPipeline `json:"allowed_pipelines,omitempty"`
}

func (s *SandboxService) Create(workspaceDomain string, query *Query, ops *SandboxOps) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes", workspaceDomain), &ops, query, &sb)
	return sb, resp, err
}

func (s *SandboxService) Update(workspaceDomain string, sandboxId string, ops *SandboxOps) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s", workspaceDomain, sandboxId), &ops, nil, &sb)
	return sb, resp, err
}

func (s *SandboxService) Delete(workspaceDomain string, sandboxId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s", workspaceDomain, sandboxId), nil, nil)
}

func (s *SandboxService) Get(workspaceDomain string, sandboxId string) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s", workspaceDomain, sandboxId), &sb, nil)
	return sb, resp, err
}

func (s *SandboxService) GetList(workspaceDomain string, query *Query) (*Sandboxes, *http.Response, error) {
	var l *Sandboxes
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes", workspaceDomain), &l, query)
	return l, resp, err
}

func (s *SandboxService) GetYaml(workspaceDomain string, sandboxId string) (*SandboxYaml, *http.Response, error) {
	var y *SandboxYaml
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/yaml", workspaceDomain, sandboxId), &y, nil)
	return y, resp, err
}

// UpdateYaml replaces the whole sandbox definition with the YAML, a key omitted from the document clears its value
func (s *SandboxService) UpdateYaml(workspaceDomain string, sandboxId string, ops *SandboxYamlOps) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/yaml", workspaceDomain, sandboxId), &ops, nil, &sb)
	return sb, resp, err
}

func (s *SandboxService) Stop(workspaceDomain string, sandboxId string) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/stop", workspaceDomain, sandboxId), &SandboxStatusOps{}, nil, &sb)
	return sb, resp, err
}

func (s *SandboxService) Start(workspaceDomain string, sandboxId string) (*Sandbox, *http.Response, error) {
	var sb *Sandbox
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/start", workspaceDomain, sandboxId), &SandboxStatusOps{}, nil, &sb)
	return sb, resp, err
}

func (s *SandboxService) GetFetches(workspaceDomain string, sandboxId string) (*SandboxFetches, *http.Response, error) {
	var l *SandboxFetches
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/fetches", workspaceDomain, sandboxId), &l, nil)
	return l, resp, err
}

func (s *SandboxService) GetFetch(workspaceDomain string, sandboxId string, fetchId string) (*SandboxFetch, *http.Response, error) {
	var f *SandboxFetch
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/fetches/%s", workspaceDomain, sandboxId, fetchId), &f, nil)
	return f, resp, err
}

func (s *SandboxService) CreateFetch(workspaceDomain string, sandboxId string, ops *SandboxFetchOps) (*SandboxFetch, *http.Response, error) {
	var f *SandboxFetch
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/fetches", workspaceDomain, sandboxId), &ops, nil, &f)
	return f, resp, err
}

func (s *SandboxService) DeleteFetch(workspaceDomain string, sandboxId string, fetchId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/fetches/%s", workspaceDomain, sandboxId, fetchId), nil, nil)
}

func (s *SandboxService) Refetch(workspaceDomain string, sandboxId string, fetchId string) (*SandboxFetch, *http.Response, error) {
	var f *SandboxFetch
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/fetches/%s/refetch", workspaceDomain, sandboxId, fetchId), &SandboxStatusOps{}, nil, &f)
	return f, resp, err
}

func (s *SandboxService) WaitForFetchStatuses(workspaceDomain string, sandboxId string, fetchId string, timeout int, statuses []string) (*SandboxFetch, error) {
	start := time.Now()
	for {
		f, _, err := s.GetFetch(workspaceDomain, sandboxId, fetchId)
		if err != nil {
			return nil, err
		}
		for _, status := range statuses {
			if f != nil && f.Status == status {
				return f, nil
			}
		}
		if int(time.Since(start).Seconds()) > timeout {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return nil, errors.New("timeout waiting for sandbox fetch")
}

func (s *SandboxService) waitUntilCondition(workspaceDomain string, sandboxId string, timeout int, condition func(sb *Sandbox) bool) (*Sandbox, error) {
	start := time.Now()
	for {
		sb, _, err := s.Get(workspaceDomain, sandboxId)
		if err != nil {
			return nil, err
		}
		if condition(sb) {
			return sb, nil
		}
		if int(time.Since(start).Seconds()) > timeout {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return nil, errors.New("timeout waiting for sandbox")
}

func (s *SandboxService) WaitForStatuses(workspaceDomain string, sandboxId string, timeout int, statuses []string) (*Sandbox, error) {
	return s.waitUntilCondition(workspaceDomain, sandboxId, timeout, func(sb *Sandbox) bool {
		if sb == nil {
			return false
		}
		for _, status := range statuses {
			if sb.Status == status {
				return true
			}
		}
		return false
	})
}

func (s *SandboxService) WaitForSetupStatuses(workspaceDomain string, sandboxId string, timeout int, statuses []string) (*Sandbox, error) {
	return s.waitUntilCondition(workspaceDomain, sandboxId, timeout, func(sb *Sandbox) bool {
		if sb == nil {
			return false
		}
		for _, status := range statuses {
			if sb.SetupStatus == status {
				return true
			}
		}
		return false
	})
}

func (s *SandboxService) WaitForAppStatuses(workspaceDomain string, sandboxId string, timeout int, statuses []string) (*Sandbox, error) {
	return s.waitUntilCondition(workspaceDomain, sandboxId, timeout, func(sb *Sandbox) bool {
		if sb == nil {
			return false
		}
		for _, app := range sb.Apps {
			hasStatus := false
			for _, status := range statuses {
				if app.AppStatus == status {
					hasStatus = true
					break
				}
			}
			if !hasStatus {
				return false
			}
		}
		return true
	})
}
