package building

type Building struct {
	Name     string
	Nickname string
	Address  string
	Email    string
	Picture  string
	PayData RecollectionAccount
	Budget           string

	Towers []Tower

	FirstColumn  []string
	SecondColumn []string

	HaveWater bool
}

type Tower struct {
	Name string
	Account RecollectionAccount
}

type RecollectionAccount struct {
	Number string
	Bank string
	Owner string
	CCI string
}


func (b *Building) GetBuildingData(name string) {
	switch name {
	case "gpr":
		b.Name = "CONDOMINIO GRAN PARQUE ROMA"
		b.Nickname = "GPR"
		b.Address = "LEONARDO ARIETA 825 - CERCADO DE LIMA"
		b.Email = "granparqueroma@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "3059864512041",
			Bank: "BCP",
			Owner: "C. RECAUDADORA GRAN PARQUE ROMA",
		}
		b.Picture = "files/parque-roma-logo.jpg"
		b.HaveWater = false
	case "tampu":
		b.Name = "TAMPUMACHAY"
		b.Nickname = "TAMPU"
		b.Address = "Jr. Tampumachay 229 - Santiago de Surco"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19196858198060 / CCI: OO219119685819806051",
			Bank: "BCP",
			Owner: "MICHELLE SWAYNE",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	case "gpl":
		b.Name = "GRAN PLAZA LORETO"
		b.Nickname = "GPL"
		b.Address = "JIRÓN LORETO 1590 BREÑA - PUEBLO LIBRE"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "3059864512041",
			Bank: "BCP",
			Owner: "LIMA 2000 - EL MOLIO",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	case "arenaycampo":
		b.Name = "ARENA Y CAMPO"
		b.Nickname = "arenaycampo"
		b.Address = "CHOCAYA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "305-8123356-0-07",
			Bank: "BCP",
			Owner: "EL MOLIO",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "jardines":
		b.Name = "LOS JARDINES DE CHORILLOS"
		b.Nickname = "JARDINES"
		b.Address = "AV. GUARDIA CIVIL 665 - CHORRILLOS"
		b.Email = "administradorlimaoeste@elmolio.com"
		b.Picture = "files/default.png"
		b.HaveWater = true
		b.Towers = []Tower{
			{
				Name: "TORRE A",
				Account: RecollectionAccount{
					Number: "0011-0137-0200665517",
					CCI: "011-137-000200665517-75",
					Bank: "BBVA",
					Owner: "MARLON GUTIERREZ GUTIERREZ",
				},
			},
			{
				Name: "TORRE B",
				Account: RecollectionAccount{
					Number: "0011-0659-0200167604",
					Bank: "BBVA",
					CCI: "011-659-000200167604-04",
					Owner: "JEANCARLO G. ZEBALLOS GONZALES/ SOFIA ROSALBA AVALOS",
				},
			},
			{
				Name: "TORRE C",
				Account: RecollectionAccount{
					Number: "0011-0137-0200665355",
					Bank: "BBVA",
					Owner: "NERY PEÑA",
					CCI: "011-137-000200665355-70",
				},
			},
		}
	case "elite":
		b.Name = "ÉLITE CASUARINAS"
		b.Nickname = "ELITE"
		b.Address = "LOS LIRIOS - SANTIAGO DE SURCO"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "001101860100065358",
			Bank: "BBVA",
			Owner: "LOTE 25 DE LA MANZANA G1 JIRON LOS LIRIOS",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "belmonte":
		b.Name = "EDIFICIO BELMONTE"
		b.Nickname = "BELMONTE"
		b.Address = "JIRON DANIEL OLAECHEA 246"
		b.Email = "administradorlimaeste@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19103422218085",
			CCI: "00219110342221808551",
			Bank: "BCP",
			Owner: "Eric Arancibia",
		}
		b.Picture = "files/belmonte.jpeg"
		b.HaveWater = true
	case "rio":
		b.Name = "EDIFICIO RÍO DE JANEIRO"
		b.Nickname = "RIO"
		b.Address = "Calle Río de Janeiro - 256 - Miraflores"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "OO1101870200433449",
			CCI: "011 187 000200433449 22",
			Bank: "BBVA",
			Owner: "JUNTA DE PROPIETARIOS EDIFICIO RESIDENCIAL RÍO DE JANEIRO",
		}
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = true
	case "torrereal":
		b.Name = "EDIFICIO TORRE REAL"
		b.Nickname = "TORREREAL"
		b.Address = "JIRON DANIEL OLAECHEA 175 - JESÚS MARÍA"
		b.Email = "administradorlimaeste@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "1937118402049",
			CCI: "00219300711840204914",
			Bank: "BCP",
			Owner: "Junta de Propietarios",
		}
		b.Picture = "files/torrereal.jpeg"
		b.HaveWater = true
	case "valera":
		b.Name = "EDIFICIO VARELA III"
		b.Nickname = "VALERA"
		b.Address = "Jr. Gral. Varela 871-879 - Breña"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "OO1101100200477594",
			Bank: "BBVA",
			Owner: "HENRY ALVAREZ / SIGUAS GALVEZ CORY NATHALI",
		}
		b.Picture = "files/valera.jpeg"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "sanjose":
		b.Name = "CABALLERIZAS DE SAN JOSE"
		b.Nickname = "sanjose"
		b.Address = "Dist. El Carmen - Chincha"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19391258783086",
			Bank: "BCP",
			Owner: "PATRICIA VERA / MÓNICA ENRIQUEZ",
			CCI: "OO219319125878308615",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	case "mirador":
		b.Name = "EDIFICIO MIRADOR 2"
		b.Nickname = "MIRADOR"
		b.Address = "AV. PARQUE SUR #446 URB.CORPAC - SAN ISIDRO"
		b.Email = "administradorlimaeste@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19374206534093",
			Bank: "BCP",
			Owner: "ROSA CACEDA/GONZALO CERDA",
		}
		b.Picture = "files/mirador.jpeg"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "nitoa":
		b.Name = "EDIFICIO NITOA I"
		b.Nickname = "NITOA"
		b.Address = "C. LAS PALOMAS 204 - LIMATAMBO - SURQUILLO"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19492161547070",
			Bank: "BCP",
			Owner: "VEGA GABRIELA-O-CHERO AMELIA",
			CCI: "00219419216154707091",
		}
		b.Picture = "files/nitoa.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = true
	case "golf":
		b.Name = "EDIFICIO GOLF PARK"
		b.Nickname = "GOLF"
		b.Address = "CERROS DE CAMACHO 417-421 - SURCO"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19100437985007",
			Bank: "BCP",
			Owner: "XIMENA MURO FELMAN",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "mora":
		b.Name = "EDIFICIO MORA 454"
		b.Nickname = "MORA"
		b.Address = "FEDERICO VILLAREAL 454 - MIRAFLORES"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "0011 0426 0200302091",
			Bank: "BBVA",
			Owner: "ANA LUISA JOVE / GISELLA INCIO",
			CCI: "011 426 000200302091 40",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "alayza":
		b.Name = "CARLOS ALAYZA"
		b.Nickname = "ALAYZA"
		b.Address = "CARLOS ALAYZA Y ROEL 2561 - 2555 - LINCE"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "305-2646967-0-46",
			Bank: "BCP",
			Owner: "EL MOLIO - CARLOS ALAYZA",
			CCI: "OO230500264696704612",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	case "avila":
		b.Name = "PARQUE AVILA"
		b.Nickname = "AVILA"
		b.Address = "JIRON MALAGA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "0453456596152",
			Bank: "INTERBANK",
			Owner: "César Nuñovero y David",
			CCI: "00304501345659615228",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "sbs":
		b.Name = "SAN BORJA SUR"
		b.Nickname = "SBS"
		b.Address = "SAN BORJA SUR 1069 - SAN BORJA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "3059864513051",
			Bank: "BCP",
			Owner: "EL MOLIO - SAN BORJA SUR",
		}
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "montereal":
		b.Name = "MONTE REAL"
		b.Nickname = "MONTEREAL"
		b.Address = "JR. MONTE REAL 490-492 URB. CHACARILLA DEL ESTANQUE"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "305-9864565-0-76",
			Bank: "BCP",
			Owner: "EL MOLIO",
			CCI: "",
		}
		b.Picture = "files/default.png"
		b.FirstColumn = []string{}
		b.SecondColumn = []string{}
		b.HaveWater = false
	case "tomasal":
		b.Name = "TOMASAL"
		b.Nickname = "TOMASAL"
		b.Address = "Jr. Tomasal 753 - Santiago de Surco"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "305-8969463-0-42",
			Bank: "BCP",
			Owner: "EL MOLIO - TOMASAL",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "balcones":
		b.Name = "LOS BALCONES DE SAN BLAS"
		b.Nickname = "BALCONES"
		b.Address = "LOS FAISANES 379-375 - CHORRILLOS"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "1948287857024",
			Bank: "BCP",
			Owner: "LOS BALCONES DE SAN BLAS",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "killa":
		b.Name = "EDIFICIO RESIDENCIAL KILLA"
		b.Nickname = "KILLA"
		b.Address = "MALECON SUPERIOR 1201 - PUNTA HERMOSA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "194-9354551-0-12",
			Bank: "BCP",
			Owner: "MALECON SUPERIOR 1201 Y AVENIDA COSTA PERUANA",
			CCI: "",
		}
		b.Picture = "files/default.png"
		b.HaveWater = false
	case "gcc":
		b.Name = "GRAN CENTRAL COLONIAL"
		b.Nickname = "GCC"
		b.Address = "GENERAL OSCAR R BENAVIDES 2703 - CERCADO DE LIMA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "3058123355097",
			Bank: "BCP",
			Owner: "EL MOLIO",
			CCI: "OO230500812335509719",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	case "huascar":
		b.Name = "EDIFICIO HUASCAR"
		b.Nickname = "HUASCAR"
		b.Address = "JIRON VARELA Y ORBEGOZO 439 - SURQUILLO"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "19310371405066",
			Bank: "BCP",
			Owner: "ARACELI VALDERRAMA",
		}
		b.Picture = "files/default.png"
		b.FirstColumn = []string{"NOMBRE: "}
		b.SecondColumn = []string{"AGUA: "}
		b.HaveWater = false
	case "rosapark":
		b.Name = "LA ROSA PARK"
		b.Nickname = "ROSAPARK"
		b.Address = "PASAJE QUIÑONES 195 - JESUS MARÍA"
		b.Email = "administracion@elmolio.com"
		b.PayData = RecollectionAccount{
			Number: "0011-0184-0200750219",
			Bank: "BBVA",
			Owner: "Guillermo Omar González Cucho / Cesar Camilo Soto Vidal",
			CCI: "01118400020075021994",
		}
		b.Picture = "files/default.png"
		b.HaveWater = true
	}
}
