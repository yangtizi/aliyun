package oss

import (
	"fmt"
	"strings"

	oossss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oossss.Client

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oossss.ProgressEvent) {
	switch event.EventType {
	// case oossss.TransferStartedEvent:
	// 	fmt.Printf("传输开始, 已完成: %d, 总字节 %d.\n",
	// 		event.ConsumedBytes, event.TotalBytes)
	case oossss.TransferDataEvent:
		fmt.Printf("\r%d%%.        ",
			event.ConsumedBytes*100/event.TotalBytes)
	case oossss.TransferCompletedEvent:
		// fmt.Printf("\n传输完毕, 已完成: %d, 总字节 %d.\n",
		// 	event.ConsumedBytes, event.TotalBytes)
	case oossss.TransferFailedEvent:
		// fmt.Printf("\n传输失败, 已完成: %d, 总字节 %d.\n",
		// 	event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

// InitSDK 初始化SDK
func InitSDK(strAccessKeyID, strAccessKeySecret, strEndpoint string) {
	c, err := oossss.New(strEndpoint, strAccessKeyID, strAccessKeySecret)

	if err != nil {
		fmt.Println("初始化SDK错误:", err)
		return
	}

	client = c
}

// PutObject 快速上传
func PutObject(strBucketName, strObjectName, strLocalFile string) bool {
	if client == nil {
		fmt.Println("没有对应的桶指针")
		return false
	}

	bucket, err := client.Bucket(strBucketName)
	if err != nil {
		fmt.Println("打开桶[", strBucketName, "]错误: ", err)
		return false
	}

	err = bucket.PutObjectFromFile(strObjectName, strLocalFile, oossss.Progress(&OssProgressListener{}))
	if err != nil {
		fmt.Println("上传文件错误: ", err)
		return false
	}

	return true
}

// PutDir 上传文件夹
func PutDir(strBucketName, strObjectName string) bool {
	if client == nil {
		fmt.Println("没有对应的桶指针")
		return false
	}

	bucket, err := client.Bucket(strBucketName)
	if err != nil {
		fmt.Println("打开桶[", strBucketName, "]错误: ", err)
		return false
	}

	err = bucket.PutObject(strObjectName, strings.NewReader(""))
	if err != nil {
		fmt.Println("上传文件错误: ", err)
		return false
	}

	return true
}

// SetCORSRule 跨域
func SetCORSRule(strBucketName string) bool {
	if client == nil {
		fmt.Println("没有对应的桶指针")
		return false
	}

	rule1 := oossss.CORSRule{
		AllowedOrigin: []string{"*"},
		AllowedMethod: []string{"PUT", "GET", "POST", "DELETE", "HEAD"},
		AllowedHeader: []string{"*"},
		ExposeHeader:  []string{},
		MaxAgeSeconds: 200,
	}

	err := client.SetBucketCORS(strBucketName, []oossss.CORSRule{rule1})
	if err != nil {
		fmt.Println("跨域设置错误: ", err)
	}

	return true
}

// ReleaseSDK 释放内存
func ReleaseSDK() {
	client = nil
}
