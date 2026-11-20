package buddy

import "net/http"

const (
	RulesetBypassAccessAllowed = "ALLOWED"
	RulesetBypassAccessDenied  = "DENIED"

	RulesetCheckSourceTypePipeline         = "PIPELINE"
	RulesetCheckSourceTypeUnitTestsSuite   = "UNIT_TESTS_SUITE"
	RulesetCheckSourceTypeVisualTestsSuite = "VISUAL_TESTS_SUITE"

	RulesetRefDefault = ":default"
)

type RulesetService struct {
	client *Client
}

type RulesetBypassEntry struct {
	Id     int    `json:"id"`
	Access string `json:"access"`
}

type RulesetBypass struct {
	Admins *bool                 `json:"admins,omitempty"`
	Others string                `json:"others,omitempty"`
	Users  []*RulesetBypassEntry `json:"users,omitempty"`
	Groups []*RulesetBypassEntry `json:"groups,omitempty"`
}

type RulesetRequiredChecks struct {
	Pipelines         []string `json:"pipelines,omitempty"`
	UnitTestsSuites   []string `json:"unit_tests_suites,omitempty"`
	VisualTestsSuites []string `json:"visual_tests_suites,omitempty"`
	RequireUpToDate   bool     `json:"require_up_to_date"`
}

type RulesetRules struct {
	RequiredChecks     *RulesetRequiredChecks `json:"required_checks,omitempty"`
	RequirePullRequest bool                   `json:"require_pull_request"`
	BlockForcePush     bool                   `json:"block_force_push"`
	RestrictDeletion   bool                   `json:"restrict_deletion"`
	RestrictCreation   bool                   `json:"restrict_creation"`
	RestrictUpdate     bool                   `json:"restrict_update"`
}

type RulesetSource struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type Ruleset struct {
	Url            string           `json:"url"`
	HtmlUrl        string           `json:"html_url"`
	Id             string           `json:"id"`
	Identifier     string           `json:"identifier"`
	Name           string           `json:"name"`
	Disabled       bool             `json:"disabled"`
	Refs           []string         `json:"refs"`
	Bypass         *RulesetBypass   `json:"bypass"`
	Rules          *RulesetRules    `json:"rules"`
	Note           string           `json:"note"`
	AgentNote      string           `json:"agent_note"`
	MissingSources []*RulesetSource `json:"missing_sources"`
}

type Rulesets struct {
	Url      string     `json:"url"`
	HtmlUrl  string     `json:"html_url"`
	Rulesets []*Ruleset `json:"rulesets"`
}

// RulesetOps - Update changes only the fields that are set. Refs, Bypass and Rules replace the current value as a
// whole: a field left out inside them takes its empty value, not the stored one. Name, Note and AgentNote set to ""
// clear the value
type RulesetOps struct {
	Identifier *string        `json:"identifier,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Disabled   *bool          `json:"disabled,omitempty"`
	Refs       *[]string      `json:"refs,omitempty"`
	Bypass     *RulesetBypass `json:"bypass,omitempty"`
	Rules      *RulesetRules  `json:"rules,omitempty"`
	Note       *string        `json:"note,omitempty"`
	AgentNote  *string        `json:"agent_note,omitempty"`
}

func (s *RulesetService) Create(workspaceDomain string, projectName string, ops *RulesetOps) (*Ruleset, *http.Response, error) {
	var r *Ruleset
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/projects/%s/rulesets", workspaceDomain, projectName), &ops, nil, &r)
	return r, resp, err
}

func (s *RulesetService) Get(workspaceDomain string, projectName string, rulesetId string) (*Ruleset, *http.Response, error) {
	var r *Ruleset
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/rulesets/%s", workspaceDomain, projectName, rulesetId), &r, nil)
	return r, resp, err
}

func (s *RulesetService) GetList(workspaceDomain string, projectName string) (*Rulesets, *http.Response, error) {
	var r *Rulesets
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/projects/%s/rulesets", workspaceDomain, projectName), &r, nil)
	return r, resp, err
}

func (s *RulesetService) Update(workspaceDomain string, projectName string, rulesetId string, ops *RulesetOps) (*Ruleset, *http.Response, error) {
	var r *Ruleset
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/projects/%s/rulesets/%s", workspaceDomain, projectName, rulesetId), &ops, nil, &r)
	return r, resp, err
}

func (s *RulesetService) Delete(workspaceDomain string, projectName string, rulesetId string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/projects/%s/rulesets/%s", workspaceDomain, projectName, rulesetId), nil, nil)
}
