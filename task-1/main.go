package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLine(prompt string) []int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	text := scanner.Text()
	parts := strings.Fields(text)
	var result []int
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		result = append(result, num)
	}
	return result
}

func checkLeaderboard(leaderboard []int, viska []int) []int {
	scoreMap := make(map[int]bool)
	for _, score := range leaderboard {
		scoreMap[score] = true
	}

	var uniqueScores []int
	for score := range scoreMap {
		uniqueScores = append(uniqueScores, score)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(uniqueScores)))

	var rankings []int

	for _, score := range viska {
		index := len(uniqueScores) - 1
		for index >= 0 && score >= uniqueScores[index] {
			index--
		}
		rankings = append(rankings, index+2)
	}

	return rankings
}

func main() {
	fmt.Println("=== VISKA LEADERBOARD DENSE RANKING ===")

	leaderboard := readLine("Masukkan skor leaderboard (pisahkan dengan spasi): ")
	viska := readLine("Masukkan skor VISKA (pisahkan dengan spasi): ")

	rankings := checkLeaderboard(leaderboard, viska)

	fmt.Println("Peringkat VISKA: ")
	for _, rank := range rankings {
		fmt.Printf("%d ", rank)
	}
	fmt.Println()
}
