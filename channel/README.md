# Go chan 的使用方法总结

## 无缓冲？

无缓冲channel,需要异步(使用goruntine)接/收消息,否则会死锁。

		c := make(chan bool)

####	deadlock

		c <- true
		rec, ok := <-c
		// deadlock

####	go func send

		go func(){
			c <- true
		}()
		rec,ok := <- c
		t.Log(rec, ok)
		// true, true

####	go func rec

		go func() {
			select {
			case rec, ok := <-c:
				t.Log(rec, ok)
			}
		}()
		c <- true

## 缓冲

带缓冲chan,适当的接收消息可以避免死锁.

		c := make(chan bool, 1)

		c <- true
		t.Log(<-c)
		// true


		c <- true
		c <- true
		t.Log(<-c)
		// deadlock

## 带缓冲的chan是异步的，不带缓冲的chan是同步的。

### 	主动关闭

关闭后,chan只可收消息,不能发消息(发消息会panic).
		
		close(c)
		c <- true
		// panic: send on closed channel

关闭chan,消息全部被取走后,试图继续接收,结果中第2个参数ok为false,此时应该结束接收.

异步接收,没有发消息,chan被关闭：

		go func() {
			select {
			case rec, ok := <-c:
				t.Log(rec, ok)
			}
		}()
		close(c)
		time.Sleep(1e9)
		// false, false

异步接收,发了消息,chan被关闭：

		go func() {
			select {
			case rec, ok := <-c:
				t.Log(rec, ok)
			}
		}()
		c <- true
		close(c)
		rec, ok := <-c
		t.Log(rec, ok)
		// true, true
		// false, false

异步、同步接收消息：

		c := make(chan bool, 2)
		go func() {
			select {
			case rec, ok := <-c:
				t.Log(rec, ok)
			}
		}()
		c <- true
		c <- true
		close(c)
		rec, ok := <-c
		t.Log(rec, ok)
		time.Sleep(1e9)
		// true, true
		// true, true

异步发送消息:

		c := make(chan bool, 2)
		go func() {
			select {
			case c <- false:
				c <- true
			}
		}()
		rec, ok := <-c
		t.Log(rec, ok)
		time.Sleep(1e9)
		close(c)
		rec, ok = <-c
		t.Log(rec, ok)
		// false, true
		// true, true

## select


#### 超时

select,chan组合

		func Foo() {
			time.Sleep(1e9)
			// time.Sleep(3e9)
		}

		func TestT(t *testing.T) {

			c := make(chan bool)
			go func() {
				Foo()
				c <- true
			}()
			select {
			case <-time.After(2e9):
				t.Log("timeout")
			case <-c:
				t.Log("Foo done")
			}
		}

#### 任务完成notify

		func Bar() chan bool {
			c := make(chan bool)
			defer func() {
				c <- true
			}()
			// TODO

			return c
		}

#### 接收多个chan

for,select,chan组合

		for{
			select{
				case <-c1:
					...
				case <-c2:
					...
			}
		}