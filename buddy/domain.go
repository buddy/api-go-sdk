package buddy

import "net/http"

const (
	DomainRecordRoutingGeolocation = "Geolocation"
	DomainRecordRoutingSimple      = "Simple"

	DomainRecordContinentNorthAmerica = "NorthAmerica"
	DomainRecordContinentEurope       = "Europe"
	DomainRecordContinentAsia         = "Asia"
	DomainRecordContinentOceania      = "Oceania"
	DomainRecordContinentSouthAmerica = "SouthAmerica"
	DomainRecordContinentAfrica       = "Africa"
	DomainRecordContinentAntarctica   = "Antarctica"

	DomainRecordCountryAfghanistan                  = "AF"
	DomainRecordCountryAlbania                      = "AL"
	DomainRecordCountryAlgeria                      = "DZ"
	DomainRecordCountryAndorra                      = "AD"
	DomainRecordCountryAngola                       = "AO"
	DomainRecordCountryAntiguaAndBarbuda            = "AG"
	DomainRecordCountryArgentina                    = "AR"
	DomainRecordCountryArmenia                      = "AM"
	DomainRecordCountryAustralia                    = "AU"
	DomainRecordCountryAustria                      = "AT"
	DomainRecordCountryAzerbaijan                   = "AZ"
	DomainRecordCountryBahamas                      = "BS"
	DomainRecordCountryBahrain                      = "BH"
	DomainRecordCountryBangladesh                   = "BD"
	DomainRecordCountryBarbados                     = "BB"
	DomainRecordCountryBelarus                      = "BY"
	DomainRecordCountryBelgium                      = "BE"
	DomainRecordCountryBelize                       = "BZ"
	DomainRecordCountryBenin                        = "BJ"
	DomainRecordCountryBermuda                      = "BM"
	DomainRecordCountryBhutan                       = "BT"
	DomainRecordCountryBolivia                      = "BO"
	DomainRecordCountryBosniaAndHerzegovina         = "BA"
	DomainRecordCountryBotswana                     = "BW"
	DomainRecordCountryBrazil                       = "BR"
	DomainRecordCountryBrunei                       = "BN"
	DomainRecordCountryBulgaria                     = "BG"
	DomainRecordCountryBurkinaFaso                  = "BF"
	DomainRecordCountryBurundi                      = "BI"
	DomainRecordCountryCaboVerde                    = "CV"
	DomainRecordCountryCambodia                     = "KH"
	DomainRecordCountryCameroon                     = "CM"
	DomainRecordCountryCanada                       = "CA"
	DomainRecordCountryCaymanIslands                = "KY"
	DomainRecordCountryCentralAfricanRepublic       = "CF"
	DomainRecordCountryChad                         = "TD"
	DomainRecordCountryChile                        = "CL"
	DomainRecordCountryChina                        = "CN"
	DomainRecordCountryChristmasIsland              = "CX"
	DomainRecordCountryColombia                     = "CO"
	DomainRecordCountryComoros                      = "KM"
	DomainRecordCountryDemocraticRepublicOfTheCongo = "CD"
	DomainRecordCountryCongo                        = "CG"
	DomainRecordCountryCostaRica                    = "CR"
	DomainRecordCountryCroatia                      = "HR"
	DomainRecordCountryCuba                         = "CU"
	DomainRecordCountryCzechia                      = "CZ"
	DomainRecordCountryIvoryCoast                   = "CI"
	DomainRecordCountryDenmark                      = "DK"
	DomainRecordCountryDjibouti                     = "DJ"
	DomainRecordCountryDominica                     = "DM"
	DomainRecordCountryDominicanRepublic            = "DO"
	DomainRecordCountryEcuador                      = "EC"
	DomainRecordCountryEgypt                        = "EG"
	DomainRecordCountryElSalvador                   = "SV"
	DomainRecordCountryEquatorialGuinea             = "GQ"
	DomainRecordCountryEritrea                      = "ER"
	DomainRecordCountryEstonia                      = "EE"
	DomainRecordCountryEswatini                     = "SZ"
	DomainRecordCountryEthiopia                     = "ET"
	DomainRecordCountryFiji                         = "FJ"
	DomainRecordCountryFinland                      = "FI"
	DomainRecordCountryFrance                       = "FR"
	DomainRecordCountryGabon                        = "GA"
	DomainRecordCountryGambia                       = "GM"
	DomainRecordCountryGeorgia                      = "GE"
	DomainRecordCountryGermany                      = "DE"
	DomainRecordCountryGhana                        = "GH"
	DomainRecordCountryGreece                       = "GR"
	DomainRecordCountryGreenland                    = "GL"
	DomainRecordCountryGrenada                      = "GD"
	DomainRecordCountryGuatemala                    = "GT"
	DomainRecordCountryGuinea                       = "GN"
	DomainRecordCountryGuineaBissau                 = "GW"
	DomainRecordCountryGuyana                       = "GY"
	DomainRecordCountryHaiti                        = "HT"
	DomainRecordCountryHolySee                      = "VA"
	DomainRecordCountryHonduras                     = "HN"
	DomainRecordCountryHungary                      = "HU"
	DomainRecordCountryIceland                      = "IS"
	DomainRecordCountryIndia                        = "IN"
	DomainRecordCountryIndonesia                    = "ID"
	DomainRecordCountryIran                         = "IR"
	DomainRecordCountryIraq                         = "IQ"
	DomainRecordCountryIreland                      = "IE"
	DomainRecordCountryIsrael                       = "IL"
	DomainRecordCountryItaly                        = "IT"
	DomainRecordCountryJamaica                      = "JM"
	DomainRecordCountryJapan                        = "JP"
	DomainRecordCountryJordan                       = "JO"
	DomainRecordCountryKazakhstan                   = "KZ"
	DomainRecordCountryKenya                        = "KE"
	DomainRecordCountryKiribati                     = "KI"
	DomainRecordCountryNorthKorea                   = "KP"
	DomainRecordCountrySouthKorea                   = "KR"
	DomainRecordCountryKuwait                       = "KW"
	DomainRecordCountryKyrgyzstan                   = "KG"
	DomainRecordCountryLaoPeopleDemocraticRepublic  = "LA"
	DomainRecordCountryLatvia                       = "LV"
	DomainRecordCountryLebanon                      = "LB"
	DomainRecordCountryLesotho                      = "LS"
	DomainRecordCountryLiberia                      = "LR"
	DomainRecordCountryLibya                        = "LY"
	DomainRecordCountryLiechtenstein                = "LI"
	DomainRecordCountryLithuania                    = "LT"
	DomainRecordCountryLuxembourg                   = "LU"
	DomainRecordCountryMacao                        = "MO"
	DomainRecordCountryMadagascar                   = "MG"
	DomainRecordCountryMalawi                       = "MW"
	DomainRecordCountryMalaysia                     = "MY"
	DomainRecordCountryMaldives                     = "MV"
	DomainRecordCountryMali                         = "ML"
	DomainRecordCountryMalta                        = "MT"
	DomainRecordCountryMauritania                   = "MR"
	DomainRecordCountryMauritius                    = "MU"
	DomainRecordCountryMexico                       = "MX"
	DomainRecordCountryMicronesia                   = "FM"
	DomainRecordCountryMoldova                      = "MD"
	DomainRecordCountryMonaco                       = "MC"
	DomainRecordCountryMongolia                     = "MN"
	DomainRecordCountryMontenegro                   = "ME"
	DomainRecordCountryMontserrat                   = "MS"
	DomainRecordCountryMorocco                      = "MA"
	DomainRecordCountryMozambique                   = "MZ"
	DomainRecordCountryMyanmar                      = "MM"
	DomainRecordCountryNamibia                      = "NA"
	DomainRecordCountryNauru                        = "NR"
	DomainRecordCountryNepal                        = "NP"
	DomainRecordCountryNetherlands                  = "NL"
	DomainRecordCountryNewZealand                   = "NZ"
	DomainRecordCountryNicaragua                    = "NI"
	DomainRecordCountryNiger                        = "NE"
	DomainRecordCountryNigeria                      = "NG"
	DomainRecordCountryNorway                       = "NO"
	DomainRecordCountryOman                         = "OM"
	DomainRecordCountryPakistan                     = "PK"
	DomainRecordCountryPalestine                    = "PS"
	DomainRecordCountryPanama                       = "PA"
	DomainRecordCountryPapuaNewGuinea               = "PG"
	DomainRecordCountryParaguay                     = "PY"
	DomainRecordCountryPeru                         = "PE"
	DomainRecordCountryPhilippines                  = "PH"
	DomainRecordCountryPoland                       = "PL"
	DomainRecordCountryPortugal                     = "PT"
	DomainRecordCountryPuertoRico                   = "PR"
	DomainRecordCountryQatar                        = "QA"
	DomainRecordCountryRepublicOfNorthMacedonia     = "MK"
	DomainRecordCountryRomania                      = "RO"
	DomainRecordCountryRussianFederation            = "RU"
	DomainRecordCountryRwanda                       = "RW"
	DomainRecordCountrySaintHelena                  = "SH"
	DomainRecordCountrySaintKittsAndNevis           = "KN"
	DomainRecordCountrySaintLucia                   = "LC"
	DomainRecordCountrySaintVincentAndTheGrenadines = "VC"
	DomainRecordCountrySamoa                        = "WS"
	DomainRecordCountrySanMarino                    = "SM"
	DomainRecordCountrySaoTomeAndPrincipe           = "ST"
	DomainRecordCountrySaudiArabia                  = "SA"
	DomainRecordCountrySenegal                      = "SN"
	DomainRecordCountrySerbia                       = "RS"
	DomainRecordCountrySeychelles                   = "SC"
	DomainRecordCountrySierraLeone                  = "SL"
	DomainRecordCountrySingapore                    = "SG"
	DomainRecordCountrySlovakia                     = "SK"
	DomainRecordCountrySlovenia                     = "SI"
	DomainRecordCountrySolomonIslands               = "SB"
	DomainRecordCountrySomalia                      = "SO"
	DomainRecordCountrySouthAfrica                  = "ZA"
	DomainRecordCountrySouthSudan                   = "SS"
	DomainRecordCountrySpain                        = "ES"
	DomainRecordCountrySriLanka                     = "LK"
	DomainRecordCountrySudan                        = "SD"
	DomainRecordCountrySuriname                     = "SR"
	DomainRecordCountrySweden                       = "SE"
	DomainRecordCountrySwitzerland                  = "CH"
	DomainRecordCountrySyrianArabRepublic           = "SY"
	DomainRecordCountryTaiwan                       = "TW"
	DomainRecordCountryTajikistan                   = "TJ"
	DomainRecordCountryTanzania                     = "TZ"
	DomainRecordCountryThailand                     = "TH"
	DomainRecordCountryTimorLeste                   = "TL"
	DomainRecordCountryTogo                         = "TG"
	DomainRecordCountryTonga                        = "TO"
	DomainRecordCountryTrinidadAndTobago            = "TT"
	DomainRecordCountryTunisia                      = "TN"
	DomainRecordCountryTurkey                       = "TR"
	DomainRecordCountryTurkmenistan                 = "TM"
	DomainRecordCountryUganda                       = "UG"
	DomainRecordCountryUkraine                      = "UA"
	DomainRecordCountryUnitedArabEmirates           = "AE"
	DomainRecordCountryUnitedKingdom                = "GB"
	DomainRecordCountryUnitedStatesOfAmerica        = "US"
	DomainRecordCountryUruguay                      = "UY"
	DomainRecordCountryUzbekistan                   = "UZ"
	DomainRecordCountryVanuatu                      = "VU"
	DomainRecordCountryVenezuela                    = "VE"
	DomainRecordCountryVietNam                      = "VN"
	DomainRecordCountryYemen                        = "YE"
	DomainRecordCountryZambia                       = "ZM"
	DomainRecordCountryZimbabwe                     = "ZW"
)

type DomainService struct {
	client *Client
}

type Record struct {
	Name      string              `json:"name"`
	Type      string              `json:"type"`
	Ttl       int                 `json:"ttl"`
	Routing   string              `json:"routing"`
	Country   map[string][]string `json:"country"`
	Continent map[string][]string `json:"continent"`
	Values    []string            `json:"values"`
}

type Domain struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type DomainCreateOps struct {
	Name *string `json:"name"`
}

type RecordUpsertOps struct {
	Routing   *string              `json:"routing,omitempty"`
	Ttl       *int                 `json:"ttl,omitempty"`
	Values    *[]string            `json:"values,omitempty"`
	Country   *map[string][]string `json:"country,omitempty"`
	Continent *map[string][]string `json:"continent,omitempty"`
}

type Domains struct {
	Domains []*Domain `json:"domains"`
}

type Records struct {
	Records []*Record `json:"records"`
}

func (s *DomainService) GetList(workspaceDomain string) (*Domains, *http.Response, error) {
	var d *Domains
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains", workspaceDomain), &d, nil)
	return d, resp, err
}

func (s *DomainService) Create(workspaceDomain string, ops *DomainCreateOps) (*Domain, *http.Response, error) {
	var d *Domain
	resp, err := s.client.Create(s.client.NewUrlPath("/workspaces/%s/domains", workspaceDomain), &ops, nil, &d)
	return d, resp, err
}

func (s *DomainService) GetRecords(workspaceDomain string, domain string) (*Records, *http.Response, error) {
	var r *Records
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains/%s/records", workspaceDomain, domain), &r, nil)
	return r, resp, err
}

func (s *DomainService) GetRecord(workspaceDomain string, domain string, typ string) (*Record, *http.Response, error) {
	var r *Record
	resp, err := s.client.Get(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), &r, nil)
	return r, resp, err
}

func (s *DomainService) UpsertRecord(workspaceDomain string, domain string, typ string, ops *RecordUpsertOps) (*Record, *http.Response, error) {
	var r *Record
	resp, err := s.client.Patch(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), &ops, nil, &r)
	return r, resp, err
}

func (s *DomainService) DeleteRecord(workspaceDomain string, domain string, typ string) (*http.Response, error) {
	return s.client.Delete(s.client.NewUrlPath("/workspaces/%s/domains/%s/records/%s", workspaceDomain, domain, typ), nil, nil)
}
