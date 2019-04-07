package utility

import "fmt"

func PrintTitle(title string) {
	title = " " + title + " "
	LineLength := 50
	dashLength := (LineLength - len(title)) / 2

	// 标题前换一行
	out := "\n"

	// 前部------------
	for i := 0; i < dashLength; i++ {
		out += "-"
	}

	// 标题
	out += title

	// 后部----------
	for i := 0; i < dashLength; i++ {
		out += "-"
	}

	// 两个换行
	out += "\n\n"
	fmt.Printf(out)
}
