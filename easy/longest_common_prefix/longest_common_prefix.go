package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
    
    	var (
        	currentPrefix byte
        	commonPrefix string
    	)

	low := len(strs[0])

	for _, str := range strs {
		if len(str) < low {
			low = len(str)
		}
	}
    
    	for i := 0; i < low; i++ {        
        	currentPrefix = strs[0][i]
    
        	for _, str := range strs {
            		if str[i] != currentPrefix {
                		return commonPrefix
            		}
        	}
        
        	commonPrefix += string(currentPrefix)
    	}
    
    	return commonPrefix
}
