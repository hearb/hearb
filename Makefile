.PHONY: more \
  run test build

ORG_PATH=$(GOPATH)/src/github.com/hearb
PROJ_PATH=$(ORG_PATH)/hearb

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
more:
	@awk 'BEGIN {FS = ":.*?## "} /^[\$$\(\)\/a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

mklink:                                 ## GOPATHにシンボリックリンクを作成する
	mkdir -p $(ORG_PATH)
	ln -fhs $(CURDIR) $(PROJ_PATH)

deps:
	cd $(PROJ_PATH); go get -d -t -v ./...

run: mklink deps                        ## サーバ起動
	cd $(PROJ_PATH); go run hearb.go

test: mklink deps                       ## テスト
	cd $(PROJ_PATH); go test -v ./...

build: mklink deps                      ## ビルド
	cd $(PROJ_PATH); go build hearb.go
