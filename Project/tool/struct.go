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

func CreateComment(author string, writer string, comment string) model.Comment {
	var C model.Comment
	C.Author = author
	C.Writer = writer
	C.Comment = comment
	return C
}

func CreateMessage(username string, age int, birthday float64, constellation string, sex string) model.Message {
	var C model.Message
	C.Username = username
	C.Age = age
	C.Birthday = birthday
	C.Constellation = constellation
	C.Sex = sex
	return C
}

func CreatePraise(author string, reader string, good bool) model.Praise {
	var P model.Praise
	P.Author = author
	P.Reader = reader
	P.Good = good
	return P
}

func CreateGood(author string, member int) model.Good {
	var G model.Good
	G.Author = author
	G.Member = member
	return G
}
