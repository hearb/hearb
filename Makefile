.PHONY: more \
  run test build

PROJ_PATH=$(GOPATH)/src/github.com/hearb/hearb

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
more:
	@awk 'BEGIN {FS = ":.*?## "} /^[\$$\(\)\/a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

mklink:                                 ## GOPATHにシンボリックリンクを作成する
	ln -fhs $(CURDIR)/server $(PROJ_PATH)

deps:
	cd $(PROJ_PATH); go get -d -t -v ./...

run:                                    ## サーバ起動
	cd $(PROJ_PATH); go run hearb.go

test:                                   ## テスト
	cd $(PROJ_PATH); go test -v ./...

build:                                  ## ビルド
	cd $(PROJ_PATH); go build hearb.go
