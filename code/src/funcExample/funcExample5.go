package funcExample

func Square() func()int{
	var x int
	return func () int {
		square := x*x
		x++
		return square
	}
}
