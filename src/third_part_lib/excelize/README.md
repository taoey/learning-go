# go操作Excel

## 前言

我选用的是[Excelize](https://github.com/360EntSecGroup-Skylar/excelize)这个库，当前时间为2019年9月23日，该库的关注情况如下，可以看出还是比较热门的库，因此我们就使用该库进行excel的相关操作。

![1569224062432](README.assets/1569224062432.png)

## 使用

1.创建excel文件

```go
// 创建一个excel
func TestCreate(t *testing.T) {
    f := excelize.NewFile()
    // 模式Sheet1 已经建立好
    f.SetCellValue("Sheet1", "B2", "hello excelize")
    f.SaveAs("./Book1.xlsx")
}
```

是的，没错，三行建立一个excel，创建excel就是如此简单，下面是结果：

![1569224381650](README.assets/1569224381650.png)



用于将excel表格中列索引转成列号字母，从A对应1开始 

```go
// 10进制->26进制：用于获取excel横坐标
func num2excelNum(index int) string{
    index--
    column := ""
    for index >0{
        if len(column)>0 {
            index--
        }
        column = string(index % 26 + 65)+column
        index = (index - index % 26) / 26
    } 
    return column
}
```

写一个小例子：

```go
// 多行数据填充
func TestCreateMdata(t *testing.T) {
    f := excelize.NewFile()
    for i:=1;i<=20;i++{
        f.SetCellValue("Sheet1",fmt.Sprintf("%s%d",index2column(i),i),i)
    }
    f.SaveAs("./data2.xlsx")
}
```

执行效果如下图：

![1569302108563](README.assets/1569302108563.png)

2.读取Excel文件

根据官网给出的例子，我们把刚刚创建的data2.xlsx复制一份，命名为readdata.xlsx

```go
// 读取excel
func TestReadData(t *testing.T) {
    file, _ := excelize.OpenFile("./data/readdata.xlsx")
    // 读取一个单元格中的信息
    value := file.GetCellValue("Sheet1", "B2")
    fmt.Println(value)

    //读取整个sheet中的信息
    rows := file.GetRows("Sheet1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}
```

打印结果如下：

```
=== RUN   TestReadData
2
1																				
	2																			
		3																		
			4																	
				5																
					6															
						7														
							8													
								9												
									10											
										11										
											12									
												13								
													14							
														15						
															16					
																17				
																	18			
																		19		
																			20	
--- PASS: TestReadData (0.03s)
PASS

```

