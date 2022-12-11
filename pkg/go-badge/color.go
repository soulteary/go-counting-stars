package badge

// Color represents color of the badge.
type Color string

// ColorScheme contains named colors that could be used to render the badge.
// https://github.com/badges/shields/blob/b6be37d277b64a90bde98ca10446c9d433a56681/badge-maker/lib/color.js#L6
var ColorScheme = map[string]string{
	"brightgreen": "#4c1",
	"green":       "#97ca00",
	"yellow":      "#dfb317",
	"yellowgreen": "#a4a61d",
	"orange":      "#fe7d37",
	"red":         "#e05d44",
	"blue":        "#007ec6",
	"grey":        "#555",
	"lightgrey":   "#9f9f9f",
	// aliases
	"gray":          "#555",
	"lightgray":     "#9f9f9f",
	"critical":      "#e05d44",
	"important":     "#fe7d37",
	"success":       "#4c1",
	"informational": "#007ec6",
	"inactive":      "#9f9f9f",
}

// Standard colors.
const (
	ColorBrightgreen = Color("brightgreen")
	ColorGreen       = Color("green")
	ColorYellow      = Color("yellow")
	ColorYellowgreen = Color("yellowgreen")
	ColorOrange      = Color("orange")
	ColorRed         = Color("red")
	ColorBlue        = Color("blue")
	ColorGrey        = Color("grey")
	ColorLightgrey   = Color("lightgrey")
	// aliases
	ColorGray          = Color("gray")
	ColorLightgray     = Color("lightgray")
	ColorCritical      = Color("critical")
	ColorImportant     = Color("important")
	ColorSuccess       = Color("success")
	ColorInformational = Color("informational")
	ColorInactive      = Color("inactive")
)

func (c Color) String() string {
	color, ok := ColorScheme[string(c)]
	if ok {
		return color
	} else {
		return string(c)
	}
}
