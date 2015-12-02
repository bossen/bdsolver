package earthmover
 
import (
    "errors"
    "math"
)

// This was taken completely from:
// http://rosettacode.org/wiki/Gaussian_elimination#Go

func GaussPartial(a0 [][]float64, b0 []float64) ([]float64, error) {
    // make augmented matrix
    m := len(b0)
    a := make([][]float64, m)
    for i, ai := range a0 {
        row := make([]float64, m+1)
        copy(row, ai)
        row[m] = b0[i]
        a[i] = row
    }
    // WP algorithm from Gaussian elimination page
    // produces row-eschelon form
    for k := range a {
        // Find pivot for column k:
        iMax := k
        max := math.Abs(a[k][k])
        for i := k + 1; i < m; i++ {
            if abs := math.Abs(a[i][k]); abs > max {
                iMax = i
                max = abs
            }
        }
        if a[iMax][k] == 0 {
            return nil, errors.New("singular")
        }
        // swap rows(k, i_max)
        a[k], a[iMax] = a[iMax], a[k]
        // Do for all rows below pivot:
        for i := k + 1; i < m; i++ {
            // Do for all remaining elements in current row:
            for j := k + 1; j <= m; j++ {
                a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
            }
            // Fill lower triangular matrix with zeros:
            a[i][k] = 0
        }
    }
    // end of WP algorithm.
    // now back substitute to get result.
    x := make([]float64, m)
    for i := m - 1; i >= 0; i-- {
        x[i] = a[i][m]
        for j := i + 1; j < m; j++ {
            x[i] -= a[i][j] * x[j]
        }
        x[i] /= a[i][i]
    }
    return x, nil
}
