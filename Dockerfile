
FROM golang:1.15.1

# Grab the source code and add it to the workspace.
ADD . /go/src/label

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

# Use the revel CLI to start up our application.
CMD ["revel","run", "src/label"]

# Open up the port where the app is running.
EXPOSE 9000
