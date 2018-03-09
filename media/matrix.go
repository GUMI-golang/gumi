package media

import (
	"math"
	"fmt"
)

type Matrix [][]float64

func NewMatrix(w, h int) Matrix {
	mat := make([][]float64, w)
	for i := range mat {
		mat[i] = make([]float64, h)
	}
	return Matrix(mat)
}
func (s Matrix) String() (res string) {
	var w, h = s.Size()
	for y := 0; y < h; y ++{
		for x := 0; x < w; x ++{
			res += fmt.Sprintf("%6.2f ", s[x][y])
		}
		res += "\n"
	}
	return res
}
func (s *Matrix) Clear(clearValue float64){
	w, h := s.Size()
	for x := 0; x < w; x ++{
		for y := 0; y < h; y ++{
			(*s)[x][y] = clearValue
		}
	}
}
func (s Matrix) Size() (w, h int) {
	return len(s), len(s[0])
}
func (s Matrix) Normal() Matrix {
	var mat = NewMatrix(s.Size())
	var sum = s.AbsSum()
	for x, vert := range s {
		for y, elem := range vert {
			mat[x][y] = elem / sum
		}
	}
	return mat
}
func (s Matrix) Sum() (sum float64) {
	for _, hori := range s {
		for _, elem := range hori {
			sum += elem
		}
	}
	return sum
}
func (s Matrix) AbsSum() (sum float64) {
	for _, hori := range s {
		for _, elem := range hori {
			sum += math.Abs(elem)
		}
	}
	return sum
}