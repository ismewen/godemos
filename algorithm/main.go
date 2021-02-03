package main

import (
	"algorithm/sort"
	alstring "algorithm/string"
	"fmt"
	"log"
	"math/rand"
	sort2 "sort"
	"time"
)

func GetRandomArray(length int64) []int64 {
	data := make([]int64, length)
	for i := int64(0); i < length; i++ {
		data[i] = rand.Int63()
	}
	return data
}

func GetRandomIntArray(length int) sort2.IntSlice {
	data := make(sort2.IntSlice, length)
	for i := 0; i < length; i++ {
		data[i] = rand.Int()
	}
	return data
}

func min(a []int) int {
	m := a[0]
	for _, i := range a {
		if i < m {
			m = i
		}
	}
	return m
}

func max(a []int) int {
	m := a[0]
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
}
func sortCompare() {
	//length := int64(100000)
	//start := time.Now()
	//data := GetRandomArray(length)
	//spend := time.Since(start)
	//log.Println(fmt.Sprintf("init data spend %v", spend))

	//start = time.Now()
	//sort.MaopaoSort(data, true)
	//spend = time.Since(start)
	//log.Println(fmt.Sprintf(" maopao spend %v", spend))
	//
	//start = time.Now()
	//data = GetRandomArray(length)
	//sort.InsertSortFirst(data)
	//spend = time.Since(start)
	//log.Println(fmt.Sprintf("insert spend %v", spend))
	//
	//start = time.Now()
	//data = GetRandomArray(length)
	//sort.SelectSort(data, true)
	//spend = time.Since(start)
	//log.Println(fmt.Sprintf("select spend %v", spend))

	start := time.Now()
	sdata := GetRandomIntArray(100 * 10000)
	sort.ShellSort(sdata)
	spend := time.Since(start)
	log.Println(fmt.Sprintf("shell spend %v", spend))

	start = time.Now()
	idata := GetRandomIntArray(1000000)
	sort2.Sort(idata)
	spend = time.Since(start)
	log.Println(fmt.Sprintf("inner spend %v", spend))

	start = time.Now()
	gdata := GetRandomIntArray(1000000)
	gdata = sort.GuibinSort(gdata)
	spend = time.Since(start)
	log.Println(fmt.Sprintf("guibin pend %v", spend))
	log.Println(gdata[0], gdata[len(gdata)-1])
	log.Println(min(gdata), max(gdata))

	start = time.Now()
	gdata = GetRandomIntArray(1000000)
	sort.QuickSort(gdata)
	spend = time.Since(start)
	log.Println(fmt.Sprintf("quick spend %v", spend))
	log.Println(gdata[0], gdata[len(gdata)-1])
	log.Println(min(gdata), max(gdata))

	ddata := GetRandomIntArray(1000000)
	bh := sort.NewBinaryHeap(ddata)
	bh.Sort()

}
func TestKmp() {
	//s := alstring.Kmp("abcasdfasdfasbf", "sbfbasdfasdfasdfasdfasdfasdf")
	next1 := alstring.GetNext("abcasdfasdfasbf")
	next2 := alstring.GetNext1("abcasdfasdfasbf")
	log.Println(next1, next2)
}

func main() {
	TestKmp()
	//search.Test()

	//data := []int{4, 1, 5, 2, 3, 6, 12, 89, 75}
	//sort.QuickSort(data)
	//log.Println("after sort", data)
	//position := search.BinarySearchIter(data, 99)
	//log.Println("iter sort position", position)
	//position = search.BinarySearchRecursion(data, 99)
	//log.Println("recusion sort position", position)
	//sort.QuickSort(data)
	//log.Println(data)
	//log.Println(data)
	//data = sort.GuibinSort(data)
	//log.Println(data)

	//TestKmp()
	//a := sort.InsertSortFirst([]int64{8, 9, 1, 7, 2, 3, 5, 4, 0, 6})
	//b := sort.InsertSortSecond([]int64{8, 9, 1, 7, 2, 3, 5, 4, 0, 6})
	//c := sort.ShellSort([]int64{8, 9, 1, 7, 2, 3, 5, 4, 0, 6})
	//log.Println(a, b, c)
}
