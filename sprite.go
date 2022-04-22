package ge

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/ge/gemath"
)

type Sprite struct {
	image *ebiten.Image

	Pos *gemath.Vec

	Rotation *gemath.Rad

	Scaling float64

	Centered bool

	Offset gemath.Vec
	Width  float64
	Height float64

	disposed bool
}

func NewSprite(img *ebiten.Image) *Sprite {
	w, h := img.Size()
	sprite := &Sprite{
		image:    img,
		Width:    float64(w),
		Height:   float64(h),
		Centered: true,
		Scaling:  1,
	}
	return sprite
}

func (s *Sprite) ImageWidth() float64 {
	w, _ := s.image.Size()
	return float64(w)
}

func (s *Sprite) ImageHeight() float64 {
	_, h := s.image.Size()
	return float64(h)
}

func (s *Sprite) IsDisposed() bool {
	return s.disposed
}

func (s *Sprite) Dispose() {
	s.disposed = true
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	var drawOptions ebiten.DrawImageOptions

	var origin gemath.Vec
	if s.Centered {
		origin = gemath.Vec{X: s.Width / 2, Y: s.Height / 2}
	}

	drawOptions.GeoM.Translate(-origin.X, -origin.Y)
	if s.Rotation != nil {
		drawOptions.GeoM.Rotate(float64(*s.Rotation))
	}
	if s.Scaling != 1 {
		drawOptions.GeoM.Scale(s.Scaling, s.Scaling)
	}
	drawOptions.GeoM.Translate(origin.X, origin.Y)

	drawOptions.GeoM.Translate(s.Pos.X-origin.X, s.Pos.Y-origin.Y)

	subImage := s.image.SubImage(image.Rectangle{
		Min: image.Point{
			X: int(s.Offset.X),
			Y: int(s.Offset.Y),
		},
		Max: image.Point{
			X: int(s.Offset.X + s.Width),
			Y: int(s.Offset.Y + s.Height),
		},
	}).(*ebiten.Image)
	screen.DrawImage(subImage, &drawOptions)
}