package main

// Basic const
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
	White  = "\033[97m"
	Black  = "\033[30m"
	Orange = "\033[38;5;208m"
	Pink   = "\033[38;5;200m"
	Gold   = "\033[38;5;214m"
)

// Red Const
const (
	BloodRed = "\033[38;5;88m"
	WarRed1  = "\033[38;5;124m"
	WarRed2  = "\033[38;5;160m"
	DarkRed  = "\033[38;5;52m"
	RubyRed  = "\033[38;5;196m"
)

// Green Const
const (
	ForestGreen = "\033[38;5;22m"
	GrassGreen  = "\033[38;5;34m"
	LimeGreen   = "\033[38;5;154m"
	Mint        = "\033[38;5;155m"
)

// Blue Const
const (
	SkyBlue   = "\033[38;5;39m"
	NavyBlue  = "\033[38;5;21m"
	OceanBlue = "\033[38;5;32m"
	IceBlue   = "\033[38;5;74m"
	Teal      = "\033[38;5;30m"
)

// Orange Const
const (
	BrightOrange = "\033[38;5;208m" // Orange vif
	DarkOrange   = "\033[38;5;202m" // Orange foncé
	MediumOrange = "\033[38;5;214m" // Orange moyen
	PaleOrange   = "\033[38;5;202m" // Orange pâle
	RustyOrange  = "\033[38;5;172m" // Orange rouillé
)

// Purple Const
const (
	Lavender  = "\033[38;5;183m"
	Lavender1 = "\033[38;5;93m"  // Lavande clair
	Lavender2 = "\033[38;5;129m" // Lavande moyen
	Lavender3 = "\033[38;5;141m" // Lavande foncé
	Lavender4 = "\033[38;5;177m" // Lavande très clair
	Lavender5 = "\033[38;5;183m" // Lavande très foncé
	Orchid1   = "\033[38;5;95m"  // Orchidée claire
	Orchid2   = "\033[38;5;134m" // Orchidée moyenne
	Orchid3   = "\033[38;5;170m" // Orchidée foncée
	Orchid4   = "\033[38;5;175m" // Orchidée très claire
	Orchid5   = "\033[38;5;176m" // Orchidée très foncée
)

// Beige Const
const (
	LightBeige     = "\033[38;5;180m" // Beige clair
	MediumBeige    = "\033[38;5;223m" // Beige moyen
	DarkBeige      = "\033[38;5;223m" // Beige foncé
	VeryLightBeige = "\033[38;5;224m" // Beige très clair
	VeryDarkBeige  = "\033[38;5;223m" // Beige très foncé
)

// Gray Const
const (
	Gray1 = "\033[38;5;232m" // Gris très clair
	Gray2 = "\033[38;5;240m" // Gris clair
	Gray3 = "\033[38;5;248m" // Gris moyen
	Gray4 = "\033[38;5;254m" // Gris très clair
)

// Special Const
const (
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Reverse   = "\033[7m"
)

/*
func main() {
	fmt.Println(SkyBlue + "Ce texte est coloré en SkyBlue !" + Reset)
	fmt.Println(NavyBlue + "Ce texte est coloré en NavyBlue !" + Reset)
	fmt.Println(OceanBlue + "Ce texte est coloré en OceanBlue !" + Reset)
	fmt.Println(IceBlue + "Ce texte est coloré en IceBlue !" + Reset)
	fmt.Println(Teal + "Ce texte est coloré en Teal !" + Reset)

	fmt.Println(ForestGreen + "Ce texte est coloré en ForestGreen !" + Reset)
	fmt.Println(GrassGreen + "Ce texte est coloré en GrassGreen !" + Reset)
	fmt.Println(LimeGreen + "Ce texte est coloré en LimeGreen !" + Reset)
	fmt.Println(Mint + "Ce texte est coloré en Mint !" + Reset)

	fmt.Println(LightBeige + "Ce texte est coloré en Beige clair !" + Reset)
	fmt.Println(MediumBeige + "Ce texte est coloré en Beige moyen !" + Reset)
	fmt.Println(DarkBeige + "Ce texte est coloré en Beige foncé !" + Reset)
	fmt.Println(VeryLightBeige + "Ce texte est coloré en Beige très clair !" + Reset)
	fmt.Println(VeryDarkBeige + "Ce texte est coloré en Beige très foncé !" + Reset)

	fmt.Println(BrightOrange + "Ce texte est coloré en Orange vif !" + Reset)
	fmt.Println(RustyOrange + "Ce texte est coloré en Orange rouillé !" + Reset)
	fmt.Println(MediumOrange + "Ce texte est coloré en Orange moyen !" + Reset)
	fmt.Println(DarkOrange + "Ce texte est coloré en Orange foncé !" + Reset)
	fmt.Println(PaleOrange + "Ce texte est coloré en Orange pâle !" + Reset)

	fmt.Println(Gray1 + "Ce texte est coloré en Gris1 !" + Reset)
	fmt.Println(Gray2 + "Ce texte est coloré en Gris2 !" + Reset)
	fmt.Println(Gray3 + "Ce texte est coloré en Gris3 !" + Reset)
	fmt.Println(Gray4 + "Ce texte est coloré en Gris4 !" + Reset)

	// Print new colors
	fmt.Println(Lavender + "Ce texte est coloré en Lavender !" + Reset)
	fmt.Println(Orchid1 + "Ce texte est coloré en Orchid1 !" + Reset)
	fmt.Println(Gray4 + "Ce texte est coloré en Gray4 !" + Reset)

	// Print shades of red
	fmt.Println(BloodRed + "Ce texte est coloré en BloodRed !" + Reset)
	fmt.Println(WarRed1 + "Ce texte est coloré en WarRed1 !" + Reset)
	fmt.Println(WarRed2 + "Ce texte est coloré en WarRed2 !" + Reset)
	fmt.Println(DarkRed + "Ce texte est coloré en DarkRed !" + Reset)
	fmt.Println(RubyRed + "Ce texte est coloré en RubyRed !" + Reset)
}
*/
