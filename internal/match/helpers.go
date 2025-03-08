package match

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func parse_vector(vec_string string) []float64 {
	parts := strings.Split(vec_string, ",")
	vector := make([]float64, len(parts))
	for i, part := range parts {
		val, err := strconv.ParseFloat(part, 64)
		if err != nil {
			log.Fatalf("error parsing string")
		}
		vector[i] = val
	}
	return vector
}

func distance(x []float64, y []float64) (float64, error) {

	if len(x) != len(y) {
		return 0, fmt.Errorf("vectors must be same size")
	}

	var sum float64
	for i := 0; i < len(x); i++ {
		diff := x[i] - y[i]
		sum += diff * diff
	}
	return math.Sqrt(sum), nil

}
