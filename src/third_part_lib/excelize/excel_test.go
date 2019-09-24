package excelize

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "testing"
)

// 创建一个excel（默认重写）
func TestCreate(t *testing.T) {
    f := excelize.NewFile()
    // 模式Sheet1 已经建立好
    f.SetCellValue("Sheet1", "B2", "hello excelize")
    f.SaveAs("./data1.xlsx")
}

// 多行数据填充（默认重写）
func TestCreateMdata(t *testing.T) {
    f := excelize.NewFile()
    for i:=1;i<=20;i++{
        f.SetCellValue("Sheet1",fmt.Sprintf("%s%d",index2column(i),i),i)
    }
    f.SaveAs("./data2.xlsx")
}

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


// 10进制->26进制：用于获取excel横坐标 1->A
func index2column(index int) string{
    if index ==1 {
        return "A"
    }
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

func TestNum2excelNum(t *testing.T) {
    num := index2column(2)
    fmt.Println(num)
}