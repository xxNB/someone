package main 

import (
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 32
const SKIPLIST_P = 4

type Node struct {
	Forward []Node 
	Value interface{}
}

func (s1 *SkipList) Search(score float64) (element *Element, ok bool){
	x := s1.header
	for i := s1.level -1; i >=0; i--{
		for x.forward[i] !=nil && x.forward[i].Score < score{
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score{
		return x, true
	}
	return nil, false
}