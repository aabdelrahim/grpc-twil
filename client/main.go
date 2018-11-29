package twilio_client

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/aabdelrahim/grpc-twil/api"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"google.golang.org/grpc"
)

const (
	welcomeMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Gather numDigits="1">
			<Say voice="woman">Please select a language.</Say>
			<Say voice="woman">Press 1 for English.</Say>
			<Say voice="woman">Press 2 for German.</Say>
			<Say voice="woman">Press 3 for Arabic.</Say>
			<Say voice="woman">Press star to replay this message.</Say>
		</Gather>
	</Response>
	`
	recordMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Say> You have selected %s</Say>
		<Say> What message would you like to have transcribed </Say>
		<Record maxLength="20" timeout="3"/>
	</Response>
		`
	repeatMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Say> I'm sorry I didn't catch that, can you repeat your message?</Say>
		<Record maxLength="20" timeout="3"/>
	</Response>
	`
	sendMsg = `<?xml version="1.0" encoding="UTF-8"?>
	<Response>
		<Sms>%s</Sms>
	</Response>
	`
)

var Backend = flag.String("b", "35.232.199.229:8080", "address of the grpc backend")

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	flag.Parse()
	ctx := appengine.NewContext(req)

	var language string
	var recUrl string

	_ = req.ParseForm() // TODO handle error

	digits := req.FormValue("Digits")
	recUrl = req.FormValue("RecordingUrl")

	log.Infof(ctx, "digit pressed: %s", digits)

	switch digits {
	case "1":
		language = "English"
	case "2":
		language = "German"
	case "3":
		language = "Arabic"
	default:
		if recUrl != "" {
			break
		}
		fmt.Fprint(w, welcomeMsg)
		return
	}

	fmt.Fprintf(w, recordMsg, language)

	log.Infof(ctx, "recordingUrl: %s", recUrl)

	if recUrl == "" {
		return
	}

	log.Infof(ctx, "Fetching Audio\n")
	audio, err := fetchAudio(ctx, recUrl)
	if err != nil {
		log.Criticalf(ctx, "could not fetch audio file: %v", err)
	}

	log.Infof(ctx, "Dialing Backend\n")
	conn, err := grpc.Dial(*Backend, grpc.WithInsecure())
	if err != err {
		log.Criticalf(ctx, "could not dial backend %s: %v", *Backend, err)
	}
	defer conn.Close()

	client := pb.NewSpeechToTextClient(conn)
	request := &pb.RecognizeRequest{Audio: audio, Language: language}
	log.Infof(ctx, "Sending Request\n")
	res, err := client.Recognize(ctx, request)
	if err != nil {
		log.Criticalf(ctx, "could not recognize request: %v", err)
	}
	fmt.Fprintf(w, sendMsg, res.Text)
	log.Infof(ctx, "Got response: %s", res.Text)
}

func fetchAudio(ctx context.Context, recUrl string) ([]byte, error) {
	fetchClient := urlfetch.Client(ctx)
	res, err := fetchClient.Get(recUrl)
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
