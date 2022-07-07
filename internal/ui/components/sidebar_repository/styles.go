package sidebar_repository

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	labelDockerOfficial    = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#2E7F74"}
	labelVerifiedPublisher = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#086DD7"}
	labelOpenSourceProgram = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#7D2EFF"}
	statsDownloads         = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#00BBF9"}
	statsStars             = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#FFB703"}
	archTagFg              = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#EDEDE9"}
	archTagBg              = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#0077b6"}
	dockerImageTagFg       = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#22223b"}
	dockerImageTagBg       = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#EDEDE9"}
	dockerPullCmdBoxBg     = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#8ac926"}
	dockerPullCmdBoxFg     = lipgloss.AdaptiveColor{Light: "#FFF", Dark: "#22223b"}

	label = lipgloss.NewStyle().Copy().
		Padding(0, 1).
		Bold(false)

	archTitle = label.Copy().
			Underline(true).
			Margin(0, 1)

	archTag = label.Copy().
		Background(archTagBg).
		Foreground(archTagFg).
		Margin(0, 1)

	dockerImageTitle = label.Copy().
				Underline(true).
				Margin(0, 1)

	dockerImageTag = label.Copy().
			Background(dockerImageTagBg).
			Foreground(dockerImageTagFg).
			Margin(0, 1)

	dockerPullCmdTitle = label.Copy().
				Underline(true).
				Margin(0, 1)

	dockerPullCmdBox = label.Copy().
				Background(dockerPullCmdBoxBg).
				Foreground(dockerPullCmdBoxFg).
				Padding(2).
				Margin(2)
)
