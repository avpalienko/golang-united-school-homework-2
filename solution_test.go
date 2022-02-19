package square

import (
    "fmt"
    "math"
    "testing"
)

func TestCalcSquare(t *testing.T) {
    type args struct {
        sideLen  float64
        sidesNum SidesCount
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        // TODO: Add test cases.
        {"square", args{10, 4}, 100},
        {"triangle", args{10, 3}, 43.3012702},
        {"circle", args{10, 0}, math.Pi * 100},
        {"anything else", args{10, -1}, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CalcSquare(tt.args.sideLen, tt.args.sidesNum); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
                t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
            }
        })
    }
}
