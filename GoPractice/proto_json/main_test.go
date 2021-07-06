package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"learning-note/GoPractice/proto_json/entity"
	lProto "learning-note/GoPractice/proto_json/proto"
	"testing"
	"time"
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
}

func TestClickFirstCache(t *testing.T) {
	str := ""
	md5 := GetMD5LowerHash(str)
	fmt.Println(len(md5), md5)
	cfc := entity.ClickFirstCache{
		Md5Key:    md5,
		ClickTime: time.Now().Format("2006-01-02 15:04:05"),
		Time:      time.Now().UnixNano() / 1e6,
		Platform:  "fdmidu",
	}
	firstCacheByte, _ := json.Marshal(cfc)
	fmt.Printf("len(firstCacheByte):%d,firstCacheByte:%s \n", len(firstCacheByte), firstCacheByte)
	/**
	len(firstCacheByte):122,firstCacheByte:{"md5_key":"d41d8cd98f00b204e9800998ecf8427e","click_time":"2021-07-05 15:22:22","time":1625469742511,"platform":"fdmidu"}
	*/

	cfcu := entity.ClickFirstCacheUpdate{
		Md5Key:    cfc.Md5Key,
		ClickTime: cfc.ClickTime,
		Time:      cfc.Time,
	}
	cfcuByte, _ := json.Marshal(cfcu)
	fmt.Printf("len(cfcuByte):%d,cfcuByte:%s \n", len(cfcuByte), cfcuByte)
	/**
	len(cfcuByte):102,cfcuByte:{"md5_key":"d41d8cd98f00b204e9800998ecf8427e","click_time":"2021-07-05 15:32:27","time":1625470347081}
	*/
	key := fmt.Sprintf("%s:superlink:click:device:%s", "mdxs", "35ae9ed22fe2d044bb3aa312b27a5a32")
	fmt.Printf("len(key):%d,key:%s \n", len(key), key)
	/**
	len(key):60,key:mdxs:superlink:click:device:35ae9ed22fe2d044bb3aa312b27a5a32
	*/

	cfcu2 := entity.ClickFirstCacheUpdate2{
		Md5Key: cfc.Md5Key,
		Time:   cfc.Time,
	}
	cfcu2Byte, _ := json.Marshal(cfcu2)
	fmt.Printf("len(cfcu2Byte):%d,cfcu2Byte:%s \n", len(cfcu2Byte), cfcu2Byte)
	/**
	len(cfcu2Byte):67,cfcu2Byte:{"md5_key":"d41d8cd98f00b204e9800998ecf8427e","time":1625477839012}
	*/
}

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Printf(fmt.Sprintf("time.UnixNano:%v,1e6:%v,time.UnixNano:%v \n", now.UnixNano(), 1e6, now.UnixNano()/1e6))
	fmt.Printf(fmt.Sprintf("len(now.Format):%v,int64.Unix:%v,IntToBytes:%v,len(IntToBytes):%v \n", len(now.Format("2006-01-02 15:04:05")), now.UnixNano()/1e6, IntToBytes(now.UnixNano()/1e6), len(IntToBytes(now.UnixNano()/1e6))))
}

func TestClickSecondCache(t *testing.T) {
	key := fmt.Sprintf("%s:superlink:click:md5key:%s", "mdxs", "35ae9ed22fe2d044bb3aa312b27a5a32")
	fmt.Printf("len(key):%d,key:%s \n", len(key), key)
	/**
	len(key):60,key:mdxs:superlink:click:md5key:35ae9ed22fe2d044bb3aa312b27a5a32
	*/
	cacheStr := `{"account_id":234,"campaign_id":"1649350659212340","adgroup_id":"1649355063218196","ad_id":"","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"","tuid":"","oaid":"","device_md5":"bd6db178bb06ca267cd4cb6ec82725f7","click_time":"2021-07-05 16:04:35","android_id":"","android_id_md5":""}`

	scsc := entity.SuperlinkClickSecondCache{}
	_ = json.Unmarshal([]byte(cacheStr), &scsc)
	scscByte, _ := json.Marshal(scsc)
	fmt.Printf("len(scscByte):%d,scscByte:%s \n", len(scscByte), scscByte)
	/**
	len(scscByte):330,scscByte:{"account_id":234,"campaign_id":"1649350659212340","adgroup_id":"1649355063218196","ad_id":"","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"","tuid":"","oaid":"","device_md5":"bd6db178bb06ca267cd4cb6ec82725f7","click_time":"2021-07-05 16:04:35","android_id":"","android_id_md5":""}
	*/
	lcmi := entity.LastClickMd5Info{
		AccountID:   int32(scsc.AccountID),
		CampaignID:  scsc.CampaignID,
		AdgroupID:   scsc.AdgroupID,
		AdID:        scsc.AdID,
		Dtu:         scsc.Dtu,
		Platform:    scsc.Platform,
		ChannelID:   int32(scsc.ChannelID),
		CategoryID:  scsc.CategoryID,
		ProductID:   scsc.ProductID,
		ExtensionID: scsc.ExtensionID,
		ClickTime:   scsc.ClickTime,
	}
	lcmiByte, _ := json.Marshal(lcmi)
	fmt.Printf("len(lcmiByte):%d,lcmiByte:%s \n", len(lcmiByte), lcmiByte)
	/**
	len(lcmiByte):226,lcmiByte:{"account_id":234,"campaign_id":"1649350659212340","adgroup_id":"1649355063218196","ad_id":"","dtu":"","platform":"toutiao","channel_id":16,"category_id":"","product_id":"","extension_id":"","click_time":"2021-07-05 16:04:35"}
	*/
}
