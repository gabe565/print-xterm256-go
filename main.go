package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func main() {
	lipgloss.SetColorProfile(termenv.ANSI256)

	style := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("#fff"))

	for i := range 256 {
		str := fmt.Sprintf("%03d", i)
		color := lipgloss.Color(strconv.Itoa(i))
		fmt.Print(
			style.Background(color).Render(str),
			style.Foreground(color).Render(str),
		)

		switch {
		case i <= 15:
			switch {
			case i == 15:
				fmt.Print("\n\n")
			case i == 7:
				fmt.Println()
			}
		default:
			switch {
			case (i-15)%36 == 0:
				fmt.Print("\n\n")
			case (i-15)%6 == 0:
				fmt.Println()
			}
		}
	}
}
