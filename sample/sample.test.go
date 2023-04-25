package sample

import (
    "fmt"

    useragentparser "github.com/hardikshah197/user-agent-parser"
)

func main() {
	for _, userAgent := range Sample_inpust {
		details := useragentparser.Parser(userAgent)

		fmt.Println("device details: ", details)
	}
}