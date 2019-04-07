package utility

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FileSize(file string) int {
	f, err := os.Stat(file)
	if os.IsNotExist(err) {
		return 0
	} else {
		return int(f.Size())
	}
}

func FileIsExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFileIfExist(file string) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		err := os.Remove(file)
		CheckErr(err)
		//fmt.Println("Delete file: ", file)
	}
}

func ReadFile(filepath string) (result [][]float64) {
	file, err := os.Open(filepath)
	CheckErr(err)
	defer func() {
		err = file.Close()
		CheckErr(err)
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var one []float64
		line := strings.Split(scanner.Text(), ",")
		for _, i := range line {
			j, err := strconv.ParseFloat(strings.TrimSpace(i), 64)
			CheckErr(err)
			one = append(one, j)
		}
		result = append(result, one)
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return
}

func SaveFile(file string, list [][]string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	CheckErr(err)
	defer func() {
		err = f.Close()
		CheckErr(err)
	}()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.WriteAll(list)
	CheckErr(err)
	//fmt.Println("Save to file: ", file)
}

func SaveAndOverwriteFile(file string, list [][]string) {
	DeleteFileIfExist(file)
	SaveFile(file, list)
}


