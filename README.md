#	Golang Try & Benchmark & Test

##	nil
 	
 	(*struct)(nil)

 	new(struct)

 	&struct{}

![New_bench](https://raw.githubusercontent.com/everfore/gotest/master/nil/nil.png "New Benchmark")

	
##	protobuf

![protobuf_bench](https://raw.githubusercontent.com/everfore/gotest/master/protobuf/OSMsg_test.png "protobuf Benchmark")

![2nd protobuf_bench](https://raw.githubusercontent.com/everfore/gotest/master/protobuf/OSMsg2_test.png "2nd protobuf Benchmark")

## b2s

	bytes <===> string

![BS_bench](https://raw.githubusercontent.com/everfore/gotest/master/b2s/test.png "BS Benchmark")

## channel

	chan bool

	chan struct{}

	chan [1<<15]byte     channel element must be less than 64kb.

![chan_bench](https://raw.githubusercontent.com/everfore/gotest/master/channel/chan.png "chan Benchmark")

##	download

	download with 1kb buf.

##	echo

	echo a site: localhost:8080/site?site=http://localhost:8080

##	gbk

[golang.org/x/text/encoding/simplifiedchinese](http://gopm.io/download) 

##	websocket

	websocket c-s

##	mail

	send mail