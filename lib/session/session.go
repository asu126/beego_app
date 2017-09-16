package mysession

import (
	"fmt"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"net/http"
)

var globalSessions *session.Manager

func InitSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:     "gosessionid",
		Gclifetime:     3600,
		ProviderConfig: "127.0.0.1:6379,100",
	}

	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

// 根据客户端传回来的session_id在服务器端校验是否成功登录
func GetSessionIdBySid(sid string) (string, error) {
	InitSession()

	sess, err := globalSessions.GetSessionStore(sid)
	fmt.Println("get username..........")
	fmt.Println(sess.Get("username"))
	return sess.SessionID(), err
}

// 首次登录，服务器端记录产生seesion，并记住用户信息
func SetSessionIdByUsername(w http.ResponseWriter, r *http.Request, username string) (string, error) {
	InitSession()

	//session
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	fmt.Println(username)
	err := sess.Set("username", username)
	fmt.Println(sess.SessionID())
	fmt.Println("set session success!!!")
	return sess.SessionID(), err
}
