package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	//扩容
	tmps := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			tmps += "  "
		}
	}
	ss := s + tmps

	bs := []byte(ss)
	j := len(bs) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if bs[i] != ' ' {
			bs[j] = bs[i]
			j--
		} else {
			bs[j] = '0'
			bs[j-1] = '2'
			bs[j-2] = '%'
			j -= 3
		}
	}

	return string(bs)
}
