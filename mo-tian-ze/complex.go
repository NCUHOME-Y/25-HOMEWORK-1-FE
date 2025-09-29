package main

import (
	"fmt"
	"math"
)

type Complex struct {
	real float64
	imag float64
}

func ResultComplex(real, imag float64) *Complex {
	return &Complex{real: real, imag: imag}
}

// 加法
func (num *Complex) Add(other *Complex) *Complex {
	return &Complex{
		real: num.real + other.real,
		imag: num.imag + other.imag,
	}
}

// 减法
func (num *Complex) Sub(other *Complex) *Complex {
	return &Complex{
		real: num.real - other.real,
		imag: num.imag - other.imag,
	}
}

// 乘法
func (num *Complex) Mul(other *Complex) *Complex {
	return &Complex{
		real: num.real*other.real - num.imag*other.imag,
		imag: num.real*other.imag + num.imag*other.real,
	}
}

// 除法
func (num *Complex) Div(other *Complex) *Complex {
	denominator := other.real*other.real + other.imag*other.imag
	if denominator == 0 {
		panic("division by zero")
	}

	return &Complex{
		real: (num.real*other.real + num.imag*other.imag) / denominator,
		imag: (num.imag*other.real - num.real*other.imag) / denominator,
	}
}

// 求模长
func (num *Complex) Mod() float64 {
	return math.Sqrt(num.real*num.real + num.imag*num.imag)
}

// 求共轭复数
func (num *Complex) Conjugate() *Complex {
	return &Complex{
		real: num.real,
		imag: -num.imag,
	}
}

// toString
func (num *Complex) String() string {
	if num.imag == 0 {
		return fmt.Sprintf("%g", num.real)
	}
	if num.real == 0 {
		return fmt.Sprintf("%gi", num.imag)
	}
	if num.imag > 0 {
		return fmt.Sprintf("%g+%gi", num.real, num.imag)
	}
	return fmt.Sprintf("%g%gi", num.real, num.imag)
}
