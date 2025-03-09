/*
Problem: LFU Cache

Design and implement a data structure for a Least Frequently Used (LFU) cache.

Implement the LFUCache class:
- LFUCache(int capacity) Initializes the object with the capacity of the data structure.
- int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
- void put(int key, int value) Update the value of the key if present, or inserts the key if not already present.
  When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item.
  For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.

To determine the least frequently used key, a use counter is maintained for each key in the cache.
The key with the smallest use counter is the least frequently used key.

When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation).
The use counter for a key in the cache is incremented either a get or put operation is called on it.

The functions get and put must each run in O(1) average time complexity.

Example:
    Input:
    ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
    [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
    Output:
    [null, null, null, 1, null, -1, 3, null, -1, 3, 4]

    Explanation:
    // cnt(x) = the use counter for key x
    // cache=[] will show the last used order for tiebreakers (leftmost element is most recent)
    LFUCache lfu = new LFUCache(2);
    lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
    lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
    lfu.get(1);      // return 1
                     // cache=[1,2], cnt(2)=1, cnt(1)=2
    lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
                     // cache=[3,1], cnt(3)=1, cnt(1)=2
    lfu.get(2);      // return -1 (not found)
    lfu.get(3);      // return 3
                     // cache=[3,1], cnt(3)=2, cnt(1)=2
    lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
                     // cache=[4,3], cnt(4)=1, cnt(3)=2
    lfu.get(1);      // return -1 (not found)
    lfu.get(3);      // return 3
                     // cache=[3,4], cnt(4)=1, cnt(3)=3
    lfu.get(4);      // return 4
                     // cache=[4,3], cnt(4)=2, cnt(3)=3

Constraints:
    - 0 <= capacity <= 10^4
    - 0 <= key <= 10^5
    - 0 <= value <= 10^9
    - At most 2 * 10^5 calls will be made to get and put.
*/

package main

import (
	"fmt"
)

// LFUCache represents a Least Frequently Used cache
type LFUCache struct {
	// TODO: Add necessary fields for LFU cache implementation
}

// Constructor initializes a new LFU cache with the given capacity
func Constructor(capacity int) LFUCache {
	// TODO: Implement constructor
	return LFUCache{}
}

// Get retrieves the value for a key and updates its frequency
func (c *LFUCache) Get(key int) int {
	// TODO: Implement Get method
	return -1
}

// Put adds or updates a key-value pair in the cache
func (c *LFUCache) Put(key int, value int) {
	// TODO: Implement Put method
}

func runTests() {
	// Test case 1
	lfu := Constructor(2)
	lfu.Put(1, 1)                        // cache=[1,_], cnt(1)=1
	lfu.Put(2, 2)                        // cache=[2,1], cnt(2)=1, cnt(1)=1
	fmt.Println(lfu.Get(1) == 1)         // return 1; cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.Put(3, 3)                        // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2; cache=[3,1], cnt(3)=1, cnt(1)=2
	fmt.Println(lfu.Get(2) == -1)        // return -1 (not found)
	fmt.Println(lfu.Get(3) == 3)         // return 3; cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.Put(4, 4)                        // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1; cache=[4,3], cnt(4)=1, cnt(3)=2
	fmt.Println(lfu.Get(1) == -1)        // return -1 (not found)
	fmt.Println(lfu.Get(3) == 3)         // return 3; cache=[3,4], cnt(4)=1, cnt(3)=3
	fmt.Println(lfu.Get(4) == 4)         // return 4; cache=[4,3], cnt(4)=2, cnt(3)=3
	
	// Test case 2 (empty cache)
	lfu2 := Constructor(0)
	lfu2.Put(0, 0)                       // Nothing happens, capacity is 0
	fmt.Println(lfu2.Get(0) == -1)       // return -1 (not found, capacity is 0)
}

func main() {
	runTests()
} 