# dailies
daily code experiments

requires https://cairographics.org/download/

requires go 1.3 or higher. uses go modules which should handle all the dependencies

makefiles make use of my `tmux_send` which I've included here. if you use tmux and a console based editor like vim, put that script in your path somewhere and set up a key binding like F5 to run `make`. this will open a side panel to the right and do the build in that, leaving your focus on the code. Or you can just type `go run main.go` in the same directory.

Each file has a line that calls `util.ViewImage`. This requires "eog" (eye of gnome) to be installed on Linux, or will use Quicktime (qlmanage) on a Mac.

The animation examples require an empty folder called `frames` in the same folder as the `main.go` file. And they require ImageMagick to be installed on your system to create the gifs. Or ffmpeg to create videos.
