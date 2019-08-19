VERSION=V0.1.0
ifndef GOPATH
	$(error GOPATH is not set)
endif
GOAPTH = $(firstword $(subst :, ,${GOPATH}))

compile:
	cd cmd/hashqd && \
	go build -ldflags "-s -w \
	-X github.com/gokultp/hashqd/internal/version.Version=${VERSION} \
	-X github.com/gokultp/hashqd/internal/version.GitVersion=`git rev-parse HEAD` \
	-X github.com/gokultp/hashqd/internal/version.BuildTime=`date +%FT%T%z` " \
	-o ../../build/hashqd

.PHONY:
	build


test:
	go test -v ./internal/session/