# Go Speedtest
Simple internet speedtest written in Go


## Go Speedtest
1. Download go1.5beta2 here (https://golang.org/dl/) to `/home/adam/Downloads/gobeta`

2. Install Gomobile using go1.5beta2. This is a temp install of GoBeta, didnt want to mess up my current env
```bash
export PATH=/home/adam/Downloads/gobeta/bin/:$PATH
export PATH=/home/adam/Downloads/gobeta/temp_path/bin/:$PATH
export GOPATH=/home/adam/Downloads/gobeta/temp_path/
export GOROOT=/home/adam/Downloads/gobeta/
go get golang.org/x/mobile/cmd/gomobile
```

Build my speedtest app
```bash
ANDROID_HOME="/home/adam/Android/Sdk/" gomobile bind github.com/dr4ke616/gospeedtest
```


## Problems
Some reason gomobile doesnt like `net/http` and I get the following errors:
```
gomobile: type net/http.Request not defined in package package nw_speedtest ("github.com/dr4ke616/gospeedtest/nw_speedtest")
type net/http.Request not defined in package package nw_speedtest ("github.com/dr4ke616/gospeedtest/nw_speedtest")
type net/http.Client not defined in package package nw_speedtest ("github.com/dr4ke616/gospeedtest/nw_speedtest")
```


## TODO
Either:
- Fix or wait for a release to fix the issue with gomobile so i can complete the binding process
- Understand OpenGL and print the results on the screen, without the need to bind to Android/iOS project


## Links
http://www.sajalkayan.com/post/android-apps-golang.html
