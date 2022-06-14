# gophet
text editor with focus on speed and compatibility

## configuring
you can edit your colourscheme, delays and tab size. by default, everything specified in `config.json` will override gophet defaults - that way, you can change the tab size while keeping the default colourscheme.

every `Color` property is an array - first element stands for background RGB value, second - foreground.

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

## building
run the build script for your platform with OS and ARCH as its arguments.

if you provide none, it will compile for your configuration.

example:
```console
$ ./script/build.sh windows amd64
```
