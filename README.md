# goexample

## getting started

1. install go and dep
  * ```brew install go dep```
2. setup go directory and vars 
  * ```mkdir ~/go```
  * ```vi ~/.bash_profile```
    * GOPATH=/Users/rthomas/go
    * GOBIN=$GOPATH/bin
    * GOROOT=/usr/local/Cellar/go/1.9.2/libexec
  * ```source ~/.bash_profile``` 
  * ```mkdir $GOPATH/bin $GOPATH/pkg $GOPATH/src```
3. get this service 
  * ```go get github.com/randith/goexample```

## developing

### add or update dependencies

* ```dep ensure```
* commit and push dependencies