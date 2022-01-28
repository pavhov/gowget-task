init:
	mkdir -p bin pkg src/service/src

mod-init:
	cd ./src/service && \
	go mod init gowget

install:
	cd ./src/service/src && \
	go get -d -v

download:
	cd ./src/service/src && \
	go mod download

update: clean
	cd ./src/service/src && \
	go get -u -d -v

clean:
	cd ./src/service/src && \
	go clean --modcache

tidy:
	cd ./src/service/src && \
	go mod tidy

build:
	cd ./src/service && \
	go build -ldflags="-w -s" -o ./pkg/gowget ./src && \
	chmod +x ./pkg/gowget

run:
	cd ./src/service && \
	./pkg/gowget "http://ipv4.download.thinkbroadband.com/1GB.zip" "http://ipv4.download.thinkbroadband.com/512MB.zip" "http://ipv4.download.thinkbroadband.com/200MB.zip" "http://ipv4.download.thinkbroadband.com/100MB.zip" "http://ipv4.download.thinkbroadband.com/50MB.zip" "http://ipv4.download.thinkbroadband.com/20MB.zip" "http://ipv4.download.thinkbroadband.com/10MB.zip" "http://ipv4.download.thinkbroadband.com/5MB.zip"
