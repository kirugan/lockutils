package main

import (
	"lockutils"
	"fmt"
	"time"
)

func main() {
	// reader not block reader
	echo ("=== RR ===")
	testRR()
	// reader block writer
	echo ("=== RW ===")
	testRW()
	// writer block writer
	echo ("=== WW ===")
	testWW()
	// writer block reader
	echo ("=== WR ===")
	testWR()
}

/* TEST CASES */
func testWR() {
	path := "/var/www/haha"
	mainEcho := factoryEcho("main")
	m1 := lockutils.NewMap()

	mainEcho("Lock")
	m1.Lock(path)


	go func () {
		goEcho := factoryEcho("go")

		goEcho("RLock")
		m1.RLock(path)

		goEcho("RUnlock")
		m1.RUnlock(path)
	}();

	mainEcho("Unlock")
	m1.Unlock(path)
	sleep(1)
}

func testWW() {
	path := "/var/www/haha"
	mainEcho := factoryEcho("main")
	m1 := lockutils.NewMap()

	mainEcho("Lock")
	m1.Lock(path)


	go func () {
		goEcho := factoryEcho("go")

		goEcho("Lock")
		m1.Lock(path)

		goEcho("Unlock")
		m1.Unlock(path)
	}();

	sleep(1)
	mainEcho("Unlock")
	m1.Unlock(path)
	sleep(1)
}

func testRW() {
	path := "/var/www/haha"
	mainEcho := factoryEcho("main")
	m1 := lockutils.NewMap()

	mainEcho("RLock")
	m1.RLock(path)


	go func () {
		goEcho := factoryEcho("go")

		goEcho("Lock")
		m1.Lock(path)

		goEcho("Unlock")
		m1.Unlock(path)
	}();

	sleep(1)
	mainEcho("RUnlock")
	m1.RUnlock(path)
	sleep(2)
}

func testRR() {
	path := "/var/www/haha"
	mainEcho := factoryEcho("main")
	m1 := lockutils.NewMap()

	mainEcho("RLock")
	m1.RLock(path)


	go func () {
		goEcho := factoryEcho("go")

		goEcho("RLock")
		m1.RLock(path)

		goEcho("RUnlock")
		m1.RUnlock(path)
	}();

	sleep(2)
	mainEcho("RUnlock")
	m1.RUnlock(path)
}

/* UTILS */
func factoryEcho(prefix string) func (string) {
	return func (msg string) {
		echo(prefix + ": " + msg)
	};
}
func echo(msg string) {
	fmt.Println(msg)
}

func sleep(num int) {
	time.Sleep(time.Second * time.Duration(num))
}