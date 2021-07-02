package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"learning-note/GoPractice/proto_json/entity"
	lProto "learning-note/GoPractice/proto_json/proto"
	"testing"
)

var (
	emptyStr   = `{"account_id":0,"campaign_id":"","adgroup_id":"","ad_id":"","dtu":"","platform":"","channel_id":0,"category_id":"","product_id":"","extension_id":"","click_time":""}`
	normalStr  = `{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"","click_time":"2021-07-02 11:50:37"}`
	englishStr = `{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"sdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasfsdfassadfasf","click_time":"2021-07-02 11:50:37"}`
	chineseStr = `{"account_id":2920,"campaign_id":"1701623169620024","adgroup_id":"1702160451818557","ad_id":"1702163495328894","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费是劳动法；按理说快递费","click_time":"2021-07-02 11:50:37"}`
	strs       = map[string]string{"emptyStr": emptyStr, "normalStr": normalStr, "englishStr": englishStr, "chineseStr": chineseStr}
)

func Benchmark(b *testing.B) {
	for k, v := range strs {
		last := &entity.LastClickMd5Info{}
		byteLast := []byte(v)
		b.Run("json.Unmarshal."+k, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = json.Unmarshal(byteLast, last)
			}
		})

		_ = json.Unmarshal(byteLast, last)
		b.Run("json.Marshal."+k, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = json.Marshal(last)
			}
		})

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
		b.Run("proto.Marshal."+k, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = proto.Marshal(pLast)
			}
		})

		pByte, _ := proto.Marshal(pLast)
		protoLast := &lProto.LastClickMd5Info{}
		b.Run("proto.Unmarshal."+k, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = proto.Unmarshal(pByte, protoLast)
			}
		})
	}

}

/**
  执行结果：
    Benchmark/json.Unmarshal.emptyStr
    Benchmark/json.Unmarshal.emptyStr-4         	  147146	      7806 ns/op
    Benchmark/json.Marshal.emptyStr
    Benchmark/json.Marshal.emptyStr-4           	  861547	      1444 ns/op
    Benchmark/proto.Marshal.emptyStr
    Benchmark/proto.Marshal.emptyStr-4          	 1934200	       614.9 ns/op
    Benchmark/proto.Unmarshal.emptyStr
    Benchmark/proto.Unmarshal.emptyStr-4        	 8919195	       131.4 ns/op
    Benchmark/json.Unmarshal.normalStr
    Benchmark/json.Unmarshal.normalStr-4        	  119592	      9706 ns/op
    Benchmark/json.Marshal.normalStr
    Benchmark/json.Marshal.normalStr-4          	  614247	      2024 ns/op
    Benchmark/proto.Marshal.normalStr
    Benchmark/proto.Marshal.normalStr-4         	 1316199	      1183 ns/op
    Benchmark/proto.Unmarshal.normalStr
    Benchmark/proto.Unmarshal.normalStr-4       	 1350600	       833.2 ns/op
    Benchmark/json.Unmarshal.englishStr
    Benchmark/json.Unmarshal.englishStr-4       	   55011	     22527 ns/op
    Benchmark/json.Marshal.englishStr
    Benchmark/json.Marshal.englishStr-4         	  176656	      7036 ns/op
    Benchmark/proto.Marshal.englishStr
    Benchmark/proto.Marshal.englishStr-4        	  603134	      1986 ns/op
    Benchmark/proto.Unmarshal.englishStr
    Benchmark/proto.Unmarshal.englishStr-4      	  617253	      1979 ns/op
    Benchmark/json.Unmarshal.chineseStr
    Benchmark/json.Unmarshal.chineseStr-4       	   35311	     34434 ns/op
    Benchmark/json.Marshal.chineseStr
    Benchmark/json.Marshal.chineseStr-4         	  106598	     11440 ns/op
    Benchmark/proto.Marshal.chineseStr
    Benchmark/proto.Marshal.chineseStr-4        	  138146	      8103 ns/op
    Benchmark/proto.Unmarshal.chineseStr
    Benchmark/proto.Unmarshal.chineseStr-4      	  144132	      8144 ns/op

  比值：
      Benchmark/json.Marshal.emptyStr-4           	  861547	      1444 ns/op
      Benchmark/proto.Marshal.emptyStr-4          	 1934200	       614.9 ns/op
      Benchmark/json.Unmarshal.emptyStr-4         	  147146	      7806 ns/op
      Benchmark/proto.Unmarshal.emptyStr-4        	 8919195	       131.4 ns/op

      Benchmark/json.Marshal.normalStr-4          	  614247	      2024 ns/op
      Benchmark/proto.Marshal.normalStr-4         	 1316199	      1183 ns/op
      Benchmark/json.Unmarshal.normalStr-4        	  119592	      9706 ns/op
      Benchmark/proto.Unmarshal.normalStr-4       	 1350600	       833.2 ns/op

      Benchmark/json.Marshal.englishStr-4         	  176656	      7036 ns/op
      Benchmark/proto.Marshal.englishStr-4        	  603134	      1986 ns/op
      Benchmark/json.Unmarshal.englishStr-4       	   55011	     22527 ns/op
      Benchmark/proto.Unmarshal.englishStr-4      	  617253	      1979 ns/op

      Benchmark/json.Marshal.chineseStr-4         	  106598	     11440 ns/op
      Benchmark/proto.Marshal.chineseStr-4        	  138146	      8103 ns/op
      Benchmark/json.Unmarshal.chineseStr-4       	   35311	     34434 ns/op
      Benchmark/proto.Unmarshal.chineseStr-4      	  144132	      8144 ns/op
  结论：
json的Unmarshal比Marshal慢很多（大约3倍）；proto在value比较少的情况反而Unmarshal更快，value多的情况Marshal和Unmarshal接近相等
proto对中文的操作耗时是英文的4倍，在大量中文情况下，两者的Marshal耗时趋于接近，但还是proto快一些
*/
