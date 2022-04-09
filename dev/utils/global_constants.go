package utils

const RESTAPIBASEURL = "/rest/cashtransfer/v1"
const PAYPALRESTAPIBASUEURL = "/rest/paypal/v1"
const MOMORESTAPIBASUEURL = "/rest/momo/v1"

const DBURI = "mongodb://0.0.0.0:27017"
const CASHTRANSFERDBNAME = "cashtransfer"
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
	PayPalID      string  `bson:"_id,omitempty" json:"paypalid"`
	PayPalName    string  `bson:"paypalname,omitempty" json:"paypalname"`
	PayPalBalance float32 `bson:"paypalbalance,omitempty" json:"paypalbalance"`
}

type MomoAccount struct {
	MomoID      string  `bson:"_id,omitempty" json:"momoid"`
	MomoName    string  `bson:"momoname,omitempty" json:"momoname"`
	MomoBalance float32 `bson:"momobalance,omitempty" json:"momobalance"`
}

type TransferInfo struct {
	DestUserID     string  `bson:"destuserid,omitempty" json:"destuserid"`
	TransferAmount float32 `bson:"amount,omitempty" json:"amount"`
}
