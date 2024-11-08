package easy

import (
	"fmt"
	"strings"
)

var strs = []string{"flower", "flow", "flight"}

// longestCommonPrefixFirstCase...
func longestCommonPrefixFirstCase(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// longestCommonPrefixSecondCase...
func longestCommonPrefixSecondCase(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// longestCommonPrefixFirstCaseRunner...
func LongestCommonPrefixFirstCaseRunner() {
	fmt.Println(longestCommonPrefixFirstCase(strs))
}

// longestCommonPrefixSecondCaseRunner...
func LongestCommonPrefixSecondCaseRunner() {
	fmt.Println(longestCommonPrefixSecondCase(strs))
}
