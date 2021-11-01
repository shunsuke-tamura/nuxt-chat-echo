package internal

import (
	"fmt"
	"encoding/json"
	"net/http"
	"nuxt-echo-chat/backend/lib"
	"time"

	"github.com/labstack/echo/v4"
	"gopkg.in/olahol/melody.v1"
)

type Msg struct {
	Msg    string `json:"message"`
	UID    string `json:"userId"`
	Time   string `json:"time"`
	RoomId string `json:"roomId"`
}

const (
	layout   = "2006/01/02 15:04:05"
	timeZone = "Asia/Tokyo"
	dit      = 9 * 60 * 60
)

func FetchChat(ctx echo.Context) error {
	roomId := ctx.Param("roomid")
	chatContents, err := lib.FetchChatContents(roomId)
	if err != nil {
		must(err)
		return err
	}
	return ctx.JSON(http.StatusOK, chatContents)
}

var m *melody.Melody

func NewMelody() {
	m = melody.New()
}

func ChatWebsocket(ctx echo.Context) error {
	m.HandleRequest(ctx.Response().Writer, ctx.Request())
	return nil
}

func chatHandler(s1 *melody.Session, msg []byte) {
	var messageObj Msg
	json.Unmarshal(msg, &messageObj)

	now := time.Now()
	jst := time.FixedZone(timeZone, dit)
	now = now.In(jst)
	messageObj.Time = now.Format(layout)

	lib.InsertUser(messageObj.UID)
	lib.InsertRoom(messageObj.RoomId)
	lib.InsertChatContents(messageObj.Msg, messageObj.Time, messageObj.UID, messageObj.RoomId)

	msg, _ = json.Marshal(messageObj)
	m.BroadcastFilter(msg, func(s2 *melody.Session) bool {
		return s1.Request.URL.Path == s2.Request.URL.Path
	})
}

func DefineMelodyBehavior() {
	m.HandleMessage(chatHandler)
}

func must(err error) {
	fmt.Println(err)
}
