"""
Problem: Merge k Sorted Lists

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.

Example 1:
    Input: lists = [[1,4,5],[1,3,4],[2,6]]
    Output: [1,1,2,3,4,4,5,6]
    Explanation: The linked-lists are:
    [
      1->4->5,
      1->3->4,
      2->6
    ]
    merging them into one sorted list:
    1->1->2->3->4->4->5->6

Example 2:
    Input: lists = []
    Output: []

Example 3:
    Input: lists = [[]]
    Output: []

Constraints:
    - k == lists.length
    - 0 <= k <= 10^4
    - 0 <= lists[i].length <= 500
    - -10^4 <= lists[i][j] <= 10^4
    - lists[i] is sorted in ascending order.
    - The sum of lists[i].length won't exceed 10^4.
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
    
    def __repr__(self):
        values = []
        current = self
        while current:
            values.append(str(current.val))
            current = current.next
        return "->".join(values)


def create_linked_list(values):
    """Helper function to create a linked list from a list of values."""
    dummy = ListNode(0)
    current = dummy
    for val in values:
        current.next = ListNode(val)
        current = current.next
    return dummy.next


def linked_list_to_list(head):
    """Helper function to convert a linked list to a regular list."""
    result = []
    current = head
    while current:
        result.append(current.val)
        current = current.next
    return result


def merge_k_lists(lists):
    """
    Merge k sorted linked lists into one sorted linked list.
    
    Args:
        lists: List[ListNode] - A list of heads of linked lists
        
    Returns:
        ListNode - The head of the merged sorted linked list
    """
    # TODO: Implement your solution here
    pass


# Test cases
def test_merge_k_lists():
    # Test case 1
    list1 = create_linked_list([1, 4, 5])
    list2 = create_linked_list([1, 3, 4])
    list3 = create_linked_list([2, 6])
    lists1 = [list1, list2, list3]
    expected1 = [1, 1, 2, 3, 4, 4, 5, 6]
    result1 = linked_list_to_list(merge_k_lists(lists1))
    print(f"Test 1: {result1 == expected1}, Output: {result1}, Expected: {expected1}")
    
    # Test case 2 (empty list of lists)
    lists2 = []
    expected2 = []
    result2 = linked_list_to_list(merge_k_lists(lists2))
    print(f"Test 2: {result2 == expected2}, Output: {result2}, Expected: {expected2}")
    
    # Test case 3 (list containing an empty list)
    lists3 = [None]
    expected3 = []
    result3 = linked_list_to_list(merge_k_lists(lists3))
    print(f"Test 3: {result3 == expected3}, Output: {result3}, Expected: {expected3}")
    
    # Test case 4 (multiple empty lists)
    lists4 = [None, None, None]
    expected4 = []
    result4 = linked_list_to_list(merge_k_lists(lists4))
    print(f"Test 4: {result4 == expected4}, Output: {result4}, Expected: {expected4}")


if __name__ == "__main__":
    test_merge_k_lists() 