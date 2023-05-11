package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"go-queue/pkg/event"
	"os"
	"strings"
)

func CreateEventEmitter() (event.Emitter, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return event.Emitter{}, err
	}

	emitter, err := event.NewEventEmitter(conn)
	if err != nil {
		return emitter, err
	}

	return emitter, nil
}

var schema = `
CREATE TABLE patients (
    name text,
    queue_number integer,
    identifier_number text,
    is_publish boolean
);
`

func NewDB() (*sqlx.DB, error) {
	// Configure the database connection
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Read the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Construct the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the da	tabase: %w", err)
	}

	err = initializeDatabaseSchema(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the database schema: %w", err)
	}

	err = executeInsertStatements(db)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert statements: %w", err)
	}
	return db, nil
}

func initializeDatabaseSchema(db *sqlx.DB) error {
	db.MustExec("DROP TABLE IF EXISTS patients")
	db.MustExec(schema)
	return nil
}

func executeInsertStatements(db *sqlx.DB) error {
	insertStatements := `INSERT INTO patients (name, queue_number, identifier_number, is_publish) VALUES
		('Haris', 1, '41251241241', false),
		('Budi', 2, '31421241242', false),
		('Joko', 3, '312512241241', false),
		('Tiara', 4, '11251321241', false),
		('Annisa', 5, '231253241241', false),
		('Morphius', 6, '61251341241', false),
		('Ani', 7, '31251241241', false),
		('Mamat', 8, '2351241241', false),
		('Rusdi Kirana', 9, '234251241241', false),
		('Joko', 10, '41252341241', false),
		('Anwar', 11, '412512412324', false),
		('Hussein', 12, '4125124136', false),
		('Rudi', 13, '12325124241', false),
		('Philips', 14, '11251241241', false),
		('Horrison', 15, '32251241241', false);`

	statements := strings.Split(insertStatements, ";")
	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement != "" {
			db.MustExec(statement)
		}
	}

	return nil
}
