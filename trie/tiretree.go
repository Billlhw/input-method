package trie

import (
	"sort"
	"sync"
)

type TrieNode struct {
	list      CharNodeList
	childNode map[rune]*TrieNode
	lock      sync.Mutex
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		list:      CharNodeList{},
		childNode: make(map[rune]*TrieNode),
		lock:      sync.Mutex{},
	}
}

func (t *TrieNode) Insert(spelling string, character rune, weight int) {
	//initialize to point to the root node
	curNode := t

	//find the location to insert
	for _, c := range spelling {
		curNode.lock.Lock()
		if _, exists := curNode.childNode[c]; !exists {
			curNode.childNode[c] = NewTrieNode()
		}
		nextNode := curNode.childNode[c]
		curNode.lock.Unlock()
		curNode = nextNode
	}

	//insert the current character into the trie node
	curNode.lock.Lock()
	curNode.list.InsertChar(character, weight)
	curNode.lock.Unlock()
}

// Search through the trie tree
func (t *TrieNode) Search(spelling string) []rune {
	curNode := t
	for _, c := range spelling {
		if _, exists := curNode.childNode[c]; !exists { //the spelling does not exist
			return nil
		}
		curNode = curNode.childNode[c]
	}

	// contains exact match of spelling
	if curNode.list.Len() > 0 {
		return curNode.list.GetCharacters()
	}

	// iteratively collect all characters in the child node
	resNodeList := CharNodeList{}
	resNodeList = curNode.collectCharNode(resNodeList)
	return resNodeList.GetCharacters()
}

// Iterate through the child nodes and collect all characters
func (t *TrieNode) collectCharNode(prevRes CharNodeList) CharNodeList {
	if t == nil {
		return prevRes
	}

	//insert characters in the current node
	res := append(prevRes, t.list...)

	//collect characters in the child node
	for _, n := range t.childNode {
		res = n.collectCharNode(res)
	}
	sort.Stable(res)
	return res
}
