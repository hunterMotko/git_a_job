package font

import (
	"os"

	"github.com/common-nighthawk/go-figure"
)

func GetFont() string {
	//	func NewFigure(phrase, fontName string, strict bool) figure
	// func NewColorFigure(phrase, fontName string, color string, strict bool) figure
	// func NewFigureWithFont(phrase string, reader io.Reader, strict bool) figure
	username := os.Getenv("USER")
	return figure.NewFigure(username, "doom", true).String()
}
