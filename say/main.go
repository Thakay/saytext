package main

import (
	"context"
	"flag"
	pb "github.com/Thakay/saytext/api/github.com/Thakay/saytext/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {

	backend := flag.String("b", "localhost:8080", "addr of the say backend")
	output := flag.String("o", "output.wav", "wav file for output")
	textData := flag.String("d", "template data", "message to send")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect", err)
	}

	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{
		Text: *textData}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(*output, res.Audio, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
