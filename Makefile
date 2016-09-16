.PHONY: more \
  run test build

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
more:
	@awk 'BEGIN {FS = ":.*?## "} /^[\$$\(\)\/a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps:
	go get -d -t -v ./...

run: deps                        ## サーバ起動
	go run hearb.go

test: deps                       ## テスト
	go test -v ./...

build: deps                      ## ビルド
	go build hearb.go
