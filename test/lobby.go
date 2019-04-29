// Copyright 2019 MuGuangyi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"

	"github.com/muguangyi/seek/seek"
)

func newLobby(wg *sync.WaitGroup) ILobby {
	return &lobby{
		wg: wg,
	}
}

type ILobby interface {
}

type lobby struct {
	seek.Signal
	wg *sync.WaitGroup
}

func (l *lobby) OnInit(s seek.ISignaler) {
	l.Signal.OnInit(s)
	l.Book("IGame")
}

func (l *lobby) OnStart() {
	l.Visit("IGame").(IGame).Start("level1")
	l.wg.Done()
}
