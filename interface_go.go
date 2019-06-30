package main

type Actor interface {
	OnConnect() error
	OnMessage(data []byte) error
	OnError(err error)
	OnClose()
	GetMsgChan() <-chan []byte
	Done() <-chan struct{}
}

type TCPCtrl struct {
	Actor
}

func (t *TCPCtrl) Scan(){
	defer t.OnClose()
	err := t.OnConnect()
	if err != nil{
		t.OnError(err)
	}
Circle: 
		for{
			select{
			case <-t.Done():
				break Circle
			case msg := <- t.GetMsgChan():
					err := t.OnMessage(msg)
					if err != nil{
							t.OnError(err)
					}
			}
		}
}

func main(){
	//然后如果你要使用这个库，可以随便定义自己的ActorType，只需要保证你的 Actor 有四个回调函数，一个 MsgChannel，和一个Done（使用Context），然后就可
	ctrl := &TCPCtrl{NewActorType()}
	go ctrl.Scan()
}