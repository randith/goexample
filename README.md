# goexample

## notes
1. I included a build script to make a docker container from the PoC I did recently.  It is very early in development but a good first step to a real ci/cd build of a go service.  It is not complete, missing among other things versioning."
2. I started running short on time so I did not add tests for 4-6.  I would add these as well as a round-trip test using the endpoints.
3. I interpreted step 6 "hash requests" to include creating and getting the hashes
4. I left TODO comments in the code to shed light on my thoughts.  I do not like leaving them in the code normally as they grow stale fast.

## running
1. prerequisites
  * docker
  * this repository installed properly in the GOPATH
  * being in the root directory of this repository 
2. execute build script
  * ```./build.sh```
3. run container
  * ```docker run -p 8080:8080 pwhash```

## developing

### setting up go environment

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

### add or update dependencies
Of course there are none yet since only using standard libraries

* ```dep ensure```
* commit and push dependencies