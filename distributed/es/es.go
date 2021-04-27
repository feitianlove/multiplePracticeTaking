package es

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v3"
)

type ElasticClient struct {
	EsClient *elastic.Client
}

func NewESClient(host string, port string) (*ElasticClient, error) {
	var esClient = new(ElasticClient)
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%s", host, port)),
		elastic.SetMaxRetries(3),
	)

	if err != nil {
		return nil, err
	}
	esClient.EsClient = client
	return esClient, nil
}

func (client *ElasticClient) InsertDocument(db string, table string, obj map[string]interface{}) (
	*elastic.IndexResponse, error) {
	id, ok := obj["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id :%s is not a string", id)
	}
	var indexName, typeName string
	// 数据库中的 database/table 概念，可以简单映射到 es 的 index 和 type
	// 不过需要注意，因为 es 中的 _type 本质上只是 document 的一个字段
	// 所以单个 index 内容过多会导致性能问题
	// 在新版本中 type 已经废弃
	// 为了让不同表的数据落入不同的 index，这里我们用 table+name 作为 index 的名字
	indexName = fmt.Sprintf("%v_%v", db, table)
	typeName = table

	// 正常情况
	index, err := client.EsClient.Index().Index(indexName).Type(typeName).Id(id).BodyJson(obj).Do()
	if err != nil {
		return nil, err
	} else {
		return index, nil
	}
}

// 通过 bool must 和 bool should 添加 bool 查询条件
//q := elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("id", 1),
//	elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("male", "m")))
//
//q = q.Should(
//	elastic.NewMatchPhraseQuery("name", "alex"),
//	elastic.NewMatchPhraseQuery("name", "xargin"),
//)

func (client *ElasticClient) Query(query elastic.Query, indexName string, typeName string) (*elastic.SearchResult, error) {

	searchService := client.EsClient.Search(indexName).Type(typeName)
	res, err := searchService.Query(query).Do()
	if err != nil {
		// log error
		return nil, err
	}

	return res, nil
}

func (client *ElasticClient) DeleteDocument(
	indexName string, typeName string, obj map[string]interface{}) (*elastic.DeleteResponse, error) {
	id, ok := obj["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id :%s is not a string", id)
	}
	res, err := client.EsClient.Delete().Index(indexName).Type(typeName).Id(id).Do()
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
