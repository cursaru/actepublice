FROM golang:latest
RUN mkdir /app
WORKDIR /app
ADD . /app
EXPOSE 7000
RUN go build -o main
CMD ./main
