package models

import "time"

type Message struct {
	Sender    string    `bson:"sender"`    // AccountSID ou BotKit
	Text      string    `bson:"text"`      // Conteúdo da mansagem
	Timestamp time.Time `bson:"timestamp"` // Hora da mensagem
}

type Conversation struct {
	ID        string     `bson:"_id,omitempty"`
	UserID    string     `bson:"user_id"`
	StartTime time.Time  `bson:"start_time"`
	EndTime   *time.Time `bson:"end_time,omitempty"` // ponteiro para permitir ser nil
	Messages  []Message  `bson:"messages"`
}
