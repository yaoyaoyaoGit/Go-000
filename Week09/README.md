用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出

命令行 cheatsheet

netstat -ano | findstr "8080"， 查询端口占用 windows
https://blog.csdn.net/shmnh/article/details/12092699


参考资料
https://www.linode.com/docs/guides/developing-udp-and-tcp-clients-and-servers-in-go/

https://colobu.com/2014/12/02/go-socket-programming-TCP/

https://juejin.cn/post/6844903609138692110