# Teapot

A simple teapot server written in Go. [teapot.now.sh](https://teapot.now.sh)

![teapot server](https://user-images.githubusercontent.com/7525670/29248631-3298c1c4-801d-11e7-8619-2caef2ff15aa.png)

## Usage

```
$ curl -i https://teapot.now.sh
  HTTP/1.1 418 I'm a teapot
  The requested entity body is short and stout.

$ curl -X BREW -i https://teapot.now.sh
  HTTP/1.1 200 OK
  Brewing tea.
```

## Get your own teapot server

```sh
# Download the code
git clone git@github.com:mxstbr/teapot.git

# Deploy your own version (requires https://now.sh)
now
```

## License

Licensed under the MIT License, Copyright 2017 Maximilian Stoiber.
