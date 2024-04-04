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
