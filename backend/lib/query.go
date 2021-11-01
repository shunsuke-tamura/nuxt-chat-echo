package lib

type ChatContent struct {
	UserId  string `json:"userId"`
	RoomId  string `json:"roomId"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type ChatContents []ChatContent

func InsertUser(userId string) error {
	query := `INSERT INTO users (uid) VALUES ($1)`
	err := Conn.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func InsertRoom(roomId string) error {
	query := `INSERT INTO room (rid) VALUES ($1)`
	err := Conn.Exec(query, roomId)
	if err != nil {
		return err
	}
	return nil
}

func InsertChatContents(message, time, userId, roomId string) error {
	query := `INSERT INTO chat (
		uid,
		rid,
		message,
		time
	) values (
		$1,
		$2,
		$3,
		$4
	)`
	err := Conn.Exec(query, userId, roomId, message, time)
	if err != nil {
		return err
	}
	return nil
}

func FetchChatContents(roomId string) (ChatContents, error) {
	query := `select uid, rid, message, time from chat where rid=$1`
	rows, err := Conn.GetRow(query, roomId)
	var chatContents ChatContents
	for _, row := range rows {
		chatContents = append(chatContents, ChatContent{
			UserId:  row[0].(string),
			RoomId:  row[1].(string),
			Message: row[2].(string),
			Time:    row[3].(string),
		})
	}
	return chatContents, err
}