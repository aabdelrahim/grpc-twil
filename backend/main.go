package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/aabdelrahim/grpc-twil/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var SpeechUrl = "https://speech.googleapis.com/v1/speech:recognize?key=" + os.Getenv("GCLOUD_SPEECH_KEY")

// TODO replace twilio api key with google speech api key
type recognizeReq struct {
	Config struct {
		Encoding        string `json:"encoding"`
		SampleRateHertz int    `json:"sampleRateHertz"`
		LanguageCode    string `json:"languageCode"`
	} `json:"config"`
	Audio struct {
		Content string `json:"content"`
	} `json:"audio"`
}

type recognizeRes struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
	Results []struct {
		Alternatives []struct {
			Transcript string  `json:"transcript"`
			Confidence float64 `json:"confidence"`
		} `json:"alternatives"`
	}
}

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.Parse()

	log.Printf("listening on port %d \n", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("could not listen to port %d: %v", *port, err)
	}
	fmt.Printf("Check it out, it works: %s\n", os.Getenv("GCLOUD_SPEECH_KEY"))
	s := grpc.NewServer()

	pb.RegisterSpeechToTextServer(s, server{})

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve %v", err)
	}
}

type server struct {
}

func (server) Recognize(ctx context.Context, req *pb.RecognizeRequest) (*pb.RecognizeResponse, error) {
	text, err := transcribe(ctx, req.Audio, req.Language)
	if err != nil {
		// http.Error("could not transcribe", http.StatusInternalServerError)
		// log.Errorf("could not transcribe: %v", err)
		return nil, err
	}

	return &pb.RecognizeResponse{Text: text}, nil
}

func transcribe(ctx context.Context, audio []byte, language string) (string, error) {
	// rec, err := fetchAudio(ctx, recUrl)
	// if err != nil {
	// 	return "", err
	// }

	text, err := fetchTranscription(ctx, audio, language)
	if err != nil {
		return "", err
	}

	return text, nil
}

func fetchTranscription(ctx context.Context, rec []byte, language string) (string, error) {
	var req recognizeReq
	req.Config.Encoding = "LINEAR16"
	req.Config.SampleRateHertz = 8000

	switch language {
	case "Arabic":
		req.Config.LanguageCode = "ar-JO"
	case "German":
		req.Config.LanguageCode = "de-DE"
	default:
		req.Config.LanguageCode = "en-US"
	}

	req.Audio.Content = base64.StdEncoding.EncodeToString(rec)

	j, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("could not encode speech request: %v", err)
	}

	// fetchClient := urlfetch.Client(ctx)
	resp, err := http.Post(SpeechUrl, "application/json", bytes.NewReader(j))
	if err != nil {
		return "", fmt.Errorf("could not transcribe: %v", err)
	}

	var data recognizeRes

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("could not decode response: %v", err)
	}
	if data.Error.Code != 0 {
		return "", fmt.Errorf("speech api error: %d %s %s. Given key: %s", data.Error.Code, data.Error.Status, data.Error.Message, os.Getenv("GCLOUD_SPEECH_KEY"))
	}

	if len(data.Results) == 0 || len(data.Results[0].Alternatives) == 0 {
		return "", fmt.Errorf("no transcription found")
	}
	return data.Results[0].Alternatives[0].Transcript, nil
}
