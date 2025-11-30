package storage

import (
	"context"
	"log"
)

func Migration(){

	conn, err := DBConnection()
	if err != nil {
		log.Fatalf(err)
	}
	defer conn.Close(context.Background)
	
	err := conn.Ping(context.Background()); err != nil {
		log.Fatalf(err)
	}

	createEnrollmentsTable := `
		Create Table If Not Exists Enrollments(
			id BIGSERIAL Primary Key,
			token varchar(100),
			accountType varchar(100)
			institution varchar(50),
			enrollmentId varchar(100),
			status bool,
			accountName varchar (50),
	)`

}
