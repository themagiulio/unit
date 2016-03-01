package unit

// Length represents a SI unit of length (in meters, m)
type Length float64

// ...
const (
	// SI
	Yoctometer              = Meter * 1e-24
	Zeptometer              = Meter * 1e-21
	Attometer               = Meter * 1e-18
	Femtometer              = Meter * 1e-15
	Picometer               = Meter * 1e-12
	Nanometer               = Meter * 1e-9
	Micrometer              = Meter * 1e-6
	Millimeter              = Meter * 1e-3
	Centimeter              = Meter * 1e-2
	Decimeter               = Meter * 1e-1
	Meter            Length = 1e0
	Decameter               = Meter * 1e1
	Hectometer              = Meter * 1e2
	Kilometer               = Meter * 1e3
	ScandinavianMile        = Meter * 1e4
	Megameter               = Meter * 1e6
	Gigameter               = Meter * 1e9
	Terameter               = Meter * 1e12
	Petameter               = Meter * 1e15
	Exameter                = Meter * 1e18
	Zettameter              = Meter * 1e21
	Yottameter              = Meter * 1e24

	// US
	Inch         = Meter * 0.0254
	Feet         = Inch * 12
	Yard         = Feet * 3
	Rod          = Yard * 5.5
	Chain        = Rod * 4
	Furlong      = Chain * 10
	Mile         = Meter * 1609.344
	NauticalMile = Meter * 1852

	// space
	LunarDistance    = Kilometer * 384400
	AstronomicalUnit = Meter * 149597870700
)

// Yoctometers returns the length in ym
func (l Length) Yoctometers() float64 {
	return float64(l / Yoctometer)
}

// Zeptometers returns the length in zm
func (l Length) Zeptometers() float64 {
	return float64(l / Zeptometer)
}

// Attometers returns the length in am
func (l Length) Attometers() float64 {
	return float64(l / Attometer)
}

// Femtometers returns the length in fm
func (l Length) Femtometers() float64 {
	return float64(l / Femtometer)
}

// Picometers returns the length in pm
func (l Length) Picometers() float64 {
	return float64(l / Picometer)
}

// Nanometers returns the length in nm
func (l Length) Nanometers() float64 {
	return float64(l / Nanometer)
}

// Micrometers returns the length in microms
func (l Length) Micrometers() float64 {
	return float64(l / Micrometer)
}

// Millimeters returns the length in mm
func (l Length) Millimeters() float64 {
	return float64(l / Millimeter)
}

// Centimeters returns the length in cm
func (l Length) Centimeters() float64 {
	return float64(l / Centimeter)
}

// Decimeters returns the length in dm
func (l Length) Decimeters() float64 {
	return float64(l / Decimeter)
}

// Meters returns the length in m
func (l Length) Meters() float64 {
	return float64(l)
}

// Decameters returns the length in dam
func (l Length) Decameters() float64 {
	return float64(l / Decameter)
}

// Hectometers returns the length in hm
func (l Length) Hectometers() float64 {
	return float64(l / Hectometer)
}

// Kilometers returns the length in km
func (l Length) Kilometers() float64 {
	return float64(l / Kilometer)
}

// ScandinavianMiles returns the length in in scandinavian miles (1 scandmile = 10 km)
func (l Length) ScandinavianMiles() float64 {
	return float64(l / ScandinavianMile)
}

// Megameters returns the length in in Mm
func (l Length) Megameters() float64 {
	return float64(l / Megameter)
}

// Gigameters returns the length in in Gm
func (l Length) Gigameters() float64 {
	return float64(l / Gigameter)
}

// Terameters returns the length in in Tm
func (l Length) Terameters() float64 {
	return float64(l / Terameter)
}

// Petameters returns the length in in Pm
func (l Length) Petameters() float64 {
	return float64(l / Petameter)
}

// Exameters returns the length in in Em
func (l Length) Exameters() float64 {
	return float64(l / Exameter)
}

// Zettameters returns the length in in Zm
func (l Length) Zettameters() float64 {
	return float64(l / Zettameter)
}

// Yottameters returns the length in in Ym
func (l Length) Yottameters() float64 {
	return float64(l / Yottameter)
}

// Inches returns the length in in
func (l Length) Inches() float64 {
	return float64(l / Inch)
}

// Feets returns the length in ft
func (l Length) Feets() float64 {
	return float64(l / Feet)
}

// Yards returns the length in yd
func (l Length) Yards() float64 {
	return float64(l / Yard)
}

// Rods returns the length in rod
func (l Length) Rods() float64 {
	return float64(l / Rod)
}

// Chains returns the length in chain
func (l Length) Chains() float64 {
	return float64(l / Chain)
}

// Furlongs returns the length in furlong
func (l Length) Furlongs() float64 {
	return float64(l / Furlong)
}

// Miles returns the length in mi
func (l Length) Miles() float64 {
	return float64(l / Mile)
}

// NauticalMiles returns the length in nm
func (l Length) NauticalMiles() float64 {
	return float64(l / NauticalMile)
}

// LunarDistances returns the length in ld
func (l Length) LunarDistances() float64 {
	return float64(l / LunarDistance)
}

// AstronomicalUnits returns the length in au
func (l Length) AstronomicalUnits() float64 {
	return float64(l / AstronomicalUnit)
}