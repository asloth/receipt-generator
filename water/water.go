package water

import "strings"

type WaterMonthData struct {
	LastMonth              float64
	CurrentMonth           float64
	WaterConsumedThisMonth float64
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
			Consume:     176.41,
			Consumo_rec: 198.00,
			Rec_soles:   713.20,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
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
			Consume:     212.17,
			Consumo_rec: 301.00, //casi no sirve para nada
			Rec_soles:   1090.30,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
	case "valera":
		gpr := WaterByMonth{
			Consume:     416.00,
			Consumo_rec: 353.00, //casi no sirve para nada
			Rec_soles:   1265.00,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "mora":
		gpr := WaterByMonth{
			Consume:     310.40,
			Consumo_rec: 284.88, //casi no sirve para nada
			Rec_soles:   1022.30,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "alayza":
		gpr := WaterByMonth{
			Consume:     362.63,
			Consumo_rec: 358.00, //casi no sirve para nada
			Rec_soles:   1290.80,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consume
		return &gpr
	case "sbs":
		gpr := WaterByMonth{
			Consume:     60.384,
			Consumo_rec: 69, //casi no sirve para nada
			Rec_soles:   253.30,
		}
		gpr.Soles_m3 = gpr.Rec_soles / gpr.Consumo_rec
		return &gpr
	}
	return &WaterByMonth{}

}
