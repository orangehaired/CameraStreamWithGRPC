package main

import (
	CV_Server "github.com/orangehaired/CameraStreamWithGRPC/CV/server"
	pb "github.com/orangehaired/CameraStreamWithGRPC/my_proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"runtime"
)

const PORT = ":1337"


//return status.Errorf(codes.Unimplemented, "not implemented")
type Server struct {}



func (s *Server) Analyse (stream pb.Camera_AnalyseServer) (error) {
	log.Println("Stream Started..")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}
		if err != nil {
			log.Println("ERROR: Incoming from client:", err)
			return err
		}
		howmanyface, err := CV_Server.HowManyFace(in.Image)
		if err != nil {
			log.Fatalln("How many face error", err)
		}
		stream.Send(&pb.ImageReply{Reply: howmanyface})
		log.Println("Sended faces.")

		/*
		//Save photos
		timestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
		writePath := fmt.Sprintf("%s/image%s.jpg", CV.SavedImagesPath, timestamp)
		//log.Println("Saved:", CV.SavedImagesPath, "DataPath:", CV.DataPath)
		CV_Server.Save(writePath, in.Image)
		*/

	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	s := grpc.NewServer()
	pb.RegisterCameraServer(s, &Server{})

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Listening error: %v", err)
	}

	defer CV_Server.CloseServer()

	log.Println("Starting..")
	if err := s.Serve(listen); err != nil{
		log.Fatalf("Serving error: %v", err)
	}

}
