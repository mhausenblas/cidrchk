release_version:= v0.1

export GO111MODULE=on

.PHONY: bin
bin:
	go build -o bin/cidrchk github.com/mhausenblas/cidrchk

.PHONY: release
release:
	curl -sL https://git.io/goreleaser | bash -s -- --rm-dist --config .goreleaser.yml

.PHONY: publish
publish:
	git tag ${release_version}
	git push --tags