FROM golang:1.22
LABEL authors="renato druzian"

WORKDIR /crawler
COPY go.mod .
RUN ["go", "mod", "tidy"]

COPY . .
CMD ["go", "run", "main.go"]