.PHONY: all test

all: test

release:
	cd api/cmd/formation && make release VERSION=$(VERSION)
	docker build -t convox/api:$(VERSION) .
	docker push convox/api:$(VERSION)
	docker tag convox/api:$(VERSION) 922560784203.dkr.ecr.us-east-1.amazonaws.com/rack:$(VERSION)
	$(aws ecr get-login)
	docker push 922560784203.dkr.ecr.us-east-1.amazonaws.com/rack:$(VERSION)
	mkdir -p /tmp/release/$(VERSION)
	cd /tmp/release/$(VERSION)
	jq '.Parameters.Version.Default |= "$(VERSION)"' api/dist/kernel.json > kernel.json
	aws s3 cp kernel.json s3://convox/release/$(VERSION)/formation.json --acl public-read

test-deps:
	go get -t -u ./...

test:
	docker info >/dev/null
	go get -t ./...
	go test -v -cover ./...

vendor:
	godep save -r ./...
