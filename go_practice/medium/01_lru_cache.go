/*
Problem: LRU Cache

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:
- LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
- int get(int key) Return the value of the key if the key exists, otherwise return -1.
- void put(int key, int value) Update the value of the key if the key exists. 
  Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity, 
  evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.

Example:
    Input:
    ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
    [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
    Output:
    [null, null, null, 1, null, -1, null, -1, 3, 4]

    Explanation:
    LRUCache lRUCache = new LRUCache(2);
    lRUCache.put(1, 1); // cache is {1=1}
    lRUCache.put(2, 2); // cache is {1=1, 2=2}
    lRUCache.get(1);    // return 1
    lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
    lRUCache.get(2);    // returns -1 (not found)
    lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
    lRUCache.get(1);    // return -1 (not found)
    lRUCache.get(3);    // return 3
    lRUCache.get(4);    // return 4

Constraints:
    - 1 <= capacity <= 3000
    - 0 <= key <= 10^4
    - 0 <= value <= 10^5
    - At most 2 * 10^5 calls will be made to get and put.
*/

package main

import (
	"fmt"
)

// LRUCache represents a Least Recently Used cache
type LRUCache struct {
	// TODO: Add necessary fields for LRU cache implementation
}

// Constructor initializes a new LRU cache with the given capacity
func Constructor(capacity int) LRUCache {
	// TODO: Implement constructor
	return LRUCache{}
}

// Get retrieves the value for a key and updates it as recently used
func (c *LRUCache) Get(key int) int {
	// TODO: Implement Get method
	return -1
}

// Put adds or updates a key-value pair in the cache
func (c *LRUCache) Put(key int, value int) {
	// TODO: Implement Put method
}

func runTests() {
	// Test case 1
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)                    // cache is {1=1}
	lRUCache.Put(2, 2)                    // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1) == 1)     // return 1
	lRUCache.Put(3, 3)                    // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2) == -1)    // returns -1 (not found)
	lRUCache.Put(4, 4)                    // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1) == -1)    // return -1 (not found)
	fmt.Println(lRUCache.Get(3) == 3)     // return 3
	fmt.Println(lRUCache.Get(4) == 4)     // return 4
	
	// Test case 2
	lRUCache2 := Constructor(1)
	lRUCache2.Put(2, 1)                   // cache is {2=1}
	fmt.Println(lRUCache2.Get(2) == 1)    // return 1
	lRUCache2.Put(3, 2)                   // LRU key was 2, evicts key 2, cache is {3=2}
	fmt.Println(lRUCache2.Get(2) == -1)   // returns -1 (not found)
	fmt.Println(lRUCache2.Get(3) == 2)    // returns 2
}

func main() {
	runTests()
} 