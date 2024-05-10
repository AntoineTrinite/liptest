package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func main() {
	var healthStyle2 = lipgloss.NewStyle().
		Width(50).
		Height(1).
		BorderStyle(lipgloss.NormalBorder()).
		Background(lipgloss.Color("#F4A4E4"))

	fmt.Println(healthStyle2.Render())

	rows := [][]string{
		{"Chinese", "您好", "你好"},
		{"Japanese", "こんにちは", "やあ"},
		{"Arabic", "أهلين", "أهلا"},
		{"Russian", "Здравствуйте", "Привет"},
		{"Spanish", "Hola", "¿Qué tal?"},
	}

	var HeaderStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("227"))

	var EvenRowStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("227"))

	var OddRowStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("227"))

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Headers("LANGUAGE", "FORMAL", "INFORMAL").
		Rows(rows...)

	// You can also add tables row-by-row
	t.Row("English", "You look absolutely fabulous.", "How's it going?")

	fmt.Println(t)
}
