package building

type Building struct {
	Name     string
	Nickname string
	Address  string
	Email    string
	Picture  string

	Bank             string
	BankAccount      string
	BankAccountOwner string
	Budget           string

	FirstColumn  []string
	SecondColumn []string
}

func (b *Building) GetBuildingData(name string) {
	switch name {
	case "gpr":
		b.Name = "CONDOMINIO GRAN PARQUE ROMA"
		b.Nickname = "GPR"
		b.Address = "LEONARDO ARIETA 825 - CERCADO DE LIMA"
		b.Email = "granparqueroma@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "3059864512041"
		b.BankAccountOwner = "C. RECAUDADORA GRAN PARQUE ROMA"
		b.Picture = "files/parque-roma-logo.jpg"
	case "belmonte":
		b.Name = "EDIFICIO BELMONTE"
		b.Nickname = "BELMONTE"
		b.Address = "JIRON DANIEL OLAECHEA 246"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19306790451040"
		b.BankAccountOwner = "NITZIA ROJAS / ODARIS LUCENA"
		b.Picture = "files/belmonte.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA: ", "MAN. PREVENTIVO: ", "MAN. ASCENSOR: "}
		b.SecondColumn = []string{"MATERIALES LIMPIEZA: ", "MANTENIMIENTO JARDINES: ", "LUZ SSGG: ", "LUZ BCI: ", "ADMINISTRACION Y PERSONAL: "}
	case "torrereal":
		b.Name = "EDIFICIO TORRE REAL"
		b.Nickname = "TORREREAL"
		b.Address = "JIRON DANIEL OLAECHEA 175 - JESÚS MARÍA"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "1939621755042"
		b.BankAccountOwner = "ESCOBAR BARRIONUEVO DIEGO"
		b.Picture = "files/torrereal.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "PARTICIPACION: ", "ESTACIONAMIENTO: ", "DEPOSITO: ", "AGUA: ", "FONDO DE RESERVA: ", "FONDO MANTENIMIENTO: "}
		b.SecondColumn = []string{"MAN. ASCENSOR: ", "MATERIALES LIMPIEZA: ", "LUZ SSGG: ", "LUZ BCI: ", "ADMINISTRACION Y PERSONAL: ", "MULTA: ", "SALDO A FAVOR/CONTRA: ", "AGUA DIC.: "}
	case "mirador":
		b.Name = "EDIFICIO MIRADOR 2"
		b.Nickname = "MIRADOR"
		b.Address = "AV. PARQUE SUR #446 URB.CORPAC - SAN ISIDRO"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19374206534093"
		b.BankAccountOwner = "ROSA CACEDA/GONZALO CERDA"
		b.Picture = "files/mirador.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA TOTAL: ", "AGUA X DPTO: ", "MAN. ASCENSOR: ", "PORTERO JERSON: ", "PORTERO ROBERTO: "}
		b.SecondColumn = []string{"MATERIALES LIMPIEZA: ", "PERSONAL LIMPIEZA: ", "DESCANSERO: ", "MANTENIMIENTO JARDINES: ", "LUZ SSGG: ", "LUZ BCI: ", "ADMINISTRACION: "}
	}

}
