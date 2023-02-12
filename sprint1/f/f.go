// Package main - решение задачи "Палиндром".
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	str = re.ReplaceAllString(str, "")
	str = strings.ToLower(str)
	strR := []rune(str)
	for i := 0; i < len(strR)/2; i++ {
		if strR[i] != strR[len(strR)-i-1] {
			fmt.Println("False")
			return
		}
	}
	fmt.Println("True")
}
