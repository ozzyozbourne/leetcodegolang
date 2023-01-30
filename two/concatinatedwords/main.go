package main

import "fmt"

func main() {
	in := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	in = findAllConcatenatedWordsInADict(in)
	fmt.Println(in)
}

func findAllConcatenatedWordsInADict(words []string) []string {
	var exists = struct{}{}
	cacheMap := make(map[string]bool)
	var res []string
	dp := make(map[string]struct{})
	for _, v := range words {
		if _, ok := dp[v]; !ok {
			dp[v] = exists
		}
	}
	for _, k := range words {
		if dfs(k, dp, cacheMap) {
			res = append(res, k)
		}
	}
	return res
}

func dfs(word string, dp map[string]struct{}, cacheMap map[string]bool) bool {
	if _, ok := cacheMap[word]; ok {
		return cacheMap[word]
	}

	for i := 1; i < len(word); i++ {
		prefix := word[:i]
		suffix := word[i:]
		_, okPrefix := dp[prefix]
		_, okSuffix := dp[suffix]
		if okPrefix && okSuffix || okPrefix && dfs(suffix, dp, cacheMap) {
			cacheMap[word] = true
			return cacheMap[word]
		}
	}
	cacheMap[word] = false
	return cacheMap[word]
}
