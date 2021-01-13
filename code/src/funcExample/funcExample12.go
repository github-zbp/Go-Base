package funcExample

import "fmt"

func Run12() (err error){
	// 在defer中将Run12从Panic异常中断中恢复，并将这panic信息赋值给err
	defer func(){
		if p := recover(); p != nil {	// p是一个接口类型，也就是说p可能是任何类型形式的异常
			fmt.Printf("%T : %v\n\n",p,p)
			err = fmt.Errorf("发生内部错误 : %v", p)
		}
	}()
	//panic("A test panic")
	a := 0
	b := 1
	fmt.Println(b/a)
	fmt.Println("这一行打印不会被执行")
	return err
}
