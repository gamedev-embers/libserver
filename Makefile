root_dir:=$(CURDIR)


test:
	go test ./...


benchmark:
	go test ./... -bench=.


release:
	[ -n "$(tag)" ] || (echo "Please provide a tag. Usage: make tag tag=vX.Y.Z" && exit 1)
	git tag -a $(tag) -m "Release $(tag)"
	git push origin $(tag)

