package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette []color.Color

const DefaultCycles = 5.0

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 256; i++ {
		palette = append(palette, color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 0xFF,
		})
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}

			var cyclesFloat float64
			cycles, ok := r.Form["cycles"]
			if !ok {
				cyclesFloat = DefaultCycles
			} else {
				cyclesFloat, _ = strconv.ParseFloat(cycles[0], 64)
			}
			lissajous(w, cyclesFloat)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout, DefaultCycles)

}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001 //	角度分辨率
		size    = 100   //	图像画布包含 [-size..+size]
		nframes = 64    //	动画中的帧数
		delay   = 8     //	以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
