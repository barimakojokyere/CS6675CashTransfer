package utils

const RESTAPIBASEURL = "/rest/cashtransfer/v1"
const PAYPALRESTAPIBASUEURL = "/rest/paypal/v1"
const MOMORESTAPIBASUEURL = "/rest/momo/v1"

const DBURI = "mongodb://0.0.0.0:27017"
const DBNAME = "cashtransfer"
const PAYPALDBNAME = "paypal"
const MOMODBNAME = "momo"
const USERSCOLLECTION = "users"
const ACCNTSCOLLECTION = "accounts"

type Accounts struct {
	Username string `bson:"_id,omitempty" json:"username"`
	ExternalAccounts
}

type User struct {
	Username string `bson:"_id,omitempty" json:"username"`
	Name     string `bson:"name,omitempty" json:"name"`
	Country  string `bson:"country,omitempty" json:"country"`
	Email    string `bson:"email,omitempty" json:"email"`
}

type ExternalAccounts struct {
	PayPalAccount
	MomoAccount
}

type PayPalAccount struct {
	PayPalUsrName string  `bson:"paypalusrname,omitempty" json:"paypalusrname"`
	PayPalBalance float32 `bson:"paypalbalance,omitempty" json:"paypalbalance"`
}

type MomoAccount struct {
	MomoUsrName string  `bson:"momousrname,omitempty" json:"momousrname"`
	MomoBalance float32 `bson:"momobalance,omitempty" json:"momobalance"`
}
