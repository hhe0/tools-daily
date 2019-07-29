package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// 8538beb3fb7f468eb59f17ec2ba98108
	fmt.Println(Md5("hehong"))

	// 8538beb3fb7f468eb59f17ec2ba98108
	fmt.Println(MD5("hehong"))
}

func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func MD5(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
