package sprite

import (
	"image"
	"image/png"
	"os"
)

var resourceCache = make(map[string]image.Image)

func Load(path string) (image.Image, error) {
	if i, ok := resourceCache[path]; ok {
		return i, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	i, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	resourceCache[path] = i
	return i, nil
}
