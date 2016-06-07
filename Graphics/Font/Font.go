package Font

import (
	"GT/Graphics/Image"
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"math"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi float64 = 72
	//dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile string = "../Graphics/Font/Times_New_Roman_Normal.ttf"
	//fontfile = flag.String("fontfile", "./Times_New_Roman_Normal.ttf", "filename of the ttf font")
	hinting string = "none"
	//hinting  = flag.String("hinting", "none", "none | full")
	size float64 = 100
	//size     = flag.Float64("size", 100, "font size in points")
	spacing float64 = 1.5
	//spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb bool = false
	//wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

func loadFont() image.Image {

	var text = []string{
		"ACDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`0123456789-=~!@#$%^&*()_+[]\\{}|;':\",./<>?",
	}

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return nil
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	i := f.Index('l')
	hmet := f.HMetric(100, i)
	vmet := f.VMetric(100, i)

	fmt.Println("bounds")
	// fmt.Println(int(f.Bounds(100).Min.Y) - int(f.Bounds(100).Max.Y))
	fmt.Println(int(f.Bounds(100).Min.X) - int(f.Bounds(100).Max.X))
	fmt.Println("hmetric: ")
	fmt.Println(hmet)
	fmt.Println(float32(hmet.AdvanceWidth))
	fmt.Println("vmetric: ")
	fmt.Println(vmet)
	fmt.Println(int(vmet.AdvanceHeight))

	totalWidth := 0
	for _, val := range text[0] {
		i := f.Index(val)
		hmet := f.HMetric(100, i)
		totalWidth += int(hmet.AdvanceWidth)
	}

	fmt.Printf("total width: %d\n", totalWidth)
	// Draw the background and the guidelines.
	fg, bg := image.Black, image.Transparent
	// ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	// if *wonb {
	// 	fg, bg = image.White, image.Black
	// 	ruler = color.RGBA{0x22, 0x22, 0x22, 0xff}
	// }
	const imgW, imgH = 640, 480
	rgba := image.NewRGBA(image.Rect(0, 0, totalWidth, int(math.Abs(float64(f.Bounds(100).Min.Y)-float64(f.Bounds(100).Max.Y)))))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	// draw the ruler
	// for i := 0; i < 200; i++ {
	// 	rgba.Set(10, 10+i, ruler)
	// 	rgba.Set(10+i, 10, ruler)
	// }

	// Draw the text.
	h := font.HintingNone
	switch hinting {
	case "full":
		h = font.HintingFull
	}
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: h,
		}),
	}
	y := int(math.Abs(float64(f.Bounds(100).Max.Y))) //10 + int(math.Ceil(*size**dpi/72))
	dy := int(math.Ceil(size * spacing * dpi / 72))
	// d.Dot = fixed.Point26_6{
	// 	X: (fixed.I(imgW) - d.MeasureString(title)) / 2,
	// 	Y: fixed.I(y),
	// }
	// d.DrawString(title)
	// y += dy
	for _, s := range text {
		d.Dot = fixed.P(0, y)
		d.DrawString(s)
		y += dy
	}

	return rgba
}

func ReadFonts(path string) {
	Image.AggrImg.AppendImage(loadFont(), "TimesNewRoman")
	Image.AggrImg.Print("./aggr.png")
}
