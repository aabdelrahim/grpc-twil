## Personal Project 
### Purpose

Application accepts phone calls and prompts user for language and accepts input as keypresses. (1 = English, 2 = German etc.)
Application asks user for message to transcribe
Application sends the user a text message with the contents of the message transcribed in the selected language

#### API
- Define a formal API contract in protobuf 3
- Generate a go contract from the proto contract

#### Background
- Create a base docker image to use as a base for a gRPC server
- Kubernetes yaml used to define the deployments and services
  - deployments are the actual pieces running the server defined in main.go
  - services are used as a load balancer and is the static IP hit externally
- Server accepts gRPC requests that follow the proto definition of the API 
  - does the actual service work (converting audio files into text)
  
 #### Client 
 - Define the deployment to Google's App Engine
 - Define a handler that accepts TwilML (modified XML) requests 
 
