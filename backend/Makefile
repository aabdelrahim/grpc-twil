build:
	GOOS=linux go build -o app
	docker build -t gcr.io/grpc-twil/recognize .
	rm -f app

push:
	gcloud docker -- push gcr.io/grpc-twil/recognize

test:
	docker run -p 8080:8080 -e GCLOUD_SPEECH_KEY='AIzaSyCCFyeDruNfn6e65ls5ho3UQuKd0U5SDjE' gcr.io/grpc-twil/recognize 