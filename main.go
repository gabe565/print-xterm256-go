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

	for i := 0; i < 256; i += 1 {
		if i != 0 {
			if i <= 16 {
				if i%8 == 0 {
					fmt.Println()
				}
				if i%16 == 0 {
					fmt.Println()
				}
			} else {
				if (i-16)%6 == 0 {
					fmt.Println()
				}
				if (i-16)%36 == 0 {
					fmt.Println()
				}
			}
		}

		str := fmt.Sprintf("%03d", i)
		color := lipgloss.Color(strconv.Itoa(i))

		bgStyle := style.Copy().Background(color)
		fmt.Print(bgStyle.Render(str))

		fgStyle := style.Copy().Foreground(color)
		fmt.Print(fgStyle.Render(str))
	}
	fmt.Println()
}
