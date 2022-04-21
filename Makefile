PACKAGE=github.com/olden/topdeck-sentinel-sentinel
ARTIFACTS=./build

.PHONY: vendor
vendor:
	@echo "======> vendoring dependencies"
	@go mod vendor

.PHONY: clear
clear:
	@echo "======> clearing artifacts"
	@rm -fR $(ARTIFACTS)


.PHONY: build
build: clear
	go build -o $(ARTIFACTS)/cards_scrapper ./cmd/cards_scrapper
	go build -o $(ARTIFACTS)/migrations ./cmd/migrations
	go build -o $(ARTIFACTS)/topdeck_auctions ./cmd/topdeck_auctions