package renderer_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/DagerFGM/spinning-cube/internal/renderer"
)

// vector represents a 3D coordinate or axis
type vector struct {
	x, y, z float64
}

func (v vector) String() string {
	return fmt.Sprintf("(%.1f, %.1f, %.1f )", v.x, v.y, v.z)
}

// almostEqual checks if two float64 values are close enough to be considered equal, accounting for floating-point precision issues
func almostEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}

// vectorsMatch checks if two vectors are practically identical
func vectorsMatch(v1, v2 vector) bool {
	return almostEqual(v1.x, v2.x) && almostEqual(v1.y, v2.y) && almostEqual(v1.z, v2.z)
}

// checkAxis is a helper function to compare the computed axis vector with the expected one and report any discrepancies in the test output
func checkAxis(t *testing.T, axisName string, got, want vector) {
	t.Helper()
	if !vectorsMatch(got, want) {
		t.Errorf("%s axis missmatch: Got %s, but want %s", axisName, got, want)
	}
}

func TestUpdateAngles(t *testing.T) {
	tests := []struct {
		name    string
		A, B, C float64 // Angles
		// vector axis
		wantX, wantY, wantZ vector
	}{
		{
			name: "Zero rotation (A = 0, B = 0, C = 0)",
			A:    0, B: 0, C: 0,
			wantX: vector{1, 0, 0},
			wantY: vector{0, 1, 0},
			wantZ: vector{0, 0, 1},
		},
		{
			name: "90 degree rotation on X axis (A = Pi/2)",
			A:    math.Pi / 2, B: 0, C: 0,
			wantX: vector{1, 0, 0},
			wantY: vector{0, 0, -1},
			wantZ: vector{0, 1, 0},
		},
		{
			name: "90 degree rotation on Y axis (B = Pi/2)",
			A:    0, B: math.Pi / 2, C: 0,
			wantX: vector{0, 0, 1},
			wantY: vector{0, 1, 0},
			wantZ: vector{-1, 0, 0},
		},
	}

	renderer := renderer.NewRenderer()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			renderer.UpdateAngles(tc.A, tc.B, tc.C)

			gotX := vector{renderer.X1, renderer.X2, renderer.X3}
			gotY := vector{renderer.Y1, renderer.Y2, renderer.Y3}
			gotZ := vector{renderer.Z1, renderer.Z2, renderer.Z3}

			checkAxis(t, "X", gotX, tc.wantX)
			checkAxis(t, "Y", gotY, tc.wantY)
			checkAxis(t, "Z", gotZ, tc.wantZ)
		})
	}
}
