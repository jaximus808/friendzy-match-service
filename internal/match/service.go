package match

// Add takes two integers and returns their sum.
import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//deprecate, maybe change to insertusers

//grab users based on date added, so from last 2 weeks, add those people

func createDBConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgress_url := os.Getenv("PQ_URL")
	db, err := sql.Open("postgres", postgress_url)
	if err != nil {
		return nil, fmt.Errorf("Error opening database")
	}
	defer db.Close()

	//keep later
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging")
	}
	fmt.Print("server ok!")
	return db, nil

}

func GetNewUsers(db *sql.DB) []User {

	users := []User{}

	rows, err := db.Query("SELECT id, survey_vector FROM users.user_survey")
	if err != nil {
		fmt.Print(err)
		log.Fatal("Error loading qeurerying")
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		var survey_vector string

		// Scan the row into variables
		err := rows.Scan(&user.user_id, &survey_vector)
		if err != nil {
			log.Println("Error scanning row: ", err)
			continue
		}

		user.trait_vector = parse_vector(survey_vector)

		// Process the data
		fmt.Printf("user of id: %s loaded", user.user_id)
		_ = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error after iterating rows: ", err)
	}
	return users
}

// grab users that want to be deleted

func GetDeletedUsers(db *sql.DB) {

}
