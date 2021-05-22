.PHONY: build

build:
	sam build

package:
	sam package

deploy:
	sam deploy --guided
