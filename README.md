# Crop Viewport
Crops view from string given position and size.  
**With ANSI support**

![ansitext](https://github.com/ogios/clipviewport/assets/96933655/c95e3338-b297-45f2-9da4-00be3ed5e3ff)

For **Usage** checkout [`examples`](https://github.com/ogios/clipviewport/tree/master/examples)

## Keymap

Keymaps are inside the model struct, can be changed if you want to.

default keymaps:
- `j` move down 1 line
- `k` move up 1 line
- `ctrl+d` move down half page
- `ctrl+u` move up half page
- `h` move left 1 column
- `k` move right 1 column
- `H` move left half row
- `L` move right half row

## wcwidth related

- if not enough place for a char width over 1, the char will not display and replaced to (multiple) white space.
- `TAB(\t)` will be replaced by 4 `SPACE( )`, and it's hard coded, not sure if [`processRune`](https://github.com/ogios/clipviewport/blob/master/process/process.go#L42) should be available to access from the outside



## Others


So, this is actually the first time i use `go` to make stuffs for others to use. Any suggestion is welcomed here
