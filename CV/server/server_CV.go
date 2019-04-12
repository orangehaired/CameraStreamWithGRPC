package server

import (
    "fmt"
    "github.com/orangehaired/CameraStreamWithGRPC/CV"
    "gocv.io/x/gocv"
    "io/ioutil"
    "log"
    "os"
    "errors"
)

//Server components FOR GOCV
var (
    classifier gocv.CascadeClassifier

)
func init()  {
    os.MkdirAll(CV.SavedImagesPath, os.ModePerm)


    classifier = gocv.NewCascadeClassifier()  // Preparing classifier for recognize faces

    if !classifier.Load(CV.DataPath+ "/data/haarcascades/haarcascade_frontalface_default.xml") {
        log.Println("Error reading cascade file: data/haarcascade_frontalface_default.xml")
        //return 0, errors.New("Error reading cascade file: data/haarcascade_frontalface_default.xml")
    }
}

func CloseServer() {
    defer classifier.Close()
}
func Save(filename string, buf []byte) (error) {
    if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
        return err
    }
    return nil
}

func HowManyFace(buf []byte) (int32, error) {
    img, err := gocv.IMDecode(buf,gocv.IMReadColor)
    if err != nil {
        return 0, errors.New(fmt.Sprintf("Buffer bytes decoding to matris: %v", err))
    }

    rects := classifier.DetectMultiScale(img)
    //fmt.Printf("Found %d faces\n", len(rects))
    return int32(len(rects)), nil
}