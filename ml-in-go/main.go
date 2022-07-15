package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Describe() {
	advertFile, err := os.Open("advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()

	advertDF := dataframe.ReadCSV(advertFile)
	advertSummary := advertDF.Describe()

	fmt.Println(advertSummary)
}

func Histogram() {
	advertFile, err := os.Open("advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()

	advertDF := dataframe.ReadCSV(advertFile)
	for _, colName := range advertDF.Names() {
		plotVals := make(plotter.Values, advertDF.Nrow())

		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}
		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	Describe()
	Histogram()
}
