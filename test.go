package main

import "fmt"

func main() {
	output := "@debug1212 @CaroBays \"あああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああネコが寝転んだ\"から\"ネコ\"を検出しました。\n本日1回目。"
	fmt.Println(len([]rune(output)))
	fmt.Println(output)
}
