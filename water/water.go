package water

import "strings"

type WaterMonthData struct {
	LastMonth              float64
	CurrentMonth           float64
	WaterConsumedThisMonth float64
  CommonWater float64
}

type WaterByMonth struct {
	Consume     float64 //consumo calculado
	Consumo_rec float64 //consumo en m3 que viene en el recibo
	Rec_soles   float64 //el monto del recibo
	Soles_m3    float64 //el monto de cuanto cuesta 1 m3
}

func GetWaterDataByBuilding(name string) *WaterByMonth {
	switch strings.ToLower(name) {
	case "gpr":
		gpr := WaterByMonth{
			Consume:     1693,
			Consumo_rec: 1753,
			Rec_soles:   6265.7,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "belmonte":
		gpr := WaterByMonth{
			Consume:     249.00,
			Consumo_rec: 277.00,
			Rec_soles:   994.10,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
	case "torrereal":
		gpr := WaterByMonth{
			Consume:     190.65,
			Consumo_rec: 186.00,
			Rec_soles:   669.40,
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
	case "nitoa":
		gpr := WaterByMonth{
			Consume:     295.02,
			Consumo_rec: 292.00, //casi no sirve para nada
			Rec_soles:   1056.70,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "valera":
		gpr := WaterByMonth{
			Consume:     375.00,
			Consumo_rec: 318.00, //casi no sirve para nada
			Rec_soles:   1182.00,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "mora":
		gpr := WaterByMonth{
			Consume:     359.64,
			Consumo_rec: 338.00, //casi no sirve para nada
			Rec_soles:   1213.10,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "alayza":
		gpr := WaterByMonth{
			Consume:     337.74,
			Consumo_rec: 335.00, //casi no sirve para nada
			Rec_soles:   1200.70,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "sbs":
		gpr := WaterByMonth{
			Consume:     64.086,
			Consumo_rec: 73, //casi no sirve para nada
			Rec_soles:   268.80,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
	case "montereal":
		gpr := WaterByMonth{
			Consume:     235991,
			Consumo_rec: 235991, //casi no sirve para nada
			Rec_soles:   961.71,
		}
		gpr.Soles_m3 = gpr.Rec_soles / (gpr.Consume / 1000)
		return &gpr
	case "tomasal":
		gpr := WaterByMonth{
			Consume:     151.807,
			Consumo_rec: 160.000, //casi no sirve para nada
			Rec_soles:   578.60,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
  case "balcones":
		gpr := WaterByMonth{
			Consume:     1065.61,
			Consumo_rec: 1226, //casi no sirve para nada
			Rec_soles:   4436.7,
		}
		gpr.Soles_m3 = 3.69
		return &gpr
	}
	return &WaterByMonth{}

}
