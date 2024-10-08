package handlers

import (
    "fmt"
    "github.com/jerryharbour/futugg"
    "github.com/jerryharbour/futugg/pb/Qot_GetTicker"

    "github.com/golang/protobuf/proto"
    "github.com/golang/protobuf/jsonpb"
)

func init() {
    futugg.SetHandlerId(uint32(3010), "Qot_GetTicker")
    var err error
    err = futugg.On("send.Qot_GetTicker", QotGetTickerSend)
    err = futugg.On("recv.Qot_GetTicker", QotGetTickerRecv)
    if err != nil {
        fmt.Println(err)
    }
}

func QotGetTickerSend(conn *futugg.FutuGG, stockCode string, maxRetNum int32) error {
    pack := &futugg.FutuPack{}
    pack.SetProto(uint32(3010))

    security := transStockCode(stockCode)

    reqData := &Qot_GetTicker.Request{
        C2S: &Qot_GetTicker.C2S{
            Security: security,
            MaxRetNum: &maxRetNum,
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

func QotGetTickerRecv(data []byte) error {
    resp := &Qot_GetTicker.Response{}
    err := proto.Unmarshal(data, resp)
    if err != nil {
        return fmt.Errorf("marshal error: %s", err)
    }

    m := jsonpb.Marshaler{}
    result, err := m.MarshalToString(resp)
    fmt.Println(result)
    return err
}
