// Package main renders an image or video
package main

import (
	"log"
	"os"

	flag "github.com/spf13/pflag"

	cairo "github.com/bit101/blcairo"
	"github.com/bit101/manifesto"
)

//revive:disable:unused-parameter

var (
	// input/output
	input   *string = flag.StringP("input", "i", "", "A text file containing the text of the manifesto.")
	output  *string = flag.StringP("output", "o", "manifesto.png", "The name of the output png file.")
	verbose *bool   = flag.Bool("verbose", false, "Enables extra logging output.")

	// canvas
	width    *int     = flag.IntP("width", "w", 800, "The width of the manifesto, in pixels.")
	height   *int     = flag.IntP("height", "h", 800, "The height of the manifesto, in pixels.")
	margin   *float64 = flag.Float64P("margin", "m", 50, "The margin around the text, in pixels.")
	rotation *float64 = flag.Float64P("rotation", "r", 0, "Rotation in degrees, if copy effects are applied.")

	// text
	fontFace      *string  = flag.String("font-face", "courier", "The font used to render the text. Monospaced fonts are best.")
	fontHeight    *float64 = flag.Float64("font-height", 24, "The height of the font in pixels.")
	fontWidth     *float64 = flag.Float64("font-width", 13, "The width of a single character, in pixels.")
	tabSize       *float64 = flag.Float64("tab-size", 4, "How many spaces is a single tab character.")
	lineSpacing   *float64 = flag.Float64("line-spacing", 1, "How much space one line of text takes up. 1 is single-spaced, 2 is double-spaced, etc.")
	textCase      *string  = flag.String("case", "", "Transform text case. Valid options: 'upper', 'lower'")
	bold          *bool    = flag.Bool("bold", false, "Uses a bold font.")
	darkInk       *float64 = flag.Float64("dark-ink", 0, "The darkest ink used. From 0.0 (black) to 1.0 (white).")
	lightInk      *float64 = flag.Float64("light-ink", 0.5, "The lightest ink used. From 0.0 (black) to 1.0 (white).")
	randomSpacing *float64 = flag.Float64("random-spacing", 1, "How much to randomize space between characters.")
	randomY       *float64 = flag.Float64("random-y", 2, "How much to randomize the vertical position of characters.")
	randomYChance *float64 = flag.Float64("random-y-chance", 0.07, "The chance a character will be randomly vertically positioned.")

	noCopyEffects *bool    = flag.Bool("no-copy", false, "Do not apply copy effects.")
	copyCount     *int     = flag.Int("copy-count", 3, "How many times the manifest is re-copied.")
	dither        *bool    = flag.Bool("dither", false, "Applied dithering to further gitch the text during the copy effects phase.")
	blur          *int     = flag.IntP("blur", "b", 3, "How much blurring is applied in the copy effects stage.")
	sharpCount    *int     = flag.Int("sharp-count", 2, "How much sharpening is applied after blurring during the copy effects stage.")
	brightness    *float64 = flag.Float64("brightness", -0.2, "Brightens or darkens the image in the copy effects phase.")
	contrast      *float64 = flag.Float64("contrast", 0.4, "Increases or decreases the contrast of the image in the copy effects phase.")

	help *bool = flag.Bool("help", false, "shows this help")
)

func main() {
	flag.CommandLine.SortFlags = false
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	context := makeContext()
	text := readText()

	man := manifesto.NewManifesto(context)
	man.SetBlur(*blur)
	man.SetBold(*bold)
	man.SetBrightness(*brightness)
	man.SetContrast(*contrast)
	man.SetCopy(!*noCopyEffects)
	man.SetCopyCount(*copyCount)
	man.SetDarkness(*darkInk, *lightInk)
	man.SetDither(*dither)
	man.SetFontFace(*fontFace)
	man.SetFontSize(*fontHeight, *fontWidth)
	man.SetLineHeight(*lineSpacing * *fontHeight)
	man.SetMargin(*margin)
	man.SetRandomSpacing(*randomSpacing)
	man.SetRandomY(*randomY, *randomYChance)
	man.SetRotation(*rotation)
	man.SetSharpCount(*sharpCount)
	man.SetTabSize(*tabSize)
	man.SetTextCase(*textCase)
	man.SetVerbose(*verbose)
	man.Render(text)

	context.Surface.WriteToPNG(*output)
}

func makeContext() *cairo.Context {
	surface := cairo.NewSurface(*width, *height)
	context := cairo.NewContext(surface)
	context.BlackOnWhite()
	return context
}

func readText() string {
	bytes, err := os.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	return text
}
