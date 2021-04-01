package demo1

// driver 下载地址 http://npm.taobao.org/mirrors/chromedriver/

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

const (
	//设置常量 分别设置chromedriver.exe的地址和本地调用端口
	seleniumPath = `F:\head_less_browser\chromedriver.exe`
	port         = 9515
)

func main() {
	//1.开启selenium服务
	//设置selium服务的选项,设置为空。根据需要设置。
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//延迟关闭服务
	defer service.Stop()

	//2.调用浏览器
	//设置浏览器兼容性，我们设置浏览器名称为chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//调用浏览器urlPrefix: 测试参考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//延迟退出chrome
	defer wd.Quit()

	//3.对页面元素进行操作
	//获取百度页面
	if err := wd.Get("https://www.baidu.com/"); err != nil {
		panic(err)
	}
	//找到百度输入框id
	we, err := wd.FindElement(selenium.ByID, "kw")
	if err != nil {
		panic(err)
	}
	//向输入框发送文字
	err = we.SendKeys("天下第一")
	if err != nil {
		panic(err)
	}

	//找到百度提交按钮id
	we, err = wd.FindElement(selenium.ByID, "su")
	if err != nil {
		panic(err)
	}
	//点击提交
	err = we.Click()
	if err != nil {
		panic(err)
	}

	//睡眠20秒后退出
	time.Sleep(20 * time.Second)
}
