OUTPUT_DIR=bin

bd:
	go build -v -o bin/handystuff handystuff/cmd/api

run: bd
	bin/handystuff

run-ihandy: run
	-c config.ihandy.yaml
