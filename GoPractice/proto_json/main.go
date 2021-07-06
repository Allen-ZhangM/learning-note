package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"learning-note/GoPractice/proto_json/entity"
	lProto "learning-note/GoPractice/proto_json/proto"
)

func main() {
	//序列化json
	lastStr := `{"account_id":0,"campaign_id":"","adgroup_id":"","ad_id":"","dtu":"","platform":"","channel_id":0,"category_id":"","product_id":"","extension_id":"","click_time":""}`
	//lastStr:=`{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"","click_time":"2021-07-02 11:50:37"}`
	//lastStr:=`{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"sdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasf","click_time":"2021-07-02 11:50:37"}`
	//lastStr:=`{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费","click_time":"2021-07-02 11:50:37"}`
	var lastByte []byte
	last := &entity.LastClickMd5Info{}
	_ = json.Unmarshal([]byte(lastStr), last)
	lastByte, _ = json.Marshal(last)
	fmt.Printf("len(lastByte):%d,len([]byte(lastStr)):%d,len(lastStr):%d,last:%v\n", len(lastByte), len([]byte(lastStr)), len(lastStr), last)

	//序列化proto
	pLast := &lProto.LastClickMd5Info{
		AccountId:   last.AccountID,
		CampaignId:  last.CampaignID,
		AdgroupId:   last.AdgroupID,
		AdId:        last.AdID,
		Dtu:         last.Dtu,
		Platform:    last.Platform,
		ChannelId:   last.ChannelID,
		CategoryId:  last.CategoryID,
		ProductId:   last.ProductID,
		ExtensionId: last.ExtensionID,
		ClickTime:   last.ClickTime,
	}
	pByte, _ := proto.Marshal(pLast)
	fmt.Printf("len(pByte):%d,pByte:%s,pLast:%s \n", len(pByte), pByte, pLast)

	//结论
	//json序列化后byte长度为(空结构，常规数据，大量英文，大量中文)：165,243,1095,1761；
	//proto序列化后byte长度为(空结构，常规数据，大量英文，大量中文)：0,89,944,1610；
	//序列化后len差值为：165,154,151,151；
	//可见proto节省的就是结构体key的长度，key越长，value越小，proto节省空间的效果越好；
}

//返回一个32位md5加密后的小写字符串
func GetMD5LowerHash(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//整形转换成字节
func IntToBytes(n int64) []byte {
	x := n

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
