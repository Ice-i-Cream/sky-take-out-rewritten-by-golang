package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketServer WebSocket 服务结构体
type WebSocketServer struct {
	sessions   map[string]*websocket.Conn // 存储会话对象
	sessionMux sync.Mutex                 // 用于同步访问 sessions 的互斥锁
}

// NewWebSocketServer 创建新的 WebSocket 服务实例
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		sessions: make(map[string]*websocket.Conn),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源，生产环境应根据需要调整
	},
}

// ServeWebSocket 处理 WebSocket 连接
func (s *WebSocketServer) ServeWebSocket(c *gin.Context) {
	sid := c.Param("sid")
	if sid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sid parameter is required"})
		return
	}

	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Printf("客户端: %s 建立连接\n", sid)

	// 将连接添加到会话映射
	s.sessionMux.Lock()
	s.sessions[sid] = conn
	s.sessionMux.Unlock()

	// 移除断开连接的客户端
	defer func() {
		s.sessionMux.Lock()
		delete(s.sessions, sid)
		s.sessionMux.Unlock()
		log.Printf("连接断开: %s\n", sid)
	}()

	// 处理来自客户端的消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("收到来自客户端: %s 的信息: %s\n", sid, message)
	}
}

// SendToAllClients 向所有客户端广播消息
func (s *WebSocketServer) SendToAllClients(message string) {
	s.sessionMux.Lock()
	defer s.sessionMux.Unlock()

	for _, conn := range s.sessions {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Write error:", err)
			continue
		}
	}
}
