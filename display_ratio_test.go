package ge

import (
	"fmt"
	"testing"
)

func TestInferDisplayRatio(t *testing.T) {
	tests := map[[2]int]string{
		{1024, 768}: "12:9",
		{768, 1024}: "12:9",
		{640, 480}:  "12:9",
		{800, 600}:  "12:9",

		{1080, 1920}: "16:9",
		{1920, 1080}: "16:9",
		{1600, 900}:  "16:9",
		{2560, 1440}: "16:9",
		{3840, 2160}: "16:9",

		{1680, 1050}: "16:10",
		{1280, 800}:  "16:10",
		{2560, 1600}: "16:10",

		{1440, 2880}: "18:9",
		{1080, 2160}: "18:9",

		{1920, 910}: "19:9",
		{1920, 912}: "19:9",
		{1280, 608}: "19:9",

		{2408, 1080}: "20:9",
		{1080, 2408}: "20:9",
		{2418, 1080}: "20:9",
		{1080, 2418}: "20:9",
		{2408, 1100}: "20:9",
		{1100, 2408}: "20:9",

		{1644, 3840}: "21:9",
		{3840, 1644}: "21:9",
		{720, 1680}:  "21:9",
		{1680, 720}:  "21:9",

		{1644, 4140}: "21:9",
		{4140, 1644}: "21:9",

		{1360, 768}: "16:9",
		{1366, 768}: "16:9",
	}

	for sizes, want := range tests {
		w := sizes[0]
		h := sizes[1]
		t.Run(fmt.Sprintf("%dx%d", w, h), func(t *testing.T) {
			ratioX, ratioY := inferDisplayRatio(w, h)
			have := fmt.Sprintf("%d:%d", ratioX, ratioY)
			if have != want {
				t.Fatalf("result mismatch: have %s, want %s", have, want)
			}
		})
	}
}
