package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var count int = 0

func main() {
	//running := true
	log.Println("please input file path:")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)
	err := ReadLine(command, handler)
	if err == nil {
		log.Println("path=", command)
	} else {
		log.Println("err=", err)
	}

}
func handler(s string) {
	count++
	if strings.Contains(s, "发送验证码") {
		var res string
		time_start := strings.Index(s, "[")
		time_end := strings.Index(s, "]")
		res = SubString(s, time_start, time_end-time_start+1)
		res += "  "
		a := strings.Index(s, "给手机")
		if a > -1 {
			len := strings.Count(s, "")
			res += SubString(s, a, len-a)
			log.Println(res)
		}
	}
}

func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		//line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
