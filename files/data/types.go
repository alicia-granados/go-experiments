package data

import "fmt"

/* type integer  = int
type json = map[string]string

var x integer */

type distanceMiles float64
type distanceKm float64

/* func ToKm(miles distanceMiles) distanceKm {
	return distanceKm(1.60934 * miles)
}*/

// Mtehods
func (miles distanceMiles) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distanceMiles {
	return distanceMiles(km / 1.60934)
}
func Types() {
	d := 34.4
	fmt.Println(d)

	d1 := distanceMiles(4.05)
	fmt.Println(d1)

	distanceKm := d1.ToKm()
	fmt.Println(distanceKm)

	distanceMiles := distanceKm.ToMiles()
	fmt.Println(distanceMiles)
}

type location string

// global function
/* func DistanceTo(origin location, destination location) distanceMiles {
	fmt.Printf("origin %v destination %v", origin, destination)
	return 10
} */

// method
func (origin location) DistanceTo(destination location) distanceMiles {
	fmt.Printf("origin %v destination %v", origin, destination)
	return 10.00
}

func locationTest() {
	nyc := location("33.42 , 34.34")
	// DistanceTo(nyc, location("23, -24")) global function
	nyc.DistanceTo(location("-23,-44"))
	print(nyc)
}
