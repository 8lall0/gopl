package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA{0xFF, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xFF, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xFF, 0xff},
}

const (
	whiteIndex = 0
	blackIndex = 1
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	var cycles int
	cyclesStr, ok := r.Form["cycles"]
	if ok {
		cycles, _ = strconv.Atoi(cyclesStr[0])
	} else {
		cycles = 5
	}

	lissajous(w, cycles)
}

func lissajous(out io.Writer, cyc int) {
	var cycles = float64(cyc)

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{
		LoopCount: nframes,
	}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i)%4+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}
