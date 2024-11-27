root_dir:=$(CURDIR)


test:
	go test ./...


benchmark:
	go test ./... -bench=.
