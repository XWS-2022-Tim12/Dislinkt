package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetClient(username, password, uri string) (*neo4j.Session, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	if err != nil {
		fmt.Println("connection established")
	}

	return &session, err
}
