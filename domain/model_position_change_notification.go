package domain

type PositionChangeNotification struct {
	Title    string    `json:"title"`
	Position *Position `json:"position"`
}

func PositionChangeNotificationFromTransport(transport *Transport) *PositionChangeNotification {
	return &PositionChangeNotification{Title: transport.ID, Position: transport.CurrentPosition}
}
