package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dannywolfmx/sorting-algorithms/merge"
)

func main() {

	list := randomList(10000000)

	c := make(chan []int, len(list))

	go merge.Concurrent(list, c)

	list = <-c

	//list = merge.Mergesort(list)

	fmt.Printf("%d, %d \n", len(list), cap(list))
}

func randomList(len int) []int {
	rand.Seed(time.Now().UnixNano())

	return make(UniqueNumber).generateList(len)
}

var Empty struct{}

type UniqueNumber map[int]struct{}

func (u UniqueNumber) generateList(lenList int) []int {
	list := make([]int, 0, lenList)
	for len(list) != lenList {
		num := rand.Int()
		if _, ok := u[num]; !ok {
			u[num] = Empty
			list = append(list, num)
		}
	}

	return list
}
