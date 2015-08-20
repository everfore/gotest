#	Golang test Benchmark&Test

##	nil [bench]
 	
 	(*struct)(nil)

 	new(struct)

 	&struct{}

![New_bench](https://raw.githubusercontent.com/everfore/gotest/master/nil/nil.png "New Benchmark")


##	mail

	send mail
	
##	protobuf

![protobuf_bench](https://raw.githubusercontent.com/everfore/gotest/master/protobuf/OSMsg_test.png "protobuf Benchmark")

![2nd protobuf_bench](https://raw.githubusercontent.com/everfore/gotest/master/protobuf/OSMsg2_test.png "2nd protobuf Benchmark")

## b2s

	bytes <===> string

![BS_bench](https://raw.githubusercontent.com/everfore/gotest/master/b2s/test.png "BS Benchmark")


##	go1.5 build @windows

[gcc-TMD](http://sourceforge.net/projects/tdm-gcc/)

### env

export CC=C:\tmd-gcc\bin\gcc.exe

export CGO_ENABLED=1

#build go1.5 with go1.4

export GOROOT_BOOTSTRAP=go1.4

export GOROOT=...

export GOBIN=...

export GOPATH=...

PATH=%...%:$PATH
