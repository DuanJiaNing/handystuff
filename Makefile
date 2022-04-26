OUTPUT_DIR=bin

run:
	go build -v -o bin/handystuff handystuff/cmd/api && bin/handystuff -c config.dev.yaml

run-ihandy:
	go build -v -o bin/handystuff handystuff/cmd/api && bin/handystuff -c config.ihandy.yaml
