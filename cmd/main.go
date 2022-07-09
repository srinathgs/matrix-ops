package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Matrix() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix(" "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix(" "))
	fmt.Printf("a^-1 = %v\n\n", fi)

}

func SepalLengthMeanMedianMode() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calculate the Median of the variable.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
}

func PetalLengthMeanMedianMode() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	petalLength := irisDF.Col("petal_length").Float()

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(petalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(petalLength, nil)

	// Calculate the Median of the variable.
	medianVal, err := stats.Median(petalLength)
	if err != nil {
		log.Fatal(err)
	}

	// Output the results to standard out.
	fmt.Printf("\nPetal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
}

func SepalLengthVarianceQuantiles() {
	// Open the CSV file.
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()
	// Calculate the Max of the variable.
	minVal := floats.Min(sepalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}

func PetalLengthVarianceQuantiles() {
	// Open the CSV file.
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	petalLength := irisDF.Col("petal_length").Float()
	// Calculate the Max of the variable.
	minVal := floats.Min(petalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(petalLength)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(petalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(petalLength, nil)

	// Sort the values.
	inds := make([]int, len(petalLength))
	floats.Argsort(petalLength, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, petalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, petalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, petalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nPetal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}

func PlotHistogram() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)
	for _, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			h, err := plotter.NewHist(v, 16)
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
}

func PlotBox() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	p := plot.New()
	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"
	//p.Y.Padding = 10 * vg.Inch
	w := vg.Points(50)

	for idx, colName := range irisDF.Names() {
		if colName != "species" {
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")
	if err := p.Save(6*vg.Inch, 9*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}

func BivariateSepalLW() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	sepalLength := irisDF.Col("sepal_length").Float()
	sepalWidth := irisDF.Col("sepal_width").Float()
	corr, _ := stats.Correlation(sepalLength, sepalWidth)
	fmt.Println("Correlation", corr)
}

func ScatterPlot() {
	irisFile, err := os.Open("../data/iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()
	irisDF := dataframe.ReadCSV(irisFile)
	pts := make(plotter.XYs, irisDF.Nrow())
	yVals := irisDF.Col("sepal_width").Float()
	for i, floatVal := range irisDF.Col("sepal_length").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}
	p := plot.New()
	p.X.Label.Text = "sepallength"
	p.Y.Label.Text = "sepalwidth"
	p.Add(plotter.NewGrid())
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "scatter.png"); err != nil {
		log.Fatal(err)
	}
}

func ChiSquare() {
	observed := []float64{260, 135, 105}
	totalObserved := 500.0
	expected := []float64{totalObserved * 0.6, totalObserved * 0.25, totalObserved * 0.15}
	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Println("Χ-square: ", chiSquare)
}

func CalcPVal() {
	observed := []float64{260, 135, 105}
	totalObserved := 500.0
	expected := []float64{totalObserved * 0.6, totalObserved * 0.25, totalObserved * 0.15}
	chiSquare := stat.ChiSquare(observed, expected)
	chiDist := distuv.ChiSquared{
		K:   2.0, // degrees of freedom = number of possible categories - 1
		Src: nil,
	}
	pVal := chiDist.Prob(chiSquare)
	fmt.Printf("P Val: %0.04f\n\n", pVal)
}

func main() {
	// Create a new matrix a.
	Matrix()
	SepalLengthMeanMedianMode()
	PetalLengthMeanMedianMode()
	SepalLengthVarianceQuantiles()
	PetalLengthVarianceQuantiles()
	PlotHistogram()
	PlotBox()
	BivariateSepalLW()
	ScatterPlot()
	ChiSquare()
	CalcPVal()
}
