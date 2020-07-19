package qiniu

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"testing"
	"time"
)

const (
	AK string = ""
	SK string = ""
)

// 通过表单方式进行上传
func TestUploadByForm(t *testing.T) {
	localFile := "wx.png"
	bucket := "python_wx_pic"
	key := "wx-by-form.png"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "wx logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)

}

// 七牛云大文件分片上传
func Test02(t *testing.T) {
	localFile := "wx.png"
	bucket := "python_wx_pic"
	key := "wx-by-form.png"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(AK, SK)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

// 七牛云私有仓库下载
func Test03(t *testing.T) {
	policy := storage.PutPolicy{}
	policy.ForceSaveKey = true

	mac := qbox.NewMac(AK, SK)
	deadline := time.Now().Add(time.Second * 3600).Unix()
	privateURL := storage.MakePrivateURL(mac, "http://hwtblog.cn", "wx-by-form.png", deadline)
	fmt.Println(privateURL)
}
