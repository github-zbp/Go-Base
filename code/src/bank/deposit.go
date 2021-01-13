package bank

var Deposit chan int	// 存钱或取钱用到的channel
var Search chan int    // 查看月用到的channel
var amount int		// 余额

func init() {
	Deposit = make(chan int, 100000)
	Search = make(chan int)
	go func(){		// 必须开一个goroutine跑否则会阻塞init方法，也会阻塞main可执行文件
		for {	// 银行会不停的提供存钱和取钱的业务服务
			select {
			case money := <-Deposit: // 当客户端存钱或取钱时才会走到这个case
				amount += money
			case Search <- amount:	// 将余额告诉客户（当客户端要查看余额时才会走到这个case）

			}
			//  如果没有客户存钱或者取钱，那么这个goroutine就会被select阻塞，银行就闲下来没事干了
		}
	}()

}
