package models

type PrInfo struct {
	ID int `bson:"id"`

	PrNum      int    `bson:"pr_num"`
	Title      string `bson:"title"`
	NodeId     string `bson:"node_id"`
	UserLogin  string `bson:"user_login"`
	UserNodeID string `bson:"user_node_id"`
	Folder     string `bson:"folder"`
	PrType     string `bson:"pr_type"`

	LastCommitHash string `bson:"last_commit_hash"`
	Status         string `bson:"status"` //Draft, Open, Closed
}

type TitleInfo struct {
	PrType  string
	Folder  string
	Version string
}
