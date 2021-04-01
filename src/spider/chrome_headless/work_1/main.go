package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	_ "github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"strings"
)

const (
	//设置常量 分别设置chromedriver.exe的地址和本地调用端口
	seleniumPath = `F:\head_less_browser\chromedriver.exe`
	port         = 9515

	TARGET_URL = `https://blog.csdn.net/orangefly0214/article/details/81462245`
)

func main() {
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)

	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
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
	wd.Get(TARGET_URL)

	//wd.ExecuteScript(`alert("请在1分钟内输入您的筛选条件，1分钟后程序会进行执行")`,nil)

	source, err := wd.PageSource()

	nodes, err := htmlquery.Parse(strings.NewReader(source))
	title := htmlquery.FindOne(nodes, `//*[@id="articleContentId"]`)
	//fmt.Println(one.FirstChild.Data)
	fmt.Println(htmlquery.InnerText(title))

}
