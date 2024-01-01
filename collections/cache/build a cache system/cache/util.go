// cache/util.go
package cache

import (
	"regexp"
	"strconv"
	"strings"
	"encoding/json"
	"log"
)


const (
	B = 1 << (iota * 10) // 1
	KB // 1024
	MB // 1048576
	GB // 1073741824
	TB
	PB
)


// 解析设置存储大小字符串
func ParseSize(size string) ( int64, string ){

	// 默认大小为 100

	// 利用正则表达式匹配一个或者多个数字
	re, _ := regexp.Compile("[0-9]+")
	// 获取单位: 使用编译好的正则表达式 re，将 size 字符串中匹配的数字字符替换为空字符串
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	// 获取数字: 将 size 字符串中的单位部分 unit 用空字符串替换，即可获取数字部分。最后再将数字转换为 int64 类型
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)

	// 单位转换为大写
	unit = strings.ToUpper(unit)

	var byteNum int64 =0
	// 1KB 100KB 1MB 2MB 1GB, 单位统一为 byte
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}

	if num == 0 {
		log.Println("[WARN] parseSize 仅支持 B、KB、MB、GB、TB、PB")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}

	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}


// GetValSize 计算 value 值大小
func GetValSize(val interface{}) int64 {
	// 野路子：利用 json 包，将 val 序列化为字节数组，然后求字节数组的长度，就知道占用了多少字节了
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}