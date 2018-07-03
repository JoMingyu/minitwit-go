package models

type User struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Pw       string `bson:"pw" json:"pw"`
}

type Follow struct {
	Follower User `bson:"follower"`
	Followee User `bson:"followee"`
}
