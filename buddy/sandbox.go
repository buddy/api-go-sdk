package buddy

import (
  "errors"
  "net/http"
  "time"
)

const (
  SandboxOsUbuntu2204 = "ubuntu:22.04"
  SandboxOsUbuntu2404 = "ubuntu:24.04"

  SandboxResource1X2   = "1x2"
  SandboxResource2X4   = "2x4"
  SandboxResource3X6   = "3x6"
  SandboxResource4X8   = "4x8"
  SandboxResource5X10  = "5x10"
  SandboxResource6X12  = "6x12"
  SandboxResource7X14  = "7x14"
  SandboxResource8X16  = "8x16"
  SandboxResource9X18  = "9x18"
  SandboxResource10X20 = "10x20"
  SandboxResource11X22 = "11x22"
  SandboxResource12X24 = "12x24"

  SandboxAppTypeCmd     = "CMD"
  SandboxAppTypeService = "SERVICE"

  SandboxEndpointTypeHttp = "HTTP"
  SandboxEndpointTypeTls  = "TLS"
  SandboxEndpointTypeTcp  = "TCP"

  SandboxEndpointRegionEu = "EU"
  SandboxEndpointRegionUs = "US"

  SandboxStatusStarting  = "STARTING"
  SandboxStatusStopping  = "STOPPING"
  SandboxStatusFailed    = "FAILED"
  SandboxStatusRunning   = "RUNNING"
  SandboxStatusStopped   = "STOPPED"
  SandboxStatusRestoring = "RESTORING"

  SandboxSetupStatusConfiguring = "CONFIGURING"
  SandboxSetupStatusDone        = "DONE"
  SandboxSetupStatusFailed      = "FAILED"

  SandboxAppStatusNone    = "NONE"
  SandboxAppStatusEnded   = "ENDED"
  SandboxAppStatusRunning = "RUNNING"
  SandboxAppStatusFailed  = "FAILED"

  SandboxEndpointTlsTerminateAtRegion = "REGION"
  SandboxEndpointTlsTerminateAtAgent  = "AGENT"
  SandboxEndpointTlsTerminateAtTarget = "TARGET"
)

type SandboxService struct {
  client *Client
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
  CircuitBreaker      *int               `json:"circuit_breaker,omitempty"`
  TlsCa               *string            `json:"tls_ca,omitempty"`
}

type SandboxEndpointTls struct {
  TerminateAt   *string `json:"terminate_at,omitempty"`
  PrivateKey    *string `json:"private_key,omitempty"`
  Certificate   *string `json:"certificate,omitempty"`
  CaCertificate *string `json:"ca_certificate,omitempty"`
}

type SandboxEndpoint struct {
  Name      *string              `json:"name,omitempty"`
  Endpoint  *string              `json:"endpoint,omitempty"`
  Type      *string              `json:"type,omitempty"`
  Region    *string              `json:"region,omitempty"`
  Whitelist *[]string            `json:"whitelist,omitempty"`
  Timeout   *int                 `json:"timeout,omitempty"`
  Http      *SandboxEndpointHttp `json:"http,omitempty"`
  Tls       *SandboxEndpointTls  `json:"tls,omitempty"`
}

type Sandboxes struct {
  Url       string     `json:"url"`
  HtmlUrl   string     `json:"html_url"`
  Sandboxes []*Sandbox `json:"sandboxes"`
}

type Sandbox struct {
  Url             string            `json:"url"`
  HtmlUrl         string            `json:"html_url"`
  Id              string            `json:"id"`
  Identifier      string            `json:"identifier"`
  Name            string            `json:"name"`
  Project         *Project          `json:"project"`
  Status          string            `json:"status"`
  Os              string            `json:"os"`
  Resources       string            `json:"resources"`
  InstallCommands string            `json:"install_commands"`
  RunCommand      string            `json:"run_command"`
  AppDir          string            `json:"app_dir"`
  AppType         string            `json:"app_type"`
  AppStatus       string            `json:"app_status"`
  SetupStatus     string            `json:"setup_status"`
  Tags            []string          `json:"tags"`
  Endpoints       []SandboxEndpoint `json:"endpoints"`
}

type SandboxOps struct {
  Name            *string            `json:"name,omitempty"`
  Identifier      *string            `json:"identifier,omitempty"`
  Os              *string            `json:"os,omitempty"`
  Resources       *string            `json:"resources,omitempty"`
  InstallCommands *string            `json:"install_commands,omitempty"`
  RunCommand      *string            `json:"run_command,omitempty"`
  AppDir          *string            `json:"app_dir,omitempty"`
  AppType         *string            `json:"app_type,omitempty"`
  Tags            *[]string          `json:"tags,omitempty"`
  Endpoints       *[]SandboxEndpoint `json:"endpoints,omitempty"`
}

func (s *SandboxService) Create(workspaceDomain string, projectName string, ops *SandboxOps) (*Sandbox, *http.Response, error) {
  var sb *Sandbox
  resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes", workspaceDomain), &ops, &Query{
    ProjectName: &projectName,
  }, &sb)
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

func (s *SandboxService) GetList(workspaceDomain string, query Query) (*Sandboxes, *http.Response, error) {
  var l *Sandboxes
  resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/sandboxes", workspaceDomain), &l, &query)
  return l, resp, err
}

func (s *SandboxService) Stop(workspaceDomain string, sandboxId string) (*Sandbox, *http.Response, error) {
  var sb *Sandbox
  resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/stop", workspaceDomain, sandboxId), nil, nil, &sb)
  return sb, resp, err
}

func (s *SandboxService) Start(workspaceDomain string, sandboxId string) (*Sandbox, *http.Response, error) {
  var sb *Sandbox
  resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/sandboxes/%s/start", workspaceDomain, sandboxId), nil, nil, &sb)
  return sb, resp, err
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
    for _, status := range statuses {
      if sb.AppStatus == status {
        return true
      }
    }
    return false
  })
}
