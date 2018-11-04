package web

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/hawkwithwind/chat-bot-hub/proto/chatbothub"
	"golang.org/x/net/context"
	grctx "github.com/gorilla/context"
	"google.golang.org/grpc"

	"github.com/hawkwithwind/chat-bot-hub/server/domains"
)

type GRPCWrapper struct {
	conn    *grpc.ClientConn
	client  pb.ChatBotHubClient
	context context.Context
	cancel  context.CancelFunc
}

func (w *GRPCWrapper) Cancel() {
	if w == nil {
		return
	}

	if w.cancel != nil {
		w.cancel()
	}

	if w.conn != nil {
		w.conn.Close()
	}
}

func (ctx *ErrorHandler) GRPCConnect(target string) *GRPCWrapper {
	if ctx.Err != nil {
		return nil
	}

	if conn, err := grpc.Dial(target, grpc.WithInsecure()); err != nil {
		ctx.Err = err
		return nil
	} else {
		client := pb.NewChatBotHubClient(conn)
		gctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		return &GRPCWrapper{
			conn:    conn,
			client:  client,
			context: gctx,
			cancel:  cancel,
		}
	}
}

func (ctx *ErrorHandler) GetBots(w *GRPCWrapper, req *pb.BotsRequest) *pb.BotsReply {
	if ctx.Err != nil {
		return nil
	}

	if botsreply, err := w.client.GetBots(w.context, &pb.BotsRequest{Secret: "secret"}); err != nil {
		ctx.Err = err
		return nil
	} else {
		return botsreply
	}
}

func (ctx *ErrorHandler) LoginQQ(w *GRPCWrapper, req *pb.LoginQQRequest) *pb.LoginQQReply {
	if ctx.Err != nil {
		return nil
	}

	if loginreply, err := w.client.LoginQQ(w.context, req); err != nil {
		ctx.Err = err
		return nil
	} else {
		return loginreply
	}
}

func (ctx *ErrorHandler) LoginWechat(w *GRPCWrapper, req *pb.LoginWechatRequest) *pb.LoginWechatReply {
	if ctx.Err != nil {
		return nil
	}

	if loginreply, err := w.client.LoginWechat(w.context, req); err != nil {
		ctx.Err = err
		return nil
	} else {
		return loginreply
	}
}

func findDevice(bots []domains.Bot, login string) *domains.Bot {
	for _, bot := range bots {
		if bot.Login == login {
			return &bot
		}
	}

	return nil
}

func (ctx *WebServer) getBots(w http.ResponseWriter, r *http.Request) {
	type BotsInfo struct {
		pb.BotsInfo
		BotName string `json:"botName"`
		CreateAt int64 `json:"createAt"`
	}
	
	o := ErrorHandler{}
	defer o.WebError(w)

	var login string
	if loginptr := grctx.Get(r, "login"); loginptr != nil {
		login = loginptr.(string)
	}
	
	bots := o.GetBotsByAccountName(ctx.db.Conn, login)
	if o.Err == nil && len(bots) == 0 {
		o.ok(w, "", []BotsInfo{})
		return
	}
	
	wrapper := o.GRPCConnect(fmt.Sprintf("127.0.0.1:%s", ctx.Hubport))
	defer wrapper.Cancel()	
	
	bs := []BotsInfo{}
	
	if botsreply := o.GetBots(wrapper, &pb.BotsRequest{Secret: "secret"}); botsreply != nil {
		for _, info := range botsreply.BotsInfo {
			if b := findDevice(bots, info.Login); b != nil {
				bs = append(bs, BotsInfo{
					BotsInfo: *info,
					BotName: b.BotName,
					CreateAt: b.CreateAt.Time.Unix(),
				})
			}
		}
	}
	
	o.ok(w, "", bs)
}

func (ctx *WebServer) loginQQ(w http.ResponseWriter, r *http.Request) {
	o := ErrorHandler{}
	defer o.WebError(w)

	r.ParseForm()
	qqnumstr := o.getStringValue(r.Form, "qqnum")
	pass := o.getStringValue(r.Form, "password")
	clientId := o.getStringValue(r.Form, "clientId")
	qqnum := o.ParseUint(qqnumstr, 10, 64)

	wrapper := o.GRPCConnect(fmt.Sprintf("127.0.0.1:%s", ctx.Hubport))
	defer wrapper.Cancel()

	loginreply := o.LoginQQ(wrapper, &pb.LoginQQRequest{
		ClientId: clientId,
		QQNum:    qqnum,
		Password: pass,
	})
	o.ok(w, "", loginreply)
}

func (ctx *WebServer) loginWechat(w http.ResponseWriter, r *http.Request) {
	o := ErrorHandler{}
	defer o.WebError(w)

	r.ParseForm()
	clientId := o.getStringValue(r.Form, "clientId")

	wrapper := o.GRPCConnect(fmt.Sprintf("127.0.0.1:%s", ctx.Hubport))
	defer wrapper.Cancel()

	loginreply := o.LoginWechat(wrapper, &pb.LoginWechatRequest{
		ClientId: clientId,
	})
	o.ok(w, "", loginreply)
}

func (ctx *WebServer) getConsts(w http.ResponseWriter, r *http.Request) {
	o := ErrorHandler{}
	defer o.WebError(w)

	o.ok(w, "", map[string]interface{}{
		"types": map[string]string{
			"QQBOT":     "QQ机器人",
			"WECHATBOT": "微信机器人",
		},
		"status": map[int]string{
			1:   "初始化",
			100: "准备登录",
			150: "等待扫码",
			151: "登录失败",
			200: "已登录",
			500: "连接断开",
		},
	})
}
