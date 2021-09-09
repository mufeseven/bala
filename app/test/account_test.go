package test

import (
	"bala/app/mid"
	"bala/app/protocol"
	"bala/app/protocol/protocode"
	"fmt"
	"testing"
)

func TestAccountCreate(t *testing.T) {
	/*var wg sync.WaitGroup
	for i := 1; i <= 1; i++ {
		wg.Add(1)
		go func(idx int) {
			req := &protocol.AccountCreateRequest{}
			req.Protocol = protocode.Account_Create
			req.SessionKey.AccountServerId = 1
			req.SessionKey.MxToken = "token"
			req.DevId = fmt.Sprintf("test_%04d", idx)
			body, _ := mid.EncodeRequestBody(req.Protocol, req)
			resp := PostForm("/api/account/create", body)
			t.Log(resp)
			wg.Done()
		}(i)
	}
	wg.Wait()*/
	req := &protocol.AccountCreateRequest{}
	req.Protocol = protocode.Account_Create
	req.DevId = fmt.Sprintf("test_%04d", 1)
	body, _ := mid.EncodeRequestBody(req.Protocol, req)
	resp := PostForm("/api/account/create", body)
	t.Log(resp)
}
