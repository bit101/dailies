default:
	tmux_send "make build"

build: clean
	go run main.go
	open 200226.mp4

clean:
	rm -f frames/*
	rm -f *.mp4
