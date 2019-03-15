# Web-rotate [![Build Status](https://travis-ci.com/ArthurHlt/web-rotate.svg?branch=master)](https://travis-ci.com/ArthurHlt/web-rotate)

Rotate website view on a chromium browser for monitoring purpose.

**Requirements**: A chromium based browser installed on your system (chrome is one of them)

## install

### On *nix system

You can install this via the command-line with either `curl` or `wget`.

#### via curl

```bash
$ sh -c "$(curl -fsSL https://raw.github.com/ArthurHlt/web-rotate/master/bin/install.sh)"
```

#### via wget

```bash
$ sh -c "$(wget https://raw.github.com/ArthurHlt/web-rotate/master/bin/install.sh -O -)"
```

### On windows

You can install it by downloading the `.exe` corresponding to your cpu from releases page: https://github.com/ArthurHlt/web-rotate/releases .
Alternatively, if you have terminal interpreting shell you can also use command line script above, it will download file in your current working dir.

### From go command line

Simply run in terminal:

```bash
$ go get github.com/ArthurHlt/web-rotate
```

## Usage

```
Usage of web-rotate:
  -config string
    	set config path (default "config.yml")
  -version
    	see version
  -help
        this help
```

## Configuration

```yaml
# Set to true to run chromium in fullscreen
fullscreen: false
# Pages to rotate
pages:
  # cartridge is the text set on the cartridge
  # this is optional, let empty if you don't want cartridge
  - cartridge: Github
  # Url to target, this is mandatory
    url: http://github.com/
  # You can pass headers when making request
    headers:
      Authorization: toto
  # Show this page during this time before going to next one
  # if not set, default to 30s
    duration: 10s
  # Change css style of your cartridge
  # see default style below
    style_cartridge:
    # change css style for the box
    # Configuration below is the default configuration when not overwrite/set
      box:
        background: white
        border: 1px solid black
        border-radius: 5px
        font-size: xx-large
        height: 120px
        left: 0
        line-height: 120px
        margin: 10px
        opacity: 0.7
        position: absolute
        text-align: center
        top: 0
        vertical-align: middle
        width: 30%
        z-index: 50000
      # change css style for the text
      # Configuration below is the default configuration when not overwrite/set
      text:
        color: blue
        opacity: 1
  - cartridge: Google
    url: http://google.com/
    headers:
      Authorization: google
    duration: 10s
```