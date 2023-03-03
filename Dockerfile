FROM golang:alpine 

COPY . /app
#set up that directory as working directory
WORKDIR /app
#open port 8080
EXPOSE 8080
#to build linux bin
RUN CGO_ENABLED=0 GOOS=linux go build -o main
#CMD to start our web app
CMD ["./main"]