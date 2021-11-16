.PHONY: debug
debug:
	go build -gcflags="all=-N -l" -o hash .
