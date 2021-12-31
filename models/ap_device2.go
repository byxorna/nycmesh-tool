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

// ApDevice2 ap device 2
//
// swagger:model apDevice 2
type ApDevice2 struct {

	// id
	ID string `json:"id,omitempty"`

	// Short names, for example UF-OLT.
	// Enum: [UF-Nano UF-Loco UF-Wifi UF-Instant UF-OLT UF-OLT4 UISP-R-Pro UISP-R-Lite UNMS-S-Lite UISP-S-Lite UISP-S-Pro UISP-P-Lite UISP-LTE ER-X ER-X-SFP ERLite-3 ERPoe-5 ERPro-8 ER-8 ER-8-XG ER-4 ER-6P ER-12 ER-12P ER-10X EP-R8 EP-R6 EP-S16 ES-12F ES-16-150W ES-24-250W ES-24-500W ES-24-Lite ES-48-500W ES-48-750W ES-48-Lite ES-8-150W ES-16-XG ES-10XP ES-10X ES-18X ES-26X EP-54V-150W EP-24V-72W EP-54V-72W TSW-PoE TSW-PoE PRO ACB-AC ACB-ISP ACB-LOCO AF11FX AF24 AF24HD AF2X AF3X AF4X AF5 AF5U AF5X AF-5XHD AF-LTU LTU-LITE AF-LTU5 LTU-Rocket AFLTULR AF60 AF60-LR WaveAP WaveCPE GBE-LR GBE GBE-Plus GBE-AP R2N R2T R5N R6N R36-GPS RM3-GPS R2N-GPS R5N-GPS R9N-GPS R5T-GPS RM3 R36 R9N N2N N5N N6N NS3 N36 N9N N9S LM2 LM5 B2N B2T B5N B5T BAC AG2 AG2-HP AG5 AG5-HP p2N p5N M25 P2B-400 P5B-300 P5B-300-ISO P5B-400 P5B-400-ISO P5B-620 LB5-120 LB5 N5B N5B-16 N5B-19 N5B-300 N5B-400 N5B-Client N2B N2B-13 N2B-400 PAP LAP-HP LAP AGW AGW-LR AGW-Pro AGW-Installer PB5 PB3 P36 PBM10 NB5 NB2 NB3 B36 NB9 SM5 WM5 IS-M5 Loco5AC NS-5AC R5AC-PTMP R5AC-PTP R5AC-Lite R5AC-PRISM R2AC-Prism R2AC-Gen2 RP-5AC-Gen2 NBE-2AC-13 NBE-5AC-16 NBE-5AC-19 NBE-5AC-Gen2 PBE-5AC-300 PBE-5AC-300-ISO PBE-5AC-400 PBE-5AC-400-ISO PBE-5AC-500 PBE-5AC-500-ISO PBE-5AC-620 PBE-5AC-620-ISO PBE-2AC-400 PBE-2AC-400-ISO PBE-5AC-X-Gen2 PBE-5AC-Gen2 PBE-5AC-ISO-Gen2 PBE-5AC-400-ISO-Gen2 LBE-5AC-16-120 LAP-120 LBE-5AC-23 LBE-5AC-Gen2 LBE-5AC-LR LAP-GPS IS-5AC PS-5AC SolarSwitch SolarPoint BulletAC-IP67 B-DB-AC UNKNOWN]
	Model string `json:"model,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// series
	Series string `json:"series,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this ap device 2
func (m *ApDevice2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var apDevice2TypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","UISP-R-Pro","UISP-R-Lite","UNMS-S-Lite","UISP-S-Lite","UISP-S-Pro","UISP-P-Lite","UISP-LTE","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AFLTULR","AF60","AF60-LR","WaveAP","WaveCPE","GBE-LR","GBE","GBE-Plus","GBE-AP","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apDevice2TypeModelPropEnum = append(apDevice2TypeModelPropEnum, v)
	}
}

const (

	// ApDevice2ModelUFDashNano captures enum value "UF-Nano"
	ApDevice2ModelUFDashNano string = "UF-Nano"

	// ApDevice2ModelUFDashLoco captures enum value "UF-Loco"
	ApDevice2ModelUFDashLoco string = "UF-Loco"

	// ApDevice2ModelUFDashWifi captures enum value "UF-Wifi"
	ApDevice2ModelUFDashWifi string = "UF-Wifi"

	// ApDevice2ModelUFDashInstant captures enum value "UF-Instant"
	ApDevice2ModelUFDashInstant string = "UF-Instant"

	// ApDevice2ModelUFDashOLT captures enum value "UF-OLT"
	ApDevice2ModelUFDashOLT string = "UF-OLT"

	// ApDevice2ModelUFDashOLT4 captures enum value "UF-OLT4"
	ApDevice2ModelUFDashOLT4 string = "UF-OLT4"

	// ApDevice2ModelUISPDashRDashPro captures enum value "UISP-R-Pro"
	ApDevice2ModelUISPDashRDashPro string = "UISP-R-Pro"

	// ApDevice2ModelUISPDashRDashLite captures enum value "UISP-R-Lite"
	ApDevice2ModelUISPDashRDashLite string = "UISP-R-Lite"

	// ApDevice2ModelUNMSDashSDashLite captures enum value "UNMS-S-Lite"
	ApDevice2ModelUNMSDashSDashLite string = "UNMS-S-Lite"

	// ApDevice2ModelUISPDashSDashLite captures enum value "UISP-S-Lite"
	ApDevice2ModelUISPDashSDashLite string = "UISP-S-Lite"

	// ApDevice2ModelUISPDashSDashPro captures enum value "UISP-S-Pro"
	ApDevice2ModelUISPDashSDashPro string = "UISP-S-Pro"

	// ApDevice2ModelUISPDashPDashLite captures enum value "UISP-P-Lite"
	ApDevice2ModelUISPDashPDashLite string = "UISP-P-Lite"

	// ApDevice2ModelUISPDashLTE captures enum value "UISP-LTE"
	ApDevice2ModelUISPDashLTE string = "UISP-LTE"

	// ApDevice2ModelERDashX captures enum value "ER-X"
	ApDevice2ModelERDashX string = "ER-X"

	// ApDevice2ModelERDashXDashSFP captures enum value "ER-X-SFP"
	ApDevice2ModelERDashXDashSFP string = "ER-X-SFP"

	// ApDevice2ModelERLiteDash3 captures enum value "ERLite-3"
	ApDevice2ModelERLiteDash3 string = "ERLite-3"

	// ApDevice2ModelERPoeDash5 captures enum value "ERPoe-5"
	ApDevice2ModelERPoeDash5 string = "ERPoe-5"

	// ApDevice2ModelERProDash8 captures enum value "ERPro-8"
	ApDevice2ModelERProDash8 string = "ERPro-8"

	// ApDevice2ModelERDash8 captures enum value "ER-8"
	ApDevice2ModelERDash8 string = "ER-8"

	// ApDevice2ModelERDash8DashXG captures enum value "ER-8-XG"
	ApDevice2ModelERDash8DashXG string = "ER-8-XG"

	// ApDevice2ModelERDash4 captures enum value "ER-4"
	ApDevice2ModelERDash4 string = "ER-4"

	// ApDevice2ModelERDash6P captures enum value "ER-6P"
	ApDevice2ModelERDash6P string = "ER-6P"

	// ApDevice2ModelERDash12 captures enum value "ER-12"
	ApDevice2ModelERDash12 string = "ER-12"

	// ApDevice2ModelERDash12P captures enum value "ER-12P"
	ApDevice2ModelERDash12P string = "ER-12P"

	// ApDevice2ModelERDash10X captures enum value "ER-10X"
	ApDevice2ModelERDash10X string = "ER-10X"

	// ApDevice2ModelEPDashR8 captures enum value "EP-R8"
	ApDevice2ModelEPDashR8 string = "EP-R8"

	// ApDevice2ModelEPDashR6 captures enum value "EP-R6"
	ApDevice2ModelEPDashR6 string = "EP-R6"

	// ApDevice2ModelEPDashS16 captures enum value "EP-S16"
	ApDevice2ModelEPDashS16 string = "EP-S16"

	// ApDevice2ModelESDash12F captures enum value "ES-12F"
	ApDevice2ModelESDash12F string = "ES-12F"

	// ApDevice2ModelESDash16Dash150W captures enum value "ES-16-150W"
	ApDevice2ModelESDash16Dash150W string = "ES-16-150W"

	// ApDevice2ModelESDash24Dash250W captures enum value "ES-24-250W"
	ApDevice2ModelESDash24Dash250W string = "ES-24-250W"

	// ApDevice2ModelESDash24Dash500W captures enum value "ES-24-500W"
	ApDevice2ModelESDash24Dash500W string = "ES-24-500W"

	// ApDevice2ModelESDash24DashLite captures enum value "ES-24-Lite"
	ApDevice2ModelESDash24DashLite string = "ES-24-Lite"

	// ApDevice2ModelESDash48Dash500W captures enum value "ES-48-500W"
	ApDevice2ModelESDash48Dash500W string = "ES-48-500W"

	// ApDevice2ModelESDash48Dash750W captures enum value "ES-48-750W"
	ApDevice2ModelESDash48Dash750W string = "ES-48-750W"

	// ApDevice2ModelESDash48DashLite captures enum value "ES-48-Lite"
	ApDevice2ModelESDash48DashLite string = "ES-48-Lite"

	// ApDevice2ModelESDash8Dash150W captures enum value "ES-8-150W"
	ApDevice2ModelESDash8Dash150W string = "ES-8-150W"

	// ApDevice2ModelESDash16DashXG captures enum value "ES-16-XG"
	ApDevice2ModelESDash16DashXG string = "ES-16-XG"

	// ApDevice2ModelESDash10XP captures enum value "ES-10XP"
	ApDevice2ModelESDash10XP string = "ES-10XP"

	// ApDevice2ModelESDash10X captures enum value "ES-10X"
	ApDevice2ModelESDash10X string = "ES-10X"

	// ApDevice2ModelESDash18X captures enum value "ES-18X"
	ApDevice2ModelESDash18X string = "ES-18X"

	// ApDevice2ModelESDash26X captures enum value "ES-26X"
	ApDevice2ModelESDash26X string = "ES-26X"

	// ApDevice2ModelEPDash54VDash150W captures enum value "EP-54V-150W"
	ApDevice2ModelEPDash54VDash150W string = "EP-54V-150W"

	// ApDevice2ModelEPDash24VDash72W captures enum value "EP-24V-72W"
	ApDevice2ModelEPDash24VDash72W string = "EP-24V-72W"

	// ApDevice2ModelEPDash54VDash72W captures enum value "EP-54V-72W"
	ApDevice2ModelEPDash54VDash72W string = "EP-54V-72W"

	// ApDevice2ModelTSWDashPoE captures enum value "TSW-PoE"
	ApDevice2ModelTSWDashPoE string = "TSW-PoE"

	// ApDevice2ModelTSWDashPoEPRO captures enum value "TSW-PoE PRO"
	ApDevice2ModelTSWDashPoEPRO string = "TSW-PoE PRO"

	// ApDevice2ModelACBDashAC captures enum value "ACB-AC"
	ApDevice2ModelACBDashAC string = "ACB-AC"

	// ApDevice2ModelACBDashISP captures enum value "ACB-ISP"
	ApDevice2ModelACBDashISP string = "ACB-ISP"

	// ApDevice2ModelACBDashLOCO captures enum value "ACB-LOCO"
	ApDevice2ModelACBDashLOCO string = "ACB-LOCO"

	// ApDevice2ModelAF11FX captures enum value "AF11FX"
	ApDevice2ModelAF11FX string = "AF11FX"

	// ApDevice2ModelAF24 captures enum value "AF24"
	ApDevice2ModelAF24 string = "AF24"

	// ApDevice2ModelAF24HD captures enum value "AF24HD"
	ApDevice2ModelAF24HD string = "AF24HD"

	// ApDevice2ModelAF2X captures enum value "AF2X"
	ApDevice2ModelAF2X string = "AF2X"

	// ApDevice2ModelAF3X captures enum value "AF3X"
	ApDevice2ModelAF3X string = "AF3X"

	// ApDevice2ModelAF4X captures enum value "AF4X"
	ApDevice2ModelAF4X string = "AF4X"

	// ApDevice2ModelAF5 captures enum value "AF5"
	ApDevice2ModelAF5 string = "AF5"

	// ApDevice2ModelAF5U captures enum value "AF5U"
	ApDevice2ModelAF5U string = "AF5U"

	// ApDevice2ModelAF5X captures enum value "AF5X"
	ApDevice2ModelAF5X string = "AF5X"

	// ApDevice2ModelAFDash5XHD captures enum value "AF-5XHD"
	ApDevice2ModelAFDash5XHD string = "AF-5XHD"

	// ApDevice2ModelAFDashLTU captures enum value "AF-LTU"
	ApDevice2ModelAFDashLTU string = "AF-LTU"

	// ApDevice2ModelLTUDashLITE captures enum value "LTU-LITE"
	ApDevice2ModelLTUDashLITE string = "LTU-LITE"

	// ApDevice2ModelAFDashLTU5 captures enum value "AF-LTU5"
	ApDevice2ModelAFDashLTU5 string = "AF-LTU5"

	// ApDevice2ModelLTUDashRocket captures enum value "LTU-Rocket"
	ApDevice2ModelLTUDashRocket string = "LTU-Rocket"

	// ApDevice2ModelAFLTULR captures enum value "AFLTULR"
	ApDevice2ModelAFLTULR string = "AFLTULR"

	// ApDevice2ModelAF60 captures enum value "AF60"
	ApDevice2ModelAF60 string = "AF60"

	// ApDevice2ModelAF60DashLR captures enum value "AF60-LR"
	ApDevice2ModelAF60DashLR string = "AF60-LR"

	// ApDevice2ModelWaveAP captures enum value "WaveAP"
	ApDevice2ModelWaveAP string = "WaveAP"

	// ApDevice2ModelWaveCPE captures enum value "WaveCPE"
	ApDevice2ModelWaveCPE string = "WaveCPE"

	// ApDevice2ModelGBEDashLR captures enum value "GBE-LR"
	ApDevice2ModelGBEDashLR string = "GBE-LR"

	// ApDevice2ModelGBE captures enum value "GBE"
	ApDevice2ModelGBE string = "GBE"

	// ApDevice2ModelGBEDashPlus captures enum value "GBE-Plus"
	ApDevice2ModelGBEDashPlus string = "GBE-Plus"

	// ApDevice2ModelGBEDashAP captures enum value "GBE-AP"
	ApDevice2ModelGBEDashAP string = "GBE-AP"

	// ApDevice2ModelR2N captures enum value "R2N"
	ApDevice2ModelR2N string = "R2N"

	// ApDevice2ModelR2T captures enum value "R2T"
	ApDevice2ModelR2T string = "R2T"

	// ApDevice2ModelR5N captures enum value "R5N"
	ApDevice2ModelR5N string = "R5N"

	// ApDevice2ModelR6N captures enum value "R6N"
	ApDevice2ModelR6N string = "R6N"

	// ApDevice2ModelR36DashGPS captures enum value "R36-GPS"
	ApDevice2ModelR36DashGPS string = "R36-GPS"

	// ApDevice2ModelRM3DashGPS captures enum value "RM3-GPS"
	ApDevice2ModelRM3DashGPS string = "RM3-GPS"

	// ApDevice2ModelR2NDashGPS captures enum value "R2N-GPS"
	ApDevice2ModelR2NDashGPS string = "R2N-GPS"

	// ApDevice2ModelR5NDashGPS captures enum value "R5N-GPS"
	ApDevice2ModelR5NDashGPS string = "R5N-GPS"

	// ApDevice2ModelR9NDashGPS captures enum value "R9N-GPS"
	ApDevice2ModelR9NDashGPS string = "R9N-GPS"

	// ApDevice2ModelR5TDashGPS captures enum value "R5T-GPS"
	ApDevice2ModelR5TDashGPS string = "R5T-GPS"

	// ApDevice2ModelRM3 captures enum value "RM3"
	ApDevice2ModelRM3 string = "RM3"

	// ApDevice2ModelR36 captures enum value "R36"
	ApDevice2ModelR36 string = "R36"

	// ApDevice2ModelR9N captures enum value "R9N"
	ApDevice2ModelR9N string = "R9N"

	// ApDevice2ModelN2N captures enum value "N2N"
	ApDevice2ModelN2N string = "N2N"

	// ApDevice2ModelN5N captures enum value "N5N"
	ApDevice2ModelN5N string = "N5N"

	// ApDevice2ModelN6N captures enum value "N6N"
	ApDevice2ModelN6N string = "N6N"

	// ApDevice2ModelNS3 captures enum value "NS3"
	ApDevice2ModelNS3 string = "NS3"

	// ApDevice2ModelN36 captures enum value "N36"
	ApDevice2ModelN36 string = "N36"

	// ApDevice2ModelN9N captures enum value "N9N"
	ApDevice2ModelN9N string = "N9N"

	// ApDevice2ModelN9S captures enum value "N9S"
	ApDevice2ModelN9S string = "N9S"

	// ApDevice2ModelLM2 captures enum value "LM2"
	ApDevice2ModelLM2 string = "LM2"

	// ApDevice2ModelLM5 captures enum value "LM5"
	ApDevice2ModelLM5 string = "LM5"

	// ApDevice2ModelB2N captures enum value "B2N"
	ApDevice2ModelB2N string = "B2N"

	// ApDevice2ModelB2T captures enum value "B2T"
	ApDevice2ModelB2T string = "B2T"

	// ApDevice2ModelB5N captures enum value "B5N"
	ApDevice2ModelB5N string = "B5N"

	// ApDevice2ModelB5T captures enum value "B5T"
	ApDevice2ModelB5T string = "B5T"

	// ApDevice2ModelBAC captures enum value "BAC"
	ApDevice2ModelBAC string = "BAC"

	// ApDevice2ModelAG2 captures enum value "AG2"
	ApDevice2ModelAG2 string = "AG2"

	// ApDevice2ModelAG2DashHP captures enum value "AG2-HP"
	ApDevice2ModelAG2DashHP string = "AG2-HP"

	// ApDevice2ModelAG5 captures enum value "AG5"
	ApDevice2ModelAG5 string = "AG5"

	// ApDevice2ModelAG5DashHP captures enum value "AG5-HP"
	ApDevice2ModelAG5DashHP string = "AG5-HP"

	// ApDevice2ModelP2N captures enum value "p2N"
	ApDevice2ModelP2N string = "p2N"

	// ApDevice2ModelP5N captures enum value "p5N"
	ApDevice2ModelP5N string = "p5N"

	// ApDevice2ModelM25 captures enum value "M25"
	ApDevice2ModelM25 string = "M25"

	// ApDevice2ModelP2BDash400 captures enum value "P2B-400"
	ApDevice2ModelP2BDash400 string = "P2B-400"

	// ApDevice2ModelP5BDash300 captures enum value "P5B-300"
	ApDevice2ModelP5BDash300 string = "P5B-300"

	// ApDevice2ModelP5BDash300DashISO captures enum value "P5B-300-ISO"
	ApDevice2ModelP5BDash300DashISO string = "P5B-300-ISO"

	// ApDevice2ModelP5BDash400 captures enum value "P5B-400"
	ApDevice2ModelP5BDash400 string = "P5B-400"

	// ApDevice2ModelP5BDash400DashISO captures enum value "P5B-400-ISO"
	ApDevice2ModelP5BDash400DashISO string = "P5B-400-ISO"

	// ApDevice2ModelP5BDash620 captures enum value "P5B-620"
	ApDevice2ModelP5BDash620 string = "P5B-620"

	// ApDevice2ModelLB5Dash120 captures enum value "LB5-120"
	ApDevice2ModelLB5Dash120 string = "LB5-120"

	// ApDevice2ModelLB5 captures enum value "LB5"
	ApDevice2ModelLB5 string = "LB5"

	// ApDevice2ModelN5B captures enum value "N5B"
	ApDevice2ModelN5B string = "N5B"

	// ApDevice2ModelN5BDash16 captures enum value "N5B-16"
	ApDevice2ModelN5BDash16 string = "N5B-16"

	// ApDevice2ModelN5BDash19 captures enum value "N5B-19"
	ApDevice2ModelN5BDash19 string = "N5B-19"

	// ApDevice2ModelN5BDash300 captures enum value "N5B-300"
	ApDevice2ModelN5BDash300 string = "N5B-300"

	// ApDevice2ModelN5BDash400 captures enum value "N5B-400"
	ApDevice2ModelN5BDash400 string = "N5B-400"

	// ApDevice2ModelN5BDashClient captures enum value "N5B-Client"
	ApDevice2ModelN5BDashClient string = "N5B-Client"

	// ApDevice2ModelN2B captures enum value "N2B"
	ApDevice2ModelN2B string = "N2B"

	// ApDevice2ModelN2BDash13 captures enum value "N2B-13"
	ApDevice2ModelN2BDash13 string = "N2B-13"

	// ApDevice2ModelN2BDash400 captures enum value "N2B-400"
	ApDevice2ModelN2BDash400 string = "N2B-400"

	// ApDevice2ModelPAP captures enum value "PAP"
	ApDevice2ModelPAP string = "PAP"

	// ApDevice2ModelLAPDashHP captures enum value "LAP-HP"
	ApDevice2ModelLAPDashHP string = "LAP-HP"

	// ApDevice2ModelLAP captures enum value "LAP"
	ApDevice2ModelLAP string = "LAP"

	// ApDevice2ModelAGW captures enum value "AGW"
	ApDevice2ModelAGW string = "AGW"

	// ApDevice2ModelAGWDashLR captures enum value "AGW-LR"
	ApDevice2ModelAGWDashLR string = "AGW-LR"

	// ApDevice2ModelAGWDashPro captures enum value "AGW-Pro"
	ApDevice2ModelAGWDashPro string = "AGW-Pro"

	// ApDevice2ModelAGWDashInstaller captures enum value "AGW-Installer"
	ApDevice2ModelAGWDashInstaller string = "AGW-Installer"

	// ApDevice2ModelPB5 captures enum value "PB5"
	ApDevice2ModelPB5 string = "PB5"

	// ApDevice2ModelPB3 captures enum value "PB3"
	ApDevice2ModelPB3 string = "PB3"

	// ApDevice2ModelP36 captures enum value "P36"
	ApDevice2ModelP36 string = "P36"

	// ApDevice2ModelPBM10 captures enum value "PBM10"
	ApDevice2ModelPBM10 string = "PBM10"

	// ApDevice2ModelNB5 captures enum value "NB5"
	ApDevice2ModelNB5 string = "NB5"

	// ApDevice2ModelNB2 captures enum value "NB2"
	ApDevice2ModelNB2 string = "NB2"

	// ApDevice2ModelNB3 captures enum value "NB3"
	ApDevice2ModelNB3 string = "NB3"

	// ApDevice2ModelB36 captures enum value "B36"
	ApDevice2ModelB36 string = "B36"

	// ApDevice2ModelNB9 captures enum value "NB9"
	ApDevice2ModelNB9 string = "NB9"

	// ApDevice2ModelSM5 captures enum value "SM5"
	ApDevice2ModelSM5 string = "SM5"

	// ApDevice2ModelWM5 captures enum value "WM5"
	ApDevice2ModelWM5 string = "WM5"

	// ApDevice2ModelISDashM5 captures enum value "IS-M5"
	ApDevice2ModelISDashM5 string = "IS-M5"

	// ApDevice2ModelLoco5AC captures enum value "Loco5AC"
	ApDevice2ModelLoco5AC string = "Loco5AC"

	// ApDevice2ModelNSDash5AC captures enum value "NS-5AC"
	ApDevice2ModelNSDash5AC string = "NS-5AC"

	// ApDevice2ModelR5ACDashPTMP captures enum value "R5AC-PTMP"
	ApDevice2ModelR5ACDashPTMP string = "R5AC-PTMP"

	// ApDevice2ModelR5ACDashPTP captures enum value "R5AC-PTP"
	ApDevice2ModelR5ACDashPTP string = "R5AC-PTP"

	// ApDevice2ModelR5ACDashLite captures enum value "R5AC-Lite"
	ApDevice2ModelR5ACDashLite string = "R5AC-Lite"

	// ApDevice2ModelR5ACDashPRISM captures enum value "R5AC-PRISM"
	ApDevice2ModelR5ACDashPRISM string = "R5AC-PRISM"

	// ApDevice2ModelR2ACDashPrism captures enum value "R2AC-Prism"
	ApDevice2ModelR2ACDashPrism string = "R2AC-Prism"

	// ApDevice2ModelR2ACDashGen2 captures enum value "R2AC-Gen2"
	ApDevice2ModelR2ACDashGen2 string = "R2AC-Gen2"

	// ApDevice2ModelRPDash5ACDashGen2 captures enum value "RP-5AC-Gen2"
	ApDevice2ModelRPDash5ACDashGen2 string = "RP-5AC-Gen2"

	// ApDevice2ModelNBEDash2ACDash13 captures enum value "NBE-2AC-13"
	ApDevice2ModelNBEDash2ACDash13 string = "NBE-2AC-13"

	// ApDevice2ModelNBEDash5ACDash16 captures enum value "NBE-5AC-16"
	ApDevice2ModelNBEDash5ACDash16 string = "NBE-5AC-16"

	// ApDevice2ModelNBEDash5ACDash19 captures enum value "NBE-5AC-19"
	ApDevice2ModelNBEDash5ACDash19 string = "NBE-5AC-19"

	// ApDevice2ModelNBEDash5ACDashGen2 captures enum value "NBE-5AC-Gen2"
	ApDevice2ModelNBEDash5ACDashGen2 string = "NBE-5AC-Gen2"

	// ApDevice2ModelPBEDash5ACDash300 captures enum value "PBE-5AC-300"
	ApDevice2ModelPBEDash5ACDash300 string = "PBE-5AC-300"

	// ApDevice2ModelPBEDash5ACDash300DashISO captures enum value "PBE-5AC-300-ISO"
	ApDevice2ModelPBEDash5ACDash300DashISO string = "PBE-5AC-300-ISO"

	// ApDevice2ModelPBEDash5ACDash400 captures enum value "PBE-5AC-400"
	ApDevice2ModelPBEDash5ACDash400 string = "PBE-5AC-400"

	// ApDevice2ModelPBEDash5ACDash400DashISO captures enum value "PBE-5AC-400-ISO"
	ApDevice2ModelPBEDash5ACDash400DashISO string = "PBE-5AC-400-ISO"

	// ApDevice2ModelPBEDash5ACDash500 captures enum value "PBE-5AC-500"
	ApDevice2ModelPBEDash5ACDash500 string = "PBE-5AC-500"

	// ApDevice2ModelPBEDash5ACDash500DashISO captures enum value "PBE-5AC-500-ISO"
	ApDevice2ModelPBEDash5ACDash500DashISO string = "PBE-5AC-500-ISO"

	// ApDevice2ModelPBEDash5ACDash620 captures enum value "PBE-5AC-620"
	ApDevice2ModelPBEDash5ACDash620 string = "PBE-5AC-620"

	// ApDevice2ModelPBEDash5ACDash620DashISO captures enum value "PBE-5AC-620-ISO"
	ApDevice2ModelPBEDash5ACDash620DashISO string = "PBE-5AC-620-ISO"

	// ApDevice2ModelPBEDash2ACDash400 captures enum value "PBE-2AC-400"
	ApDevice2ModelPBEDash2ACDash400 string = "PBE-2AC-400"

	// ApDevice2ModelPBEDash2ACDash400DashISO captures enum value "PBE-2AC-400-ISO"
	ApDevice2ModelPBEDash2ACDash400DashISO string = "PBE-2AC-400-ISO"

	// ApDevice2ModelPBEDash5ACDashXDashGen2 captures enum value "PBE-5AC-X-Gen2"
	ApDevice2ModelPBEDash5ACDashXDashGen2 string = "PBE-5AC-X-Gen2"

	// ApDevice2ModelPBEDash5ACDashGen2 captures enum value "PBE-5AC-Gen2"
	ApDevice2ModelPBEDash5ACDashGen2 string = "PBE-5AC-Gen2"

	// ApDevice2ModelPBEDash5ACDashISODashGen2 captures enum value "PBE-5AC-ISO-Gen2"
	ApDevice2ModelPBEDash5ACDashISODashGen2 string = "PBE-5AC-ISO-Gen2"

	// ApDevice2ModelPBEDash5ACDash400DashISODashGen2 captures enum value "PBE-5AC-400-ISO-Gen2"
	ApDevice2ModelPBEDash5ACDash400DashISODashGen2 string = "PBE-5AC-400-ISO-Gen2"

	// ApDevice2ModelLBEDash5ACDash16Dash120 captures enum value "LBE-5AC-16-120"
	ApDevice2ModelLBEDash5ACDash16Dash120 string = "LBE-5AC-16-120"

	// ApDevice2ModelLAPDash120 captures enum value "LAP-120"
	ApDevice2ModelLAPDash120 string = "LAP-120"

	// ApDevice2ModelLBEDash5ACDash23 captures enum value "LBE-5AC-23"
	ApDevice2ModelLBEDash5ACDash23 string = "LBE-5AC-23"

	// ApDevice2ModelLBEDash5ACDashGen2 captures enum value "LBE-5AC-Gen2"
	ApDevice2ModelLBEDash5ACDashGen2 string = "LBE-5AC-Gen2"

	// ApDevice2ModelLBEDash5ACDashLR captures enum value "LBE-5AC-LR"
	ApDevice2ModelLBEDash5ACDashLR string = "LBE-5AC-LR"

	// ApDevice2ModelLAPDashGPS captures enum value "LAP-GPS"
	ApDevice2ModelLAPDashGPS string = "LAP-GPS"

	// ApDevice2ModelISDash5AC captures enum value "IS-5AC"
	ApDevice2ModelISDash5AC string = "IS-5AC"

	// ApDevice2ModelPSDash5AC captures enum value "PS-5AC"
	ApDevice2ModelPSDash5AC string = "PS-5AC"

	// ApDevice2ModelSolarSwitch captures enum value "SolarSwitch"
	ApDevice2ModelSolarSwitch string = "SolarSwitch"

	// ApDevice2ModelSolarPoint captures enum value "SolarPoint"
	ApDevice2ModelSolarPoint string = "SolarPoint"

	// ApDevice2ModelBulletACDashIP67 captures enum value "BulletAC-IP67"
	ApDevice2ModelBulletACDashIP67 string = "BulletAC-IP67"

	// ApDevice2ModelBDashDBDashAC captures enum value "B-DB-AC"
	ApDevice2ModelBDashDBDashAC string = "B-DB-AC"

	// ApDevice2ModelUNKNOWN captures enum value "UNKNOWN"
	ApDevice2ModelUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *ApDevice2) validateModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apDevice2TypeModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApDevice2) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	// value enum
	if err := m.validateModelEnum("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ap device 2 based on context it is used
func (m *ApDevice2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApDevice2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApDevice2) UnmarshalBinary(b []byte) error {
	var res ApDevice2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
