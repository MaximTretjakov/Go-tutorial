package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrIntOverflow = errors.New("integer overflow")

func main() {
	// var signed int8 = math.MaxInt8
	// signed++

	// var unsigned uint8 = math.MaxUint8
	// unsigned++

	// fmt.Println(signed)
	// fmt.Println(unsigned)

	// var signed int8 = math.MaxInt8 + 1
	// var unsigned uint8 = math.MaxUint8 + 1

	// var maxInt int8 = math.MaxInt8
	// var minInt int8 = math.MinInt8
	// fmt.Println(Inc(maxInt))
	// fmt.Println(Dec(minInt))
	// fmt.Println(Add(1, 1))
	fmt.Println(Mul(26, 127))
}

func Inc(counter int8) (int8, error) {
	if counter == math.MaxInt8 {
		return 0, ErrIntOverflow
	}

	return counter + 1, nil
}

func Dec(counter int8) (int8, error) {
	if counter == math.MinInt8 {
		return 0, ErrIntOverflow
	}

	return counter - 1, nil
}

func Add(lhs, rhs int8) (int8, error) {
	if rhs > 0 && lhs > math.MaxInt8-rhs {
		return 0, ErrIntOverflow
	}

	if rhs < 0 && lhs < math.MinInt8-rhs {
		return 0, ErrIntOverflow
	}

	return lhs + rhs, nil
}

func Mul(lhs, rhs int8) (int8, error) {
	if lhs == 0 || rhs == 0 {
		return 0, nil
	}

	if lhs == 1 || rhs == 1 {
		return lhs * rhs, nil
	}

	if (lhs == -1 && rhs == math.MinInt8) || (rhs == -1 && lhs == math.MinInt8) {
		return 0, ErrIntOverflow
	}

	if lhs > math.MaxInt8/rhs {
		return 0, ErrIntOverflow
	} else if lhs < math.MinInt8/rhs {
		return 0, ErrIntOverflow
	}

	return lhs * rhs, nil
}
