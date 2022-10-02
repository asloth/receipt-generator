package water

import "strings"

type WaterByMonth struct {
	Consume     float64
	Consumo_rec float64
	Rec_soles   float64
	Soles_m3    float64
}

func GetWaterDataByBuilding(name string) *WaterByMonth {
	switch strings.ToLower(name) {
	case "gpr":
		gpr := WaterByMonth{
			Consume:     2209,
			Consumo_rec: 925,
			Rec_soles:   3171.00,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	}
	return &WaterByMonth{}

}
