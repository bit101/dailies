# dailies
daily code experiments

requires https://cairographics.org/download/

requires go 1.3 or higher. uses go modules which should handle all the dependencies

makefiles make use of my `tmux_send` which I've included here. if you use tmux and a console based editor like vim, put that script in your path somewhere and set up a key binding like F5 to run `make`. this will open a side panel to the right and do the build in that, leaving your focus on the code.
