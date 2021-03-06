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

// VisibleBy visible by
//
// swagger:model visibleBy
type VisibleBy struct {

	// id
	// Example: f7ac9cad-ea28-4390-93c8-7add010e8ee3
	// Required: true
	ID *string `json:"id"`

	// Short names, for example UF-OLT.
	// Required: true
	// Enum: [UF-Nano UF-Loco UF-Wifi UF-Instant UF-OLT UF-OLT4 UISP-R-Pro UISP-R-Lite UNMS-S-Lite UISP-S-Lite UISP-S-Pro UISP-P-Lite UISP-LTE ER-X ER-X-SFP ERLite-3 ERPoe-5 ERPro-8 ER-8 ER-8-XG ER-4 ER-6P ER-12 ER-12P ER-10X EP-R8 EP-R6 EP-S16 ES-12F ES-16-150W ES-24-250W ES-24-500W ES-24-Lite ES-48-500W ES-48-750W ES-48-Lite ES-8-150W ES-16-XG ES-10XP ES-10X ES-18X ES-26X EP-54V-150W EP-24V-72W EP-54V-72W TSW-PoE TSW-PoE PRO ACB-AC ACB-ISP ACB-LOCO AF11FX AF24 AF24HD AF2X AF3X AF4X AF5 AF5U AF5X AF-5XHD AF-LTU LTU-LITE AF-LTU5 LTU-Rocket AFLTULR AF60 AF60-LR WaveAP WaveCPE GBE-LR GBE GBE-Plus GBE-AP R2N R2T R5N R6N R36-GPS RM3-GPS R2N-GPS R5N-GPS R9N-GPS R5T-GPS RM3 R36 R9N N2N N5N N6N NS3 N36 N9N N9S LM2 LM5 B2N B2T B5N B5T BAC AG2 AG2-HP AG5 AG5-HP p2N p5N M25 P2B-400 P5B-300 P5B-300-ISO P5B-400 P5B-400-ISO P5B-620 LB5-120 LB5 N5B N5B-16 N5B-19 N5B-300 N5B-400 N5B-Client N2B N2B-13 N2B-400 PAP LAP-HP LAP AGW AGW-LR AGW-Pro AGW-Installer PB5 PB3 P36 PBM10 NB5 NB2 NB3 B36 NB9 SM5 WM5 IS-M5 Loco5AC NS-5AC R5AC-PTMP R5AC-PTP R5AC-Lite R5AC-PRISM R2AC-Prism R2AC-Gen2 RP-5AC-Gen2 NBE-2AC-13 NBE-5AC-16 NBE-5AC-19 NBE-5AC-Gen2 PBE-5AC-300 PBE-5AC-300-ISO PBE-5AC-400 PBE-5AC-400-ISO PBE-5AC-500 PBE-5AC-500-ISO PBE-5AC-620 PBE-5AC-620-ISO PBE-2AC-400 PBE-2AC-400-ISO PBE-5AC-X-Gen2 PBE-5AC-Gen2 PBE-5AC-ISO-Gen2 PBE-5AC-400-ISO-Gen2 LBE-5AC-16-120 LAP-120 LBE-5AC-23 LBE-5AC-Gen2 LBE-5AC-LR LAP-GPS IS-5AC PS-5AC SolarSwitch SolarPoint BulletAC-IP67 B-DB-AC UNKNOWN]
	Model *string `json:"model"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Device type
	// Example: erouter
	// Required: true
	// Enum: [onu olt uispp uispr uisps uispLte erouter eswitch epower airCube airMax airFiber toughSwitch solarBeam wave blackBox]
	Type *string `json:"type"`
}

// Validate validates this visible by
func (m *VisibleBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VisibleBy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var visibleByTypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UF-Nano","UF-Loco","UF-Wifi","UF-Instant","UF-OLT","UF-OLT4","UISP-R-Pro","UISP-R-Lite","UNMS-S-Lite","UISP-S-Lite","UISP-S-Pro","UISP-P-Lite","UISP-LTE","ER-X","ER-X-SFP","ERLite-3","ERPoe-5","ERPro-8","ER-8","ER-8-XG","ER-4","ER-6P","ER-12","ER-12P","ER-10X","EP-R8","EP-R6","EP-S16","ES-12F","ES-16-150W","ES-24-250W","ES-24-500W","ES-24-Lite","ES-48-500W","ES-48-750W","ES-48-Lite","ES-8-150W","ES-16-XG","ES-10XP","ES-10X","ES-18X","ES-26X","EP-54V-150W","EP-24V-72W","EP-54V-72W","TSW-PoE","TSW-PoE PRO","ACB-AC","ACB-ISP","ACB-LOCO","AF11FX","AF24","AF24HD","AF2X","AF3X","AF4X","AF5","AF5U","AF5X","AF-5XHD","AF-LTU","LTU-LITE","AF-LTU5","LTU-Rocket","AFLTULR","AF60","AF60-LR","WaveAP","WaveCPE","GBE-LR","GBE","GBE-Plus","GBE-AP","R2N","R2T","R5N","R6N","R36-GPS","RM3-GPS","R2N-GPS","R5N-GPS","R9N-GPS","R5T-GPS","RM3","R36","R9N","N2N","N5N","N6N","NS3","N36","N9N","N9S","LM2","LM5","B2N","B2T","B5N","B5T","BAC","AG2","AG2-HP","AG5","AG5-HP","p2N","p5N","M25","P2B-400","P5B-300","P5B-300-ISO","P5B-400","P5B-400-ISO","P5B-620","LB5-120","LB5","N5B","N5B-16","N5B-19","N5B-300","N5B-400","N5B-Client","N2B","N2B-13","N2B-400","PAP","LAP-HP","LAP","AGW","AGW-LR","AGW-Pro","AGW-Installer","PB5","PB3","P36","PBM10","NB5","NB2","NB3","B36","NB9","SM5","WM5","IS-M5","Loco5AC","NS-5AC","R5AC-PTMP","R5AC-PTP","R5AC-Lite","R5AC-PRISM","R2AC-Prism","R2AC-Gen2","RP-5AC-Gen2","NBE-2AC-13","NBE-5AC-16","NBE-5AC-19","NBE-5AC-Gen2","PBE-5AC-300","PBE-5AC-300-ISO","PBE-5AC-400","PBE-5AC-400-ISO","PBE-5AC-500","PBE-5AC-500-ISO","PBE-5AC-620","PBE-5AC-620-ISO","PBE-2AC-400","PBE-2AC-400-ISO","PBE-5AC-X-Gen2","PBE-5AC-Gen2","PBE-5AC-ISO-Gen2","PBE-5AC-400-ISO-Gen2","LBE-5AC-16-120","LAP-120","LBE-5AC-23","LBE-5AC-Gen2","LBE-5AC-LR","LAP-GPS","IS-5AC","PS-5AC","SolarSwitch","SolarPoint","BulletAC-IP67","B-DB-AC","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visibleByTypeModelPropEnum = append(visibleByTypeModelPropEnum, v)
	}
}

const (

	// VisibleByModelUFDashNano captures enum value "UF-Nano"
	VisibleByModelUFDashNano string = "UF-Nano"

	// VisibleByModelUFDashLoco captures enum value "UF-Loco"
	VisibleByModelUFDashLoco string = "UF-Loco"

	// VisibleByModelUFDashWifi captures enum value "UF-Wifi"
	VisibleByModelUFDashWifi string = "UF-Wifi"

	// VisibleByModelUFDashInstant captures enum value "UF-Instant"
	VisibleByModelUFDashInstant string = "UF-Instant"

	// VisibleByModelUFDashOLT captures enum value "UF-OLT"
	VisibleByModelUFDashOLT string = "UF-OLT"

	// VisibleByModelUFDashOLT4 captures enum value "UF-OLT4"
	VisibleByModelUFDashOLT4 string = "UF-OLT4"

	// VisibleByModelUISPDashRDashPro captures enum value "UISP-R-Pro"
	VisibleByModelUISPDashRDashPro string = "UISP-R-Pro"

	// VisibleByModelUISPDashRDashLite captures enum value "UISP-R-Lite"
	VisibleByModelUISPDashRDashLite string = "UISP-R-Lite"

	// VisibleByModelUNMSDashSDashLite captures enum value "UNMS-S-Lite"
	VisibleByModelUNMSDashSDashLite string = "UNMS-S-Lite"

	// VisibleByModelUISPDashSDashLite captures enum value "UISP-S-Lite"
	VisibleByModelUISPDashSDashLite string = "UISP-S-Lite"

	// VisibleByModelUISPDashSDashPro captures enum value "UISP-S-Pro"
	VisibleByModelUISPDashSDashPro string = "UISP-S-Pro"

	// VisibleByModelUISPDashPDashLite captures enum value "UISP-P-Lite"
	VisibleByModelUISPDashPDashLite string = "UISP-P-Lite"

	// VisibleByModelUISPDashLTE captures enum value "UISP-LTE"
	VisibleByModelUISPDashLTE string = "UISP-LTE"

	// VisibleByModelERDashX captures enum value "ER-X"
	VisibleByModelERDashX string = "ER-X"

	// VisibleByModelERDashXDashSFP captures enum value "ER-X-SFP"
	VisibleByModelERDashXDashSFP string = "ER-X-SFP"

	// VisibleByModelERLiteDash3 captures enum value "ERLite-3"
	VisibleByModelERLiteDash3 string = "ERLite-3"

	// VisibleByModelERPoeDash5 captures enum value "ERPoe-5"
	VisibleByModelERPoeDash5 string = "ERPoe-5"

	// VisibleByModelERProDash8 captures enum value "ERPro-8"
	VisibleByModelERProDash8 string = "ERPro-8"

	// VisibleByModelERDash8 captures enum value "ER-8"
	VisibleByModelERDash8 string = "ER-8"

	// VisibleByModelERDash8DashXG captures enum value "ER-8-XG"
	VisibleByModelERDash8DashXG string = "ER-8-XG"

	// VisibleByModelERDash4 captures enum value "ER-4"
	VisibleByModelERDash4 string = "ER-4"

	// VisibleByModelERDash6P captures enum value "ER-6P"
	VisibleByModelERDash6P string = "ER-6P"

	// VisibleByModelERDash12 captures enum value "ER-12"
	VisibleByModelERDash12 string = "ER-12"

	// VisibleByModelERDash12P captures enum value "ER-12P"
	VisibleByModelERDash12P string = "ER-12P"

	// VisibleByModelERDash10X captures enum value "ER-10X"
	VisibleByModelERDash10X string = "ER-10X"

	// VisibleByModelEPDashR8 captures enum value "EP-R8"
	VisibleByModelEPDashR8 string = "EP-R8"

	// VisibleByModelEPDashR6 captures enum value "EP-R6"
	VisibleByModelEPDashR6 string = "EP-R6"

	// VisibleByModelEPDashS16 captures enum value "EP-S16"
	VisibleByModelEPDashS16 string = "EP-S16"

	// VisibleByModelESDash12F captures enum value "ES-12F"
	VisibleByModelESDash12F string = "ES-12F"

	// VisibleByModelESDash16Dash150W captures enum value "ES-16-150W"
	VisibleByModelESDash16Dash150W string = "ES-16-150W"

	// VisibleByModelESDash24Dash250W captures enum value "ES-24-250W"
	VisibleByModelESDash24Dash250W string = "ES-24-250W"

	// VisibleByModelESDash24Dash500W captures enum value "ES-24-500W"
	VisibleByModelESDash24Dash500W string = "ES-24-500W"

	// VisibleByModelESDash24DashLite captures enum value "ES-24-Lite"
	VisibleByModelESDash24DashLite string = "ES-24-Lite"

	// VisibleByModelESDash48Dash500W captures enum value "ES-48-500W"
	VisibleByModelESDash48Dash500W string = "ES-48-500W"

	// VisibleByModelESDash48Dash750W captures enum value "ES-48-750W"
	VisibleByModelESDash48Dash750W string = "ES-48-750W"

	// VisibleByModelESDash48DashLite captures enum value "ES-48-Lite"
	VisibleByModelESDash48DashLite string = "ES-48-Lite"

	// VisibleByModelESDash8Dash150W captures enum value "ES-8-150W"
	VisibleByModelESDash8Dash150W string = "ES-8-150W"

	// VisibleByModelESDash16DashXG captures enum value "ES-16-XG"
	VisibleByModelESDash16DashXG string = "ES-16-XG"

	// VisibleByModelESDash10XP captures enum value "ES-10XP"
	VisibleByModelESDash10XP string = "ES-10XP"

	// VisibleByModelESDash10X captures enum value "ES-10X"
	VisibleByModelESDash10X string = "ES-10X"

	// VisibleByModelESDash18X captures enum value "ES-18X"
	VisibleByModelESDash18X string = "ES-18X"

	// VisibleByModelESDash26X captures enum value "ES-26X"
	VisibleByModelESDash26X string = "ES-26X"

	// VisibleByModelEPDash54VDash150W captures enum value "EP-54V-150W"
	VisibleByModelEPDash54VDash150W string = "EP-54V-150W"

	// VisibleByModelEPDash24VDash72W captures enum value "EP-24V-72W"
	VisibleByModelEPDash24VDash72W string = "EP-24V-72W"

	// VisibleByModelEPDash54VDash72W captures enum value "EP-54V-72W"
	VisibleByModelEPDash54VDash72W string = "EP-54V-72W"

	// VisibleByModelTSWDashPoE captures enum value "TSW-PoE"
	VisibleByModelTSWDashPoE string = "TSW-PoE"

	// VisibleByModelTSWDashPoEPRO captures enum value "TSW-PoE PRO"
	VisibleByModelTSWDashPoEPRO string = "TSW-PoE PRO"

	// VisibleByModelACBDashAC captures enum value "ACB-AC"
	VisibleByModelACBDashAC string = "ACB-AC"

	// VisibleByModelACBDashISP captures enum value "ACB-ISP"
	VisibleByModelACBDashISP string = "ACB-ISP"

	// VisibleByModelACBDashLOCO captures enum value "ACB-LOCO"
	VisibleByModelACBDashLOCO string = "ACB-LOCO"

	// VisibleByModelAF11FX captures enum value "AF11FX"
	VisibleByModelAF11FX string = "AF11FX"

	// VisibleByModelAF24 captures enum value "AF24"
	VisibleByModelAF24 string = "AF24"

	// VisibleByModelAF24HD captures enum value "AF24HD"
	VisibleByModelAF24HD string = "AF24HD"

	// VisibleByModelAF2X captures enum value "AF2X"
	VisibleByModelAF2X string = "AF2X"

	// VisibleByModelAF3X captures enum value "AF3X"
	VisibleByModelAF3X string = "AF3X"

	// VisibleByModelAF4X captures enum value "AF4X"
	VisibleByModelAF4X string = "AF4X"

	// VisibleByModelAF5 captures enum value "AF5"
	VisibleByModelAF5 string = "AF5"

	// VisibleByModelAF5U captures enum value "AF5U"
	VisibleByModelAF5U string = "AF5U"

	// VisibleByModelAF5X captures enum value "AF5X"
	VisibleByModelAF5X string = "AF5X"

	// VisibleByModelAFDash5XHD captures enum value "AF-5XHD"
	VisibleByModelAFDash5XHD string = "AF-5XHD"

	// VisibleByModelAFDashLTU captures enum value "AF-LTU"
	VisibleByModelAFDashLTU string = "AF-LTU"

	// VisibleByModelLTUDashLITE captures enum value "LTU-LITE"
	VisibleByModelLTUDashLITE string = "LTU-LITE"

	// VisibleByModelAFDashLTU5 captures enum value "AF-LTU5"
	VisibleByModelAFDashLTU5 string = "AF-LTU5"

	// VisibleByModelLTUDashRocket captures enum value "LTU-Rocket"
	VisibleByModelLTUDashRocket string = "LTU-Rocket"

	// VisibleByModelAFLTULR captures enum value "AFLTULR"
	VisibleByModelAFLTULR string = "AFLTULR"

	// VisibleByModelAF60 captures enum value "AF60"
	VisibleByModelAF60 string = "AF60"

	// VisibleByModelAF60DashLR captures enum value "AF60-LR"
	VisibleByModelAF60DashLR string = "AF60-LR"

	// VisibleByModelWaveAP captures enum value "WaveAP"
	VisibleByModelWaveAP string = "WaveAP"

	// VisibleByModelWaveCPE captures enum value "WaveCPE"
	VisibleByModelWaveCPE string = "WaveCPE"

	// VisibleByModelGBEDashLR captures enum value "GBE-LR"
	VisibleByModelGBEDashLR string = "GBE-LR"

	// VisibleByModelGBE captures enum value "GBE"
	VisibleByModelGBE string = "GBE"

	// VisibleByModelGBEDashPlus captures enum value "GBE-Plus"
	VisibleByModelGBEDashPlus string = "GBE-Plus"

	// VisibleByModelGBEDashAP captures enum value "GBE-AP"
	VisibleByModelGBEDashAP string = "GBE-AP"

	// VisibleByModelR2N captures enum value "R2N"
	VisibleByModelR2N string = "R2N"

	// VisibleByModelR2T captures enum value "R2T"
	VisibleByModelR2T string = "R2T"

	// VisibleByModelR5N captures enum value "R5N"
	VisibleByModelR5N string = "R5N"

	// VisibleByModelR6N captures enum value "R6N"
	VisibleByModelR6N string = "R6N"

	// VisibleByModelR36DashGPS captures enum value "R36-GPS"
	VisibleByModelR36DashGPS string = "R36-GPS"

	// VisibleByModelRM3DashGPS captures enum value "RM3-GPS"
	VisibleByModelRM3DashGPS string = "RM3-GPS"

	// VisibleByModelR2NDashGPS captures enum value "R2N-GPS"
	VisibleByModelR2NDashGPS string = "R2N-GPS"

	// VisibleByModelR5NDashGPS captures enum value "R5N-GPS"
	VisibleByModelR5NDashGPS string = "R5N-GPS"

	// VisibleByModelR9NDashGPS captures enum value "R9N-GPS"
	VisibleByModelR9NDashGPS string = "R9N-GPS"

	// VisibleByModelR5TDashGPS captures enum value "R5T-GPS"
	VisibleByModelR5TDashGPS string = "R5T-GPS"

	// VisibleByModelRM3 captures enum value "RM3"
	VisibleByModelRM3 string = "RM3"

	// VisibleByModelR36 captures enum value "R36"
	VisibleByModelR36 string = "R36"

	// VisibleByModelR9N captures enum value "R9N"
	VisibleByModelR9N string = "R9N"

	// VisibleByModelN2N captures enum value "N2N"
	VisibleByModelN2N string = "N2N"

	// VisibleByModelN5N captures enum value "N5N"
	VisibleByModelN5N string = "N5N"

	// VisibleByModelN6N captures enum value "N6N"
	VisibleByModelN6N string = "N6N"

	// VisibleByModelNS3 captures enum value "NS3"
	VisibleByModelNS3 string = "NS3"

	// VisibleByModelN36 captures enum value "N36"
	VisibleByModelN36 string = "N36"

	// VisibleByModelN9N captures enum value "N9N"
	VisibleByModelN9N string = "N9N"

	// VisibleByModelN9S captures enum value "N9S"
	VisibleByModelN9S string = "N9S"

	// VisibleByModelLM2 captures enum value "LM2"
	VisibleByModelLM2 string = "LM2"

	// VisibleByModelLM5 captures enum value "LM5"
	VisibleByModelLM5 string = "LM5"

	// VisibleByModelB2N captures enum value "B2N"
	VisibleByModelB2N string = "B2N"

	// VisibleByModelB2T captures enum value "B2T"
	VisibleByModelB2T string = "B2T"

	// VisibleByModelB5N captures enum value "B5N"
	VisibleByModelB5N string = "B5N"

	// VisibleByModelB5T captures enum value "B5T"
	VisibleByModelB5T string = "B5T"

	// VisibleByModelBAC captures enum value "BAC"
	VisibleByModelBAC string = "BAC"

	// VisibleByModelAG2 captures enum value "AG2"
	VisibleByModelAG2 string = "AG2"

	// VisibleByModelAG2DashHP captures enum value "AG2-HP"
	VisibleByModelAG2DashHP string = "AG2-HP"

	// VisibleByModelAG5 captures enum value "AG5"
	VisibleByModelAG5 string = "AG5"

	// VisibleByModelAG5DashHP captures enum value "AG5-HP"
	VisibleByModelAG5DashHP string = "AG5-HP"

	// VisibleByModelP2N captures enum value "p2N"
	VisibleByModelP2N string = "p2N"

	// VisibleByModelP5N captures enum value "p5N"
	VisibleByModelP5N string = "p5N"

	// VisibleByModelM25 captures enum value "M25"
	VisibleByModelM25 string = "M25"

	// VisibleByModelP2BDash400 captures enum value "P2B-400"
	VisibleByModelP2BDash400 string = "P2B-400"

	// VisibleByModelP5BDash300 captures enum value "P5B-300"
	VisibleByModelP5BDash300 string = "P5B-300"

	// VisibleByModelP5BDash300DashISO captures enum value "P5B-300-ISO"
	VisibleByModelP5BDash300DashISO string = "P5B-300-ISO"

	// VisibleByModelP5BDash400 captures enum value "P5B-400"
	VisibleByModelP5BDash400 string = "P5B-400"

	// VisibleByModelP5BDash400DashISO captures enum value "P5B-400-ISO"
	VisibleByModelP5BDash400DashISO string = "P5B-400-ISO"

	// VisibleByModelP5BDash620 captures enum value "P5B-620"
	VisibleByModelP5BDash620 string = "P5B-620"

	// VisibleByModelLB5Dash120 captures enum value "LB5-120"
	VisibleByModelLB5Dash120 string = "LB5-120"

	// VisibleByModelLB5 captures enum value "LB5"
	VisibleByModelLB5 string = "LB5"

	// VisibleByModelN5B captures enum value "N5B"
	VisibleByModelN5B string = "N5B"

	// VisibleByModelN5BDash16 captures enum value "N5B-16"
	VisibleByModelN5BDash16 string = "N5B-16"

	// VisibleByModelN5BDash19 captures enum value "N5B-19"
	VisibleByModelN5BDash19 string = "N5B-19"

	// VisibleByModelN5BDash300 captures enum value "N5B-300"
	VisibleByModelN5BDash300 string = "N5B-300"

	// VisibleByModelN5BDash400 captures enum value "N5B-400"
	VisibleByModelN5BDash400 string = "N5B-400"

	// VisibleByModelN5BDashClient captures enum value "N5B-Client"
	VisibleByModelN5BDashClient string = "N5B-Client"

	// VisibleByModelN2B captures enum value "N2B"
	VisibleByModelN2B string = "N2B"

	// VisibleByModelN2BDash13 captures enum value "N2B-13"
	VisibleByModelN2BDash13 string = "N2B-13"

	// VisibleByModelN2BDash400 captures enum value "N2B-400"
	VisibleByModelN2BDash400 string = "N2B-400"

	// VisibleByModelPAP captures enum value "PAP"
	VisibleByModelPAP string = "PAP"

	// VisibleByModelLAPDashHP captures enum value "LAP-HP"
	VisibleByModelLAPDashHP string = "LAP-HP"

	// VisibleByModelLAP captures enum value "LAP"
	VisibleByModelLAP string = "LAP"

	// VisibleByModelAGW captures enum value "AGW"
	VisibleByModelAGW string = "AGW"

	// VisibleByModelAGWDashLR captures enum value "AGW-LR"
	VisibleByModelAGWDashLR string = "AGW-LR"

	// VisibleByModelAGWDashPro captures enum value "AGW-Pro"
	VisibleByModelAGWDashPro string = "AGW-Pro"

	// VisibleByModelAGWDashInstaller captures enum value "AGW-Installer"
	VisibleByModelAGWDashInstaller string = "AGW-Installer"

	// VisibleByModelPB5 captures enum value "PB5"
	VisibleByModelPB5 string = "PB5"

	// VisibleByModelPB3 captures enum value "PB3"
	VisibleByModelPB3 string = "PB3"

	// VisibleByModelP36 captures enum value "P36"
	VisibleByModelP36 string = "P36"

	// VisibleByModelPBM10 captures enum value "PBM10"
	VisibleByModelPBM10 string = "PBM10"

	// VisibleByModelNB5 captures enum value "NB5"
	VisibleByModelNB5 string = "NB5"

	// VisibleByModelNB2 captures enum value "NB2"
	VisibleByModelNB2 string = "NB2"

	// VisibleByModelNB3 captures enum value "NB3"
	VisibleByModelNB3 string = "NB3"

	// VisibleByModelB36 captures enum value "B36"
	VisibleByModelB36 string = "B36"

	// VisibleByModelNB9 captures enum value "NB9"
	VisibleByModelNB9 string = "NB9"

	// VisibleByModelSM5 captures enum value "SM5"
	VisibleByModelSM5 string = "SM5"

	// VisibleByModelWM5 captures enum value "WM5"
	VisibleByModelWM5 string = "WM5"

	// VisibleByModelISDashM5 captures enum value "IS-M5"
	VisibleByModelISDashM5 string = "IS-M5"

	// VisibleByModelLoco5AC captures enum value "Loco5AC"
	VisibleByModelLoco5AC string = "Loco5AC"

	// VisibleByModelNSDash5AC captures enum value "NS-5AC"
	VisibleByModelNSDash5AC string = "NS-5AC"

	// VisibleByModelR5ACDashPTMP captures enum value "R5AC-PTMP"
	VisibleByModelR5ACDashPTMP string = "R5AC-PTMP"

	// VisibleByModelR5ACDashPTP captures enum value "R5AC-PTP"
	VisibleByModelR5ACDashPTP string = "R5AC-PTP"

	// VisibleByModelR5ACDashLite captures enum value "R5AC-Lite"
	VisibleByModelR5ACDashLite string = "R5AC-Lite"

	// VisibleByModelR5ACDashPRISM captures enum value "R5AC-PRISM"
	VisibleByModelR5ACDashPRISM string = "R5AC-PRISM"

	// VisibleByModelR2ACDashPrism captures enum value "R2AC-Prism"
	VisibleByModelR2ACDashPrism string = "R2AC-Prism"

	// VisibleByModelR2ACDashGen2 captures enum value "R2AC-Gen2"
	VisibleByModelR2ACDashGen2 string = "R2AC-Gen2"

	// VisibleByModelRPDash5ACDashGen2 captures enum value "RP-5AC-Gen2"
	VisibleByModelRPDash5ACDashGen2 string = "RP-5AC-Gen2"

	// VisibleByModelNBEDash2ACDash13 captures enum value "NBE-2AC-13"
	VisibleByModelNBEDash2ACDash13 string = "NBE-2AC-13"

	// VisibleByModelNBEDash5ACDash16 captures enum value "NBE-5AC-16"
	VisibleByModelNBEDash5ACDash16 string = "NBE-5AC-16"

	// VisibleByModelNBEDash5ACDash19 captures enum value "NBE-5AC-19"
	VisibleByModelNBEDash5ACDash19 string = "NBE-5AC-19"

	// VisibleByModelNBEDash5ACDashGen2 captures enum value "NBE-5AC-Gen2"
	VisibleByModelNBEDash5ACDashGen2 string = "NBE-5AC-Gen2"

	// VisibleByModelPBEDash5ACDash300 captures enum value "PBE-5AC-300"
	VisibleByModelPBEDash5ACDash300 string = "PBE-5AC-300"

	// VisibleByModelPBEDash5ACDash300DashISO captures enum value "PBE-5AC-300-ISO"
	VisibleByModelPBEDash5ACDash300DashISO string = "PBE-5AC-300-ISO"

	// VisibleByModelPBEDash5ACDash400 captures enum value "PBE-5AC-400"
	VisibleByModelPBEDash5ACDash400 string = "PBE-5AC-400"

	// VisibleByModelPBEDash5ACDash400DashISO captures enum value "PBE-5AC-400-ISO"
	VisibleByModelPBEDash5ACDash400DashISO string = "PBE-5AC-400-ISO"

	// VisibleByModelPBEDash5ACDash500 captures enum value "PBE-5AC-500"
	VisibleByModelPBEDash5ACDash500 string = "PBE-5AC-500"

	// VisibleByModelPBEDash5ACDash500DashISO captures enum value "PBE-5AC-500-ISO"
	VisibleByModelPBEDash5ACDash500DashISO string = "PBE-5AC-500-ISO"

	// VisibleByModelPBEDash5ACDash620 captures enum value "PBE-5AC-620"
	VisibleByModelPBEDash5ACDash620 string = "PBE-5AC-620"

	// VisibleByModelPBEDash5ACDash620DashISO captures enum value "PBE-5AC-620-ISO"
	VisibleByModelPBEDash5ACDash620DashISO string = "PBE-5AC-620-ISO"

	// VisibleByModelPBEDash2ACDash400 captures enum value "PBE-2AC-400"
	VisibleByModelPBEDash2ACDash400 string = "PBE-2AC-400"

	// VisibleByModelPBEDash2ACDash400DashISO captures enum value "PBE-2AC-400-ISO"
	VisibleByModelPBEDash2ACDash400DashISO string = "PBE-2AC-400-ISO"

	// VisibleByModelPBEDash5ACDashXDashGen2 captures enum value "PBE-5AC-X-Gen2"
	VisibleByModelPBEDash5ACDashXDashGen2 string = "PBE-5AC-X-Gen2"

	// VisibleByModelPBEDash5ACDashGen2 captures enum value "PBE-5AC-Gen2"
	VisibleByModelPBEDash5ACDashGen2 string = "PBE-5AC-Gen2"

	// VisibleByModelPBEDash5ACDashISODashGen2 captures enum value "PBE-5AC-ISO-Gen2"
	VisibleByModelPBEDash5ACDashISODashGen2 string = "PBE-5AC-ISO-Gen2"

	// VisibleByModelPBEDash5ACDash400DashISODashGen2 captures enum value "PBE-5AC-400-ISO-Gen2"
	VisibleByModelPBEDash5ACDash400DashISODashGen2 string = "PBE-5AC-400-ISO-Gen2"

	// VisibleByModelLBEDash5ACDash16Dash120 captures enum value "LBE-5AC-16-120"
	VisibleByModelLBEDash5ACDash16Dash120 string = "LBE-5AC-16-120"

	// VisibleByModelLAPDash120 captures enum value "LAP-120"
	VisibleByModelLAPDash120 string = "LAP-120"

	// VisibleByModelLBEDash5ACDash23 captures enum value "LBE-5AC-23"
	VisibleByModelLBEDash5ACDash23 string = "LBE-5AC-23"

	// VisibleByModelLBEDash5ACDashGen2 captures enum value "LBE-5AC-Gen2"
	VisibleByModelLBEDash5ACDashGen2 string = "LBE-5AC-Gen2"

	// VisibleByModelLBEDash5ACDashLR captures enum value "LBE-5AC-LR"
	VisibleByModelLBEDash5ACDashLR string = "LBE-5AC-LR"

	// VisibleByModelLAPDashGPS captures enum value "LAP-GPS"
	VisibleByModelLAPDashGPS string = "LAP-GPS"

	// VisibleByModelISDash5AC captures enum value "IS-5AC"
	VisibleByModelISDash5AC string = "IS-5AC"

	// VisibleByModelPSDash5AC captures enum value "PS-5AC"
	VisibleByModelPSDash5AC string = "PS-5AC"

	// VisibleByModelSolarSwitch captures enum value "SolarSwitch"
	VisibleByModelSolarSwitch string = "SolarSwitch"

	// VisibleByModelSolarPoint captures enum value "SolarPoint"
	VisibleByModelSolarPoint string = "SolarPoint"

	// VisibleByModelBulletACDashIP67 captures enum value "BulletAC-IP67"
	VisibleByModelBulletACDashIP67 string = "BulletAC-IP67"

	// VisibleByModelBDashDBDashAC captures enum value "B-DB-AC"
	VisibleByModelBDashDBDashAC string = "B-DB-AC"

	// VisibleByModelUNKNOWN captures enum value "UNKNOWN"
	VisibleByModelUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *VisibleBy) validateModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, visibleByTypeModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VisibleBy) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	// value enum
	if err := m.validateModelEnum("model", "body", *m.Model); err != nil {
		return err
	}

	return nil
}

func (m *VisibleBy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var visibleByTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["onu","olt","uispp","uispr","uisps","uispLte","erouter","eswitch","epower","airCube","airMax","airFiber","toughSwitch","solarBeam","wave","blackBox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		visibleByTypeTypePropEnum = append(visibleByTypeTypePropEnum, v)
	}
}

const (

	// VisibleByTypeOnu captures enum value "onu"
	VisibleByTypeOnu string = "onu"

	// VisibleByTypeOlt captures enum value "olt"
	VisibleByTypeOlt string = "olt"

	// VisibleByTypeUispp captures enum value "uispp"
	VisibleByTypeUispp string = "uispp"

	// VisibleByTypeUispr captures enum value "uispr"
	VisibleByTypeUispr string = "uispr"

	// VisibleByTypeUisps captures enum value "uisps"
	VisibleByTypeUisps string = "uisps"

	// VisibleByTypeUispLte captures enum value "uispLte"
	VisibleByTypeUispLte string = "uispLte"

	// VisibleByTypeErouter captures enum value "erouter"
	VisibleByTypeErouter string = "erouter"

	// VisibleByTypeEswitch captures enum value "eswitch"
	VisibleByTypeEswitch string = "eswitch"

	// VisibleByTypeEpower captures enum value "epower"
	VisibleByTypeEpower string = "epower"

	// VisibleByTypeAirCube captures enum value "airCube"
	VisibleByTypeAirCube string = "airCube"

	// VisibleByTypeAirMax captures enum value "airMax"
	VisibleByTypeAirMax string = "airMax"

	// VisibleByTypeAirFiber captures enum value "airFiber"
	VisibleByTypeAirFiber string = "airFiber"

	// VisibleByTypeToughSwitch captures enum value "toughSwitch"
	VisibleByTypeToughSwitch string = "toughSwitch"

	// VisibleByTypeSolarBeam captures enum value "solarBeam"
	VisibleByTypeSolarBeam string = "solarBeam"

	// VisibleByTypeWave captures enum value "wave"
	VisibleByTypeWave string = "wave"

	// VisibleByTypeBlackBox captures enum value "blackBox"
	VisibleByTypeBlackBox string = "blackBox"
)

// prop value enum
func (m *VisibleBy) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, visibleByTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VisibleBy) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this visible by based on context it is used
func (m *VisibleBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VisibleBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisibleBy) UnmarshalBinary(b []byte) error {
	var res VisibleBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
