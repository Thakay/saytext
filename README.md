A simple gRPC + Go application that will send text to server and gets back the speech data of it. It uses Flite for making the speech from the text.

The app uses MakeFile and Docker to compile the binary and docker together.

To Run: 
  First Run MakeFile in backend directory.
  then docker run with the port assigned 8080:8080 
  then go run main in client with -d "message
The Say is the client that will be run seperately.
