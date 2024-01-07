package water

type WaterMonthData struct {
	LastMonth              float64
	CurrentMonth           float64
	WaterConsumedThisMonth float64
	CommonWater            float64
}

type WaterByMonth struct {
	Consume     string //consumo calculado
	Consumo_rec string //consumo en m3 que viene en el recibo
	Rec_soles   string //el monto del recibo
	Soles_m3    string //el monto de cuanto cuesta 1 m3
}
