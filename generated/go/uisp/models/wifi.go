// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Wifi wifi
//
// swagger:model Wifi
type Wifi struct {

	// authentication
	// Enum: [psk psk2 ent none]
	Authentication string `json:"authentication,omitempty"`

	// available
	Available bool `json:"available,omitempty"`

	// channel
	Channel float64 `json:"channel,omitempty"`

	// channel width
	// Minimum: 0
	ChannelWidth *int64 `json:"channelWidth,omitempty"`

	// country
	// Enum: [XX XY AF AX AL DZ AS AD AO AI AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BQ BA BW BV BR IO BN BG BF BI CV KH CM CA KY CF TD CL CN CX CC CO KM CK CR CI HR CW CY CZ CD DK DJ DM DO EC EG SV GQ ER EE SZ ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GG GN GW GY HT HM HN HK HU IS IN ID IQ IE IM IL IT JM JP JE JO KZ KE KI KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NL NC NZ NI NE NG NU NF MP NO OM PK PW PA PG PY PE PH PN PL PT PR QA KR RS SC CG RE RO RU RW BL SH KN LC MF PM VC WS SM ST SA SN SL SG SX SK SI SB SO ZA GS SS ES LK PS SR SJ SE CH TW TJ TZ TH TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UY UZ VU VA VE VN VG VI WF EH YE ZM ZW]
	Country string `json:"country,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// encryption
	// Enum: [wep wpa wpa-psk wpa2 enabled none]
	Encryption string `json:"encryption,omitempty"`

	// frequency
	// Example: 5400
	// Minimum: 0
	Frequency *float64 `json:"frequency,omitempty"`

	// is channel auto
	IsChannelAuto bool `json:"isChannelAuto,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// mac
	// Pattern: ^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$
	Mac string `json:"mac,omitempty"`

	// mode
	// Enum: [ap mesh]
	Mode string `json:"mode,omitempty"`

	// SSID
	Ssid string `json:"ssid,omitempty"`

	// stations count
	StationsCount float64 `json:"stationsCount,omitempty"`

	// tx power
	TxPower float64 `json:"txPower,omitempty"`
}

// Validate validates this wifi
func (m *Wifi) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var wifiTypeAuthenticationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["psk","psk2","ent","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiTypeAuthenticationPropEnum = append(wifiTypeAuthenticationPropEnum, v)
	}
}

const (

	// WifiAuthenticationPsk captures enum value "psk"
	WifiAuthenticationPsk string = "psk"

	// WifiAuthenticationPsk2 captures enum value "psk2"
	WifiAuthenticationPsk2 string = "psk2"

	// WifiAuthenticationEnt captures enum value "ent"
	WifiAuthenticationEnt string = "ent"

	// WifiAuthenticationNone captures enum value "none"
	WifiAuthenticationNone string = "none"
)

// prop value enum
func (m *Wifi) validateAuthenticationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiTypeAuthenticationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wifi) validateAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationEnum("authentication", "body", m.Authentication); err != nil {
		return err
	}

	return nil
}

func (m *Wifi) validateChannelWidth(formats strfmt.Registry) error {
	if swag.IsZero(m.ChannelWidth) { // not required
		return nil
	}

	if err := validate.MinimumInt("channelWidth", "body", *m.ChannelWidth, 0, false); err != nil {
		return err
	}

	return nil
}

var wifiTypeCountryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["XX","XY","AF","AX","AL","DZ","AS","AD","AO","AI","AQ","AG","AR","AM","AW","AU","AT","AZ","BS","BH","BD","BB","BY","BE","BZ","BJ","BM","BT","BO","BQ","BA","BW","BV","BR","IO","BN","BG","BF","BI","CV","KH","CM","CA","KY","CF","TD","CL","CN","CX","CC","CO","KM","CK","CR","CI","HR","CW","CY","CZ","CD","DK","DJ","DM","DO","EC","EG","SV","GQ","ER","EE","SZ","ET","FK","FO","FJ","FI","FR","GF","PF","TF","GA","GM","GE","DE","GH","GI","GR","GL","GD","GP","GU","GT","GG","GN","GW","GY","HT","HM","HN","HK","HU","IS","IN","ID","IQ","IE","IM","IL","IT","JM","JP","JE","JO","KZ","KE","KI","KW","KG","LA","LV","LB","LS","LR","LY","LI","LT","LU","MO","MK","MG","MW","MY","MV","ML","MT","MH","MQ","MR","MU","YT","MX","FM","MD","MC","MN","ME","MS","MA","MZ","MM","NA","NR","NP","NL","NC","NZ","NI","NE","NG","NU","NF","MP","NO","OM","PK","PW","PA","PG","PY","PE","PH","PN","PL","PT","PR","QA","KR","RS","SC","CG","RE","RO","RU","RW","BL","SH","KN","LC","MF","PM","VC","WS","SM","ST","SA","SN","SL","SG","SX","SK","SI","SB","SO","ZA","GS","SS","ES","LK","PS","SR","SJ","SE","CH","TW","TJ","TZ","TH","TL","TG","TK","TO","TT","TN","TR","TM","TC","TV","UG","UA","AE","GB","US","UY","UZ","VU","VA","VE","VN","VG","VI","WF","EH","YE","ZM","ZW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiTypeCountryPropEnum = append(wifiTypeCountryPropEnum, v)
	}
}

const (

	// WifiCountryXX captures enum value "XX"
	WifiCountryXX string = "XX"

	// WifiCountryXY captures enum value "XY"
	WifiCountryXY string = "XY"

	// WifiCountryAF captures enum value "AF"
	WifiCountryAF string = "AF"

	// WifiCountryAX captures enum value "AX"
	WifiCountryAX string = "AX"

	// WifiCountryAL captures enum value "AL"
	WifiCountryAL string = "AL"

	// WifiCountryDZ captures enum value "DZ"
	WifiCountryDZ string = "DZ"

	// WifiCountryAS captures enum value "AS"
	WifiCountryAS string = "AS"

	// WifiCountryAD captures enum value "AD"
	WifiCountryAD string = "AD"

	// WifiCountryAO captures enum value "AO"
	WifiCountryAO string = "AO"

	// WifiCountryAI captures enum value "AI"
	WifiCountryAI string = "AI"

	// WifiCountryAQ captures enum value "AQ"
	WifiCountryAQ string = "AQ"

	// WifiCountryAG captures enum value "AG"
	WifiCountryAG string = "AG"

	// WifiCountryAR captures enum value "AR"
	WifiCountryAR string = "AR"

	// WifiCountryAM captures enum value "AM"
	WifiCountryAM string = "AM"

	// WifiCountryAW captures enum value "AW"
	WifiCountryAW string = "AW"

	// WifiCountryAU captures enum value "AU"
	WifiCountryAU string = "AU"

	// WifiCountryAT captures enum value "AT"
	WifiCountryAT string = "AT"

	// WifiCountryAZ captures enum value "AZ"
	WifiCountryAZ string = "AZ"

	// WifiCountryBS captures enum value "BS"
	WifiCountryBS string = "BS"

	// WifiCountryBH captures enum value "BH"
	WifiCountryBH string = "BH"

	// WifiCountryBD captures enum value "BD"
	WifiCountryBD string = "BD"

	// WifiCountryBB captures enum value "BB"
	WifiCountryBB string = "BB"

	// WifiCountryBY captures enum value "BY"
	WifiCountryBY string = "BY"

	// WifiCountryBE captures enum value "BE"
	WifiCountryBE string = "BE"

	// WifiCountryBZ captures enum value "BZ"
	WifiCountryBZ string = "BZ"

	// WifiCountryBJ captures enum value "BJ"
	WifiCountryBJ string = "BJ"

	// WifiCountryBM captures enum value "BM"
	WifiCountryBM string = "BM"

	// WifiCountryBT captures enum value "BT"
	WifiCountryBT string = "BT"

	// WifiCountryBO captures enum value "BO"
	WifiCountryBO string = "BO"

	// WifiCountryBQ captures enum value "BQ"
	WifiCountryBQ string = "BQ"

	// WifiCountryBA captures enum value "BA"
	WifiCountryBA string = "BA"

	// WifiCountryBW captures enum value "BW"
	WifiCountryBW string = "BW"

	// WifiCountryBV captures enum value "BV"
	WifiCountryBV string = "BV"

	// WifiCountryBR captures enum value "BR"
	WifiCountryBR string = "BR"

	// WifiCountryIO captures enum value "IO"
	WifiCountryIO string = "IO"

	// WifiCountryBN captures enum value "BN"
	WifiCountryBN string = "BN"

	// WifiCountryBG captures enum value "BG"
	WifiCountryBG string = "BG"

	// WifiCountryBF captures enum value "BF"
	WifiCountryBF string = "BF"

	// WifiCountryBI captures enum value "BI"
	WifiCountryBI string = "BI"

	// WifiCountryCV captures enum value "CV"
	WifiCountryCV string = "CV"

	// WifiCountryKH captures enum value "KH"
	WifiCountryKH string = "KH"

	// WifiCountryCM captures enum value "CM"
	WifiCountryCM string = "CM"

	// WifiCountryCA captures enum value "CA"
	WifiCountryCA string = "CA"

	// WifiCountryKY captures enum value "KY"
	WifiCountryKY string = "KY"

	// WifiCountryCF captures enum value "CF"
	WifiCountryCF string = "CF"

	// WifiCountryTD captures enum value "TD"
	WifiCountryTD string = "TD"

	// WifiCountryCL captures enum value "CL"
	WifiCountryCL string = "CL"

	// WifiCountryCN captures enum value "CN"
	WifiCountryCN string = "CN"

	// WifiCountryCX captures enum value "CX"
	WifiCountryCX string = "CX"

	// WifiCountryCC captures enum value "CC"
	WifiCountryCC string = "CC"

	// WifiCountryCO captures enum value "CO"
	WifiCountryCO string = "CO"

	// WifiCountryKM captures enum value "KM"
	WifiCountryKM string = "KM"

	// WifiCountryCK captures enum value "CK"
	WifiCountryCK string = "CK"

	// WifiCountryCR captures enum value "CR"
	WifiCountryCR string = "CR"

	// WifiCountryCI captures enum value "CI"
	WifiCountryCI string = "CI"

	// WifiCountryHR captures enum value "HR"
	WifiCountryHR string = "HR"

	// WifiCountryCW captures enum value "CW"
	WifiCountryCW string = "CW"

	// WifiCountryCY captures enum value "CY"
	WifiCountryCY string = "CY"

	// WifiCountryCZ captures enum value "CZ"
	WifiCountryCZ string = "CZ"

	// WifiCountryCD captures enum value "CD"
	WifiCountryCD string = "CD"

	// WifiCountryDK captures enum value "DK"
	WifiCountryDK string = "DK"

	// WifiCountryDJ captures enum value "DJ"
	WifiCountryDJ string = "DJ"

	// WifiCountryDM captures enum value "DM"
	WifiCountryDM string = "DM"

	// WifiCountryDO captures enum value "DO"
	WifiCountryDO string = "DO"

	// WifiCountryEC captures enum value "EC"
	WifiCountryEC string = "EC"

	// WifiCountryEG captures enum value "EG"
	WifiCountryEG string = "EG"

	// WifiCountrySV captures enum value "SV"
	WifiCountrySV string = "SV"

	// WifiCountryGQ captures enum value "GQ"
	WifiCountryGQ string = "GQ"

	// WifiCountryER captures enum value "ER"
	WifiCountryER string = "ER"

	// WifiCountryEE captures enum value "EE"
	WifiCountryEE string = "EE"

	// WifiCountrySZ captures enum value "SZ"
	WifiCountrySZ string = "SZ"

	// WifiCountryET captures enum value "ET"
	WifiCountryET string = "ET"

	// WifiCountryFK captures enum value "FK"
	WifiCountryFK string = "FK"

	// WifiCountryFO captures enum value "FO"
	WifiCountryFO string = "FO"

	// WifiCountryFJ captures enum value "FJ"
	WifiCountryFJ string = "FJ"

	// WifiCountryFI captures enum value "FI"
	WifiCountryFI string = "FI"

	// WifiCountryFR captures enum value "FR"
	WifiCountryFR string = "FR"

	// WifiCountryGF captures enum value "GF"
	WifiCountryGF string = "GF"

	// WifiCountryPF captures enum value "PF"
	WifiCountryPF string = "PF"

	// WifiCountryTF captures enum value "TF"
	WifiCountryTF string = "TF"

	// WifiCountryGA captures enum value "GA"
	WifiCountryGA string = "GA"

	// WifiCountryGM captures enum value "GM"
	WifiCountryGM string = "GM"

	// WifiCountryGE captures enum value "GE"
	WifiCountryGE string = "GE"

	// WifiCountryDE captures enum value "DE"
	WifiCountryDE string = "DE"

	// WifiCountryGH captures enum value "GH"
	WifiCountryGH string = "GH"

	// WifiCountryGI captures enum value "GI"
	WifiCountryGI string = "GI"

	// WifiCountryGR captures enum value "GR"
	WifiCountryGR string = "GR"

	// WifiCountryGL captures enum value "GL"
	WifiCountryGL string = "GL"

	// WifiCountryGD captures enum value "GD"
	WifiCountryGD string = "GD"

	// WifiCountryGP captures enum value "GP"
	WifiCountryGP string = "GP"

	// WifiCountryGU captures enum value "GU"
	WifiCountryGU string = "GU"

	// WifiCountryGT captures enum value "GT"
	WifiCountryGT string = "GT"

	// WifiCountryGG captures enum value "GG"
	WifiCountryGG string = "GG"

	// WifiCountryGN captures enum value "GN"
	WifiCountryGN string = "GN"

	// WifiCountryGW captures enum value "GW"
	WifiCountryGW string = "GW"

	// WifiCountryGY captures enum value "GY"
	WifiCountryGY string = "GY"

	// WifiCountryHT captures enum value "HT"
	WifiCountryHT string = "HT"

	// WifiCountryHM captures enum value "HM"
	WifiCountryHM string = "HM"

	// WifiCountryHN captures enum value "HN"
	WifiCountryHN string = "HN"

	// WifiCountryHK captures enum value "HK"
	WifiCountryHK string = "HK"

	// WifiCountryHU captures enum value "HU"
	WifiCountryHU string = "HU"

	// WifiCountryIS captures enum value "IS"
	WifiCountryIS string = "IS"

	// WifiCountryIN captures enum value "IN"
	WifiCountryIN string = "IN"

	// WifiCountryID captures enum value "ID"
	WifiCountryID string = "ID"

	// WifiCountryIQ captures enum value "IQ"
	WifiCountryIQ string = "IQ"

	// WifiCountryIE captures enum value "IE"
	WifiCountryIE string = "IE"

	// WifiCountryIM captures enum value "IM"
	WifiCountryIM string = "IM"

	// WifiCountryIL captures enum value "IL"
	WifiCountryIL string = "IL"

	// WifiCountryIT captures enum value "IT"
	WifiCountryIT string = "IT"

	// WifiCountryJM captures enum value "JM"
	WifiCountryJM string = "JM"

	// WifiCountryJP captures enum value "JP"
	WifiCountryJP string = "JP"

	// WifiCountryJE captures enum value "JE"
	WifiCountryJE string = "JE"

	// WifiCountryJO captures enum value "JO"
	WifiCountryJO string = "JO"

	// WifiCountryKZ captures enum value "KZ"
	WifiCountryKZ string = "KZ"

	// WifiCountryKE captures enum value "KE"
	WifiCountryKE string = "KE"

	// WifiCountryKI captures enum value "KI"
	WifiCountryKI string = "KI"

	// WifiCountryKW captures enum value "KW"
	WifiCountryKW string = "KW"

	// WifiCountryKG captures enum value "KG"
	WifiCountryKG string = "KG"

	// WifiCountryLA captures enum value "LA"
	WifiCountryLA string = "LA"

	// WifiCountryLV captures enum value "LV"
	WifiCountryLV string = "LV"

	// WifiCountryLB captures enum value "LB"
	WifiCountryLB string = "LB"

	// WifiCountryLS captures enum value "LS"
	WifiCountryLS string = "LS"

	// WifiCountryLR captures enum value "LR"
	WifiCountryLR string = "LR"

	// WifiCountryLY captures enum value "LY"
	WifiCountryLY string = "LY"

	// WifiCountryLI captures enum value "LI"
	WifiCountryLI string = "LI"

	// WifiCountryLT captures enum value "LT"
	WifiCountryLT string = "LT"

	// WifiCountryLU captures enum value "LU"
	WifiCountryLU string = "LU"

	// WifiCountryMO captures enum value "MO"
	WifiCountryMO string = "MO"

	// WifiCountryMK captures enum value "MK"
	WifiCountryMK string = "MK"

	// WifiCountryMG captures enum value "MG"
	WifiCountryMG string = "MG"

	// WifiCountryMW captures enum value "MW"
	WifiCountryMW string = "MW"

	// WifiCountryMY captures enum value "MY"
	WifiCountryMY string = "MY"

	// WifiCountryMV captures enum value "MV"
	WifiCountryMV string = "MV"

	// WifiCountryML captures enum value "ML"
	WifiCountryML string = "ML"

	// WifiCountryMT captures enum value "MT"
	WifiCountryMT string = "MT"

	// WifiCountryMH captures enum value "MH"
	WifiCountryMH string = "MH"

	// WifiCountryMQ captures enum value "MQ"
	WifiCountryMQ string = "MQ"

	// WifiCountryMR captures enum value "MR"
	WifiCountryMR string = "MR"

	// WifiCountryMU captures enum value "MU"
	WifiCountryMU string = "MU"

	// WifiCountryYT captures enum value "YT"
	WifiCountryYT string = "YT"

	// WifiCountryMX captures enum value "MX"
	WifiCountryMX string = "MX"

	// WifiCountryFM captures enum value "FM"
	WifiCountryFM string = "FM"

	// WifiCountryMD captures enum value "MD"
	WifiCountryMD string = "MD"

	// WifiCountryMC captures enum value "MC"
	WifiCountryMC string = "MC"

	// WifiCountryMN captures enum value "MN"
	WifiCountryMN string = "MN"

	// WifiCountryME captures enum value "ME"
	WifiCountryME string = "ME"

	// WifiCountryMS captures enum value "MS"
	WifiCountryMS string = "MS"

	// WifiCountryMA captures enum value "MA"
	WifiCountryMA string = "MA"

	// WifiCountryMZ captures enum value "MZ"
	WifiCountryMZ string = "MZ"

	// WifiCountryMM captures enum value "MM"
	WifiCountryMM string = "MM"

	// WifiCountryNA captures enum value "NA"
	WifiCountryNA string = "NA"

	// WifiCountryNR captures enum value "NR"
	WifiCountryNR string = "NR"

	// WifiCountryNP captures enum value "NP"
	WifiCountryNP string = "NP"

	// WifiCountryNL captures enum value "NL"
	WifiCountryNL string = "NL"

	// WifiCountryNC captures enum value "NC"
	WifiCountryNC string = "NC"

	// WifiCountryNZ captures enum value "NZ"
	WifiCountryNZ string = "NZ"

	// WifiCountryNI captures enum value "NI"
	WifiCountryNI string = "NI"

	// WifiCountryNE captures enum value "NE"
	WifiCountryNE string = "NE"

	// WifiCountryNG captures enum value "NG"
	WifiCountryNG string = "NG"

	// WifiCountryNU captures enum value "NU"
	WifiCountryNU string = "NU"

	// WifiCountryNF captures enum value "NF"
	WifiCountryNF string = "NF"

	// WifiCountryMP captures enum value "MP"
	WifiCountryMP string = "MP"

	// WifiCountryNO captures enum value "NO"
	WifiCountryNO string = "NO"

	// WifiCountryOM captures enum value "OM"
	WifiCountryOM string = "OM"

	// WifiCountryPK captures enum value "PK"
	WifiCountryPK string = "PK"

	// WifiCountryPW captures enum value "PW"
	WifiCountryPW string = "PW"

	// WifiCountryPA captures enum value "PA"
	WifiCountryPA string = "PA"

	// WifiCountryPG captures enum value "PG"
	WifiCountryPG string = "PG"

	// WifiCountryPY captures enum value "PY"
	WifiCountryPY string = "PY"

	// WifiCountryPE captures enum value "PE"
	WifiCountryPE string = "PE"

	// WifiCountryPH captures enum value "PH"
	WifiCountryPH string = "PH"

	// WifiCountryPN captures enum value "PN"
	WifiCountryPN string = "PN"

	// WifiCountryPL captures enum value "PL"
	WifiCountryPL string = "PL"

	// WifiCountryPT captures enum value "PT"
	WifiCountryPT string = "PT"

	// WifiCountryPR captures enum value "PR"
	WifiCountryPR string = "PR"

	// WifiCountryQA captures enum value "QA"
	WifiCountryQA string = "QA"

	// WifiCountryKR captures enum value "KR"
	WifiCountryKR string = "KR"

	// WifiCountryRS captures enum value "RS"
	WifiCountryRS string = "RS"

	// WifiCountrySC captures enum value "SC"
	WifiCountrySC string = "SC"

	// WifiCountryCG captures enum value "CG"
	WifiCountryCG string = "CG"

	// WifiCountryRE captures enum value "RE"
	WifiCountryRE string = "RE"

	// WifiCountryRO captures enum value "RO"
	WifiCountryRO string = "RO"

	// WifiCountryRU captures enum value "RU"
	WifiCountryRU string = "RU"

	// WifiCountryRW captures enum value "RW"
	WifiCountryRW string = "RW"

	// WifiCountryBL captures enum value "BL"
	WifiCountryBL string = "BL"

	// WifiCountrySH captures enum value "SH"
	WifiCountrySH string = "SH"

	// WifiCountryKN captures enum value "KN"
	WifiCountryKN string = "KN"

	// WifiCountryLC captures enum value "LC"
	WifiCountryLC string = "LC"

	// WifiCountryMF captures enum value "MF"
	WifiCountryMF string = "MF"

	// WifiCountryPM captures enum value "PM"
	WifiCountryPM string = "PM"

	// WifiCountryVC captures enum value "VC"
	WifiCountryVC string = "VC"

	// WifiCountryWS captures enum value "WS"
	WifiCountryWS string = "WS"

	// WifiCountrySM captures enum value "SM"
	WifiCountrySM string = "SM"

	// WifiCountryST captures enum value "ST"
	WifiCountryST string = "ST"

	// WifiCountrySA captures enum value "SA"
	WifiCountrySA string = "SA"

	// WifiCountrySN captures enum value "SN"
	WifiCountrySN string = "SN"

	// WifiCountrySL captures enum value "SL"
	WifiCountrySL string = "SL"

	// WifiCountrySG captures enum value "SG"
	WifiCountrySG string = "SG"

	// WifiCountrySX captures enum value "SX"
	WifiCountrySX string = "SX"

	// WifiCountrySK captures enum value "SK"
	WifiCountrySK string = "SK"

	// WifiCountrySI captures enum value "SI"
	WifiCountrySI string = "SI"

	// WifiCountrySB captures enum value "SB"
	WifiCountrySB string = "SB"

	// WifiCountrySO captures enum value "SO"
	WifiCountrySO string = "SO"

	// WifiCountryZA captures enum value "ZA"
	WifiCountryZA string = "ZA"

	// WifiCountryGS captures enum value "GS"
	WifiCountryGS string = "GS"

	// WifiCountrySS captures enum value "SS"
	WifiCountrySS string = "SS"

	// WifiCountryES captures enum value "ES"
	WifiCountryES string = "ES"

	// WifiCountryLK captures enum value "LK"
	WifiCountryLK string = "LK"

	// WifiCountryPS captures enum value "PS"
	WifiCountryPS string = "PS"

	// WifiCountrySR captures enum value "SR"
	WifiCountrySR string = "SR"

	// WifiCountrySJ captures enum value "SJ"
	WifiCountrySJ string = "SJ"

	// WifiCountrySE captures enum value "SE"
	WifiCountrySE string = "SE"

	// WifiCountryCH captures enum value "CH"
	WifiCountryCH string = "CH"

	// WifiCountryTW captures enum value "TW"
	WifiCountryTW string = "TW"

	// WifiCountryTJ captures enum value "TJ"
	WifiCountryTJ string = "TJ"

	// WifiCountryTZ captures enum value "TZ"
	WifiCountryTZ string = "TZ"

	// WifiCountryTH captures enum value "TH"
	WifiCountryTH string = "TH"

	// WifiCountryTL captures enum value "TL"
	WifiCountryTL string = "TL"

	// WifiCountryTG captures enum value "TG"
	WifiCountryTG string = "TG"

	// WifiCountryTK captures enum value "TK"
	WifiCountryTK string = "TK"

	// WifiCountryTO captures enum value "TO"
	WifiCountryTO string = "TO"

	// WifiCountryTT captures enum value "TT"
	WifiCountryTT string = "TT"

	// WifiCountryTN captures enum value "TN"
	WifiCountryTN string = "TN"

	// WifiCountryTR captures enum value "TR"
	WifiCountryTR string = "TR"

	// WifiCountryTM captures enum value "TM"
	WifiCountryTM string = "TM"

	// WifiCountryTC captures enum value "TC"
	WifiCountryTC string = "TC"

	// WifiCountryTV captures enum value "TV"
	WifiCountryTV string = "TV"

	// WifiCountryUG captures enum value "UG"
	WifiCountryUG string = "UG"

	// WifiCountryUA captures enum value "UA"
	WifiCountryUA string = "UA"

	// WifiCountryAE captures enum value "AE"
	WifiCountryAE string = "AE"

	// WifiCountryGB captures enum value "GB"
	WifiCountryGB string = "GB"

	// WifiCountryUS captures enum value "US"
	WifiCountryUS string = "US"

	// WifiCountryUY captures enum value "UY"
	WifiCountryUY string = "UY"

	// WifiCountryUZ captures enum value "UZ"
	WifiCountryUZ string = "UZ"

	// WifiCountryVU captures enum value "VU"
	WifiCountryVU string = "VU"

	// WifiCountryVA captures enum value "VA"
	WifiCountryVA string = "VA"

	// WifiCountryVE captures enum value "VE"
	WifiCountryVE string = "VE"

	// WifiCountryVN captures enum value "VN"
	WifiCountryVN string = "VN"

	// WifiCountryVG captures enum value "VG"
	WifiCountryVG string = "VG"

	// WifiCountryVI captures enum value "VI"
	WifiCountryVI string = "VI"

	// WifiCountryWF captures enum value "WF"
	WifiCountryWF string = "WF"

	// WifiCountryEH captures enum value "EH"
	WifiCountryEH string = "EH"

	// WifiCountryYE captures enum value "YE"
	WifiCountryYE string = "YE"

	// WifiCountryZM captures enum value "ZM"
	WifiCountryZM string = "ZM"

	// WifiCountryZW captures enum value "ZW"
	WifiCountryZW string = "ZW"
)

// prop value enum
func (m *Wifi) validateCountryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiTypeCountryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wifi) validateCountry(formats strfmt.Registry) error {
	if swag.IsZero(m.Country) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

var wifiTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wep","wpa","wpa-psk","wpa2","enabled","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiTypeEncryptionPropEnum = append(wifiTypeEncryptionPropEnum, v)
	}
}

const (

	// WifiEncryptionWep captures enum value "wep"
	WifiEncryptionWep string = "wep"

	// WifiEncryptionWpa captures enum value "wpa"
	WifiEncryptionWpa string = "wpa"

	// WifiEncryptionWpaDashPsk captures enum value "wpa-psk"
	WifiEncryptionWpaDashPsk string = "wpa-psk"

	// WifiEncryptionWpa2 captures enum value "wpa2"
	WifiEncryptionWpa2 string = "wpa2"

	// WifiEncryptionEnabled captures enum value "enabled"
	WifiEncryptionEnabled string = "enabled"

	// WifiEncryptionNone captures enum value "none"
	WifiEncryptionNone string = "none"
)

// prop value enum
func (m *Wifi) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiTypeEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wifi) validateEncryption(formats strfmt.Registry) error {
	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *Wifi) validateFrequency(formats strfmt.Registry) error {
	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := validate.Minimum("frequency", "body", *m.Frequency, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Wifi) validateMac(formats strfmt.Registry) error {
	if swag.IsZero(m.Mac) { // not required
		return nil
	}

	if err := validate.Pattern("mac", "body", m.Mac, `^([0-9a-fA-F][0-9a-fA-F]:){5}([0-9a-fA-F][0-9a-fA-F])$|^([0-9a-fA-F]){12}$`); err != nil {
		return err
	}

	return nil
}

var wifiTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ap","mesh"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wifiTypeModePropEnum = append(wifiTypeModePropEnum, v)
	}
}

const (

	// WifiModeAp captures enum value "ap"
	WifiModeAp string = "ap"

	// WifiModeMesh captures enum value "mesh"
	WifiModeMesh string = "mesh"
)

// prop value enum
func (m *Wifi) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wifiTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Wifi) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wifi based on context it is used
func (m *Wifi) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Wifi) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wifi) UnmarshalBinary(b []byte) error {
	var res Wifi
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
