package water

import "strings"

type WaterMonthData struct {
	LastMonth              float64
	CurrentMonth           float64
	WaterConsumedThisMonth float64
}

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
			Consume:     1632,
			Consumo_rec: 1565,
			Rec_soles:   5582.90,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "belmonte":
		gpr := WaterByMonth{
			Consume:     288,
			Consumo_rec: 280,
			Rec_soles:   1004,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "torrereal":
		gpr := WaterByMonth{
			Consume:     180.08,
			Consumo_rec: 0,
			Rec_soles:   0,
		}
		gpr.Soles_m3 = 0
		return &gpr
	}
	return &WaterByMonth{}

}
