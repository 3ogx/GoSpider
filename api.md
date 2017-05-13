# ���Ĵ�������

APIʹ���뿴����ʾ����������������������,���Ĵ���spider/spider.go�

```go
// �½�һ�����棬���ipstring��һ������IP��ַ����ʹ�ô���ͻ���
func NewSpider(ipstring interface{}) (*Spider, error) {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	if ipstring != nil {
		client, err := NewProxyClient(ipstring.(string))
		spider.Client = client
		spider.Ipstring = ipstring.(string)
		return spider, err
	} else {
		client, err := NewClient()
		spider.Client = client
		spider.Ipstring = "localhost"
		return spider, err
	}

}
```

���Դ���ipstring����ʾʹ�ô���Ĭ�Ͽ���cookie��¼��cookie��һֱ���ڴ��и��£�Ĭ����ͷ�������Ҫ�Զ���http client�ͻ���,ʹ�ã�

```go
// ͨ���ٷ�Client���½����棬�����������
func NewSpiderByClient(client *http.Client) *Spider {
	spider := new(Spider)
	spider.SpiderConfig = new(SpiderConfig)
	spider.Header = http.Header{}
	spider.Data = url.Values{}
	spider.BData = []byte{}
	spider.Client = client
	return spider
}
```

�ٷ���http.Client����ô�õģ���spider/client.go

```go
//cookie record
// ��¼Cookie
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//default client to ask get or post
	// Ĭ�ϵĹٷ��ͻ��ˣ���cookie,����ʹ�ã�û�г�ʱʱ�䣬����cookie�Ŀͻ��˲��ṩ
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Logger.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
)
```

�ÿͻ����ض����ӡ��־��֧��cookie�־ã���Ҳ�������ó�ʱʱ�䣬����SSH�ȡ�������