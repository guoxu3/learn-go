package main

import (
	"fmt"
	"reflect"
	"os"
	"strings"
	"strconv"
	"regexp"
	 iconv "github.com/djimenez/iconv-go"

	//"bufio"
)

// 将字符串赋值给str,并打印字符串
func create_string() {
	str := "abcdefg"
	fmt.Println(str)
}

// 从控制台输入字符串并打印
func print_string()  {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Printf(str)
}

// 将数字类型转换为字符串
func change_num_to_string()  {
	var (
		num1 int
		num2 int8
		num3 int32
		num4 int64
		num5 float32
		num6 float64
		num7 complex64
	)
	num1 = 1
	num2 = 88
	num3 = 90900
	num4 = 32334324234342
	num5 = 1.003
	num6 = 213232132321.435453453453454353543543433232243243242342342
	num7 = 12.65 + 7i

	fmt.Println(num1,num2,num3,num4,num5,num6,num7)

	fmt.Println("+++++++++++++++++++")
	fmt.Println(reflect.TypeOf(num1))
	fmt.Println(reflect.TypeOf(num2))
	fmt.Println(reflect.TypeOf(num3))
	fmt.Println(reflect.TypeOf(num4))
	fmt.Println(reflect.TypeOf(num5))
	fmt.Println(reflect.TypeOf(num6))
	fmt.Println(reflect.TypeOf(num7))

	num1_1 := strconv.FormatInt(int64(num1), 10)
	num2_1 := strconv.FormatInt(int64(num2), 10)
	num3_1 := strconv.FormatInt(int64(num3), 10)
	num4_1 := strconv.FormatInt(int64(num4), 10)
	num5_1 := strconv.FormatFloat(float64(num5),'f',32,64)
	num6_1 := strconv.FormatFloat(num6,'f',64,64)
	num7_1 := fmt.Sprintf("%f", num7)

	/*
	num1_1 := fmt.Sprintf("%d", num1)
	num2_1 := fmt.Sprintf("%d", num2)
	num3_1 := fmt.Sprintf("%d", num3)
	num4_1 := fmt.Sprintf("%d", num4)
	num5_1 := fmt.Sprintf("%f", num5)
	num6_1 := fmt.Sprintf("%f", num6)
	num7_1 := fmt.Sprintf("%f", num7)
	*/


	fmt.Println("+++++++++++++++++++")
	fmt.Println(reflect.TypeOf(num1_1))
	fmt.Println(reflect.TypeOf(num2_1))
	fmt.Println(reflect.TypeOf(num3_1))
	fmt.Println(reflect.TypeOf(num4_1))
	fmt.Println(reflect.TypeOf(num5_1))
	fmt.Println(reflect.TypeOf(num6_1))
	fmt.Println(reflect.TypeOf(num7_1))

	fmt.Println("+++++++++++++++++++")
	fmt.Println(num1_1,num2_1,num3_1,num4_1,num5_1,num6_1,num7_1)
}

// 将bool类型转换为字符串
func change_bool_to_string()  {
	var (
		bool1 bool
		bool2 bool
	)
	bool1 = true
	bool2 = false
	fmt.Println(reflect.TypeOf(bool1))
	fmt.Println(reflect.TypeOf(bool2))

	bool1_1 := strconv.FormatBool(bool1)
	bool2_1 := strconv.FormatBool(bool2)
	//bool1_1 := fmt.Sprintf("%t", bool1)
	//bool2_1:= fmt.Sprintf("%t", bool2)

	fmt.Println(reflect.TypeOf(bool1_1))
	fmt.Println(reflect.TypeOf(bool2_1))

	fmt.Println(bool1_1,bool2_1)
}


// 遍历输出字符串
func print_each_string()  {
	str := "sffdsdsfsdf"
	for _, v := range str {
		p := string(v)
		fmt.Println(p)
	}
}

// 给定任意变量v,判定其是否是字符串
func judge_string()  {
	v := 111
	//v := "fdsffs"
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		fmt.Println(v," is string")
	} else {
		fmt.Println(v, "is not string")
	}
}

// 字符串的操作
func string_action() {
	str1 := "abcfrsetwsggregsg"
	str2 := "tws"

	// 抽取出指定位置的子串
	str3 := str1[2:5]
	fmt.Println(str3)

	// 将两个字符串拼接在一起
	str4 := str1 + str2
	fmt.Println(str4)

	// 将 B 字符串插入到 A 字符串的某个位置
	n := 5
	str5 := str1[:n] + str2 + str1[n:]
	fmt.Println(str5)

	// 移除指定位置的子串
	str6 := strings.Replace(str1, "abc", "", 3)
	fmt.Println(str6)

	// 将字符串逆序
	var str7 string
	for i := len(str1)-1; i >= 0; i-- {
		str7 += string(str1[i])
	}
	fmt.Println(str7)

	// 在 A 字符串中搜索 B 字符串
	result := strings.Contains(str1, str2)
	if result == true {
		fmt.Println("str1 中可以搜索到 str2")
	} else {
		fmt.Println("str1 中搜索不到 str2")
	}
}

// 正则表达式
func regular_expression()  {
	str1 := "abcfrsetwsggregsg"
	//str2 := "aaa bbb ccc"
	// 匹配指定的字符串
	rlt,err := regexp.MatchString("tws", str1)
	if err == nil {
		if rlt == true {
		fmt.Println("可以在str1中匹配到指定字符串")
		}
	} else {
		fmt.Println(err)
	}

	// 指定字符串替换
	str2 := "i am learning Golang, i like it very much!"
	str2Rx := regexp.MustCompile("i ")
	str3 := str2Rx.ReplaceAllString(str2, "I ")
	fmt.Println(str3)

}

// 低级化
func byte_string()  {
	str := "hahaha"

	// string to byte
	str2 := []byte(str)
	fmt.Println(str2)

	// byte to string
	str3 := string(str2)
	fmt.Println(str3)
}


// 编码转换
func encoding()  {
	str := "hello,世界!"

	// GBK 转换为 UTF8 以及反向转换
	str1,err1 := iconv.ConvertString(str, "utf-8", "gbk")
	if err1 == nil {
		fmt.Println(str1)
	}

	str2,err2 := iconv.ConvertString(str1, "gbk", "utf-8")
	if err2 ==nil {
		fmt.Println(str2)
	}
	fmt.Println("+++++++++++++++++++++")

	// UTF8 与 UTF16LE相互转换
	str3,err3 := iconv.ConvertString(str, "utf-8", "utf-16le")
	if err3 == nil {
		fmt.Println(str3)
	}

	str4,err4 := iconv.ConvertString(str3, "utf-16le", "utf-8")
	if err4 == nil {
		fmt.Println(str4)
	}
	fmt.Println("+++++++++++++++++++++")

	// UTF8 与 UTF16BE相互转换
	str5,err5 := iconv.ConvertString(str, "utf-8", "utf-16be")
	if err5 == nil {
		fmt.Println(str5)
	}

	str6,err6 := iconv.ConvertString(str5, "utf-16be", "utf-8")
	if err6 == nil {
		fmt.Println(str6)
	}
	fmt.Println("+++++++++++++++++++++")

	// UTF8 与 UTF32LE相互转换
	str7,err7 := iconv.ConvertString(str, "utf-8", "utf-32le")
	if err7 == nil {
		fmt.Println(str7)
	}

	str8,err8 := iconv.ConvertString(str7, "utf-32le", "utf-8")
	if err8 == nil {
		fmt.Println(str8)
	}
	fmt.Println("+++++++++++++++++++++")

	// UTF8 与 UTF32BE相互转换
	str9,err9 := iconv.ConvertString(str, "utf-8", "utf-32be")
	if err9 == nil {
		fmt.Println(str9)
	}

	str10,err10 := iconv.ConvertString(str9, "utf-32be", "utf-8")
	if err10 == nil {
		fmt.Println(str10)
	}
	fmt.Println("+++++++++++++++++++++")
}

// 编码转换函数
func encoding_func(s, src, dst string)  {
	str,err := iconv.ConvertString(s, src, dst)
	if err == nil {
		fmt.Println(str)
	}
}

func encoding1()  {
	str := "hello,世界"
	encoding_func(str, "utf-8", "gbk")
}

/*
文件读写：
读取（UTF8编码的）文本文件到字符串变量 str
将字符串变量 str 中的内容写入到文本文件（存储为UTF8编码）
如何读取非 UTF8 编码的文本文件（例如 UTF16LE 或 GBK）？
如何将字符串变量 str 中的内容写入到文本文件，但存储时采用非 UTF8 编码（例如 UTF16LE 或 GBK）？
*/
// 文件读写
func file_read_and_write()  {

	// 读文件
	file := "/Users/guoxu/code/go/string/test.txt"
	fin,err := os.Open(file)
	defer fin.Close()
	if err != nil {
		panic("file open failed")
	}

	buff := make([]byte, 1024)
	for n, err := fin.Read(buff); err == nil; n, err = fin.Read(buff) {
		fmt.Print(string(buff[:n]))
		// fmt.Println(reflect.TypeOf(string(buff[:n])))
	}

	// 写文件
	str1 := "hello world\nhello"
	fout, err := os.Create("/Users/guoxu/code/go/string/test1.txt")
	defer fout.Close()
	if err == nil {
		fout.WriteString(str1)
	}


}

/*
了解文本格式化输出
了解格式化文本的输入
 */
func format_input_output()  {
	fmt.Printf("%d\n", 123)
	fmt.Printf("%s\n", string("123"))
	fmt.Printf("%f\n", 123)
}



func main()  {
	// create_string()
	// print_string()
	// change_num_to_string()
	// change_bool_to_string()
	// print_each_string()
	// judge_string()
	// string_action()
	 regular_expression()
	// byte_string()
	// encoding()
	// encoding1()
	// file_read_and_write()
	// format_input_output()
}