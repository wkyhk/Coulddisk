package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUploadByFilePath(t *testing.T) {
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶 region 可以在 COS 控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://disk-1316937931.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKIDmNmvoRowpSpxqHVXTA6wgCUWI9pHiz85", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "nkcmjSxBAuMfkY7lANXUDgtrYyCE0X5L", // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	key := "cloud-disk/wailpaper.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/wailpaper.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}
func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://disk-1316937931.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKIDmNmvoRowpSpxqHVXTA6wgCUWI9pHiz85", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "nkcmjSxBAuMfkY7lANXUDgtrYyCE0X5L", // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	key := "cloud-disk/wailpaper2.jpg"
	f, err := os.ReadFile("./img/wailpaper.jpg")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://disk-1316937931.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDmNmvoRowpSpxqHVXTA6wgCUWI9pHiz85",
			SecretKey: "nkcmjSxBAuMfkY7lANXUDgtrYyCE0X5L",
		},
	})
	key := "cloud-disk/test.mp4"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	//UploadID := v.UploadID //1679985924aab87bb92dd8f4d9cd6a7e92263491620406d247b1fda2163d7ca5c39fa13baa
	UploadID := v.UploadID //16800018274e2a827d7487f36c0c5f6644bed29f9504fb8bdbff43500c660d5ffd14b6b7eb
	fmt.Println(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://disk-1316937931.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDmNmvoRowpSpxqHVXTA6wgCUWI9pHiz85",
			SecretKey: "nkcmjSxBAuMfkY7lANXUDgtrYyCE0X5L",
		},
	})

	key := "cloud-disk/test.mp4"

	UploadID := "16800018274e2a827d7487f36c0c5f6644bed29f9504fb8bdbff43500c660d5ffd14b6b7eb"
	//f, err := os.ReadFile("./img/0.chunk") //md5 : "dd7816e89d4662aa3fba6834673ab3aa"
	f, err := os.ReadFile("./img/1.chunk") //md5 : "ff043a0db38d2392d14cab8fc772253c"
	if err != nil {
		return
	}
	// opt 可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 2, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
	PartETag := resp.Header.Get("ETag")
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://disk-1316937931.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDmNmvoRowpSpxqHVXTA6wgCUWI9pHiz85",
			SecretKey: "nkcmjSxBAuMfkY7lANXUDgtrYyCE0X5L",
		},
	})
	key := "cloud-disk/test.mp4"
	UploadID := "16800018274e2a827d7487f36c0c5f6644bed29f9504fb8bdbff43500c660d5ffd14b6b7eb"
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "dd7816e89d4662aa3fba6834673ab3aa"}, cos.Object{
		PartNumber: 2, ETag: "ff043a0db38d2392d14cab8fc772253c"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}

}
