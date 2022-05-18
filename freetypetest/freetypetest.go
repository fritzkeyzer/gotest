package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"log"
	"path"
)

func main() {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 28})

	genImage("a", "train", face)
	genImage("b", "train", face)
	genImage("c", "train", face)
	genImage("d", "train", face)
	genImage("e", "train", face)
	genImage("f", "train", face)

}

func genImage(char string, outputFolder string, face font.Face) {
	dc := gg.NewContext(28, 28)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(char, 13, 13, 0.5, 0.5)
	dc.SavePNG(path.Join(outputFolder, char+".png"))
}
