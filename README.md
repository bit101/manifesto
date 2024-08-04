# manifesto

manifesto is a command line tool that takes an input text file and renders an image that looks like it was typed on an old typewriter and photocopied multiple times.

There are many command line parameters you can use to customize how the output appears. 

Rendering is done in two steps. 

First, the text is written to a surface. Some randomness is applied to simulate the analogue process of typing on a manual typewriter. This includes slightly randomizing the position of each character and the lightness/darkness of the "ink".

Then, there is the "copy effects" stage, which mimics photocopying the document. This can be turned off with the `--no-copy` parameter. Copy effects consist of multiple rounds of blurring and sharpening, followed by adjustments to brightness and contrast. There is also an optional dithering step with further distresses the quality of the output text. All of these are configurable with the various parameters.  

By default, a file called `manifesto.png` is created as the output, but the file name can be configured with the `-o` or `--output` parameter.

## Parameters

      -i, --input string            A text file containing the text of the manifesto.
      -o, --output string           The name of the output png file. (default "manifesto.png")
          --verbose                 Enables extra logging output.
      -w, --width int               The width of the manifesto, in pixels. (default 800)
      -h, --height int              The height of the manifesto, in pixels. (default 800)
      -m, --margin float            The margin around the text, in pixels. (default 50)
      -r, --rotation float          Rotation in degrees, if copy effects are applied.
          --font-face string        The font used to render the text. Monospaced fonts are best. (default "courier")
          --font-height float       The height of the font in pixels. (default 24)
          --font-width float        The width of a single character, in pixels. (default 13)
          --tab-size float          How many spaces is a single tab character. (default 4)
          --line-spacing float      How much space one line of text takes up. 1 is single-spaced, 2 is double-spaced, etc. (default 1)
          --case string             Transform text case. Valid options: 'upper', 'lower'
          --bold                    Uses a bold font.
          --dark-ink float          The darkest ink used. From 0.0 (black) to 1.0 (white).
          --light-ink float         The lightest ink used. From 0.0 (black) to 1.0 (white). (default 0.5)
          --random-spacing float    How much to randomize space between characters. (default 1)
          --random-y float          How much to randomize the vertical position of characters. (default 2)
          --random-y-chance float   The chance a character will be randomly vertically positioned. (default 0.07)
          --no-copy                 Do not apply copy effects.
          --copy-count int          How many times the manifest is re-copied. (default 3)
          --dither                  Applied dithering to further gitch the text during the copy effects phase.
      -b, --blur int                How much blurring is applied in the copy effects stage. (default 3)
          --sharp-count int         How much sharpening is applied after blurring during the copy effects stage. (default 2)
          --brightness float        Brightens or darkens the image in the copy effects phase. (default -0.2)
          --contrast float          Increases or decreases the contrast of the image in the copy effects phase. (default 0.4)
          --help                    shows this help
      -i, --input string            A text file containing the text of the manifesto.
      -o, --output string           The name of the output png file. (default "manifesto.png")
          --verbose                 Enables extra logging output.
      -w, --width int               The width of the manifesto, in pixels. (default 800)
      -h, --height int              The height of the manifesto, in pixels. (default 800)
      -m, --margin float            The margin around the text, in pixels. (default 50)
      -r, --rotation float          Rotation in degrees, if copy effects are applied.
          --font-face string        The font used to render the text. Monospaced fonts are best. (default "courier")
          --font-height float       The height of the font in pixels. (default 24)
          --font-width float        The width of a single character, in pixels. (default 13)
          --tab-size float          How many spaces is a single tab character. (default 4)
          --line-spacing float      How much space one line of text takes up. 1 is single-spaced, 2 is double-spaced, etc. (default 1)
          --case string             Transform text case. Valid options: 'upper', 'lower'
          --bold                    Uses a bold font.
          --dark-ink float          The darkest ink used. From 0.0 (black) to 1.0 (white).
          --light-ink float         The lightest ink used. From 0.0 (black) to 1.0 (white). (default 0.5)
          --random-spacing float    How much to randomize space between characters. (default 1)
          --random-y float          How much to randomize the vertical position of characters. (default 2)
          --random-y-chance float   The chance a character will be randomly vertically positioned. (default 0.07)
          --no-copy                 Do not apply copy effects.
          --copy-count int          How many times the manifest is re-copied. (default 3)
          --dither                  Applied dithering to further gitch the text during the copy effects phase.
      -b, --blur int                How much blurring is applied in the copy effects stage. (default 3)
          --sharp-count int         How much sharpening is applied after blurring during the copy effects stage. (default 2)
          --brightness float        Brightens or darkens the image in the copy effects phase. (default -0.2)
          --contrast float          Increases or decreases the contrast of the image in the copy effects phase. (default 0.4)
          --help                    shows this help

## Releases

Binaries for Linux and MacOS arm64 (Silicon/M-series) are included in the release. Building cross platform with external platform-specific dependencies is hard, so I'm mainly just going to build natively on the machines I have. I'll try to get a Windows build going soon.

## Building

If you want a build on something else, these steps should work:

1. Check out the repo.
2. Run `go mod tidy` which should install the go dependencies.
3. Install the [cairographics](https://cairographics.org) library for your platform and system.
4. You will also need a supported C compiler such as `gcc`. This probably means installing `mingw` on Windows.
5. The `build/Makefile` script has scripts for each platform that should get you started.
