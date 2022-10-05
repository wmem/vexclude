package main

import (
	"fmt"
	"os"

	"github.com/wmem/go-cmlib/strpkg"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(os.Getwd())

	test := strpkg.NewStringSet()
	test.Add("hhh")

	// project_dir := "/home/mem/wsl-kernel-src/wsl-5.10.16.3"
	// if os.IsExist() {

	// }

	// info, err := os.Stat(project_dir + "/.vscode")
	// if err != nil {

	// }

}
