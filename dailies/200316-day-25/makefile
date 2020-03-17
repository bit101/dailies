default:
	@tmux_send "make build"

build: clean
	@go run main.go

clean:
	@rm -rf frames
	@mkdir frames
