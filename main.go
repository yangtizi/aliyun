package main

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yangtizi/aliyun/instances"
)

const strEndpoint = "oss-cn-hangzhou.aliyuncs.com"
const strAccessKeyID = "LTAIbkRnr6mkDoCI"
const strAccessKeySecret = "DEFfqXmalyGtNK4GcVTe0djbq9B6cB"
const strBucketName = "789test"

func main() {
	instances.DescribeInstances()
	// securitygroup.CreateAuthorizeSecurityGroupRequest(securitygroup.CreateCreateSecurityGroupRequest())
}

func copyFile() {
	// 创建OSSClient实例。
	client, err := oss.New(strEndpoint, strAccessKeyID, strAccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketName := strBucketName
	objectName := "余额宝客户端协议.json"
	destObjectName := "Backup/20190831/余额宝客户端协议.json"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 拷贝文件到同一个存储空间的另一个文件。
	r, err := bucket.CopyObject(objectName, destObjectName)
	fmt.Println(r, err)
}

func updateFile() {
	// var strObjectName string
	// var strLocalFile string
	// flag.StringVar(&strObjectName, "o", "ObjectName", "ObjectName")
	// flag.StringVar(&strLocalFile, "l", "LocalFileX", "LocalFileY")
	// flag.Parse()

	strObjectName := "余额宝客户端协议.json"
	strLocalFile := "E:/ALL.sql"

	// 创建OSSClient实例。
	client, err := oss.New(strEndpoint, strAccessKeyID, strAccessKeySecret)
	if err != nil {
		fmt.Println("oss.New Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(strBucketName)
	if err != nil {
		fmt.Println("Bucket Error:", err)
		os.Exit(-1)
	}

	fmt.Println("上传的文件名是", strLocalFile)

	// 上传本地文件。
	err = bucket.PutObjectFromFile(strObjectName, strLocalFile)
	if err != nil {
		fmt.Println("PutObjectFromFile Error:", err)
		os.Exit(-1)
	}
}
