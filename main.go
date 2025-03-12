package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func main() {
	_, _ = io.WriteString(os.Stdout, Generate())
}

const bufCap = 20_000

func Generate() string {
	lipgloss.SetColorProfile(termenv.ANSI256)
	style := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("15"))

	var buf strings.Builder
	buf.Grow(bufCap)
	for i := range int64(256) {
		str := fmt.Sprintf("%03d", i)
		color := lipgloss.Color(strconv.FormatInt(i, 10))
		buf.WriteString(style.Background(color).Render(str))
		buf.WriteString(style.Foreground(color).Render(str))

		switch {
		case i <= 15:
			switch {
			case i == 15:
				buf.WriteString("\n\n")
			case i == 7:
				buf.WriteByte('\n')
			}
		default:
			switch {
			case (i-15)%36 == 0:
				buf.WriteString("\n\n")
			case (i-15)%6 == 0:
				buf.WriteByte('\n')
			}
		}
	}
	return buf.String()
}
