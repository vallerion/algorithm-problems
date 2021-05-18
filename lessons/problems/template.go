package problems

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainT() {
	reader := bufio.NewReader(os.Stdin)
	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimRight(firstLine, "\n")
	secondLine, _ := reader.ReadString('\n')
	secondLine = strings.TrimRight(secondLine, "\n")
	fmt.Println(firstLine, secondLine)

	length, _ := strconv.Atoi(firstLine)
	strs := strings.Split(secondLine, " ")
	nums := strToInt(strs)
	fmt.Println(length, nums)
}
