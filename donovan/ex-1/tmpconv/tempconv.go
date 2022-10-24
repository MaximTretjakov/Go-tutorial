package tmpconv

import "fmt"

type Kelvin float64
type Celsius float64

const AbsoluteZero Kelvin = -273.15

func (c Celsius) String() string { return fmt.Sprintf("%g", c) }
func (k Kelvin) String() string  { return fmt.Sprintf("%g", k) }
