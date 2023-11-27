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

	HaveWater bool
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
		b.HaveWater = true
	case "belmonte":
		b.Name = "EDIFICIO BELMONTE"
		b.Nickname = "BELMONTE"
		b.Address = "JIRON DANIEL OLAECHEA 246"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "193-90343163-0-18"
		b.BankAccountOwner = "DE LA PUENTE MARIA / ODARIS LUCENA"
		b.Picture = "files/belmonte.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA: ", "MAN. PREVENTIVO: ", "MAN. ASCENSOR: "}
		b.SecondColumn = []string{"MATERIALES LIMPIEZA: ", "MANTENIMIENTO JARDINES: ", "LUZ SSGG: ", "LUZ BCI: ", "ADMINISTRACION Y PERSONAL: "}
		b.HaveWater = true
	case "torrereal":
		b.Name = "EDIFICIO TORRE REAL"
		b.Nickname = "TORREREAL"
		b.Address = "JIRON DANIEL OLAECHEA 175 - JESÚS MARÍA"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "INTERBANK"
		b.BankAccount = "6393372392454"
		b.BankAccountOwner = "NOLES COTITO MARIELA INES"
		b.Picture = "files/torrereal.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "PARTICIPACION: ", "ESTACIONAMIENTO: ", "DEPOSITO: ", "AGUA: ", "FONDO DE RESERVA: ", "FONDO MANTENIMIENTO: "}
		b.SecondColumn = []string{"MAN. ASCENSOR: ", "INSUMOS LIMPIEZA: ", "LUZ SSGG: ", "LUZ BCI: ", "ADMINISTRACION Y PERSONAL: ", "MULTA: ", "SALDO A FAVOR: ", "MORA: "}
		b.HaveWater = true
	case "valera":
		b.Name = "EDIFICIO VARELA III"
		b.Nickname = "VALERA"
		b.Address = "Jr. Gral. Varela 871-879 - Breña"
		b.Email = "administrador@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "191-71949491-0-50"
		b.BankAccountOwner = "SIGUAS GALVEZ CORY NATHALI"
		b.Picture = "files/valera.jpeg"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "PARTICIPACION(%): ", "ESTACIONAMIENTO (%): ", "DPTO (%): ", "LUZ BCI: ", "LUZ SSGG: "}
		b.SecondColumn = []string{"AGUA X DPTO.: ", "RECARGA EXTINTORES: ", "REPARACIÓN ASC. DISCAPACITADOS: ", "SERV. ADMINISTRACION: ", "INSUMOS FERRETERIA: ", "CUOTA ESTACIONAMIENTO: ", "CUOTA X DPTO.: "}
		b.HaveWater = true
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
		b.HaveWater = true
	case "nitoa":
		b.Name = "EDIFICIO NITOA I"
		b.Nickname = "NITOA"
		b.Address = "C. LAS PALOMAS 204 - LIMATAMBO - SURQUILLO"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19492161547070"
		b.BankAccountOwner = "VEGA GABRIELA-O-CHERO AMELIA"
		b.Picture = "files/nitoa.png"
		b.FirstColumn = []string{"PROPIETARIO: ", "DPTO: ", "AGUA COMUN: ", "AGUA X DPTO: ", "LUZ SSGG: ", "LUZ BCI: ", "TOTAL AGUA Y LUZ: ", "INCREMENTO APROBADO ASAMBLEA 2023: "}
		b.SecondColumn = []string{"SANEAMIENTO Y LIMPIEZA: ", "MAN. PREVENTIVOS: ", "MAN. CORRECTIVO: ", "SEGURIDAD: ", "ADMINISTRACION: ", "MORA: ", "REDONDEO DEL MES: ", "SUBTOTAL: "}
		b.HaveWater = true
	case "golf":
		b.Name = "EDIFICIO GOLF PARK"
		b.Nickname = "GOLF"
		b.Address = "CERROS DE CAMACHO 417-421 - SURCO"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19100437985007"
		b.BankAccountOwner = "XIMENA MURO FELMAN"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"DEPARTMENTO: "}
		b.HaveWater = false
	case "mora":
		b.Name = "EDIFICIO MORA 454"
		b.Nickname = "MORA"
		b.Address = "FEDERICO VILLAREAL 454 - MIRAFLORES"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "OO1104260200287653"
		b.BankAccountOwner = "DIAZ MAMANI NATALY YULIANA Y/O BENZAQUEN VASQUEZ YAIR"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA X DPTO.: ", "AGUA COMUN: "}
		b.SecondColumn = []string{"LUZ SSGG: ", "MANTENIMIENTOS PREVENTIVOS: ", "MONTO CONTINGENCIAS:", "ADMINISTRACION: "}
		b.HaveWater = true
	case "alayza":
		b.Name = "CARLOS ALAYZA"
		b.Nickname = "ALAYZA"
		b.Address = "CARLOS ALAYZA Y ROEL 2561 - 2555 - LINCE"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-2646967-0-46"
		b.BankAccountOwner = "EL MOLIO - CARLOS ALAYZA"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA X DPTO.: ", "PERSONAL: ", "MAN. ALARMAS CONTRA INCENDIO: ", "MAN. PUERTA LEVADIZA: ", "REINTEGRO SEDAPAL: "}
		b.SecondColumn = []string{"LUZ SSGG: ", "LUZ BCI: ", "SERVICIOS POR HONORARIOS:", "ADMINISTRACION: ", "MANTENIMIENTOS ASCENSOR: ", "MONTO CONTINGENCIAS: ", "MAN. POZOS A TIERRA: "}
		b.HaveWater = true
	case "sbs":
		b.Name = "SAN BORJA SUR"
		b.Nickname = "SBS"
		b.Address = "SAN BORJA SUR 1069 - SAN BORJA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "3059864513051"
		b.BankAccountOwner = "EL MOLIO - SAN BORJA SUR"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "AGUA X DPTO.: ", "GASTOS BANCARIOS: ", "INTERNET: "}
		b.SecondColumn = []string{"LUZ SSGG: ", "PERSONAL LIMPIEZA: ", "SERV. ADMINISTRACION:", "INSUMOS DE LIMPIEZA: ", "FONDO CONTINGENCIA: "}
		b.HaveWater = true
	case "montereal":
		b.Name = "MONTE REAL"
		b.Nickname = "MONTEREAL"
		b.Address = "JR. MONTE REAL 490-492 URB. CHACARILLA DEL ESTANQUE"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-9864565-0-76"
		b.BankAccountOwner = "EL MOLIO"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "ESTACIONAMIENTO: "}
		b.SecondColumn = []string{"AGUA: ", "MANTENIMIENTO: ", "CUOTA EXTRAORDINARIA: "}
		b.HaveWater = true
	case "tomasal":
		b.Name = "TOMASAL"
		b.Nickname = "TOMASAL"
		b.Address = "Jr. Tomasal 753 - Santiago de Surco"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-8969463-0-42"
		b.BankAccountOwner = "EL MOLIO - TOMASAL"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: ", "DEPARTAMENTO: ", "PARTICIPACION: "}
		b.SecondColumn = []string{"AGUA: ", "CUOTA ANTERIOR: ", "4,000 MÁS 7%: "}
		b.HaveWater = true
	}
}
