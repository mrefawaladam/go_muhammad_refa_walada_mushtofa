package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'deleteProducts' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ids
 *  2. INTEGER m
 */

func deleteProducts(ids []int32, m int32) int32 {
	freqMap := make(map[int32]int32)
	for _, id := range ids {
		freqMap[id]++
	}

	freqList := make([]int32, 0, len(freqMap))
	for _, freq := range freqMap {
		freqList = append(freqList, freq)
	}

	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i] < freqList[j]
	})

	for i := 0; i < len(freqList); i++ {
		if freqList[i] <= m {
			m -= freqList[i]
			delete(freqMap, ids[i])
		} else {
			break
		}
	}

	return int32(len(freqMap))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	idsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var ids []int32

	for i := 0; i < int(idsCount); i++ {
		idsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		idsItem := int32(idsItemTemp)
		ids = append(ids, idsItem)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := m

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
