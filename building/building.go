package building

type Building struct {
	name    string
	address string
	email   string

	bank             string
	bankAccount      string
	bankAccountOwner string
}

func getBuildingData(name string) *Building {
	switch name {
	case "gpr":
		gpr := Building{
			name:             "CONDOMINIO GRAN PARQUE ROMA",
			address:          "LEONARDO ARIETA 825 - CERCADO DE LIMA",
			email:            "granparqueroma@elmolio.com",
			bank:             "BCP",
			bankAccount:      "3059864512041",
			bankAccountOwner: "C. RECAUDADORA GRAN PARQUE ROMA",
		}
		return &gpr
	}
	return &Building{}

}
