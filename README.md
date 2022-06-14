# Gophet

![GitHub repo size](https://img.shields.io/github/repo-size/b1tt0/gophet) ![Lines of code](https://img.shields.io/tokei/lines/github.com/b1tt0/gophet) ![GitHub issues](https://img.shields.io/github/issues-raw/b1tt0/gophet) ![GitHub](https://img.shields.io/github/license/b1tt0/gophet)

Text editor made in Go, focused on speed and compatibility

## Configuration
You can edit your colour scheme, delays and tab size. By default, everything specified in `config.json` will override gophet defaults - that way, you can change the tab size while keeping the default colourscheme.

every `Color` property is an array - first element stands for background RGB value, second for foreground.

every `Delay` property is provided in milliseconds.

**please note!** `config.json` has to be in the same folder as the executable.
### default
```json
{
  "Color": {
    "InfoBar":    [[0  , 0  , 0  ], [255, 255, 255]],
    "TextField":  [[0  , 0  , 0  ], [255, 255, 255]],
    "Cursor":     [[255, 255, 255], [0  , 0  , 0  ]],
    "Ruler":      [[0  , 0  , 0  ], [155, 155, 155]],
    "CommandBar": [[0  , 0  , 0  ], [255, 255, 255]]
  },
  "Delay": {
    "Input": 16,
    "Draw":  16
  },
  "TabSize": 4
}
```

## Build
Run the build script for your platform with OS and ARCH as its arguments.

If you provide none, it will compile for your configuration.

Example:
```console
$ ./script/build.sh windows amd64
```
