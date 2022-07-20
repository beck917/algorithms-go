package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToInt(ipstr string) int {
	numList := strings.Split(ipstr, ".")

	ip1, _ := strconv.Atoi(numList[0])
	ip2, _ := strconv.Atoi(numList[1])
	ip3, _ := strconv.Atoi(numList[2])
	ip4, _ := strconv.Atoi(numList[3])

	return ip1<<24 | ip2<<16 | ip3<<8 | ip4
}

// 因为是32位整数, 所以只要位移后, 取最后8位即可
func intToIp(ip uint32) string {
	ip1 := ip >> 24 & 0xff
	ip2 := ip >> 16 & 0xff
	ip3 := ip >> 8 & 0xff
	ip4 := ip & 0xff

	return fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
}

func main() {
	fmt.Println(ipToInt("192.168.2.11"))

	fmt.Println(intToIp(3232236043))
}
