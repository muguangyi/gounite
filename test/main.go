// Copyright 2019 MuGuangyi. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/muguangyi/gounite"
	"github.com/muguangyi/gounite/framework"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	gounite.RunHub("127.0.0.1:9999")

	gounite.Run("127.0.0.1:9999", "util",
		framework.NewUnit("math", &MathControl{}, true))

	time.Sleep(10)

	gounite.Run("127.0.0.1:9999", "logic",
		framework.NewUnit("game", &GameControl{wg: &wg}, true))

	wg.Wait()
	fmt.Println("Completed!")
}

type MathControl struct {
	unit framework.IUnit
}

func (math *MathControl) OnInit(u framework.IUnit) {
	math.unit = u
	u.BindCall("add", math.add)
}

func (math *MathControl) OnStart() {

}

func (math *MathControl) OnDestroy() {

}

func (math *MathControl) add(args []interface{}) interface{} {
	fmt.Println(fmt.Sprintf("-----add method called, %T, %T", args[0], args[1]))
	result := args[0].(float64) + args[1].(float64)
	return result
}

type GameControl struct {
	unit framework.IUnit
	wg   *sync.WaitGroup
}

func (g *GameControl) OnInit(u framework.IUnit) {
	g.unit = u
	u.Import("math")
}

func (g *GameControl) OnStart() {
	result, _ := g.unit.CallWithResult("math", "add", 1, 2)
	fmt.Println("-----Math add result:", result)
	g.wg.Done()
}

func (g *GameControl) OnDestroy() {

}
