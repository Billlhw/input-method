package trie

import "sort"

type CharNode struct {
	Character rune
	Weight    int
}

type CharNodeList []CharNode //sorted list of CharNodes

func (c CharNodeList) Len() int {
	return len(c)
}

func (c CharNodeList) Less(i, j int) bool {
	return c[i].Weight > c[j].Weight
}

func (c CharNodeList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// insert a node to CharNodeList
func (c *CharNodeList) InsertChar(ch rune, w int) {
	newNode := CharNode{
		Character: ch,
		Weight:    w,
	}
	*c = append(*c, newNode)
	sort.Stable(*c)
}

// return all characters in characterList
func (c CharNodeList) GetCharacters() []rune {
	res := make([]rune, 0)
	for _, node := range c {
		res = append(res, node.Character)
	}
	return res
}
