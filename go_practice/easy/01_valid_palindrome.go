/*
Problem: Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and 
removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
    Input: s = "A man, a plan, a canal: Panama"
    Output: true
    Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
    Input: s = "race a car"
    Output: false
    Explanation: "raceacar" is not a palindrome.

Example 3:
    Input: s = " "
    Output: true
    Explanation: s is an empty string "" after removing non-alphanumeric characters.
    Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
    - 1 <= s.length <= 2 * 10^5
    - s consists only of printable ASCII characters.
*/

package main

import (
	"fmt"
	"strings"

)

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ToLower(s)
	for i:=0;i<len(s);i++{
		if s[i] != s[len(s)-i-1]{
			return false
		}
	}
	return true
}

func runTests() {
	// Test case 1
	s1 := "A man, a plan, a canal: Panama"
	expected1 := true
	result1 := isPalindrome(s1)
	fmt.Printf("Test 1: %v, Output: %v, Expected: %v\n", result1 == expected1, result1, expected1)

	// Test case 2
	s2 := "race a car"
	expected2 := false
	result2 := isPalindrome(s2)
	fmt.Printf("Test 2: %v, Output: %v, Expected: %v\n", result2 == expected2, result2, expected2)

	// Test case 3
	s3 := " "
	expected3 := true
	result3 := isPalindrome(s3)
	fmt.Printf("Test 3: %v, Output: %v, Expected: %v\n", result3 == expected3, result3, expected3)

	// Test case 4
	s4 := "0P"
	expected4 := false
	result4 := isPalindrome(s4)
	fmt.Printf("Test 4: %v, Output: %v, Expected: %v\n", result4 == expected4, result4, expected4)
}

func main() {
	runTests()
} 