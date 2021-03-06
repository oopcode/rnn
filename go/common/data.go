package common

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/gonum/matrix/mat64"
)

//////////////////////////////////////////////////////////////////////////////
//
// This file contains various functions that get you train/test data.
//
//////////////////////////////////////////////////////////////////////////////

// GetIris returns you X (samples) and Y (labels) for the Iris
// dataset. Labels are one-hot encoded.
func GetIris() (input, expected *mat64.Dense) {
	input = LoadFromCSV("data/iris_x.csv", 4)
	expected = LoadFromCSV("data/iris_y.csv", 3)
	return
}

// LoadFromCSV returns a matrix restored from a file (@path) which must be a
// valid csv with each line representing a float64-vector of length @vectorLen.
func LoadFromCSV(path string, vectorLen int) (out *mat64.Dense) {
	// Prepare storage for points
	points := [][]float64{}
	f, _ := os.Open(path)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	i := 0
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		j := 0
		points = append(points, make([]float64, vectorLen))
		for value := range record {
			points[i][j], err = strconv.ParseFloat(record[value], 64)
			j++
		}
		i++
	}
	out = mat64.NewDense(len(points), vectorLen, nil)
	for idx, point := range points {
		out.SetRow(idx, point)
	}
	return out
}

// GetAbstractTimeSeries1Step creates a time-series dataset. The X sample
// pattern occurs twice in the dataset and predicts different things depending
// on the X-1 pattern. Thus a neural network needs to have at least a 1-step
// memory.
func GetAbstractTimeSeries1() (input, expected *mat64.Dense) {
	input = mat64.NewDense(6, 4, []float64{
		1., 0, 0, 0,
		0, 1., 0, 0, // X-1
		0, 0, 1., 0, // X
		0, 0, 0, 1., // X-1
		0, 0, 1., 0, // X
		0, 1., 0, 0,
	})
	expected = mat64.NewDense(6, 4, nil)
	expected.SetRow(0, []float64{0, 1., 0, 0})
	expected.SetRow(1, []float64{0, 0, 1., 0})
	expected.SetRow(2, []float64{0, 0, 0, 1.})
	expected.SetRow(3, []float64{0, 0, 1., 0})
	expected.SetRow(4, []float64{0, 1., 0, 0})
	expected.SetRow(5, []float64{1., 0, 0, 0})
	return
}

// GetAbstractTimeSeries2 creates a time-series dataset. The X sample
// pattern occurs twice in the dataset and predicts different things depending
// on the X-2 pattern. Thus a neural network needs to have at least a 2-step
// memory.
func GetAbstractTimeSeries2() (input, expected *mat64.Dense) {
	input = mat64.NewDense(6, 4, []float64{
		1., 0, 0, 0, // X-2 <--
		0, 1., 0, 0, // X-1
		0, 0, 0, 1., // X
		0, 1., 0, 0, // X-2 <--
		0, 1., 0, 0, // X-1
		0, 0, 0, 1., // X
	})
	expected = mat64.NewDense(6, 4, []float64{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 1., 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		1., 0, 0, 0,
	})
	return
}

func GetAbstractTimeSeries3() (input, expected *mat64.Dense) {
	input = mat64.NewDense(8, 4, []float64{
		0, 1., 0, 0, // X-3 <--
		0, 1., 0, 0, // X-2 <--
		0, 0, 1., 0, // X-1
		0, 0, 0, 1., // X
		1., 0, 0, 0, // X-3 <--
		0, 1., 0, 0, // X-2 <--
		0, 0, 1., 0, // X-1
		0, 0, 0, 1., // X
	})
	expected = mat64.NewDense(8, 4, []float64{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 1., 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		1., 0, 0, 0,
	})
	return
}
