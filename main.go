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
		switch {
		case i == 0:
		case i <= 16:
			if i%8 == 0 {
				fmt.Println()
			}
			if i%16 == 0 {
				fmt.Println()
			}
		default:
			if (i-16)%6 == 0 {
				fmt.Println()
			}
			if (i-16)%36 == 0 {
				fmt.Println()
			}
		}

		str := fmt.Sprintf("%03d", i)
		color := lipgloss.Color(strconv.Itoa(i))

		fmt.Print(style.Background(color).Render(str))
		fmt.Print(style.Foreground(color).Render(str))
	}
	fmt.Println()
}
