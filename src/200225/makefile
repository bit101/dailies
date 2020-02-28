default:
	tmux_send "make build"

build: clean
	go run main.go

clean:
	rm -f frames/*
	rm -f *.mp4
