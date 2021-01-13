package funcExample

import "log"

func SquareX(x int) (result int){
	//defer func() { result = result * x } ()
	defer func() { log.Print(123)} ()
	defer func() { log.Print(456)} ()
	return x * x
}
