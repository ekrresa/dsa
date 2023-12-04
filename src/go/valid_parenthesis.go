package main

import "fmt"

func IsValidParenthesis(str string) bool {
	if len(str)%2 == 1 || len(str) == 0 {
		return false
	}

	store := []string{}
	tokens := make(map[string]string)
	tokens["("] = ")"
	tokens["{"] = "}"
	tokens["["] = "]"

	for _, char := range str {
		if tokens[string(char)] != "" {
			store = append(store, string(char))
		} else {
			if len(store)-1 < 0 {
				return false
			}

			lastChar := store[len(store)-1]

			if tokens[lastChar] == string(char) {
				store = store[:len(store)-1]
			} else {
				return false
			}
		}
	}

	return len(store) == 0
}

func main() {
	input := "){"
	ans := IsValidParenthesis(input)

	fmt.Println(ans)
}
