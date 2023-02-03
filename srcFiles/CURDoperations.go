package user_funcs

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// Question structure
type Question struct {
	Questiontitle string
	Options       []string
	CorrectAns    string
}

func ConnectDb() *sql.DB {
	connStr := "user=kamlesh-kumar password=kamlesh host=localhost dbname= prc_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("Failed connection")
		panic(err)
	}
	return db
}

// CreateTable creates the table "questions" in the database
func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS questions (
		question TEXT NOT NULL,
		options TEXT[] NOT NULL,
		correct_answer TEXT NOT NULL
	)`)
	return err
}

// InsertData inserts a question into the table "questions"
func InsertData(db *sql.DB, question Question) error {
	_, err := db.Exec(`INSERT INTO questions (question, options, correct_answer)
		VALUES ($1, $2, $3)`, question.Questiontitle, pq.StringArray(question.Options), question.CorrectAns)
	return err
}

// ReadData reads the questions from the table "questions"
func ReadData(db *sql.DB) ([]Question, error) {
	rows, err := db.Query(`SELECT question, options, correct_answer FROM questions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quest []Question
	for rows.Next() {
		var q Question

		err := rows.Scan(&q.Questiontitle, pq.Array(&q.Options), &q.CorrectAns)
		if err != nil {
			return nil, err
		}
		quest = append(quest, q)
	}
	return quest, nil
}

// UpdateData updates the data for a question in the table "questions"
func UpdateData(db *sql.DB, question Question, previousQuestion string) error {
	_, err := db.Exec(`UPDATE questions
		SET question=$1, options = $2, correct_answer = $3
		WHERE question = $4`, question.Questiontitle, pq.Array(question.Options), question.CorrectAns, previousQuestion)
	return err
}

// DeleteData deletes a question from the table "questions"
func DeleteData(db *sql.DB, question string) error {
	_, err := db.Exec(`DELETE FROM questions WHERE question = $1`, question)
	return err
}
