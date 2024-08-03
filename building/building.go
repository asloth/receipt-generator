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
		b.HaveWater = false
	case "elite":
		b.Name = "ÉLITE CASUARINAS"
		b.Nickname = "ELITE"
		b.Address = "LOS LIRIOS - SANTIAGO DE SURCO"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "001101860100065358"
		b.BankAccountOwner = "LOTE 25 DE LA MANZANA G1 JIRON LOS LIRIOS"
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "belmonte":
		b.Name = "EDIFICIO BELMONTE"
		b.Nickname = "BELMONTE"
		b.Address = "JIRON DANIEL OLAECHEA 246"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "193-90343163-0-18"
		b.BankAccountOwner = "DE LA PUENTE MARIA / ODARIS LUCENA"
		b.Picture = "files/belmonte.jpeg"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "rio":
		b.Name = "EDIFICIO RÍO DE JANEIRO"
		b.Nickname = "RIO"
		b.Address = "Calle Río de Janeiro - 256 - Miraflores"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "OO1101870200433449 / CCI: 011 187 000200433449 22"
		b.BankAccountOwner = "JUNTA DE PROPIETARIOS EDIFICIO RESIDENCIAL RÍO DE JANEIRO"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = true
	case "torrereal":
		b.Name = "EDIFICIO TORRE REAL"
		b.Nickname = "TORREREAL"
		b.Address = "JIRON DANIEL OLAECHEA 175 - JESÚS MARÍA"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19399644215020 / CCI: 00219319964421502015"
		b.BankAccountOwner = "Reynaldo Martín Soza Martínez"
		b.Picture = "files/torrereal.jpeg"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "valera":
		b.Name = "EDIFICIO VARELA III"
		b.Nickname = "VALERA"
		b.Address = "Jr. Gral. Varela 871-879 - Breña"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "OO1101100200477594"
		b.BankAccountOwner = "HENRY ALVAREZ / SIGUAS GALVEZ CORY NATHALI"
		b.Picture = "files/valera.jpeg"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "sanjose":
		b.Name = "CABALLERIZAS DE SAN JOSE"
		b.Nickname = "sanjose"
		b.Address = "Dist. El Carmen - Chincha"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19391258783086 / CCI: OO219319125878308615"
		b.BankAccountOwner = "PATRICIA VERA / MÓNICA ENRIQUEZ"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "mirador":
		b.Name = "EDIFICIO MIRADOR 2"
		b.Nickname = "MIRADOR"
		b.Address = "AV. PARQUE SUR #446 URB.CORPAC - SAN ISIDRO"
		b.Email = "administradorlimaeste@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19374206534093"
		b.BankAccountOwner = "ROSA CACEDA/GONZALO CERDA"
		b.Picture = "files/mirador.jpeg"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "nitoa":
		b.Name = "EDIFICIO NITOA I"
		b.Nickname = "NITOA"
		b.Address = "C. LAS PALOMAS 204 - LIMATAMBO - SURQUILLO"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19492161547070 / CCI: 00219419216154707091"
		b.BankAccountOwner = "VEGA GABRIELA-O-CHERO AMELIA"
		b.Picture = "files/nitoa.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
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
		b.BankAccount = "0011 0426 0200302091 / CCI: 011 426 000200302091 40"
		b.BankAccountOwner = "ANA LUISA JOVE / GISELLA INCIO"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "alayza":
		b.Name = "CARLOS ALAYZA"
		b.Nickname = "ALAYZA"
		b.Address = "CARLOS ALAYZA Y ROEL 2561 - 2555 - LINCE"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-2646967-0-46 / CCI: OO230500264696704612"
		b.BankAccountOwner = "EL MOLIO - CARLOS ALAYZA"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = true
	case "avila":
		b.Name = "PARQUE AVILA"
		b.Nickname = "AVILA"
		b.Address = "JIRON MALAGA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "001107870200364600 / CCI: O1178700020036460093"
		b.BankAccountOwner = "Ricardo Guzmán y Cecilia Fuchs"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "sbs":
		b.Name = "SAN BORJA SUR"
		b.Nickname = "SBS"
		b.Address = "SAN BORJA SUR 1069 - SAN BORJA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "3059864513051"
		b.BankAccountOwner = "EL MOLIO - SAN BORJA SUR"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "montereal":
		b.Name = "MONTE REAL"
		b.Nickname = "MONTEREAL"
		b.Address = "JR. MONTE REAL 490-492 URB. CHACARILLA DEL ESTANQUE"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-9864565-0-76"
		b.BankAccountOwner = "EL MOLIO"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "tomasal":
		b.Name = "TOMASAL"
		b.Nickname = "TOMASAL"
		b.Address = "Jr. Tomasal 753 - Santiago de Surco"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "305-8969463-0-42"
		b.BankAccountOwner = "EL MOLIO - TOMASAL"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "balcones":
		b.Name = "LOS BALCONES DE SAN BLAS"
		b.Nickname = "BALCONES"
		b.Address = "LOS FAISANES 342 - CHORRILLOS"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "1948287857024"
		b.BankAccountOwner = "LOS BALCONES DE SAN BLAS"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "killa":
		b.Name = "EDIFICIO RESIDENCIAL KILLA"
		b.Nickname = "KILLA"
		b.Address = "MALECON SUPERIOR 1201 - PUNTA HERMOSA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "194-9354551-0-12"
		b.BankAccountOwner = "MALECON SUPERIOR 1201 Y AVENIDA COSTA PERUANA"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"AGUA: "}
		b.HaveWater = false
	case "gcc":
		b.Name = "GRAN CENTRAL COLONIAL"
		b.Nickname = "GCC"
		b.Address = "GENERAL OSCAR R BENAVIDES 2703 - CERCADO DE LIMA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "3058123355097 / CCI: OO230500812335509719"
		b.BankAccountOwner = "EL MOLIO"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"AGUA: "}
		b.HaveWater = true
	case "huascar":
		b.Name = "EDIFICIO HUASCAR"
		b.Nickname = "HUASCAR"
		b.Address = "JIRON VARELA Y ORBEGOZO 439 - SURQUILLO"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BCP"
		b.BankAccount = "19310371405066"
		b.BankAccountOwner = "ARACELI VALDERRAMA"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"AGUA: "}
		b.HaveWater = false
	case "rosapark":
		b.Name = "LA ROSA PARK"
		b.Nickname = "ROSAPARK"
		b.Address = "PASAJE QUIÑONES 195 - JESUS MARÍA"
		b.Email = "administracion@elmolio.com"
		b.Bank = "BBVA"
		b.BankAccount = "OO1101840200742364 / CCI: 01118400020074236499"
		b.BankAccountOwner = "CECILIA PAOLA RAMOS/VICENTE PEDRO GUTIERREZ"
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"AGUA: "}
		b.HaveWater = true
	}
}
