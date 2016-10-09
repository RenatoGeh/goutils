package chisquare

import (
	"fmt"
	"testing"
)

func TestChiSquare(t *testing.T) {
	chi, df := 16.2, 2
	fmt.Printf("Pr(X^2 <= %f) = %f ~ X^2(%d)\n", chi, ChiSquare(chi, df), df)
}
