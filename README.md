go mod tidy
go run main.go

Caso tenha algum erro para executar o script, rode o seguinte comando:
go run github.com/playwright-community/playwright-go/cmd/playwright install


docker build -t crawler .
docker run -d crawler