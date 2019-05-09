// Copyright 2019 MuGuangyi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"

	"github.com/muguangyi/seek"
)

type IGame interface {
	Start(level string)
}

func newGame(wg *sync.WaitGroup) IGame {
	c := new(game)
	c.wg = wg

	return c
}

type game struct {
	seek.Signal
	wg *sync.WaitGroup
}

func (g *game) OnInit(s seek.ISignaler) {
	g.Signal.OnInit(s)
	g.Book("IMath")
}

func (g *game) OnStart() {
	math := g.Visit("IMath").(IMath)
	math.Print("Hello World!")

	result := math.Add(1, 2)
	fmt.Println("add result:", result)

	g.wg.Done()
}

func (g *game) Start(level string) {
	fmt.Println("start...")
	fmt.Println("game started:", level)
}
