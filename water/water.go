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
			Consume:     1704,
			Consumo_rec: 2057,
			Rec_soles:   7336.2,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "belmonte":
		gpr := WaterByMonth{
			Consume:     285.31,
			Consumo_rec: 302,
			Rec_soles:   1082.8,
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
	case "mirador":
		gpr := WaterByMonth{
			Consume:     229,
			Consumo_rec: 0, //casi no sirve para nada
			Rec_soles:   826.2,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	}
	return &WaterByMonth{}

}
