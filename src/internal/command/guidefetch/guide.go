package guidefetch

type Guide struct {
	Name           string
	SubCommandName string
	Description    string
	GDriveUrl      string
	GHUrl          string
}

//type Guides struct {
//	Crypt     GuideDetails
//	Garden    GuideDetails
//	KingsFall GuideDetails
//	Pit       GuideDetails
//	Vault     GuideDetails
//	Vow       GuideDetails
//	Wish      GuideDetails
//}

var (
	crypt = &Guide{
		Name:           "Deep Stone Crypt",
		SubCommandName: "raid-crypt",
		Description:    "Deep Stone Crypt Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1YKU4_-hInHQ3rVEAvqIjdJaT25oQvmYc?usp=sharing",
	}

	garden = &Guide{
		Name:           "Garden of Salvation",
		SubCommandName: "raid-garden",
		Description:    "Garden of Salvation Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1pPdtAptJMaaDYRv2i-8bfaL6l3I0WTsT?usp=sharing",
	}

	kingsfall = &Guide{
		Name:           "Kings Fall",
		SubCommandName: "raid-kingsfall",
		Description:    "King's Fall Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1tsOVCy2SwP0rLUDQUJaDIFh5y-O0DoKn",
	}

	pit = &Guide{
		Name:           "Pit of Heresy",
		SubCommandName: "dungeon-pit",
		Description:    "Pit of Heresy Dungeon",
		GDriveUrl:      "https://drive.google.com/drive/folders/17lB7m9KQMwzBb6UHfoBt9ZEA82haD2Fd?usp=sharing",
	}

	vault = &Guide{
		Name:           "Vault of Glass",
		SubCommandName: "raid-vault",
		Description:    "Vault of Glass Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1HLx6nVIji_3OcwnzaLeSoksspa4pfdjD?usp=sharing",
	}

	vow = &Guide{
		Name:           "Vow of the Disciple",
		SubCommandName: "raid-vow",
		Description:    "Vow of the Disciple Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1ZAPIXYlSs7yTQEdznQAqz2rOnnvpzwr7?usp=sharing",
		GHUrl:          "https://github.com/therealvio/destiny-guides/tree/main/raids/vow-of-the-disciple",
	}

	wish = &Guide{
		Name:           "Last Wish",
		SubCommandName: "raid-lastwish",
		Description:    "Last Wish Raid",
		GDriveUrl:      "https://drive.google.com/drive/folders/1d_WEa84KuX1_9hPTwgFhl651IwywHeOg?usp=sharing",
	}
)
