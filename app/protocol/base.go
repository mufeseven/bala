package protocol

import (
	"bala/app/protocol/errcode"
	"bala/app/protocol/protocode"
	"bala/app/util"
	"encoding/json"
)

type RequestBody struct {
	Protocol string `json:"protocol"`
	Packet   string `json:"packet"`
}

type ResponseBody struct {
	Protocol string `json:"protocol"`
	Packet   string `json:"packet"`
}

type ResponseError struct {
	Protocol  errcode.ErrorCode
	Reason    string
	ErrorCode int // WebAPIErrorCode
}

type SessionKey struct {
	AccountServerId int64
	MxToken         string
}

type BaseRequest struct {
	Protocol   protocode.ProtoCode
	SessionKey SessionKey
	RoleId     int `json:"accountId"`

	ClientUpTime int
	Resendable   bool
	Hash         int64
}

type BaseResponse struct {
	Protocol           protocode.ProtoCode
	ServerTimeTicks    int64
	ServerNotification int
}

func DecodeRequest(packet string, v interface{}) error {
	bytes := util.StringToBytes(packet)
	return json.Unmarshal(bytes, v)
}
