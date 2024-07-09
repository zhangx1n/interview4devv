package main

import (
	"errors"
	"fmt"
)

/*
题目描述:
请实现一个函数,将点分十进制格式的IPv4地址字符串转换为32位无符号整数(uint32)。
示例:
输入: "172.168.5.1"
输出: 2896692481
要求:
不使用任何第三方库和标准库
仅遍历一次输入字符串!!!
数字和点之间允许有空格,视为合法输入 例如: "172 . 168.    5.1" 是合法输入,应给出正确结果
数字内部有空格视为非法输入 例如: "1 72.168.5.1" 是非法输入,应报错
请提供相应的单元测试
*/

// ConvertIPv4ToUint32 将点分十进制格式的 IPv4 地址转换为 32 位无符号整数
func ConvertIPv4ToUint32(ip string) (uint32, error) {
	var result uint32
	var segmentValue int
	var segmentCount int
	var flag bool // 跟踪当前正在解析的段中是否已经看到过有效的数字

	for i := 0; i < len(ip); i++ {
		char := ip[i]
		if char == '.' {
			if !flag {
				return 0, errors.New("无效的 IP 地址：空段")
			}
			if segmentValue < 0 || segmentValue > 255 {
				return 0, errors.New("无效的 IP 地址：段值超出范围")
			}
			result = (result << 8) | uint32(segmentValue)
			segmentValue = 0
			segmentCount++
			flag = false
		} else if char >= '0' && char <= '9' {
			if flag && ip[i-1] == ' ' {
				return 0, errors.New("无效的 IP 地址：段内有空格")
			}
			segmentValue = segmentValue*10 + int(char-'0')
			flag = true
		} else if char == ' ' {
			// 忽略数字和点之间的空格
		} else {
			return 0, errors.New("无效的 IP 地址：包含非数字字符")
		}
	}

	// 处理最后一个段
	if !flag {
		return 0, errors.New("无效的 IP 地址：空段")
	}
	if segmentValue < 0 || segmentValue > 255 {
		return 0, errors.New("无效的 IP 地址：段值超出范围")
	}
	result = (result << 8) | uint32(segmentValue)
	segmentCount++

	if segmentCount != 4 {
		return 0, errors.New("无效的 IP 地址：段数不正确")
	}

	return result, nil
}

func main() {
	ip := "172.168.5.1"
	result, err := ConvertIPv4ToUint32(ip)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("ip: %s, uint32: %d\n", ip, result)
	}
}
