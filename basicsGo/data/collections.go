package data

import "fmt"

var Countries [10]string

var Slice []int

var Codes map[int]bool

func init() {
	Countries[0] = "argentina"
	Countries[1] = "brazil"
	Countries[8] = "usa"

	qty := len(Countries)
	fmt.Println(qty)
	fmt.Println("Countries saved")
}
