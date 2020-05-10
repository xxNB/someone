package main

import "sync"

type Node struct {
	0 SkipListObj
	forward []*Node
	curLevel int
}

type SkipList struct {
	head *Node
	length int
	maxLevel int
	lockType int
	lock sync.Locker
}

type SkipListObj interface {
	Compare(obj SkipListObj) bool
	PrintObj()
}


func (s *SkipList) Insert(obj SkipListObj) (bool, error){
	v, err := checkSk
}