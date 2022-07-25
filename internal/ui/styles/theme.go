package styles

import "github.com/charmbracelet/lipgloss"

var indigo = lipgloss.AdaptiveColor{Light: "#ffafcc", Dark: "#383B5B"}

type Theme struct {
	MainText      lipgloss.AdaptiveColor
	SubleMainText lipgloss.AdaptiveColor

	Border          lipgloss.AdaptiveColor
	SecondaryBorder lipgloss.AdaptiveColor

	ErrorText   lipgloss.AdaptiveColor
	SuccessText lipgloss.AdaptiveColor

	FaintBorder lipgloss.AdaptiveColor
	FaintText   lipgloss.AdaptiveColor

	SelectedBackground lipgloss.AdaptiveColor
	SecondaryText      lipgloss.AdaptiveColor
}

var subtleIndigo = lipgloss.AdaptiveColor{Light: "#ffc8dd", Dark: "#242347"}

var DefaultTheme = Theme{
	MainText:      lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"},
	SubleMainText: subtleIndigo,

	Border:          lipgloss.AdaptiveColor{Light: indigo.Light, Dark: indigo.Dark},
	SecondaryBorder: lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#39386b"},

	ErrorText:   lipgloss.AdaptiveColor{Light: "#F23D5C", Dark: "#F23D5C"},
	SuccessText: lipgloss.AdaptiveColor{Light: "#3DF294", Dark: "#3DF294"},

	FaintBorder: lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#2b2b40"},
	FaintText:   lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#3E4057"},

	SelectedBackground: lipgloss.AdaptiveColor{Light: subtleIndigo.Light, Dark: "#39386b"},
	SecondaryText:      lipgloss.AdaptiveColor{Light: indigo.Light, Dark: "#666CA6"},
}

var (
	SingleRuneWidth    = 4
	MainContentPadding = 1
)
