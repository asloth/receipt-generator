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
			Consume:     1692,
			Consumo_rec: 1751,
			Rec_soles:   6246.50,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "belmonte":
		gpr := WaterByMonth{
			Consume:     294.36,
			Consumo_rec: 301.00,
			Rec_soles:   1080.10,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
	case "torrereal":
		gpr := WaterByMonth{
			Consume:     172.67,
			Consumo_rec: 0,
			Rec_soles:   706.30,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	}
	return &WaterByMonth{}

}
