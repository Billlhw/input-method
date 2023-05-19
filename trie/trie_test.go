package trie

import (
	"testing"
)

func cmpLists(t *testing.T, res []rune, ans []rune) {
	if len(res) != len(ans) {
		t.Errorf("result length is different from answer, result length is %d, answer length is %d",
			len(res), len(ans))
		return
	}

	for i, r := range res {
		if r != ans[i] {
			t.Errorf("result at index %d does not match with answer: %q and %q", i, r, ans[i])
		}
	}
}

func TestSort(t *testing.T) {
	nodeList := CharNodeList{}
	nodeList.InsertChar([]rune("章")[0], 7) //convert string to rune
	nodeList.InsertChar([]rune("张")[0], 9)
	nodeList.InsertChar([]rune("杖")[0], 6)
	res := nodeList.GetCharacters()
	ans := []rune{[]rune("张")[0], []rune("章")[0], []rune("杖")[0]}
	cmpLists(t, res, ans)
}

func TestSearch(t *testing.T) {
	rootNode := NewTrieNode()
	rootNode.Insert("zhang", []rune("张")[0], 9)
	rootNode.Insert("zhang", []rune("章")[0], 7)
	rootNode.Insert("zhan", []rune("站")[0], 10)
	rootNode.Insert("zha", []rune("扎")[0], 6)

	res := rootNode.Search("z")
	ans := []rune{[]rune("站")[0], []rune("张")[0], []rune("章")[0], []rune("扎")[0]}
	cmpLists(t, res, ans)

	res = rootNode.Search("zhang")
	ans = []rune{[]rune("张")[0], []rune("章")[0]}
	cmpLists(t, res, ans)
}
