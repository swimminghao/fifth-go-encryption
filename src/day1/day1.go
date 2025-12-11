package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

// 测试文件
func main() {
	fmt.Println("des 加解密")
	key := []byte("1234abdd")
	src := []byte("特点: 密文没有规律,  明文分组是和一个数据流进行的按位异或操作, 最终生成了密文")
	cipherText := desEncrypt(src, key)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("解密之后的数据: %s\n", string(plainText))
	fmt.Printf("解密之后的数据: %s\n", plainText)

	fmt.Println("aes 加解密 ctr模式 ... ")
	key1 := []byte("1234abdd12345678")
	cipherText = aesEncrypt(src, key1)
	plainText = aesDecrypt(cipherText, key1)
	fmt.Printf("解密之后的数据: %s\n", string(plainText))
	fmt.Printf("解密之后的数据: %d\n", plainText)
	// 正确显示字节和字符对应关系
	fmt.Println("\n=== UTF-8 编码分析 ===")
	fmt.Printf("总长度: %d 字节\n", len(plainText))

	// 将字节转换为字符串，然后按rune遍历（这才是正确的方式）
	str := string(plainText)
	fmt.Printf("字符数（rune）: %d\n", utf8.RuneCountInString(str))

	fmt.Println("\n索引 | 字符 | UTF-8 字节 | Unicode码点")
	fmt.Println("-----|------|------------|------------")

	byteIndex := 0
	for i, r := range str {
		// 获取当前rune的UTF-8字节
		runeBytes := []byte(string(r))

		fmt.Printf("%4d |  %c  | ", i, r)

		// 显示UTF-8字节
		for j, b := range runeBytes {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%3d(0x%02X)", b, b)
		}

		// 对齐输出
		if len(runeBytes) == 1 {
			fmt.Print("          ") // 单字节字符
		} else if len(runeBytes) == 2 {
			fmt.Print("       ") // 双字节字符
		}
		// 三字节字符不需要额外空格

		// 显示Unicode码点
		fmt.Printf(" | U+%04X\n", r)

		byteIndex += len(runeBytes)
	}
	fmt.Println("maxInt:%d\n", math.MaxInt)
	fmt.Println("maxInt:%d\n", strconv.IntSize)
}
