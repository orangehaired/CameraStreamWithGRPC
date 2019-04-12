package CV

import (
	"path/filepath"
)



var DataPath, _ = filepath.Abs("../CV/") // /Users/ahmet/go/src/github.com/orangehaired/CameraStreamWithGRPC/CV
var SavedImagesPath = DataPath + "/" + filepath.Join(".", "/saved_images") // /Users/ahmet/go/src/github.com/orangehaired/CameraStreamWithGRPC/CV/saved_images
