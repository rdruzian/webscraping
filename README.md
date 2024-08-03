go mod tidy
go run main.go


docker build -t crawler .
docker run -d crawler