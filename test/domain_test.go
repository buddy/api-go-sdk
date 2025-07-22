package test

import (
	"fmt"
	"github.com/buddy/api-go-sdk/buddy"
	"testing"
)

func testDomainCreate(client *buddy.Client, workspace *buddy.Workspace, out *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		name := RandDomain()
		ops := buddy.DomainCreateOps{
			Name: &name,
		}
		domain, _, err := client.DomainService.Create(workspace.Domain, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.Create", err))
		}
		err = CheckDomain(domain, name)
		if err != nil {
			t.Fatal(err)
		}
		*out = *domain
	}
}

func testDomainList(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		domains, _, err := client.DomainService.GetList(workspace.Domain)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.Create", err))
		}
		err = CheckDomains(domains, domain)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainRecordGet(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, record *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		r, _, err := client.DomainService.GetRecord(workspace.Domain, fmt.Sprintf("%s.%s", record.Name, domain.Name), "A")
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.GetRecord", err))
		}
		err = CheckRecord(r, record.Name, record.Type, record.Ttl, buddy.DomainRecordRoutingSimple, record.Values[0], "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func testDomainRecordDelete(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, record *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := client.DomainService.DeleteRecord(workspace.Domain, fmt.Sprintf("%s.%s", record.Name, domain.Name), "A")
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.DeleteRecord", err))
		}
	}
}

func testDomainRecordGetList(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain) func(t *testing.T) {
	return func(t *testing.T) {
		list, _, err := client.DomainService.GetRecords(workspace.Domain, domain.Name)
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
		record, _, err := client.DomainService.UpsertRecord(workspace.Domain, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, typ, ttl, routing, val, "", "", countryName, countryValue)
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
		record, _, err = client.DomainService.UpsertRecord(workspace.Domain, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, typ, ttl, routing, val, continentName, continentValue, "", "")
		if err != nil {
			t.Fatal(err)
		}
		*out = *record
	}
}

func testDomainRecordUpsert(client *buddy.Client, workspace *buddy.Workspace, domain *buddy.Domain, out *buddy.Record) func(t *testing.T) {
	return func(t *testing.T) {
		name := UniqueString()
		fullName := fmt.Sprintf("%s.%s", name, domain.Name)
		val := "1.1.1.1"
		vals := []string{val}
		ttl := 300
		typ := "A"
		ops := buddy.RecordUpsertOps{
			Ttl:    &ttl,
			Values: &vals,
		}
		record, _, err := client.DomainService.UpsertRecord(workspace.Domain, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, typ, ttl, buddy.DomainRecordRoutingSimple, val, "", "", "", "")
		if err != nil {
			t.Fatal(err)
		}
		newVal := "2.2.2.2"
		newValues := []string{newVal}
		newTtl := 3600
		ops = buddy.RecordUpsertOps{
			Ttl:    &newTtl,
			Values: &newValues,
		}
		record, _, err = client.DomainService.UpsertRecord(workspace.Domain, fullName, typ, &ops)
		if err != nil {
			t.Fatal(ErrorFormatted("DomainService.UpsertRecord", err))
		}
		err = CheckRecord(record, name, typ, newTtl, buddy.DomainRecordRoutingSimple, newVal, "", "", "", "")
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
}
