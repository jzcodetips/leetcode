package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// get lowest elem
	low := len(strs[0])

	for _, str := range strs {
		if len(str) < low {
			low = len(str)
		}
	}

	var (
		currentPrefix  byte
		commonPrefix   string
		isCommonPrefix bool
	)

	for i := 0; i < low; i++ {
		isCommonPrefix = true
		currentPrefix = strs[0][i]

		for _, str := range strs {
			if str[i] != currentPrefix {
				isCommonPrefix = false

				break
			}
		}

		if !isCommonPrefix {
			return commonPrefix
		}

		commonPrefix += string(currentPrefix)
	}

	return commonPrefix
}
