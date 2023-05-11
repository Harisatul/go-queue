package sender

type Patient struct {
	Name             string `db:"name" json:"name"`
	QueueNumber      int    `db:"queue_number" json:"queueNumber"`
	IdentifierNumber string `db:"identifier_number" json:"identifierNumber"`
	IsPublish        bool   `db:"is_publish" json:"isPublish"`
}
