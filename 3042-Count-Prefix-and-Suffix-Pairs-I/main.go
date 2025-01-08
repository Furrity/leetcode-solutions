package main

import "strings"

func countPrefixSuffixPairs(words []string) int {
    var ans int
    for i := 0; i < len(words); i++ {
        for j := i+1; j < len(words); j++ {
            if isPrefixSuffix(words[i], words[j]) {
                ans++
            }

        }
    }
    return ans
}

func isPrefixSuffix(str1, str2 string) bool {
    return strings.HasPrefix(str2 , str1) && strings.HasSuffix(str2, str1)
}
