package conv

// convert liters to gallons
func LToGal(v Liters) Gallons {
	return Gallons(v * 0.2641720524)
}

// convert gallons to liters
func GalToL(v Gallons) Liters {
	return Liters(v / 0.2641720524)
}
