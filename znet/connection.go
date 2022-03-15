package znet

import (
	"net"
	"zinx/ziface"
)

type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前的链接状态
	isClosed bool

	// 当前链接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc

	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool
}
