package funcExample

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Request(url string) (content string, err error){
	// 发生错误是重试
	resp, err := http.Get(url)

	// 正常情况下，如果err不为空而且还要判定具体是哪种错误才能够进行重试，这里只是为了演示，因此只要发生错误就重试
	if err != nil {
		// 定义重试的时间范围是1分钟，1分钟内请求失败会重复请求
		scope := time.Minute		// 是一个Duration对象
		dealine := time.Now().Add(scope)	// 这是1分钟后的Time时间对象

		// 如果超过一分钟就会结束循环
		for tries := 0; time.Now().Before(dealine); tries++ {
			fmt.Printf("重试请求 %s \n", url)
			resp, err = http.Get(url)		// 不要用 :=，因为这里的err要覆盖上面的err才行

			// 重试失败就再重试
			if err != nil {
				// 睡一段时间，睡的时间随着重试次数增加而加长
				time.Sleep(time.Second << tries)
				continue
			}

			// 重试成功
			break
		}
	}

	//  如果重试1分钟还是失败则返回错误
	if err != nil {
		return content, fmt.Errorf("请求url %s 失败: %v", url, err)
	}

	// 读取失败
	var res_bytes []byte
	res_bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return string(res_bytes), fmt.Errorf("读取url %s 的响应失败: %v", url, err)
	}

	return string(res_bytes), err
}