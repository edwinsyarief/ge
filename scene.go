package ge

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge/collision"
)

type SceneController interface {
	Init(*Scene)

	Update(delta float64)
}

type Disposable interface {
	IsDisposed() bool

	Dispose()
}

type SceneObject interface {
	// Init is called once when object is added to the scene.
	//
	// It's a good time to initialize all dependent objects
	// and attach sprites to the scene.
	Init(*Scene)

	// IsDisposed reports whether scene object was disposed.
	//
	// Disposed objects are removed from the scene before their
	// Update method is called for the current frame.
	IsDisposed() bool

	// Update is called for every object during every logical game frame.
	// Delta specifies how many seconds have passed from the previous frame.
	Update(delta float64)
}

type SceneGraphics interface {
	Draw(dst *ebiten.Image)

	IsDisposed() bool
}

type Scene struct {
	context *Context

	Name string

	controller      SceneController
	objects         []SceneObject
	addedObjects    []SceneObject
	tmpObjectsQueue []SceneObject

	collisionEngine collision.Engine

	graphics []SceneGraphics
}

func newScene() *Scene {
	return &Scene{
		objects:         make([]SceneObject, 0, 32),
		addedObjects:    make([]SceneObject, 0, 8),
		tmpObjectsQueue: make([]SceneObject, 0, 8),
		graphics:        make([]SceneGraphics, 0, 24),
	}
}

func (s *Scene) Context() *Context {
	return s.context
}

func (s *Scene) AddBody(b *collision.Body) {
	s.collisionEngine.AddBody(b)
}

func (s *Scene) AddGraphics(g SceneGraphics) {
	s.graphics = append(s.graphics, g)
}

func (scene *Scene) AddObject(o SceneObject) {
	scene.addedObjects = append(scene.addedObjects, o)
}

func (scene *Scene) addQueuedObjects() {
	// New objects could be added while we add already queued objects.
	// We'll add them in waves, until all objects are in place.
	for len(scene.addedObjects) != 0 {
		scene.tmpObjectsQueue = scene.tmpObjectsQueue[:0]
		for _, o := range scene.addedObjects {
			scene.tmpObjectsQueue = append(scene.tmpObjectsQueue, o)
		}
		scene.addedObjects = scene.addedObjects[:0]
		for _, o := range scene.tmpObjectsQueue {
			o.Init(scene)
			scene.objects = append(scene.objects, o)
		}
	}
}

func (scene *Scene) update(delta float64) {
	scene.controller.Update(delta)

	liveObjects := scene.objects[:0]
	for _, o := range scene.objects {
		if o.IsDisposed() {
			continue
		}
		o.Update(delta)
		liveObjects = append(liveObjects, o)
	}
	scene.objects = liveObjects

	scene.addQueuedObjects()

	scene.collisionEngine.Calculate()
}