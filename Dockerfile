FROM golang:1.15.7-alpine
RUN mkdir /go/blogbe
## We copy everything in the root directory
## into our /app directory
ADD . /go/blogbe
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
ENV APP_HOME /go/blogbe
WORKDIR $APP_HOME
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main 
EXPOSE 4000
CMD ["/go/blogbe/main"]