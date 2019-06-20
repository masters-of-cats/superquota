all: setquota setpid
.PHONY: setquota setpid clean

setquota: setquota/main.go
	CGO_ENABLED=1 go build -o setquota/setquota setquota/main.go
setpid: setpid/main.go
	CGO_ENABLED=1 go build -o setpid/setpid setpid/main.go
clean:
	rm setquota/setquota
	rm setpid/setpid

