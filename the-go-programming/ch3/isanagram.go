package main

func isAnagram(str string) bool {
	length := len(str)
	for i := 0; i < length; i++ {
		isAnagram := str[i] == str[length-i-1]
		if !isAnagram {
			return false
		}
	}
	return true
}
