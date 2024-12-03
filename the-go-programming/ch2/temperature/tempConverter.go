package temperature

type Celcius float64
type Fahrenheit float64

func CtoF(c Celcius) Fahrenheit {
	return Fahrenheit((c * 9.0 / 5.0) + 32)
}

func FtoC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5.0 / 9.0)
}
