default: install

APP_NAME=go-ytm

build:
	go build -o ${APP_NAME}

install: build
	go install

clean:
	rm -f $(APP_NAME)