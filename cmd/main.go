package main

import (
	"bufio"
	"os"
	"time"

	rr "github.com/DagerFGM/spinning-cube/internal/renderer"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("\x1b[2J")
	var A, B, C float64

	rend := rr.NewRenderer()
	for {
		// Clear the ASCII buffer and z-buffer for the new frame
		for i := range rend.Buffer {
			rend.Buffer[i] = rr.BgASCIIcode
		}
		clear(rend.ZBuffer)

		// Update the precomputed values for the cube's vertices based on the current rotation angles
		rend.UpdateAngles(A, B, C)
		for cubeX := -rr.CubeWidth; cubeX < rr.CubeWidth; cubeX += rr.IncrementSpeed {
			for cubeY := -rr.CubeWidth; cubeY < rr.CubeWidth; cubeY += rr.IncrementSpeed {
				rend.CalculateForSurface(cubeX, cubeY, -rr.CubeWidth, '@')
				rend.CalculateForSurface(cubeX, cubeY, rr.CubeWidth, '*')
				rend.CalculateForSurface(rr.CubeWidth, cubeY, cubeX, '&')
				rend.CalculateForSurface(-rr.CubeWidth, cubeY, -cubeX, '%')
				rend.CalculateForSurface(cubeX, rr.CubeWidth, cubeY, '#')
				rend.CalculateForSurface(cubeX, -rr.CubeWidth, -cubeY, '$')
			}
		}
		writer.WriteString("\x1b[H")
		// Write the ASCII buffer to the terminal, line by line
		for row := range rr.Height {
			s := row * rr.Width
			e := s + rr.Width
			writer.Write(rend.Buffer[s:e])
			writer.WriteByte('\n')
		}
		writer.Flush()
		A += .005
		B += .005
		C += .001
		time.Sleep(6 * time.Millisecond)
	}
}
