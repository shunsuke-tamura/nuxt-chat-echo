package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/olahol/melody.v1"

)

type Msg struct {
	Msg		string `json:"message"`
	UID 	string `json:"userId"`
	Time 	string `json:"time"`
}

const (
	layout		= "2006/01/02 15:04:05"
	timeZone 	= "Asia/Tokyo"
	dit 		= 9 * 60 * 60
)

var m *melody.Melody

// websoketのハンドラの初期化
func NewMelody() {
	m = melody.New()
}

// chatとクライアントとサーバをつなげる
func ChatWebsocket(ctx echo.Context) error {
	m.HandleRequest(ctx.Response().Writer, ctx.Request())
	return nil
}

// chatでデータが送られた時の処理
func chatHandler(s1 *melody.Session, msg []byte) {
	var messageObj Msg
	json.Unmarshal(msg, &messageObj)

	now := time.Now()
	jst := time.FixedZone(timeZone, dit)
	now = now.In(jst)
	messageObj.Time = now.Format(layout)

	msg, _ = json.Marshal(messageObj)
	fmt.Println(string(msg))
	m.BroadcastFilter(msg, func(s2 *melody.Session) bool {
		return s1.Request.URL.Path == s2.Request.URL.Path
	})
}

// websocketのハンドラの定義
func DefineMelodyBehavior() {
	m.HandleMessage(chatHandler)
}
