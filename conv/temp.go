package conv

// convert celsius to fahrenheit
func CToF(t Celsius) Fahrenheit {
	return Fahrenheit((t * 9 / 5) + 32)
}

// convert celsius to kelvin
func CToK(t Celsius) Kelvin {
	return Kelvin(t + 273.15)
}

// convert fahrenheit to celsius
func FToC(t Fahrenheit) Celsius {
	return Celsius((t - 32) * 5 / 9)
}

// convert fahrenheit to kelvin
func FToK(t Fahrenheit) Kelvin {
	return Kelvin((t-32)*5/9 + 273.15)
}

// convert kelvin to celsius
func KToC(t Kelvin) Celsius {
	return Celsius(t - 273.15)
}

// convert kelvin to fahrenheit
func KToF(t Kelvin) Fahrenheit {
	return Fahrenheit((t-273.15)*9/5 + 32)
}
