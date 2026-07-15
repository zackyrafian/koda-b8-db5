package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
func Input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func InputInt(prompt string) (int, error) {
	text := Input(prompt)
	return strconv.Atoi(text)
}

func WaitEnter() {
	fmt.Print("\nPress Enter...")
	reader.ReadString('\n')
}