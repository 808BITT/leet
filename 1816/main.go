package main

import "strings"

func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}

func main() {
	ans := truncateSentence("param is sentence and this is very fun", 5)
	println(ans) // expect "param is"
}
