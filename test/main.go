package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pb "github.com/aabdelrahim/grpc-twil/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var recordingUrl = "https://api.twilio.com/2010-04-01/Accounts/AC858243a09aa1f5d7f9f659dd0c561242/Recordings/RE084dd6b771bc13e922b102ed6aa278f5"

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	flag.Parse()

	ctx := context.Background()

	audio, err := fetchAudio(ctx, recordingUrl)
	if err != nil {
		log.Fatalf("could not fetch audio file: %v", err)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != err {
		log.Fatalf("could not dial backend %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewSpeechToTextClient(conn)
	fmt.Printf("Client made for grpc server\n")
	request := &pb.RecognizeRequest{Audio: audio, Language: "English"}
	fmt.Printf("Sending request to grpc server\n")
	response, err := client.Recognize(context.Background(), request)
	fmt.Printf("Got response from grpc server\n")
	if err != nil {
		log.Fatalf("could not trasncribe link %v", err)
	}
	fmt.Printf("Response: %s\n", response.Text)
}

func fetchAudio(ctx context.Context, recUrl string) ([]byte, error) {
	res, err := http.Get(recUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to GET recording: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetched with status: %s", res.Status)
	}
	defer res.Body.Close()
	rec, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("could not read fetched response body: %v", err)
	}
	return rec, nil
}
