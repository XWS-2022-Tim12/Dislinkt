package persistence

import (
	"github.com/XWS-2022-Tim12/Dislinkt/back/user_suggestions_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Register(session neo4j.Session, user *domain.User) (int64, error) {
	var userId int64
	session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"CREATE (user:USER {username: $username, firstName: $firstName, email: $email, interests: $interests})"+
				"RETURN ID(user), user.username",
			map[string]interface{}{
				"username": user.Username, "firstName": user.FirstName,
				"email": user.Email, "interests": user.Interests})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			userId = result.Record().Values[0].(int64)

		}
		return userId, nil
	})
	return userId, nil
}

func GetAll(session neo4j.Session) (users []*domain.User, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (user:USER) RETURN ID(user), user.username, user.firstName, user.email, user.interests", map[string]interface{}{})

		for records.Next() {
			user := domain.User{
				Id:        records.Record().Values[0].(int64),
				Username:  records.Record().Values[1].(string),
				FirstName: records.Record().Values[2].(string),
				Email:     records.Record().Values[3].(string),
				Interests: records.Record().Values[4].(string),
			}
			users = append(users, &user)
		}

		if err != nil {
			return nil, err
		}

		return users, nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil

}
