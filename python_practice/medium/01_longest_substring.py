"""
Problem: Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
    Input: s = "abcabcbb"
    Output: 3
    Explanation: The answer is "abc", with the length of 3.

Example 2:
    Input: s = "bbbbb"
    Output: 1
    Explanation: The answer is "b", with the length of 1.

Example 3:
    Input: s = "pwwkew"
    Output: 3
    Explanation: The answer is "wke", with the length of 3.
    Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
    - 0 <= s.length <= 5 * 10^4
    - s consists of English letters, digits, symbols and spaces.
"""

def length_of_longest_substring(s):
    """
    Find the length of the longest substring without repeating characters.
    
    Args:
        s: str - The input string
        
    Returns:
        int - Length of the longest substring without repeating characters
    """
    # TODO: Implement your solution here
    pass


# Test cases
def test_length_of_longest_substring():
    # Test case 1
    s1 = "abcabcbb"
    expected1 = 3
    result1 = length_of_longest_substring(s1)
    print(f"Test 1: {result1 == expected1}, Output: {result1}, Expected: {expected1}")
    
    # Test case 2
    s2 = "bbbbb"
    expected2 = 1
    result2 = length_of_longest_substring(s2)
    print(f"Test 2: {result2 == expected2}, Output: {result2}, Expected: {expected2}")
    
    # Test case 3
    s3 = "pwwkew"
    expected3 = 3
    result3 = length_of_longest_substring(s3)
    print(f"Test 3: {result3 == expected3}, Output: {result3}, Expected: {expected3}")
    
    # Test case 4 (empty string)
    s4 = ""
    expected4 = 0
    result4 = length_of_longest_substring(s4)
    print(f"Test 4: {result4 == expected4}, Output: {result4}, Expected: {expected4}")
    
    # Test case 5 (with spaces and special characters)
    s5 = "ab c!d"
    expected5 = 5
    result5 = length_of_longest_substring(s5)
    print(f"Test 5: {result5 == expected5}, Output: {result5}, Expected: {expected5}")


if __name__ == "__main__":
    test_length_of_longest_substring() 