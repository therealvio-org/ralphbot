package guidefetch

type guide struct {
	Name           string
	SubCommandName string
	Description    string
	GDriveLink     string
	GHLink         string
}

var (
	crypt = &guide{
		Name:           "Deep Stone Crypt",
		SubCommandName: "raid-crypt",
		Description:    "Deep Stone Crypt Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1YKU4_-hInHQ3rVEAvqIjdJaT25oQvmYc?usp=sharing",
	}

	garden = &guide{
		Name:           "Garden of Salvation",
		SubCommandName: "raid-garden",
		Description:    "Garden of Salvation Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1pPdtAptJMaaDYRv2i-8bfaL6l3I0WTsT?usp=sharing",
	}

	kingsfall = &guide{
		Name:           "King's Fall",
		SubCommandName: "raid-kingsfall",
		Description:    "King's Fall Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1tsOVCy2SwP0rLUDQUJaDIFh5y-O0DoKn",
	}

	pit = &guide{
		Name:           "Pit of Heresy",
		SubCommandName: "dungeon-pit",
		Description:    "Pit of Heresy Dungeon",
		GDriveLink:     "https://drive.google.com/drive/folders/17lB7m9KQMwzBb6UHfoBt9ZEA82haD2Fd?usp=sharing",
	}

	ron = &guide{
		Name:           "The Root of Nightmares",
		SubCommandName: "raid-tron",
		Description:    "The Root of Nightmares Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1eR50Jt36GBegMALRkT-tmnH6Nnjc4Pqj?usp=share_link",
	}

	spire = &guide{
		Name:           "Spire of the Watcher",
		SubCommandName: "dungeon-spire",
		Description:    "Spire of the Watcher Dungeon",
		GDriveLink:     "https://drive.google.com/drive/folders/1Xu_8NfiPFnknocqdR8p9adWPF-qiHmxo?usp=share_link",
	}

	vault = &guide{
		Name:           "Vault of Glass",
		SubCommandName: "raid-vault",
		Description:    "Vault of Glass Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1HLx6nVIji_3OcwnzaLeSoksspa4pfdjD?usp=sharing",
	}

	vow = &guide{
		Name:           "Vow of the Disciple",
		SubCommandName: "raid-vow",
		Description:    "Vow of the Disciple Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1ZAPIXYlSs7yTQEdznQAqz2rOnnvpzwr7?usp=sharing",
		GHLink:         "https://github.com/therealvio/destiny-guides/tree/main/raids/vow-of-the-disciple",
	}

	wish = &guide{
		Name:           "Last Wish",
		SubCommandName: "raid-lastwish",
		Description:    "Last Wish Raid",
		GDriveLink:     "https://drive.google.com/drive/folders/1d_WEa84KuX1_9hPTwgFhl651IwywHeOg?usp=sharing",
	}

	Guides = []guide{
		*crypt,
		*garden,
		*kingsfall,
		*pit,
		*ron,
		*spire,
		*vault,
		*vow,
		*wish,
	}
)
