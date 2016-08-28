
.PHONY: all
all: install

.PHONY: install
install:
	go install .

.PHONY: vendor
vendor:
	@go get -u github.com/kardianos/govendor
	@govendor init
	@govendor fetch github.com/ChimeraCoder/anaconda
