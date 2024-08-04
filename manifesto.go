// Package manifesto renders text in a typewritten and over-copied style.
package manifesto

import (
	"fmt"
	"math"
	"strings"

	"github.com/bit101/bitlib/random"
	cairo "github.com/bit101/blcairo"
)

// CopySettings are the effect settings for simulating copying.
type CopySettings struct {
	copyEffects bool
	rotation    float64
	dither      bool
	copyCount   int
	blur        int
	sharpCount  int
	brightness  float64
	contrast    float64
}

// Manifesto defines a manifesto.
type Manifesto struct {
	context       *cairo.Context
	width         float64
	height        float64
	fontSize      float64
	charWidth     float64
	lineHeight    float64
	margin        float64
	randomSpacing float64
	randomY       float64
	randomYChance float64
	tabSize       float64
	fontFace      string
	fontWeight    int
	darkest       float64
	lightest      float64
	textCase      string
	verbose       bool
	copySettings  CopySettings
}

// NewManifesto creates a new manifesto with default values.
func NewManifesto(context *cairo.Context) *Manifesto {
	return &Manifesto{
		context:       context,
		width:         context.Width,
		height:        context.Height,
		fontSize:      24.0,
		charWidth:     13.0,
		lineHeight:    24.0,
		margin:        50.0,
		randomSpacing: 1,
		randomY:       2,
		randomYChance: 0.07,
		tabSize:       4.0,
		fontFace:      "courier",
		fontWeight:    cairo.FontWeightNormal,
		darkest:       0.0,
		lightest:      0.5,
		textCase:      "",
		verbose:       false,
		copySettings: CopySettings{
			copyEffects: true,
			rotation:    0.0,
			dither:      false,
			copyCount:   3,
			blur:        3,
			sharpCount:  2,
			brightness:  -0.2,
			contrast:    0.4,
		},
	}
}

//////////////////////////////
// Main Render Method
//////////////////////////////

// Render renders text in the manifesto style.
func (m *Manifesto) Render(text string) {
	if m.copySettings.rotation != 0.0 {
		m.context.Translate(m.width/2, m.height/2)
		m.context.Rotate(m.copySettings.rotation)
		m.context.Translate(-m.width/2, -m.height/2)
	}
	m.context.Translate(m.margin, m.margin)
	m.logAction("rendering text")
	m.renderText(text)
	if m.copySettings.copyEffects {
		m.logAction("rendering copy effects")
		m.copyEffects()
	}
}

//////////////////////////////
// Setting Methods
//////////////////////////////

// SetBlur sets how much the text is blurred while copying.
// The default is 3.
func (m *Manifesto) SetBlur(blur int) {
	m.copySettings.blur = blur
}

// SetBold sets whether or not text will be rendered bold.
// The default is false.
func (m *Manifesto) SetBold(bold bool) {
	if bold {
		m.fontWeight = cairo.FontWeightBold
	} else {
		m.fontWeight = cairo.FontWeightNormal
	}
}

// SetBrightness sets how much contrast is applied after blurring.
// By default it is darkened slightly with a setting of -0.2.
func (m *Manifesto) SetBrightness(brightness float64) {
	m.copySettings.brightness = brightness
}

// SetContrast sets how much the text is brightened or darkened after copying.
// By default a contrast of 0.4 is applied.
func (m *Manifesto) SetContrast(contrast float64) {
	m.copySettings.contrast = contrast
}

// SetCopy sets whether or not copying effects will be applied.
// By default they are applied.
func (m *Manifesto) SetCopy(b bool) {
	m.copySettings.copyEffects = b
}

// SetCopyCount sets how many times the manifesto will be copied (blurred and sharpened).
// Default is 3 copy cycles.
func (m *Manifesto) SetCopyCount(count int) {
	m.copySettings.copyCount = count
}

// SetDarkness sets the darkest and lightest shade of gray that will be used to render each char.
// Each char will get a random darkness within this range.
// This can simulate a ribbon that is low on ink.
// The brightness and contrast settings will tend to even this difference out somewhat,
// so you may want to make lightest a bit lighter if using copy effects.
// Default darkest is 0.0 (full black). Default lightest is 0.5 (50% black).
func (m *Manifesto) SetDarkness(darkest, lightest float64) {
	m.darkest = darkest
	m.lightest = lightest
}

// SetDither sets whether or not dithering will be applied before copying.
// Dithering adds more glitchiness.
// By default no dithering is applied.
func (m *Manifesto) SetDither(b bool) {
	m.copySettings.dither = b
}

// SetFontFace sets what font will be used to render text.
// Note that monospaced fonts will generally look a lot better here.
// Default is "courier".
func (m *Manifesto) SetFontFace(fontFace string) {
	m.fontFace = fontFace
}

// SetFontSize sets the height and width of each character.
// By default font height is 24 and width is 13.
func (m *Manifesto) SetFontSize(height, width float64) {
	m.fontSize = height
	m.charWidth = width
}

// SetLineHeight sets the height of each line of text.
// By default it is 24 - or the same as the font height.
// Hint: to make double spaced text, set line height to font height times two.
func (m *Manifesto) SetLineHeight(height float64) {
	m.lineHeight = height
}

// SetMargin sets the margin applied on all sides of the surface.
// The default margin is 50 pixels.
func (m *Manifesto) SetMargin(margin float64) {
	m.margin = margin
}

// SetRandomSpacing sets a random horizontal spacing jitter applied to each character.
// Default is 1.0, which is generally not very noticeable but results in realistic
// small spacing glitches now and then.
func (m *Manifesto) SetRandomSpacing(spacing float64) {
	m.randomSpacing = spacing
}

// SetRandomY sets how far and how often characters will jump out of line vertically.
// By default they can jump up to 2 pixels up or down.
// And this can happen with a chance of 0.07 (roughly once every 16 chars).
// Note: a random jump could be very close to 0.0 and thus not be very obvious.
func (m *Manifesto) SetRandomY(offset, chance float64) {
	m.randomY = offset
	m.randomYChance = chance
}

// SetRotation sets how much the text rotated during copying
// By default no rotation is done.
func (m *Manifesto) SetRotation(rotation float64) {
	m.copySettings.rotation = rotation * math.Pi / 180.0
}

// SetSharpCount sets how many times the text will be sharpened after blurring.
// Default is 2 and is mostly what you want.
func (m *Manifesto) SetSharpCount(sharpCount int) {
	m.copySettings.sharpCount = sharpCount
}

// SetTabSize sets how many spaces will be inserted for each tab character.
// Default is 4.0.
func (m *Manifesto) SetTabSize(tabSize float64) {
	m.tabSize = tabSize
}

// SetTextCase sets the case of the text to upper, lower, or no change.
func (m *Manifesto) SetTextCase(textCase string) {
	m.textCase = textCase
}

// SetVerbose sets extra logging of what is being rendered.
func (m *Manifesto) SetVerbose(verbose bool) {
	m.verbose = verbose
}

//////////////////////////////
// Private methods
//////////////////////////////

// renderText renders the text line by line.
func (m *Manifesto) renderText(text string) {
	if m.textCase == "lower" {
		text = strings.ToLower(text)
	} else if m.textCase == "upper" {
		text = strings.ToUpper(text)
	}
	m.context.SelectFontFace(m.fontFace, cairo.FontSlantNormal, m.fontWeight)
	m.context.SetFontSize(m.fontSize)
	lines := strings.Split(text, "\n")
	y := m.lineHeight
	for _, line := range lines {
		y = m.renderLine(line, y)
		if y > m.height-m.margin*2 {
			return
		}
	}
}

// renderLine renders a single line of the source text.
// May result in multiple lines of rendered text.
func (m *Manifesto) renderLine(line string, y float64) float64 {
	x := 0.0
	words := strings.Split(line, " ")

	for _, word := range words {
		if x > m.width-m.margin*2-m.charWidth*float64(len(word)) {
			x = 0
			y += m.lineHeight
			if y > m.height-m.margin*2 {
				return y
			}
		}
		for _, c := range word {
			if c == '\n' {
				x = 0
				y += m.lineHeight
				if y > m.height-m.margin*2 {
					return y
				}
				continue
			}
			if c == '\t' {
				x += m.charWidth * m.tabSize
				continue
			}
			m.context.Save()
			m.context.Translate(x, y)
			m.context.Translate(random.FloatRange(-m.randomSpacing, m.randomSpacing), 0)
			if random.WeightedBool(m.randomYChance) {
				m.context.Translate(0, random.FloatRange(-m.randomY, m.randomY))
			}
			m.context.SetSourceGray(random.FloatRange(m.darkest, m.lightest))
			m.context.FillText(string(c), 0, 0)
			m.context.Restore()
			x += m.charWidth
		}

		x += m.charWidth
	}
	return y + m.lineHeight
}

func (m *Manifesto) copyEffects() {
	if m.copySettings.dither {
		m.logAction("  dithering")
		m.context.DitherAtkinson()
	}
	for range m.copySettings.copyCount {
		m.logAction("  blur", m.copySettings.blur)
		m.context.GaussianBlur(m.copySettings.blur)
		for range m.copySettings.sharpCount {
			m.logAction("    sharpen")
			m.context.Sharpen()
		}
	}

	m.logAction("  brightness", m.copySettings.brightness)
	m.context.Brightness(m.copySettings.brightness)
	m.logAction("  contrast", m.copySettings.contrast)
	m.context.Contrast(m.copySettings.contrast)
	m.context.Grayscale()
}

func (m *Manifesto) logAction(args ...any) {
	if m.verbose {
		fmt.Println(args...)
	}
}
