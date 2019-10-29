package fs

import (
	"fmt"
	"github.com/yidane/formula/internal/exp"
	"github.com/yidane/formula/opt"
	"math"
	"strconv"
	"testing"
)

func TestAsinFunction_Evaluate(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{[]string{"1", "2", "3", "30", "0.5"}},
	}
	f := NewAsinFunction()
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			for i := 0; i < len(tt.args); i++ {
				var logicalExpression = *exp.NewFloatExpression(tt.args[i])

				result, err := f.Evaluate(nil, []*opt.LogicalExpression{&logicalExpression}...)
				if err != nil {
					t.Fatal(err)
				}

				v, err := result.Float64()
				if err != nil {
					t.Fatal(err)
				}

				v1, _ := strconv.ParseFloat(tt.args[i], 64)
				if r := math.Asin(v1); r != v {
					if !math.IsNaN(r) || !math.IsNaN(v) {
						t.Fatalf("%v!=%v", r, v)
					}
				}
			}
		})
	}
}
