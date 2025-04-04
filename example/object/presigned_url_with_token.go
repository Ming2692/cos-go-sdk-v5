package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type URLToken struct {
	SessionToken string `url:"x-cos-security-token,omitempty" header:"-"`
}

func main() {
	// 替换成您的临时密钥
	tak := os.Getenv("SECRETID")
	tsk := os.Getenv("SECRETKEY")
	token := &URLToken{
		SessionToken: "<token>",
	}
	u, _ := url.Parse("https://test-1259654469.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{})
	name := "xx/../exampleobject"
	ctx := context.Background()

	// 方法1 通过 tag 设置 x-cos-security-token
	// Get presigned
	// http Method需要和实际http请求一致，如PUT请求设置成http.MethodPut，GET请求设置成http.MethodGet
	presignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, name, tak, tsk, time.Hour, token)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(presignedURL.String())

	// 方法2 通过 PresignedURLOptions 设置 x-cos-security-token
	opt := &cos.PresignedURLOptions{
		Query:  &url.Values{},
		Header: &http.Header{},
	}
	opt.Query.Add("x-cos-security-token", "<token>")
	// Get presigned
	// http Method需要和实际http请求一致，如PUT请求设置成http.MethodPut，GET请求设置成http.MethodGet
	presignedURL, err = c.Object.GetPresignedURL(ctx, http.MethodGet, name, tak, tsk, time.Hour, opt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(presignedURL.String())
}
