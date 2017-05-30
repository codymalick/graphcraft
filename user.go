package main

type User struct {
	ID int
	Username string
	Server string
}

func GetUser(username string, realm string) *User {

	user := QueryUserByName(username, realm)

	if user != nil {
		//pretty.Printf("Found cached user, %v,%v,%v\n",user.ID, user.Username, user.Server)
		return user
	} else {
		err := InsertUser(username, realm)
		checkErr(err)

		newUser := QueryUserByName(username, realm)

		return newUser
	}
}