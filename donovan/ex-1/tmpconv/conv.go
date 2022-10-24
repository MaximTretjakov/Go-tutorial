package tmpconv

func CToK(c Celsius) Kelvin { return Kelvin(-273.15 + c) }
func KToC(k Kelvin) Celsius { return Celsius((AbsoluteZero * -1) - (k * -1)) }
