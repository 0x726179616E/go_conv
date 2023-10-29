package conv

// convert meters to feet
func MToFt(d Meters) Feet {
	return Feet(d * 3.280839895)
}

// convert feet to meters
func FtToM(d Feet) Meters {
	return Meters(d / 3.280839895)
}
