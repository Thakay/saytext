package main

import (
	"flag"
	"fmt"
	pb "github.com/Thakay/saytext/api/github.com/Thakay/saytext/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()
	log.Printf("Listening to %d...\n", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	pb.UnimplementedTextToSpeechServer
}

func (s server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := os.CreateTemp("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}
	err = f.Close()
	if err != nil {
		return nil, fmt.Errorf("could not close temp file: %v", err)
	}
	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}
	data, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read temp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
