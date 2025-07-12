package style

import "github.com/charmbracelet/lipgloss"

var (
	Title = lipgloss.NewStyle().
		Bold(true).
		Padding(1, 0, 0, 5).
		Foreground(lipgloss.Color("3")).
		BorderStyle(lipgloss.DoubleBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("8"))

	Event = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("8")).
		BorderLeft(true).
		PaddingLeft(2).
		MarginLeft(2)

	Error = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("9")). // Warna merah lebih cocok untuk error
		Padding(1, 2)
)
