package main

import "fmt"

func main() {
	fmt.Println("Anagram")

	testCases := []struct {
		input  []string
		output [][]string
	}{
		{input: []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}, output: [][]string{{"kita", "atik", "tika"}, {"aku", "kua"}, {"makan"}, {"kia"}}},
	}

	for _, v := range testCases {
		res := solveAnagram(v.input)

		fmt.Printf("input: %+v, expected output: %+v, actual output: %+v \n", v.input, v.output, res)
	}
}

// Time Complexity = O(N^2)
// Space Complexity = O(N), karena saya menggunakan hash map
func solveAnagram(s []string) [][]string {
	res := make([][]string, 0)

	m := make(map[[26]int][]string)

	for _, value := range s {
		var charCount [26]int

		for _, v := range value {
			charCount[v-'a'] = charCount[v-'a'] + 1
		}

		_, ok := m[charCount]
		if !ok {
			m[charCount] = make([]string, 0)
		}

		m[charCount] = append(m[charCount], value)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
