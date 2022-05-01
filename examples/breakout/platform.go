package main

import (
	"github.com/quasilyte/ge"
	"github.com/quasilyte/ge/gemath"
	"github.com/quasilyte/ge/physics"
)

type platform struct {
	scene  *ge.Scene
	body   physics.Body
	sprite *ge.Sprite
	ball   *ball
}

func newPlatform() *platform {
	p := &platform{}
	p.body.InitRotatedRect(p, 100, 22)
	return p
}

func (p *platform) Init(scene *ge.Scene) {
	p.scene = scene
	p.sprite = p.scene.LoadSprite("platform.png")
	p.sprite.Pos = &p.body.Pos
	p.sprite.Rotation = &p.body.Rotation
	scene.AddGraphics(p.sprite)
	scene.AddBody(&p.body)
}

func (p *platform) IsDisposed() bool { return p.body.IsDisposed() }

func (p *platform) Dispose() {
	p.body.Dispose()
	p.sprite.Dispose()
}

func (p *platform) Update(delta float64) {
	moving := false
	if p.scene.Input().ActionIsPressed(ActionLeft) {
		moving = true
		p.body.Pos.X -= 250 * delta
		p.body.Rotation = gemath.Rad(gemath.ClampMin(float64(p.body.Rotation)-1.5*delta, -0.3))
	} else if p.scene.Input().ActionIsPressed(ActionRight) {
		moving = true
		p.body.Pos.X += 250 * delta
		p.body.Rotation = gemath.Rad(gemath.ClampMax(float64(p.body.Rotation)+1.5*delta, 0.3))
	}
	if !moving {
		if p.body.Rotation < 0 {
			p.body.Rotation = gemath.Rad(gemath.ClampMax(float64(p.body.Rotation)+1.1*delta, 0))
		} else if p.body.Rotation > 0 {
			p.body.Rotation = gemath.Rad(gemath.ClampMin(float64(p.body.Rotation)-1.1*delta, 0))
		}
	}

	if p.ball != nil && p.ball.IsDisposed() {
		p.ball = nil
	}
	if p.ball == nil && p.scene.Input().ActionIsPressed(ActionFire) {
		b := newBall()
		p.ball = b
		b.velocity = gemath.Vec{X: 0, Y: -350}
		b.body.Pos = gemath.Vec{X: p.body.Pos.X, Y: p.body.Pos.Y - 40}
		p.scene.AddObject(b)
	}
}