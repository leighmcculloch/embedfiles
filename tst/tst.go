/*
MIT License

Copyright (c) 2017 Seis

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package tst implementst a Ternary Search Tree (TST), copied from https://github.com/xiaonanln/go-trie-tst
package tst

// TST can be the root, and can be a sub-tree
type TST struct {
	Left  *TST
	Right *TST
	Eq    *TST
	Eqkey byte
	Val   []byte
}

// Child returns the child subtree of the current tree
func (t *TST) Child(c byte) *TST {
	if t.Eq == nil {
		t.Eqkey = c
		t.Eq = &TST{}
		return t.Eq
	} else if c == t.Eqkey {
		return t.Eq
	} else if c < t.Eqkey {
		if t.Left == nil {
			t.Left = &TST{}
		}
		return t.Left.Child(c)
	} else { // c > t.eqkey
		if t.Right == nil {
			t.Right = &TST{}
		}
		return t.Right.Child(c)
	}
}

// Set sets the value of string in the current tree
func (t *TST) Set(s string, val []byte) {
	t.set(s, val, 0)
}

func (t *TST) set(s string, val []byte, idx int) {
	if idx < len(s) {
		t.Child(s[idx]).set(s, val, idx+1)
	} else {
		t.Val = val
	}
}

// Get returns the value of string in the current tree
func (t *TST) Get(s string) (val []byte) {
	return t.get(s, 0)
}

func (t *TST) get(s string, idx int) (val []byte) {
	if idx < len(s) {
		return t.Child(s[idx]).get(s, idx+1)
	} else {
		return t.Val
	}
}

// Sub returns the subtree of the current tree with specified prefix
func (t *TST) Sub(s string) *TST {
	return t.sub(s, 0)
}

func (t *TST) sub(s string, idx int) *TST {
	if idx < len(s) {
		return t.Child(s[idx]).sub(s, idx+1)
	} else {
		return t
	}
}

func (t *TST) ForEach(f func(s string, val []byte)) {
	var prefix []byte
	t.forEach(f, prefix)
}

func (t *TST) forEach(f func(s string, val []byte), prefix []byte) {
	if t.Val != nil {
		f(string(prefix), t.Val)
	}

	if t.Left != nil {
		t.Left.forEach(f, prefix)
	}

	if t.Eq != nil {
		t.Eq.forEach(f, append(prefix, t.Eqkey))
	}

	if t.Right != nil {
		t.Right.forEach(f, prefix)
	}
}
