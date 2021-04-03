package go_qrcode

import (
	"bytes"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"testing"
)

// 通过一个url生成一个二维码
func TestQrcodeMake(t *testing.T) {
	qrcode.WriteFile("https://github.com/Taoey", qrcode.Medium, 256, "./blog_qrcode.png")
}

// 背景色为绿色，前景色为白色的二维码
func TestQrcodeMake02(t *testing.T) {
	qr, err := qrcode.New("https://github.com/Taoey", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./blog_qrcode.png")
	}
}

func TestQrcodeSize(t *testing.T) {
	qrcode.WriteFile("http://www.flysnow.org/", qrcode.Highest, 60, "./pt.png")
}

// 裁剪边框
// 参考：golang 图片处理，剪切，base64数据转换，文件存储 ：http://www.philo.top/2015/03/05/golangImageLibrary/
func TestSubQrcode(t *testing.T) {
	data, _ := qrcode.Encode("https://github.com/Taoey", qrcode.Highest, 60)

	file, _ := os.OpenFile("./subimg.png", os.O_RDWR|os.O_CREATE, 0777)

	canvas, _, _ := image.Decode(bytes.NewReader(data))
	rgbImg := canvas.(*image.Paletted)
	subImg := rgbImg.SubImage(image.Rect(12, 12, 48, 48)).(*image.Paletted)

	png.Encode(file, subImg) //写入文件

}

// 在一个图片的右下角生成二维码
func TestSubQrCodeToPng(t *testing.T) {

	// 背景图片大小
	bgWidth := 1920
	bgHight := 1080

	minSize := math.Min(float64(bgWidth), float64(bgHight))
	// 二维码数据加载处理
	data, _ := qrcode.Encode("https://segmentfault.com/a/1190000018919986", qrcode.Highest, int(minSize*0.1)) //尺寸适应
	canvas, _, _ := image.Decode(bytes.NewReader(data))
	rgbImg := canvas.(*image.Paletted)
	subImg := rgbImg.SubImage(image.Rect(int(minSize*0.02), int(minSize*0.02), int(minSize*0.08), int(minSize*0.08))).(*image.Paletted) // 切掉白边
	subImgBounds := subImg.Bounds()

	//file, _ := os.OpenFile("./data/subimg.png", os.O_RDWR|os.O_CREATE, 0777)
	//png.Encode(file, subImg)       //写入文件

	// 背景图片加载
	imgBackground, _ := os.Open("./data/bg1.png")
	imgBack, _ := png.Decode(imgBackground)
	bgBounds := imgBack.Bounds()

	offset := image.Pt(bgBounds.Max.X-subImgBounds.Max.X-10, bgBounds.Max.Y-subImgBounds.Max.Y-10)
	b := imgBack.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, imgBack, image.ZP, draw.Src)
	draw.Draw(m, rgbImg.Bounds().Add(offset), subImg, image.ZP, draw.Over)

	imgw, _ := os.Create("./data/qr_bg1.png")
	png.Encode(imgw, m)
	defer imgw.Close()
}
