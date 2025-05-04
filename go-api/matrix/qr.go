package matrix

import (
	"gonum.org/v1/gonum/mat"
)

func QRDecomposition(input [][]float64) (qData [][]float64, rData [][]float64, err error) {
	rows := len(input)
	cols := len(input[0])

	// La variable flat guarda la matriz como un array unidimensional porque asi lo maneja la libreria mat 
	flat := make([]float64, 0, rows*cols)
	for _, row := range input {
		flat = append(flat, row...)
	}

	// Creamos la matriz pero como objeto matriz de mat
	a := mat.NewDense(rows, cols, flat)

	// Realizamos la factorización
	var qr mat.QR
	qr.Factorize(a)

	// Asignamos la descomposicion Q y R a las variables q y r
	var q mat.Dense
	var r mat.Dense
	qr.QTo(&q)
	qr.RTo(&r)

	// Como son matrices en forma array unidimensional las convertimos a matrices 2D para que sirva como input para la API en node
	qData = MatArrayToMatrix(&q)
	rData = MatArrayToMatrix(&r)

	return qData, rData, nil
}

func MatArrayToMatrix(m *mat.Dense) [][]float64 {
	r, c := m.Dims()
	data := make([][]float64, r)
	for i := 0; i < r; i++ {
		data[i] = make([]float64, c)
		for j := 0; j < c; j++ {
			data[i][j] = m.At(i, j)
		}
	}
	return data
}