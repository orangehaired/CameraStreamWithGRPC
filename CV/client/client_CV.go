package client

//Client components FOR GOCV

import (
	"fmt"
	"gocv.io/x/gocv"
	"log"
	"time"
	"errors"
)

var (
	webcam *gocv.VideoCapture
	err error
	DefaultDeviceID = 0 //Video name(string) or deviceID(int)
)

func init() {
	webcam, err = gocv.OpenVideoCapture(DefaultDeviceID)
	if err != nil {
		log.Fatalf("Can't connection on %v", DefaultDeviceID)
	}
}

func CloseClient() {
	defer webcam.Close()
	defer log.Println("Closed camera.")
}

//Deprecated.
func GetFrame() ([]byte, error) {
	webcam, err := gocv.OpenVideoCapture(DefaultDeviceID)
	if err != nil {
		log.Fatalf("Can't connection on %v", DefaultDeviceID)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	//time.Sleep(500 * time.Millisecond) //Ayağa kalksın kamera. Aksi taktirde ilk iki fotoğraf karanlık çıkıyor.
	time.Sleep(1000 * time.Millisecond)
	if ok := webcam.Read(&img); !ok {
		//fmt.Printf("Can't get frame from %v\n", DefaultDeviceID)
		return nil, errors.New(fmt.Sprintf("Can't get frame from %v\n", DefaultDeviceID))
	}
	if img.Empty() {
		log.Println("Empty Frame?")
		return nil, errors.New("Empty Frame?")
	}

	buf, err := gocv.IMEncode(".jpg", img)
	return  buf, err
}

func GetFrameWithChannel() (chan *gocv.Mat, error) {
	frames := make(chan *gocv.Mat, 1)
	go func() {
		img := gocv.NewMat()
		defer img.Close()

		for {
			if ok := webcam.Read(&img); !ok {
				fmt.Printf("cannot read device %v\n", DefaultDeviceID)
				break
			}
			if img.Empty() {
				continue
			}

			frames <- &img
		}

		close(frames)
	}()
	fmt.Println("Return is called ")
	return frames, nil

}

func MatrixToBytes(img *gocv.Mat) ([]byte, error) {
	buf, err := gocv.IMEncode(".jpg", *img)
	return buf, err
}
