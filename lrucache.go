package lrucache

import (
	"errors"
)

type LRUCache struct {

	Cache
	Capacity int
	Map map[string]*lruCacheNode
	FirstNode *lruCacheNode
	LastNode *lruCacheNode
}

type lruCacheNode struct {
	Key string
	Value interface{}
	PrevNode *lruCacheNode
	NextNode *lruCacheNode
	
}

func NewLRUCache (capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("Invalid capacity.")
	}
	return 	&LRUCache{
		Capacity: capacity,
		Map: make(map[string]*lruCacheNode),
		FirstNode: nil,
		LastNode: nil,	
	}, nil
}

func (c *LRUCache) Set(key string, value interface{}) {
	c.removeNode(key)
	c.addNodeToHead(key, value)	

	if len(c.Map) > c.Capacity {
		c.removeNode(c.LastNode.Key)
	}

}

func (c *LRUCache) Get(key string) interface{} {
	node := c.Map[key]

	if node == nil {
		return nil
	}	

	c.removeNode(key)
	c.addNodeToHead(key, node.Value)
	return node.Value
}

func (c *LRUCache) addNodeToHead(key string, value interface{}) {
	
	node := &lruCacheNode {
		Key: key,
		Value: value,
		NextNode: c.FirstNode,
		PrevNode: nil,	
	}
	
	if c.FirstNode != nil {
		c.FirstNode.PrevNode = node
	}
	if c.LastNode == nil {
		c.LastNode = node
	} 

	c.FirstNode = node
	c.Map[key] = node
}

func (c *LRUCache) removeNode(key string) {

	node := c.Map[key]
	if node == nil {
		return
	}

	// Change pointers of PrevNode and NextNode
	if node.PrevNode != nil {
		node.PrevNode.NextNode = node.NextNode
	}
	if node.NextNode != nil {
		node.NextNode.PrevNode = node.PrevNode
	}

	// Move FirstNode and LastNode if the node to be deleted is pointed by them
	if c.FirstNode == node {
		c.FirstNode = node.NextNode
	}
	if c.LastNode == node {
		c.LastNode = node.PrevNode
	}

	delete(c.Map, key)
}




