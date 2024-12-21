package handlers

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jerryharbour/futugg"
	"github.com/jerryharbour/futugg/pb/Qot_GetWarrant"
)

func init() {
	futugg.SetHandlerId(uint32(3210), "Qot_GetWarrant")
	var err error
	err = futugg.On("send.Qot_GetWarrant", QotGetWarrantSend)
	if err != nil {
		fmt.Println(err)
	}
	err = futugg.On("recv.Qot_GetWarrant", QotGetWarrantRecv)
	if err != nil {
		fmt.Println(err)
	}
}

func QotGetWarrantSend(conn *futugg.FutuGG, begin int32, owerStock string, warrType string, 
		minPriceRecoveryRatio float64, maxPriceRecoveryRatio float64,
		minCurPrice float64, maxCurPrice float64,
		minVolume uint64, maxVolume uint64) error {
	pack := &futugg.FutuPack{}
	pack.SetProto(uint32(3210))

	num := int32(200)
	sortField := int32(10) // volume
	ascend := false
	ownerSecurity := transStockCode(owerStock)
	warrantType := trannsWarrantType(warrType)
	warrantStatus := int32(1)

	reqData := &Qot_GetWarrant.Request{
		C2S:  &Qot_GetWarrant.C2S{
			Begin: &begin,
			Num: &num,
			SortField: &sortField,
			Ascend: &ascend,
			Owner:  ownerSecurity,
			TypeList:  []int32{warrantType},
			IssuerList: getAllIssuer(),
			Status: &warrantStatus,
			PriceRecoveryRatioMin: &minPriceRecoveryRatio,
			PriceRecoveryRatioMax: &maxPriceRecoveryRatio,
			CurPriceMin:  &minCurPrice,
			CurPriceMax: &maxCurPrice,
			VolMin: &minVolume,
			VolMax: &maxVolume,
		},
	}

	pbData, err := proto.Marshal(reqData)
	if err != nil {
		return fmt.Errorf("marshal error: %s", err)
	}

	pack.SetBody(pbData)
	err = conn.Send(pack)

	return err
}

func QotGetWarrantRecv(data []byte) ([]byte, error) {
	rsp := &Qot_GetWarrant.Response{}
	if err := proto.Unmarshal(data, rsp); err != nil {
		return nil, err
	}
	
	m := jsonpb.Marshaler{}
	result, err := m.MarshalToString(rsp)
	return []byte(result), err
}