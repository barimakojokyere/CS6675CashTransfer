package svrsrvc

import (
	"cashtransfer/dev/srvcs/database"
	"cashtransfer/dev/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user utils.User) {
	database.InsertIntoDB(utils.USERSCOLLECTION, user)
}

func RetrieveUser(username string) (user utils.User) {
	output := database.RetrieveFromDB(utils.USERSCOLLECTION, username)

	outputBytes, _ := bson.Marshal(output)
	bson.Unmarshal(outputBytes, &user)
	return user
}

func DeleteUser(username string) {
	database.RemoveFromDB(utils.USERSCOLLECTION, username)
}

func UpdateUser(username string, user utils.User) {
	database.UpdateInDB(utils.USERSCOLLECTION, username, user)
}
