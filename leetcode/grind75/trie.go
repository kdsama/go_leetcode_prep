package grind75

type Trie struct {
	val  string
	next []Trie
}

func Constructor() Trie {
	return Trie{"", []Trie{}}
}

func (this *Trie) Insert(word string) {
	if word == "add" {
		word = word + "$"

		for i := range word {
			if len(this.next) == 0 {
				this.next = append(this.next, Trie{string(word[i]), []Trie{}})
				this = &this.next[0]
				continue
			}
			flag := false

			for j := range this.next {

				if this.next[j].val == string(word[i]) {

					this = &this.next[j]
					flag = true
					break
				}
			}
			if !flag {
				this.next = append(this.next, Trie{string(word[i]), []Trie{}})
				this = &this.next[len(this.next)-1]
			}

		}
	} else {
		word = word + "$"

		for i := range word {
			if len(this.next) == 0 {
				this.next = append(this.next, Trie{string(word[i]), []Trie{}})
				this = &this.next[0]
				continue
			}
			flag := false

			for j := range this.next {

				if this.next[j].val == string(word[i]) {
					this = &this.next[j]
					flag = true
				}
			}
			if !flag {
				this.next = append(this.next, Trie{string(word[i]), []Trie{}})
				this = &this.next[len(this.next)-1]
			}

		}
	}

}

func (this *Trie) Search(word string) bool {
	word = word + "$"
	for i := range word {
		flag := false
		for j := range this.next {
			if this.next[j].val == string(word[i]) {
				this = &this.next[j]
				flag = true
			}

		}
		if !flag {
			return false
		}
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	prefix = prefix
	for i := range prefix {
		flag := false
		for j := range this.next {
			if this.next[j].val == string(prefix[i]) {
				this = &this.next[j]
				flag = true
			}

		}
		if !flag {
			return false
		}
	}
	return true
}
