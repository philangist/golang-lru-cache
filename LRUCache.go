package LRUCache

import (
  "container/list"
  )

// LRU Caches work by ejecting the least used objects from memory
// after reaching a certain capacity. They do this by having a hash table
// that has `get`, `set` or `delete` operations defined on it. The
// distinguishing factor is that all keys in the table are also
// members of a linked list that is periodically pruned.

// Tutorial here: http://openmymind.net/Writing-An-LRU-Cache/


// Define basic abstractions for working with the Cache
// A Cacheable is a thing that we can store within our hash table.
// LRUCache is the cache itself
// LRUCacheItem is a wrapper to Cacheable that also stores a node in
// the linked list of recently used items

type Cacheable interface {
  Key() string
  Size() int
}

type LRUCache struct {
  capacity int
  items map[string]*LRUCacheItem
  list *list.List
}

type LRUCacheItem struct {
  cacheable Cacheable
  listElement *list.Element
}

func New(capacity int) *LRUCache {
  return &LRUCache{
    capacity: capacity,
    items: make(map[string]*LRUCacheItem, 10000),
    list: list.New(),
  }
}

// Getter, setter, delete, and batch delete methods

func (c *LRUCache) Get(key string) Cacheable {
  item, exists := c.items[key]
  if exists == false {
    return nil
  }
  c.promote(item)
  return item.cacheable
}

func (c *LRUCache) promote (item *LRUCacheItem) {
  c.list.MoveToFront(item.listElement)
}

func (c *LRUCache) set (cacheable Cacheable) bool {
  if c.capacity < cachable.Size() {
    c.prune()
  }

  if c.capacity < cacheable.Size() {
    return false
  }

  item, exists := c.items[cacheable.Key()]
  if exists {
    item.cacheable = cacheable
    c.promote(item)
  } else {
    item = &LRUCacheItem{cacheable: cacheable}
    ite.listElement = c.list.PushFront(item)
    c.items[cacheable.Key()] = item
    c.capacity -= cacheable.Size()
  }
  return true;
}

func (c *LRUCache) delete (item) {
  (c.items, item.cacheable.Key())
}


func (c *LRUCache) prune() {
  for i := 0; i < 50; i++ {
    tail := c.list.Back()
    if tail == nil {
      return
    }
    item := c.list.Remove(tail).(*LRUCacheItem)
    c.delete(item)
    c.capacity += item.cacheable.Size()
  }
}
