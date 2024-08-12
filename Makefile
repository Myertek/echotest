

init:
	$(shell mkdir -p ./cmd/web)
	$(shell mkdir -p ./internal/models)
	$(shell mkdir -p ./ui/html/pages)
	$(shell mkdir -p ./ui/html/partials)
	$(shell mkdir -p ./ui/html/partials)
	$(shell mkdir -p ./ui/static/css)
	$(shell mkdir -p ./ui/static/js)
	$(shell mkdir -p ./ui/static/images)
	touch .env
	touch .gitignore
	go mod init github.com/myertek/echotest
	go get github.com/labstack/echo/v4

run:
	go run ./cmd/web/


	