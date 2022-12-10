package dayone

import (
	"container/heap"
	"fmt"
	utils "solutions/aoc_utils"
	"strconv"
)

func Solution() (int, error) {
	f, err := utils.OpenFile("./day_1.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}
	defer f.Close()

	calorie := 0

	topThreeHeap := &utils.IntHeap{}
	heap.Init(topThreeHeap)

	topThreeSum := 0

	scanner := utils.ReadLine(f)
	for scanner.Scan() {
		if scanner.Text() == utils.EmptyString {
			if topThreeHeap.Len() < 3 {
				heap.Push(topThreeHeap, calorie)
				topThreeSum += calorie
			} else {
				if calorie > topThreeHeap.Peek() {
					heap.Push(topThreeHeap, calorie)
					topThreeSum += calorie
					topThreeSum -= heap.Pop(topThreeHeap).(int)
				}
			}

			calorie = 0
			continue
		}

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		calorie += i
	}

	return topThreeSum, nil
}
