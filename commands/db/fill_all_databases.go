package main

import "log"

func main() {
	log.Println("Filling all databases...")

	if err := user.SeedUsers(db); err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}

}
