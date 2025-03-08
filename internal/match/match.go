package match

import (
	"log"
)

// this will be called in a cron job
func StartMatch() {

	//bruh i need another way of doing this

	layers_n := 10
	maxK := 5
	efLimit := 5
	mL := 3.0

	db_connect, err := createDBConnection()

	if err != nil {
		log.Fatalf("Couldn't connect to database")
		return
	}

	HSNW := create_hieracry(layers_n, maxK, efLimit, mL)

	new_users := GetNewUsers(db_connect)

	for _, user := range new_users {
		insert_user(HSNW, &user)
	}

}
