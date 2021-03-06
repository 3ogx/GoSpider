/*
Copyright 2017 by GoSpider author. Email: gdccmcm14@live.com
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	// 第一步：引入库 别名boss
	boss "github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoTool/util"
)

func init() {
	// 第二步：可选设置全局
	boss.SetLogLevel(boss.DEBUG) // 设置全局爬虫日志,可不设置,设置debug可打印出http请求轨迹
	boss.SetGlobalTimeout(3)     // 爬虫超时时间,可不设置

}
func main() {

	log := boss.Log() // 爬虫为你提供的日志工具,可不用

	// 第三步： 必须新建一个爬虫对象
	//sp, err := boss.NewSpider("http://smart:smart2016@104.128.121.46:808") // 代理IP爬虫 格式:协议://代理帐号(可选):代理密码(可选)@ip:port
	//sp, err := boss.NewSpider(nil)  // 正常爬虫 默认带Cookie
	//sp, err := boss.NewAPI() // API爬虫 默认不带Cookie
	sp, err := boss.New(nil) // NewSpider同名函数
	if err != nil {
		panic(err)
	}

	// 第四步：设置抓取方式和网站,可链式结构设置,只有SetUrl是必须的
	// SetUrl:Url必须设置
	// SetMethod:HTTP方法可以是POST或GET,可不设置,默认GET,传错值默认为GET
	// SetWaitTime:暂停时间,可不设置,默认不暂停
	sp.SetUrl("https://www.whitehouse.gov").SetMethod(boss.GET).SetWaitTime(2)
	sp.SetUa(boss.RandomUa())                 //设置随机浏览器标志
	sp.SetRefer("https://www.whitehouse.gov") // 设置Refer头
	sp.SetHeaderParm("diyheader", "lenggirl") // 自定义头部
	//sp.SetBData([]byte("file data")) // 如果你要提交JSON数据/上传文件
	//sp.SetFormParm("username","jinhan") // 提交表单
	//sp.SetFormParm("password","123")

	// 第五步：开始爬
	//sp.Get()             // 默认GET
	//sp.Post()            // POST表单请求,数据在SetFormParm()
	//sp.PostJSON()        // 提交JSON请求,数据在SetBData()
	//sp.PostXML()         // 提交XML请求,数据在SetBData()
	//sp.PostFILE()        // 提交文件上传请求,数据在SetBData()
	body, err := sp.Go() // 如果设置SetMethod(),采用,否则Get()
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("%s", string(body)) // 打印获取的数据
	}

	log.Debugf("%#v", sp) // 不设置全局log为debug是不会出现这个东西的

	sp.Clear() // 爬取完毕后可以清除POST的表单数据/文件数据/JSON数据
	//sp.ClearAll() // 爬取完毕后可以清除设置的Http头部和POST的表单数据/文件数据/JSON数据

	// 爬虫池子
	boss.Pool.Set("myfirstspider", sp)
	if poolspider, ok := boss.Pool.Get("myfirstspider"); ok {
		go func() {
			poolspider.SetUrl("http://www.baidu.com")
			data, _ := poolspider.Get()
			log.Info(string(data))
		}()
		util.Sleep(10)
	}

}
