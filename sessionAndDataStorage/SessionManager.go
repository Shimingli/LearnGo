package main

//import (
//	"sync"
//	"fmt"
//	"encoding/base64"
//	"crypto/rand"
//)
//
//func main() {
//
//}
//// 全局 唯一的 Session ID
//func  (m *Manager) sessionId() string  {
//	//make用于内建类型（map、slice 和channel）的内存分配 长度为32
//	b:=make([]byte,32)
//	fmt.Println(" old-----b=",b)
//	//这段代码对下面的结果有影响
//	if   _,error:=rand.Read(b) ;error!=nil {
//		return ""
//	}
//	//old-----b= [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
//	//new-----b= [82 253 252 7 33 130 101 79 22 63 95 15 154 98 29 114 149 102 199 77 16 3 124 77 123 187 4 7 209 226 198 73]
//	fmt.Println(" new-----b=",b)
//	fmt.Println("base64.URLEncoding.EncodeToString(b)=",base64.URLEncoding.EncodeToString(b))
//	//返回b字节数组的BASE64编码。
//	return base64.URLEncoding.EncodeToString(b)
//}
//
//
//func NewManager(provideName, cookieName string, maxLifeTime int64)(*Manager,error) {
//	// map有两个返回值，第一个返回值是指的是这个key对应的value的值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
//	provider,ok:=provides[provideName]
//	if !ok {
//		return nil,fmt.Errorf("不存在小兄弟")
//	}
//	return &Manager{cookieName:cookieName,provider:provider,maxLifeTime:maxLifeTime},nil
//
//}
//
//
//
//// Sessiong的处理基本就是设置值，读取值，删除值，以及获取当前的值
//type Session interface {
//	Set(key, value interface{}) error
//	Get(key interface{}) interface{}
//	Delete(key interface{}) error
//	SessionID() string
//
//}
//
////session 是保存在服务器端的数据，它可以以任何方式储存，比如储存在内存、数据库或者是文件中。因此抽象一个Provider接口，用来表示session管理器底层储存的结构
//type Provider interface {
//	//SessionInit 函数实现了 Session的初始化，操作成功则返回此新的Session变量
//	SessionInit(sid string)(Session,error)
//	//SessionRead函数返回sid所表示的Session变量，如果不存在，那么将以sid为参数调用SessionInit函数创建并返回一个新的Session变量
//	SessionRead(sid string) (Session,error)
//	// 用来摧毁sid对应的Session的变量
//	SessionDestroy(sid  string)error
//	//根据maxLifeTime来删除过期的数据
//	SessionGC(maxLifeTime int64)
//}
//
//
//type Manager struct {
//	cookieName string
//	//互斥锁是互斥锁。
//	//互斥体的零值是一个未锁定的互斥体。
//	//互斥体在首次使用后不得复制。
//	lock sync.Mutex
//	provider Provider
//	maxLifeTime  int64
//}


//操作值：设置、读取和删除
//SessionStart函数返回的是一个满足Session接口的变量，那么我们该如何用他来对session数据进行操作呢？
//
//上面的例子中的代码session.Get("uid")已经展示了基本的读取数据的操作，现在我们再来看一下详细的操作:
//
//func count(w http.ResponseWriter, r *http.Request) {
//	sess := globalSessions.SessionStart(w, r)
//	createtime := sess.Get("createtime")
//	if createtime == nil {
//		sess.Set("createtime", time.Now().Unix())
//	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
//		globalSessions.SessionDestroy(w, r)
//		sess = globalSessions.SessionStart(w, r)
//	}
//	ct := sess.Get("countnum")
//	if ct == nil {
//		sess.Set("countnum", 1)
//	} else {
//		sess.Set("countnum", (ct.(int) + 1))
//	}
//	t, _ := template.ParseFiles("count.gtpl")
//	w.Header().Set("Content-Type", "text/html")
//	t.Execute(w, sess.Get("countnum"))
//}

//通过上面的例子可以看到，Session的操作和操作key/value数据库类似:Set、Get、Delete等操作
//
//因为Session有过期的概念，所以我们定义了GC操作，当访问过期时间满足GC的触发条件后将会引起GC，但是当我们进行了任意一个session操作，都会对Session实体进行更新，都会触发对最后访问时间的修改，这样当GC的时候就不会误删除还在使用的Session实体。
//
//session重置
//我们知道，Web应用中有用户退出这个操作，那么当用户退出应用的时候，我们需要对该用户的session数据进行销毁操作，上面的代码已经演示了如何使用session重置操作，下面这个函数就是实现了这个功能：
//
////Destroy sessionid
//func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request){
//	cookie, err := r.Cookie(manager.cookieName)
//	if err != nil || cookie.Value == "" {
//		return
//	} else {
//		manager.lock.Lock()
//		defer manager.lock.Unlock()
//		manager.provider.SessionDestroy(cookie.Value)
//		expiration := time.Now()
//		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
//		http.SetCookie(w, &cookie)
//	}
//}

//session销毁
//我们来看一下Session管理器如何来管理销毁，只要我们在Main启动的时候启动：
//
//func init() {
//	go globalSessions.GC()
//}
//func (manager *Manager) GC() {
//	manager.lock.Lock()
//	defer manager.lock.Unlock()
//	manager.provider.SessionGC(manager.maxLifeTime)
//	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
//}