package web

import (
	"fmt"
	"net/http"
	"time"

	pb "github.com/hawkwithwind/chat-bot-hub/proto/chatbothub"
	"golang.org/x/net/context"
	grctx "github.com/gorilla/context"
	"google.golang.org/grpc"

	//"github.com/hawkwithwind/chat-bot-hub/server/domains"
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

func (ctx *ErrorHandler) LoginBot(w *GRPCWrapper, req *pb.LoginBotRequest) *pb.LoginBotReply {
	if ctx.Err != nil {
		return nil
	}

	if loginreply, err := w.client.LoginBot(w.context, req); err != nil {
		ctx.Err = err
		return nil
	} else {
		return loginreply
	}
}

func findDevice(bots []*pb.BotsInfo, login string) *pb.BotsInfo {
	for _, bot := range bots {
		if bot.Login == login {
			return bot
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
	
	wrapper := o.GRPCConnect(fmt.Sprintf("%s:%s", ctx.Hubhost, ctx.Hubport))
	defer wrapper.Cancel()	
	
	bs := []BotsInfo{}
	
	if botsreply := o.GetBots(wrapper, &pb.BotsRequest{Secret: "secret"}); botsreply != nil {
		for _, b := range bots {
			if info := findDevice(botsreply.BotsInfo, b.Login); info != nil {
				bs = append(bs, BotsInfo{
					BotsInfo: *info,
					BotName: b.BotName,
					CreateAt: b.CreateAt.Time.Unix(),
				})
			} else {
				bs = append(bs, BotsInfo{
					BotsInfo: pb.BotsInfo{
						ClientType: b.ChatbotType,
						Status: 0,
					},
					BotName: b.BotName,
					CreateAt: b.CreateAt.Time.Unix(),
				})
			}
		}
	} else {
		if o.Err == nil {
			o.Err = fmt.Errorf("grpc botsreply is null")
		}
	}
	
	o.ok(w, "", bs)
}

func (ctx *WebServer) loginBot(w http.ResponseWriter, r *http.Request) {
	o := ErrorHandler{}
	defer o.WebError(w)

	r.ParseForm()
	clientId := o.getStringValueDefault(r.Form, "clientId", "")
	clientType := o.getStringValue(r.Form, "clientType")
	login := o.getStringValueDefault(r.Form, "login", "")
	pass := o.getStringValueDefault(r.Form, "password", "")

	wrapper := o.GRPCConnect(fmt.Sprintf("%s:%s", ctx.Hubhost, ctx.Hubport))
	defer wrapper.Cancel()

	loginreply := o.LoginBot(wrapper, &pb.LoginBotRequest{
		ClientId: clientId,
		ClientType: clientType,
		Login: login,
		Password: pass,
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
			0:   "未连接",
			1:   "初始化",
			100: "准备登录",
			150: "等待扫码",
			151: "登录失败",
			200: "已登录",
			500: "连接断开",
		},
	})
}
