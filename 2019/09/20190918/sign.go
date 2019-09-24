package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	oldMap := map[string]string{"name": "hehong", "age": "1", "size": "18", "height": "180"}
	fmt.Print(oldMap)

	keys := make([]string, 0)
	for element := range oldMap {
		keys = append(keys, element)
	}

	fmt.Print(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	newMap := make(map[string]string, len(oldMap))
	for _, key := range keys {
		newMap[key] = oldMap[key]
	}

	newMapJson, _ := json.Marshal(newMap)

	paramJson, _ := json.Marshal(newMap)
	fmt.Println(string(newMapJson))

	enc := base64.StdEncoding.EncodeToString([]byte(string(paramJson) + "secret"))
	sign := strings.ToUpper(md5V2(enc))
	fmt.Print(sign)
}

func md5V2(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
