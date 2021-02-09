// @title main.go
// @description ES 工具类
// @author YC - 2021/1/8 9:42
package yc_es

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

// @Summary 追加
// @Param es *elasticsearch.Client "es 的 Client"
// @Param index string "es index 信息"
// @Param id string "es id 信息"
// @Param appendContent string "es 需要追加的信息"
// @Return error "异常信息"
func Append(es *elasticsearch.Client, index, id, appendContent string) error {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"script": map[string]interface{}{
			"params": map[string]interface{}{
				"message": appendContent,
			},
			"source": "ctx._source.message += params.message",
			"lang":   "painless",
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	res, err := es.Update(
		index,
		id,
		&buf,
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	return nil
}

// @Summary 根据 Index 删除
// @Param es *elasticsearch.Client "es 的 Client"
// @Param index string "es index 信息"
// @Param id string "es id 信息"
// @Return error "异常信息"
func Delete(es *elasticsearch.Client, index, id string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// @Summary 添加
// @Param es *elasticsearch.Client "es 的 Client"
// @Param index string "es index 信息"
// @Param id string "es id 信息"
// @Param content string "es content 信息，内容格式为 json"
// @Return error "异常信息"
func Create(es *elasticsearch.Client, index, id, content string) error {
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(content),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		errMessage := fmt.Sprintf("[%s] Error indexing document ID=%v", res.Status(), id)
		return errors.New(errMessage)
	}
	return nil
}

// @Summary 查询
// @Param es *elasticsearch.Client "es 的 Client"
// @Param index string "es index 信息"
// @Param id string "es id 信息"
// @return map[string]interface{} "es 响应信息"
// @return error "异常信息"
func Search(es *elasticsearch.Client, index, id string) (map[string]interface{}, error) {
	var r map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	}

	return r, nil
}

// @Summary 查询
// @Param hs string "es 的集群连接信息"
// @Return *elasticsearch.Client "es 的 Client"
// @Return error "异常信息"
func GetConnection(hs string) (*elasticsearch.Client, error) {
	hosts := strListBySplit(hs, ",")
	cfg := elasticsearch.Config{
		Addresses: hosts,
	}
	var err error
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s\n", err)
		return nil, err
	}

	// 判断连接是否正常
	_, err = es.Info()
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return nil, err
	}

	return es, nil
}

// @Summary 根据分隔符分割字符串生成切片
// @Param data string "字符串信息"
// @Param str string "指定分隔符"
// @Return s string "以指定分隔符分隔后的切片"
func strListBySplit(data string, str string) (s []string) {
	result := strings.Split(data, str)
	s = make([]string, 0, len(result))
	for _, v := range result {
		v = strings.TrimSpace(v)
		s = append(s, v)
	}
	return s
}
