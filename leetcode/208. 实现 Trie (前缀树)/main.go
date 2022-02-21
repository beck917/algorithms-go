package main

type Trie struct {
	children [26]*Trie // 数据结构采用26个字母的数组
	isEnd    bool      // 因为需要判断不是某个单词的前缀,所以需要加上isend标记
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	tmpNode := this
	for i := 0; i < len(word); i++ {
		char := word[i] - 'a' //这里要注意，是个小技巧，直接相减就可以得出此字幕在26数组中的位置
		if tmpNode.children[char] == nil {
			tmpNode.children[char] = &Trie{}
		}
		tmpNode = tmpNode.children[char]
	}
	tmpNode.isEnd = true
}

func (this *Trie) Search(word string) bool {
	tmpNode := this
	for _, char := range word {
		char -= 'a'
		if tmpNode.children[char] == nil {
			return false
		}
		tmpNode = tmpNode.children[char]
	}

	if tmpNode.isEnd == false {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	tmpNode := this
	for _, char := range prefix {
		char -= 'a'
		if tmpNode.children[char] == nil {
			return false
		}
		tmpNode = tmpNode.children[char]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
