package main

import (
    "strconv"
    "BlockChainTest/mycrypto"
    "github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func GetTracer(platformID, repID string) []byte {
    key := mycrypto.Hash256([]byte("tracer"), []byte(platformID), []byte(repID))
    value, _ := db.Get(key, nil)
    return value
}

func PutTracer(platformID, repID string, value []byte) error {
    key := mycrypto.Hash256([]byte("tracer"), []byte(platformID), []byte(repID))
    return db.Put(key, value, nil)
}

func IsActive(platformID, repID string) []byte {
    key := mycrypto.Hash256([]byte("active_state"), []byte(platformID), []byte(repID))
    value, _ := db.Get(key, nil)
    return value
}

func SetActive(platformID, repID string, value []byte) error {
    key := mycrypto.Hash256([]byte("active_state"), []byte(platformID), []byte(repID))
    return db.Put(key, value, nil)
}

func GetGroup(platformID string, groupID uint64) []byte {
    key := mycrypto.Hash256([]byte("group"), []byte(platformID), []byte(strconv.FormatInt(int64(groupID), 10)))
    value, _ := db.Get(key, nil)
    if value == nil {
        value = []byte("[]")
    }
    return value
}

func PutGroup(platformID string, groupID uint64, value []byte) error {
    key := mycrypto.Hash256([]byte("group"), []byte(platformID), []byte(strconv.FormatInt(int64(groupID), 10)))
    return db.Put(key, value, nil)
}