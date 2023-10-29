package conv

// convert kilograms to pounds
func KgToLbs(w Kilograms) Pounds {
	return Pounds(w * 2.20462262185)
}

// convert pounds to kilograms
func LbsToKg(w Pounds) Kilograms {
	return Kilograms(w / 2.20462262185)
}
