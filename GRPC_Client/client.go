package main

import (
	"context"

	CV_Client "github.com/orangehaired/CameraStreamWithGRPC/CV/client"

	"log"
	"runtime"

	pb "github.com/orangehaired/CameraStreamWithGRPC/my_proto"
	"google.golang.org/grpc"
)

const connection = "localhost:1337"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connecting error: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	stream, err := pb.NewCameraClient(conn).Analyse(ctx)
	if err != nil {
		log.Fatal("Stream starting: ", err)
	}

	defer CV_Client.CloseClient()
	defer log.Println("CloseClient is called")

	go func() {
		defer cancel()
		defer stream.CloseSend()

		//New
		for {
			frames, err := CV_Client.GetFrameWithChannel()
			if err != nil {
				log.Fatal(err)
			}
			for {
				select {
				case frame, _ := <-frames:
					//if !ok {
					//	close(frames)
					//}
					frameToByteArray, err := CV_Client.MatrixToBytes(frame)
					if err != nil {
						log.Fatalf("Matrix to Bytes converting error: %v", err)
					}
					message := &pb.ImageRequest{Image: frameToByteArray}
					stream.Send(message)
					log.Println("Sended frame.")
					response, err := stream.Recv()
					if err != nil {
						log.Println("Error response: ", err)
					}
					log.Printf("Result: %d", response.Reply)
					//time.Sleep(5 * time.Second)
				}
			}

		}

		//Old
		/*
			for {
				frame, err := CV.GetFrame()
				if err != nil {
					log.Fatal(err)
				}
				message := &pb.ImageRequest{Image: frame}
				stream.Send(message)
				log.Println("Sended frame.")
				response, err := stream.Recv()
				if err != nil{
					log.Println("Error response: ", err)
					}
				log.Printf("Result: %d", response.Reply)
				//time.Sleep(3 * time.Second)
			}
		*/
	}()

	select {}
}
