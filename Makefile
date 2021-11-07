
BUILD_ENV = CGO_ENABLED=0
PACKAGES = $$(go list ./... | grep -v /vendor/)

builddir:
	mkdir -p build

build: builddir
	${BUILD_ENV} go build -o build/urlshortener ./cmd/urlshortener/urlshortener.go

run: builddir build
	./build/urlshortener

test:
	${BUILD_ENV} go test ${PACKAGES}

clean:
	rm -rf build

cover: builddir
	${BUILD_ENV} go test -v -covermode=count -coverprofile=build/coverage.out $(PACKAGES)
	go tool cover -html=build/coverage.out -o build/coverage.html
	go2xunit -input build/test.out -output build/test.xml
	! grep -e "--- FAIL" -e "^FAIL" build/test.out


image:
	docker build -t shortener .

docker-run:
	docker run -p 8443:8080 --name shortener -d docker.io/library/shortener

docker-down:
	docker rm -f shortener

docker-redeploy: docker-down image docker-run

vendor:
	go mod vendor

update:
	go get -u
