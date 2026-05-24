# 🧊 ASCII Spinning Cube in Go

A highly optimized 3D rendering engine built completely from scratch in Go, outputting right to the terminal.

## Features
* 🚀 Zero allocations in the render loop (GC-friendly).
* 📐 Custom 3D-to-2D projection and Euler angle rotation matrices.
* 🧪 Fully tested math core using table-driven tests.
* 🖥️ Flicker-free rendering using buffered I/O.

## Installation & Running
```bash
git clone [https://github.com/DagerFGM/spinning-cube.git](https://github.com/DagerFGM/spinning-cube.git)
cd spinning-cube
go run ./cmd
