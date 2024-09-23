package main

func main() {
	/*
		美国: ASCII码 _bbbbbbb 存yang用4字节

		拉丁: 扩展ASCII码 bbbbbbbb  存yang用4字节

		中国: GBK bbbbbbbb bbbbbbbb  两个字节存一个符号  存yang用8字节

		unicode编码: bbbbbbbb bbbbbbbb bbbbbbbb bbbbbbbb 存yang用16个字节

		UTF8编码(动态unicode编码): 存yang用4字节, 存阳用3字节
			针对ASCII码表符号: 用一个字节
			非ASCII表符号: 三字节
	*/

}
