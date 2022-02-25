package main

import (
	"log"
	"os"
	"syscall"
	"time"
)

func main() {

	f, err := os.OpenFile("mmap.bin", os.O_RDWR|os.O_CREATE, 0644)

	if nil != err {

		log.Fatalln(err)

	}

	// extend file

	//if _, err := f.WriteAt([]byte{byte(0)}, 10); nil != err {

	//    log.Fatalln(err)

	//}
	err = syscall.Ftruncate(int(f.Fd()), 1000) // 文件读取的最小的单位为“一页”， 一页的大小一般为4k
	if err != nil {
		panic(err)
	}

	data, err := syscall.Mmap(int(f.Fd()), 0, 1<<12, syscall.PROT_WRITE, syscall.MAP_SHARED)

	if nil != err {

		log.Fatalln(err)

	}

	if err := f.Close(); nil != err {

		log.Fatalln(err)

	}

	//for i, v := range []byte("a\n") {
	//  data[i+4094] = v // 这里4096会报错，4095 就没问题

	//}
	for i := 0; i < 100; i++ {

		log.Println(string(data))

		for i, v := range []byte("i am from another process") {

			data[i] = v

		}

		time.Sleep(time.Second * 3)
	}
	if err := syscall.Munmap(data); nil != err {

		log.Fatalln(err)

	}

}
