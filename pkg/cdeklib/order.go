package cdeklib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) CreateOrder(addrFrom string, addrTo string, recipient Recipient, packages []Package, tariffCode int) (orderID string, err error) {

	addrFromL := OrderLocation{
		Address: addrFrom,
	}

	addrToL := OrderLocation{
		Address: addrTo,
	}

	baseURL, err := url.Parse(c.ApiURL)
	if err != nil {
		panic(err)
	}

	endpoint := "v2/orders"

	fullURL := baseURL.ResolveReference(&url.URL{Path: endpoint})

	rqst := OrderRequest{
		TariffCode:   tariffCode,
		FromLocation: addrFromL,
		ToLocation:   addrToL,
		Recipient:    recipient,
		Packages:     packages,
	}

	// Преобразуем JSON-объект в []byte
	requestBodyBytes, err := json.Marshal(rqst)
	if err != nil {
		panic(err)
	}

	// Создаем http.Request с указанными параметрами
	req, err := http.NewRequest("POST", fullURL.String(), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// Устанавливаем необходимые заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var responseObject map[string]interface{}
	err = json.Unmarshal(responseBody, &responseObject)
	if err != nil {
		return "", err
	}

	uuid := responseObject["entity"].(map[string]interface{})["uuid"].(string)

	return uuid, nil
}
