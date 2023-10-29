package conv

import "fmt"

// temperature types
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// weight types
type Kilograms float64
type Pounds float64

// distance types
type Meters float64
type Feet float64

// volume types
type Liters float64
type Gallons float64

// temperature methods
func (t Celsius) String() string {
	return fmt.Sprintf("%gC", t)
}

func (t Fahrenheit) String() string {
	return fmt.Sprintf("%gF", t)
}

func (t Kelvin) String() string {
	return fmt.Sprintf("%gK", t)
}

// weight methods
func (w Kilograms) String() string {
	return fmt.Sprintf("%gkg", w)
}

func (w Pounds) String() string {
	return fmt.Sprintf("%glbs", w)
}

// distance methods
func (d Meters) String() string {
	return fmt.Sprintf("%gm", d)
}

func (d Feet) String() string {
	return fmt.Sprintf("%gft", d)
}

// volume methods
func (v Liters) String() string {
	return fmt.Sprintf("%gL", v)
}

func (v Gallons) String() string {
	return fmt.Sprintf("%ggal", v)
}
