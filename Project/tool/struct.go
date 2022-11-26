package tool

import "Project/model"

func CreateUser(username string, password string) model.User {
	var U model.User
	U.Username = username
	U.Password = password
	return U
}

func CreateMess(ownerName string, writeName string, message string, respond string) model.Mess {
	var M model.Mess
	M.OwnerName = ownerName
	M.WriteName = writeName
	M.Message = message
	M.Respond = respond
	return M
}

func CreateQues(userName string, trueName string, likeFood string, age int) model.Ques {
	var Q model.Ques
	Q.UserName = userName
	Q.TrueName = trueName
	Q.LikeFood = likeFood
	Q.Age = age
	return Q
}
