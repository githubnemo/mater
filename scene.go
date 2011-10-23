package mater

import (
	"box2d"
)

type Scene struct {
	World *box2d.World
}

func (scene *Scene) Init () {
	scene.World = new(box2d.World)
	scene.World.Init()
}

func (scene *Scene) Update (dt float64) {
	scene.World.Step(dt)
}