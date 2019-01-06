# SVGSanitizer

SVG sanitizing tool developed in Golang.

## Sanitize SVG

SVG file like this

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 488 488">
  <defs>
    <style>
      .cls-1 {
        fill: rgba(255,255,255,0);
      }
    </style>
  </defs>
  <g id="my_circle_foo" data-name="my circle foo" class="cls-1" transform="translate(76 -38)">
    <circle class="cls-4 bar" cx="244" cy="244" r="244"/>
  </g>
</svg>
```

will be converted to

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 488 488">
  <defs>
    <style>
      .tcuAxhxKQFDaFpLSjFbc {
      fill: rgba(255,255,255,0);
      }
    </style>
  </defs>
  <g id="XVlBzgbaiCMRAjWwhTHc" class="tcuAxhxKQFDaFpLSjFbc" transform="translate(76 -38)">
    <circle class="XoEFfRsWxPLDnJObCsNV lgTeMaPEZQleQYhYzRyW" cx="244" cy="244" r="244" />
  </g>
</svg>
```

by Replacing id/class attrs and Removing custom data attrs. 

## Usage

```bash
$ svgsanitizer <filepaths>  
```

For example,

```bash
$ ls *.svg | xargs svgsanitizer
Generated. /path_to_svg/sanitized_hoge.svg
Generated. /path_to_svg/sanitized_fuga.svg
...
```


## Install

[Install Go](https://golang.org/) and `$ go get && go install`

## Bugs

There are many known bugs.

- critical: Map will not assure orders, so if there are `.class1` and `.class10`, replacing ids/classes may be done improperly.
- Panic will occur if invalid SVG (or any files other than SVG) paths are specified.
