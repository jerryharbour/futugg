package handlers

import (
	"fmt"

	"github.com/jerryharbour/futugg"
	"github.com/jerryharbour/futugg/pb/Qot_Common"
	"github.com/jerryharbour/futugg/pb/Qot_GetBasicQot"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func init() {
	futugg.SetHandlerId(uint32(3004), "Qot_GetBasicQot")
	var err error
	err = futugg.On("send.Qot_GetBasicQot", QotGetBasicQotSend)
	err = futugg.On("recv.Qot_GetBasicQot", QotGetBasicQotRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotGetBasicQotSend(conn *futugg.FutuGG, stockCode string) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(3004))

	var securityList []*Qot_Common.Security
	security := transStockCode(stockCode)
	securityList = append(securityList, security)

	reqData := &Qot_GetBasicQot.Request{
		C2S: &Qot_GetBasicQot.C2S{
			SecurityList: securityList,
		},
	}
	pbData, err := proto.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	pack.SetBody(pbData)
	err = conn.Send(pack)
	//fmt.Printf("send QotGetBasicQotSend for code %s\n", stockCode)

	return err
}

func QotGetBasicQotRecv(data []byte) ([]byte, error) {
	resp := &Qot_GetBasicQot.Response{}
	err := proto.Unmarshal(data, resp)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %s", err)
	}

	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(resp)
	//fmt.Println(result)
	return []byte(result), err
}
