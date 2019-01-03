# SVGSanitizer

SVG sanitizing tool developed in Golang.

## Sanitize SVG

SVG like

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

will convert to

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

## Usage

```bash
$ svgsanitizer <filepaths>  
```

For example,

```bash
$ ls *.svg | xargs svgsanitizer
generated /Users/~~~/svgsanitizer/sanitized_hoge.svg.
generated /Users/~~~/svgsanitizer/sanitized_fuga.svg.
```


## Install

[Install Go](https://golang.org/) and `$ go get && go install`

## Bugs

There are much known bugs.

- Panic will occur if invalid SVG (or any files other than SVG) path specified.
- Convert `<style>` elements inpropery in particular case. 