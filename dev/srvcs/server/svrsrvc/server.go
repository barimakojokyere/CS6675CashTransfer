package svrsrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user utils.User) (err error) {
	err = database.InsertIntoDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, user)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveUser(username string) (user utils.User, err error) {
	output, err := database.RetrieveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return user, err
	}

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &user)
	return user, nil
}

func DeleteUser(username string) (err error) {
	err = database.RemoveFromDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(username string, user utils.User) (err error) {
	err = database.UpdateInDB(utils.CASHTRANSFERDBNAME, utils.USERSCOLLECTION, username, user)
	if err != nil {
		return err
	}

	return nil
}
