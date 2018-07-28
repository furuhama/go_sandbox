// Package tutorial is for the chapter 1 of "The Go Programming Language"
package ch1

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

// MyServer2 returns more details compared to MyServer
func MyServer2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/lissajous", handler3)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler2 returns request header, host, remote addr, form
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter returns how many times this server is called except '/counter'
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler3(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles := 5
	for k, v := range r.Form {
		if k == "cycles" {
			cycles, _ = strconv.Atoi(v[0])
		}
	}

	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	const (
		whiteIndex = 0 // first color of palette
		blackIndex = 1 // next color of palette
		// cycles  = 5     // Oscillator X
		res     = 0.001 // resolution of cycle
		size    = 100   // image cambus uses range of [-size..+size]
		nframes = 64    // frame of animation
		delay   = 8     // delay between every-10ms frames
	)
	palette := []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}
	freq := rand.Float64() * 3.0 // freqency of Oscillator Y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // [caution] ignore the encode error
}
