package utils

const RESTAPIBASEURL = "/rest/cashtransfer/v1"
const PAYPALRESTAPIBASUEURL = "/rest/paypal/v1"
const MOMORESTAPIBASUEURL = "/rest/momo/v1"

const MOMOENDPOINT = "http://localhost:8081" + MOMORESTAPIBASUEURL
const PAYPALENDPOINT = "http://localhost:8082/" + PAYPALRESTAPIBASUEURL

const DBURI = "mongodb://0.0.0.0:27017"
const CASHTRANSFERDBNAME = "cashtransfer"
const PAYPALDBNAME = "paypal"
const MOMODBNAME = "momo"
const USERSCOLLECTION = "users"
const ACCNTSCOLLECTION = "accounts"
const TRANSFERREQCOLLECTION = "requests"

type User struct {
	Username        string `bson:"_id,omitempty" json:"username"`
	Name            string `bson:"name,omitempty" json:"name"`
	Country         string `bson:"country,omitempty" json:"country"`
	Email           string `bson:"email,omitempty" json:"email"`
	PayPalID        string `bson:"paypalid,omitempty" json:"paypalid"`
	MomoID          string `bson:"momoid,omitempty" json:"momoid"`
	TransferMessage `bson:"message,omitempty" json:"message"`
}

type TransferRequest struct {
	Username  string  `bson:"_id,omitempty" json:"username"`
	Amount    float32 `bson:"amount,omitempty" json:"amount"`
	Completed bool    `bson:"completed,omitempty" json:"completed"`
}

type TransferMessage struct {
	SenderUsername string `bson:"senderusername,omitempty" json:"senderusername"`
	Request        string `bson:"request,omitempty" json:"request"`
	Response       bool   `bson:"response,omitempty" json:"response"`
}

type CashTransfer struct {
	Username string  `bson:"username,omitempty" json:"username"`
	Amount   float32 `bson:"amount,omitempty" json:"amount"`
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
