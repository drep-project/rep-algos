package usage

import (
    "encoding/json"
    "strconv"
    "fmt"
    "rep-algos/repjs"
)

type Transaction struct {
    Version              int32
    Nonce                int64
    Type                 int32
    To                   string
    ChainId              int64
    Amount               []byte
    GasPrice             []byte
    GasLimit             []byte
    Timestamp            int64
    Data                 []byte
}

func executeGainTransaction(t *Transaction) {
    var records []map[string] interface{}
    err := json.Unmarshal(t.Data, &records)
    if err != nil {
        return
    }
    increments := make([]map[string] interface{}, len(records))
    platformID := strconv.FormatInt(1, 10)
    for i, r := range records {
        uid := r["Addr"].(string)
        ret := repjs.GetProfile(platformID, uid)
        if ret == nil {
            return
        }
        repID := ret["RepID"].(string)
        groupID := ret["GroupID"].(uint64)
        tracer := GetTracer(platformID, repID)
        if tracer == nil {
            repjs.RegisterUser(platformID, repID, groupID)
        }
        increments[i] = make(map[string] interface{})
        increments[i]["RepID"] = repID
        increments[i]["Day"] = r["Day"]
        increments[i]["Gain"] = r["Gain"]
    }
    repjs.AddGain(platformID, increments)
}

func Liquidate(height int64, until int) {
    platformID := strconv.FormatInt(1, 10)
    groupID := uint64(height % 5)
    ret, _ := repjs.LiquidateRepByGroup(platformID, groupID, until)
    tokens, _ := ret["Tokens"].(map[string] interface{})
    fmt.Println("tokens: ", tokens)
}
