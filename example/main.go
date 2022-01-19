package main

import (
	gsl "github.com/jtejido/ggsl"
	"github.com/jtejido/stats/dist/continuous"
	//"github.com/jtejido/stats/dist/continuous/directional"
	"go-hep.org/x/hep/hbook"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	//"math"
	// "math/rand"
	//"github.com/jtejido/stats"
	"strings"
)

type plotWindow struct {
	xmin, xmax, ymin, ymax float64
}

type internalFunc struct {
	f          func(float64) float64
	xmin, xmax float64
	w          plotWindow
}

var (
	xmin      = gsl.Float64Eps
	xmax      = 5.
	lineColor = color.RGBA{204, 119, 34, 255}
	pw        = plotWindow{xmin: xmin, xmax: xmax, ymin: gsl.Float64Eps, ymax: 1}
	nSamples  = 1e6 // number of samples
	nBins     = 200
	dist, _   = continuous.NewGompertz(.1, 1)
	// continuous.NewJohnsonSUWithSource(-2, 2, 1.1, 1.5, rand.NewSource(132214534))
	title    = "Gompertz(.1, 1)"
	filename = "Gompertz"
	format   = "jpg"
)

func main() {
	// create a 1-dim histogram of float64s
	hist := hbook.NewH1D(nBins, xmin, xmax)
	for i := 0; i < int(nSamples); i++ {
		hist.Fill(dist.Rand(), 1)
	}

	v := make(plotter.Values, 1000000)
	for i := range v {
		v[i] = dist.Rand()
	}

	pdf := &internalFunc{f: dist.Probability, xmin: xmin, xmax: xmax, w: pw}
	cdf := &internalFunc{f: dist.Distribution, xmin: xmin, xmax: xmax, w: pw}

	// render and save
	//renderHist(hist)
	renderHist2(v, pdf)
	// renderFunc(pdf, "pdf")
	renderFunc(cdf, "cdf")

}

func renderHist(h *hbook.H1D) {
	p := hplot.New() // create hplot.Plot
	p.Title.Text = title
	hh := hplot.NewH1D(h) // create a plotter for the histogram
	hh.Color = lineColor
	p.Add(hh, hplot.NewGrid())
	const (
		width  = 4 * vg.Inch
		height = 4 * vg.Inch
	)

	save(p, strings.ToLower(filename)+"."+format, width, height)
}

func renderHist2(v plotter.Values, pdf *internalFunc) {
	p := hplot.New() // create hplot.Plot
	p.Title.Text = title
	h, err := plotter.NewHist(v, 200)
	if err != nil {
		panic(err)
	}
	h.FillColor = color.Gray{200}
	h.LineStyle.Color = color.Gray{150}
	h.Normalize(1)
	p.Add(h, hplot.NewGrid())
	f := hplot.NewFunction(pdf.f) // create hplot.Function
	f.Color = lineColor

	f.XMin = pdf.xmin
	f.XMax = pdf.xmax
	// set plot window
	p.X.Min = pdf.w.xmin
	p.X.Max = pdf.w.xmax
	p.Y.Min = pdf.w.ymin
	p.Y.Max = pdf.w.ymax
	p.Add(f, hplot.NewGrid())

	const (
		width  = 4 * vg.Inch
		height = 4 * vg.Inch
	)

	save(p, strings.ToLower(filename)+"."+format, width, height)
}

func renderFunc(f *internalFunc, label string) {
	p := hplot.New() // create hplot.Plot
	p.Title.Text = title

	hf := hplot.NewFunction(f.f) // create hplot.Function
	hf.Color = lineColor
	hf.XMin = f.xmin
	hf.XMax = f.xmax

	// set plot window
	p.X.Min = f.w.xmin
	p.X.Max = f.w.xmax
	p.Y.Min = f.w.ymin
	p.Y.Max = f.w.ymax

	p.Add(hf, hplot.NewGrid())

	const (
		width  = 5 * vg.Inch
		height = 5 * vg.Inch
	)

	save(p, strings.ToLower(filename)+"_"+label+"."+format, width, height)
}

func save(hp *hplot.Plot, filename string, width, height vg.Length) {
	if err := hp.Save(width, height, filename); err != nil {
		panic(err)
	}
}
