FROM golang:1.19.2-bullseye

LABEL maintainer ="Marwen Ben Hriz <marwen.benhriz@gmail.com>"

WORKDIR /app
 
# Effectively tracks changes within your go.mod file
COPY go.mod .
 
RUN go mod download
 
# Copies your source code into the app directory
COPY main.go .
 
RUN go mod -o /gabackend
 
EXPOSE 8091
 
CMD [ “/gabackend” ]