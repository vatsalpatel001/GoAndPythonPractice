"""
Problem: Two Sum

Given an array of integers 'nums' and an integer 'target', return the indices of the two numbers such that they add up to the target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example:
    Input: nums = [2, 7, 11, 15], target = 9
    Output: [0, 1]
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1]

Constraints:
    - 2 <= nums.length <= 10^4
    - -10^9 <= nums[i] <= 10^9
    - -10^9 <= target <= 10^9
    - Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n²) time complexity?
"""

def two_sum(nums, target):
    """
    Find two numbers in nums that add up to target and return their indices.
    
    Args:
        nums: List[int] - A list of integers
        target: int - The target sum
        
    Returns:
        List[int] - Indices of the two numbers that add up to target
    """
    # TODO: Implement your solution here
    pass


# Test cases
def test_two_sum():
    # Test case 1
    nums1 = [2, 7, 11, 15]
    target1 = 9
    expected1 = [0, 1]
    result1 = two_sum(nums1, target1)
    print(f"Test 1: {result1 == expected1}")
    
    # Test case 2
    nums2 = [3, 2, 4]
    target2 = 6
    expected2 = [1, 2]
    result2 = two_sum(nums2, target2)
    print(f"Test 2: {result2 == expected2}")
    
    # Test case 3
    nums3 = [3, 3]
    target3 = 6
    expected3 = [0, 1]
    result3 = two_sum(nums3, target3)
    print(f"Test 3: {result3 == expected3}")


if __name__ == "__main__":
    test_two_sum() 