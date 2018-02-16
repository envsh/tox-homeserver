package main

type contentAreaState struct {
	isBottom bool
	curpos   int
	maxpos   int
}

var ccstate = &contentAreaState{isBottom: true}
