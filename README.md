# Go Speedtest
Simple internet speedtest written in Go for Android

![Android & Go image](doc/android_go.png)


## Installation
1. Download go1.5beta2 [here](https://golang.org/dl/)

2. Install Gomobile using go1.5beta2. This is a temp install of GoBeta, didnt want to mess up my current env
```bash
$ export PATH=/home/adam/Downloads/gobeta/bin/:$PATH
$ export PATH=/home/adam/Downloads/gobeta/temp_path/bin/:$PATH
$ export GOPATH=/home/adam/Downloads/gobeta/temp_path/
$ export GOROOT=/home/adam/Downloads/gobeta/
$ go get golang.org/x/mobile/cmd/gomobile
```

Build my speedtest app and create an APK
```bash
$ go get -d github.com/dr4ke616/gospeedtest
$ cd github.com/dr4ke616/gospeedtest
$ gomobile build . # will build an APK
```

### Code usage
To use gospeedtest as a library:
```go
import (
    ...

    "github.com/dr4ke616/gospeedtest/nw_speedtest"
)

st := nw_speedtest.Speedtest{
    FileLocation: "http://download.thinkbroadband.com/10MB.zip",
    Verbos:       true,
}
result, err := st.Start()
```


## Links
http://www.sajalkayan.com/post/android-apps-golang.html
