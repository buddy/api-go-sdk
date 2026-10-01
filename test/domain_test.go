package test

import (
	"encoding/base64"
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"strings"
	"testing"
)

func testDomainCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandDomain()
		typ := buddy.DomainTypePointed
		ops := buddy.DomainCreateOps{
			Name: &name,
			Type: &typ,
		}
		domain, _, err := client.DomainService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.Create", err))
		}
		err = CheckDomain(domain, name, false)
		if err != nil {
			t.Fatal(err)
		}
		*out = *domain
	}
}

func testDomainList(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		domains, _, err := client.DomainService.GetList(workspace.Domain, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetList", err))
		}
		err = CheckDomains(domains, domain)
		if err != nil {
			t.Fatal(err)
		}
		pointed, _, err := client.DomainService.GetList(workspace.Domain, &buddy.DomainGetListQuery{Type: buddy.DomainTypePointed})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetList", err))
		}
		err = CheckDomains(pointed, domain)
		if err != nil {
			t.Fatal(err)
		}
		private, _, err := client.DomainService.GetList(workspace.Domain, &buddy.DomainGetListQuery{Type: buddy.DomainTypePrivate})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetList", err))
		}
		err = CheckIntFieldEqual("len(Domains)", len(private.Domains), 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainYamlUpdate(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		y, _, err := client.DomainService.GetYaml(workspace.Domain, domain.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetYaml", err))
		}
		err = CheckFieldSet("DomainYaml.Url", y.Url)
		if err != nil {
			t.Fatal(err)
		}
		raw, err := base64.StdEncoding.DecodeString(y.Yaml)
		if err != nil {
			t.Fatal(ErrorFormatted("base64.DecodeString", err))
		}
		doc := string(raw)
		if !strings.HasPrefix(doc, domain.Name+":") {
			t.Fatalf("DomainYaml.Yaml should start with %s:, got %s", domain.Name, doc)
		}
		// apex SOA and NS stay as returned, only the record list grows
		name := UniqueString()
		doc = strings.TrimRight(doc, "\n") + fmt.Sprintf("\n    %s:\n    - type: A\n      ttl: 300\n      values: 3.3.3.3\n", name)
		encoded := base64.StdEncoding.EncodeToString([]byte(doc))
		updated, _, err := client.DomainService.UpdateYaml(workspace.Domain, domain.Id, &buddy.DomainYamlOps{Yaml: &encoded})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpdateYaml", err))
		}
		err = CheckDomain(updated, domain.Name, false)
		if err != nil {
			t.Fatal(err)
		}
		r, _, err := client.DomainService.GetRecord(workspace.Domain, domain.Id, fmt.Sprintf("%s.%s", name, domain.Name), "A")
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetRecord", err))
		}
		err = CheckRecord(r, name, "", "", "A", 300, buddy.DomainRecordRoutingSimple, "3.3.3.3", "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainDelete(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.DomainService.Delete(workspace.Domain, domain.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.Delete", err))
		}
		domains, _, err := client.DomainService.GetList(workspace.Domain, nil)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetList", err))
		}
		err = CheckIntFieldEqual("len(Domains)", len(domains.Domains), 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainPrivateYamlUpsert(client *buddy.Client, workspace *buddy.Workspace) func(t *testing.T) {
	return func(t *testing.T) {
		first := fmt.Sprintf("%s.lan", UniqueString())
		second := fmt.Sprintf("%s.lan", UniqueString())
		// apex SOA and NS may be omitted for a new private domain
		doc := fmt.Sprintf("%s:\n  records:\n    www:\n    - type: A\n      values: 10.0.0.1\n%s:\n  records:\n    api:\n    - type: A\n      values: 10.0.0.2\n", first, second)
		encoded := base64.StdEncoding.EncodeToString([]byte(doc))
		domains, _, err := client.DomainService.UpsertPrivateYaml(workspace.Domain, &buddy.DomainYamlOps{Yaml: &encoded})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertPrivateYaml", err))
		}
		err = CheckIntFieldEqual("len(Domains)", len(domains.Domains), 2)
		if err != nil {
			t.Fatal(err)
		}
		for i, name := range []string{first, second} {
			err = CheckFieldEqualAndSet("Domain.Name", domains.Domains[i].Name, name)
			if err != nil {
				t.Fatal(err)
			}
			err = CheckFieldEqualAndSet("Domain.Type", domains.Domains[i].Type, buddy.DomainTypePrivate)
			if err != nil {
				t.Fatal(err)
			}
		}
		// the same document again updates, it does not create
		_, _, err = client.DomainService.UpsertPrivateYaml(workspace.Domain, &buddy.DomainYamlOps{Yaml: &encoded})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertPrivateYaml", err))
		}
		private, _, err := client.DomainService.GetList(workspace.Domain, &buddy.DomainGetListQuery{Type: buddy.DomainTypePrivate})
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetList", err))
		}
		err = CheckIntFieldEqual("len(Domains)", len(private.Domains), 2)
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range private.Domains {
			_, err = client.DomainService.Delete(workspace.Domain, d.Id)
			if err != nil {
				t.Fatal(ErrorFormatted("DomainService.Delete", err))
			}
		}
	}
}

func testDomainRecordGet(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, record *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		r, _, err := client.DomainService.GetRecord(workspace.Domain, domain.Id, fmt.Sprintf("%s.%s", record.Name, domain.Name), "A")
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetRecord", err))
		}
		err = CheckRecord(r, record.Name, record.Note, record.AgentNote, record.Type, record.Ttl, buddy.DomainRecordRoutingSimple, record.Values[0], "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainRecordDelete(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, record *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.DomainService.DeleteRecord(workspace.Domain, domain.Id, fmt.Sprintf("%s.%s", record.Name, domain.Name), "A")
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.DeleteRecord", err))
		}
	}
}

func testDomainRecordGetList(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.DomainService.GetRecords(workspace.Domain, domain.Id)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetRecords", err))
		}
		err = CheckRecords(list, 3)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainGeoRecordUpsert(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, out *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		fullName := fmt.Sprintf("%s.%s", name, domain.Name)
		ttl := 600
		typ := "TXT"
		val := "Z"
		vals := []string{val}
		routing := buddy.DomainRecordRoutingGeolocation
		countryName := buddy.DomainRecordCountryNepal
		countryValue := "A"
		country := map[string][]string{
			countryName: {countryValue},
		}
		ops := buddy.RecordUpsertOps{
			Ttl:     &ttl,
			Routing: &routing,
			Country: &country,
			Values:  &vals,
		}
		record, _, err := client.DomainService.UpsertRecord(workspace.Domain, domain.Id, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, "", "", typ, ttl, routing, val, "", "", countryName, countryValue)
		if err != nil {
			t.Fatal(err)
		}
		continentName := buddy.DomainRecordContinentAsia
		continentValue := "C"
		continent := map[string][]string{
			continentName: {continentValue},
		}
		ops = buddy.RecordUpsertOps{
			Ttl:       &ttl,
			Routing:   &routing,
			Continent: &continent,
			Values:    &vals,
		}
		record, _, err = client.DomainService.UpsertRecord(workspace.Domain, domain.Id, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, "", "", typ, ttl, routing, val, continentName, continentValue, "", "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *record
	}
}

func testDomainRecordUpsert(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, out *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		note := RandString(10)
		agentNote := RandString(10)
		fullName := fmt.Sprintf("%s.%s", name, domain.Name)
		val := "1.1.1.1"
		vals := []string{val}
		ttl := 300
		typ := "A"
		ops := buddy.RecordUpsertOps{
			Ttl:       &ttl,
			Note:      &note,
			AgentNote: &agentNote,
			Values:    &vals,
		}
		record, _, err := client.DomainService.UpsertRecord(workspace.Domain, domain.Id, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, note, agentNote, typ, ttl, buddy.DomainRecordRoutingSimple, val, "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
		newVal := "2.2.2.2"
		newNote := RandString(10)
		newAgentNote := RandString(10)
		newValues := []string{newVal}
		newTtl := 3600
		ops = buddy.RecordUpsertOps{
			Note:      &newNote,
			AgentNote: &newAgentNote,
			Ttl:       &newTtl,
			Values:    &newValues,
		}
		record, _, err = client.DomainService.UpsertRecord(workspace.Domain, domain.Id, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, newNote, newAgentNote, typ, newTtl, buddy.DomainRecordRoutingSimple, newVal, "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *record
	}
}

func TestDomain(t *testing.T) {
	seed, err := SeedInitialData(&SeedOps{
		workspace: true,
	})
	if err != nil {
		t.Fatal(ErrorFormatted("SeedInitialData", err))
	}
	var domain buddy.Domain
	var record buddy.Record
	t.Run("Create", testDomainCreate(seed.Client, seed.Workspace, &domain))
	t.Run("List", testDomainList(seed.Client, seed.Workspace, &domain))
	t.Run("RecordUpsert", testDomainRecordUpsert(seed.Client, seed.Workspace, &domain, &record))
	t.Run("RecordGet", testDomainRecordGet(seed.Client, seed.Workspace, &domain, &record))
	t.Run("RecordGetList", testDomainRecordGetList(seed.Client, seed.Workspace, &domain))
	t.Run("RecordDelete", testDomainRecordDelete(seed.Client, seed.Workspace, &domain, &record))
	t.Run("GeoRecordUpsert", testDomainGeoRecordUpsert(seed.Client, seed.Workspace, &domain, &record))
	t.Run("YamlUpdate", testDomainYamlUpdate(seed.Client, seed.Workspace, &domain))
	t.Run("PrivateYamlUpsert", testDomainPrivateYamlUpsert(seed.Client, seed.Workspace))
	t.Run("Delete", testDomainDelete(seed.Client, seed.Workspace, &domain))
}
