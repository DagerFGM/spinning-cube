package renderer

import "math"

const (
	BgASCIIcode    = ' '
	CubeWidth      = 20.0
	Width          = 80
	Height         = 42
	Size           = Width * Height
	IncrementSpeed = 0.6
	DistFromCam    = 100.0
	K1             = 40.0
)

// Renderer struct to hold the z-buffer and ASCII buffer, along with precomputed values for the cube's vertices
type Renderer struct {
	ZBuffer []float64
	Buffer  []byte

	// Precomputed values for the cube's vertices
	X1, X2, X3 float64
	Y1, Y2, Y3 float64
	Z1, Z2, Z3 float64
}

func NewRenderer() *Renderer {
	return &Renderer{
		ZBuffer: make([]float64, Size),
		Buffer:  make([]byte, Size),
	}
}

// UpdateAngles precomputes the sine and cosine values for the given angles A, B, and C, which represent the rotation of the cube around the X, Y, and Z axes respectively. This method updates the precomputed values for the cube's vertices based on the current rotation angles.
func (r *Renderer) UpdateAngles(A, B, C float64) {
	sinA, cosA := math.Sincos(A)
	sinB, cosB := math.Sincos(B)
	sinC, cosC := math.Sincos(C)
	r.X1 = cosB * cosC
	r.X2 = sinA*sinB*cosC - cosA*sinC
	r.X3 = cosA*sinB*cosC + sinA*sinC

	r.Y1 = cosB * sinC
	r.Y2 = sinA*sinB*sinC + cosA*cosC
	r.Y3 = cosA*sinB*sinC - sinA*cosC

	r.Z1 = -sinB
	r.Z2 = sinA * cosB
	r.Z3 = cosA * cosB
}

// CalculateForSurface calculates the projected position of a point on the cube's surface and updates the z-buffer and ASCII buffer accordingly
func (r *Renderer) CalculateForSurface(cubeX, cubeY, cubeZ float64, sASCII byte) {
	x := cubeX*r.X1 + cubeY*r.X2 + cubeZ*r.X3
	y := cubeX*r.Y1 + cubeY*r.Y2 + cubeZ*r.Y3
	z := cubeX*r.Z1 + cubeY*r.Z2 + cubeZ*r.Z3 + DistFromCam

	ooz := 1 / z
	xp := int(Width/2 + K1*ooz*x*2)
	yp := int(Height/2 + K1*ooz*y)
	idx := xp + yp*Width
	if idx >= 0 && idx < Size {
		if ooz > r.ZBuffer[idx] {
			r.ZBuffer[idx] = ooz
			r.Buffer[idx] = sASCII
		}
	}
}
