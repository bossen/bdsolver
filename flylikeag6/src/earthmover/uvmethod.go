package earthmover

import (
  "coupling"
)

func Uvmethod(node *coupling.Node) float64 {
  first := true

  ulen := len(*node.Adj)
  u := make([]float64, ulen, ulen)
  udefined := make([]bool, ulen, ulen)

  vlen := len((*node.Adj)[0])
  v := make([]float64, vlen, vlen)
  vdefined := make([]bool, ulen, ulen)

  rows := *node.Adj
  cols := (*node.Adj)[0]

  for i := range rows {
    for j := range cols {
      if (rows[i][j].IsBasic) {
        if (first) {
          u[i] = 0
          udefined[i] = true
          first = false
        }

        if (udefined[i]) {
          v[j] = rows[i][j].Prob - u[i]
          vdefined[j] = true
        } else if (vdefined[j]) {
          u[i] = rows[i][j].Prob - v[j]
          vdefined[j] = true
        }
      }
    }
  }
  
  min := rows[0][0].Prob
  current := min
  for i := range rows {
    for j:= range cols {
      current = rows[i][j].Prob - u[i] - v[j]
      if (current < min) {
        min = current
      }
    }
  }
  return min
}
