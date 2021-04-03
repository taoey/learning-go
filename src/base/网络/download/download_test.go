package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func HttpDownload(url string, filePath string) error {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}

	err = ioutil.WriteFile(filePath, data, 0666) //写入文件
	if err != nil {
		return fmt.Errorf("HttpDownload error: %v", err)
	}

	return nil
}
func HttpDownloadStream(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("HttpDownload error: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HttpDownload error: %v", err)
	}
	return data, nil
}

func TestDownload(t *testing.T) {
	HttpDownload("http://192.168.3.18:8080/group1/MP4/220092.png", "./aa.png")
}
