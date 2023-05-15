package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func PostAnything(url string, data []byte) (*http.Response, error) {
	request, errReq := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if errReq != nil {
		return nil, errReq
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, errResp := client.Do(request)
	if errResp != nil {
		return nil, errResp
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	return response, nil
}

func GetAnything(url string) (*http.Response, error) {
	response, errResp := http.Get(url)
	if errResp != nil {
		return nil, errResp
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	return response, nil

}

func PostCouriers(couriers []Courier) ([]Courier, *http.Response, error) {
	url := fmt.Sprintf("%s/couriers", os.Getenv("host"))
	type input struct {
		Couriers []Courier `json:"couriers"`
	}
	i := input{Couriers: couriers}
	js, err := json.Marshal(i)
	if err != nil {
		return nil, nil, err
	}

	request, errReq := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if errReq != nil {
		return nil, nil, errReq
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, errResp := client.Do(request)
	if errResp != nil {
		return nil, nil, errResp
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)
	type output struct {
		Couriers []Courier `json:"couriers"`
	}
	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return nil, nil, errJs
	}
	return o.Couriers, response, nil
}

func GetCourier(ID int64) (*Courier, *http.Response, error) {
	url := fmt.Sprintf("%s/couriers/%d", os.Getenv("host"), ID)
	response, errReq := http.Get(url)
	if errReq != nil {
		return nil, nil, errReq
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)
	var courier Courier
	errJs := json.Unmarshal(body, &courier)
	if errJs != nil {
		return nil, nil, errJs
	}
	return &courier, response, nil
}

func GetCouriers(limit, offset int32) ([]Courier, *http.Response, error) {
	url := fmt.Sprintf("%s/couriers?limit=%d&offset=%d", os.Getenv("host"), limit, offset)
	response, errResp := http.Get(url)
	if errResp != nil {
		return nil, nil, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)

	type output struct {
		Couriers []Courier `json:"couriers"`
		Limit    int       `json:"limit"`
		Offset   int       `json:"offset"`
	}

	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return nil, nil, errJs
	}

	return o.Couriers, response, nil

}

func DeleteCourier(ID int64) (*http.Response, error) {
	url := fmt.Sprintf("%s/couriers/%d", os.Getenv("host"), ID)
	request, errReq := http.NewRequest("DELETE", url, nil)
	if errReq != nil {
		return nil, nil
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	return response, nil
}

func PostOrders(orders []Order) ([]Order, *http.Response, error) {
	url := fmt.Sprintf("%s/orders", os.Getenv("host"))
	type input struct {
		Orders []Order `json:"orders"`
	}
	i := input{Orders: orders}
	js, err := json.Marshal(i)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, errResp := client.Do(request)
	if errResp != nil {
		return nil, nil, errResp
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	body, _ := io.ReadAll(response.Body)
	type output struct {
		Orders []Order `json:"orders"`
	}
	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return nil, nil, errJs
	}
	return o.Orders, response, nil

}

func GetOrder(ID int64) (*Order, *http.Response, error) {
	url := fmt.Sprintf("%s/orders/%d", os.Getenv("host"), ID)
	response, errReq := http.Get(url)
	if errReq != nil {
		return nil, nil, errReq
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)
	var order Order
	errJs := json.Unmarshal(body, &order)
	if errJs != nil {
		return nil, nil, errJs
	}
	return &order, response, nil
}

func GetOrders(limit, offset int32) ([]Order, *http.Response, error) {
	url := fmt.Sprintf("%s/orders?limit=%d&offset=%d", os.Getenv("host"), limit, offset)
	response, errResp := http.Get(url)
	if errResp != nil {
		return nil, nil, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)

	type output struct {
		Orders []Order `json:"orders"`
		Limit  int     `json:"limit"`
		Offset int     `json:"offset"`
	}

	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return nil, nil, errJs
	}

	return o.Orders, response, nil

}

func DeleteOrder(ID int64) (*http.Response, error) {
	url := fmt.Sprintf("%s/orders/%d", os.Getenv("host"), ID)
	request, errReq := http.NewRequest("DELETE", url, nil)
	if errReq != nil {
		return nil, nil
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	return response, nil
}

func CompleteOrders(info []CompleteInfo) ([]Order, *http.Response, error) {
	url := fmt.Sprintf("%s/orders/complete", os.Getenv("host"))
	type input struct {
		Info []CompleteInfo `json:"complete_info"`
	}
	i := input{Info: info}
	js, err := json.Marshal(i)
	if err != nil {
		return nil, nil, err
	}

	request, errReq := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if errReq != nil {
		return nil, nil, errReq
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, errResp := client.Do(request)
	if errResp != nil {
		return nil, nil, errResp
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, response, nil
	}

	body, _ := io.ReadAll(response.Body)
	type output struct {
		Orders []Order `json:"orders"`
	}
	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return nil, nil, errJs
	}
	return o.Orders, response, nil

}

func CalculateRating(courierID int64, startDate, endDate string) (int32, int32, *http.Response, error) {
	url := fmt.Sprintf("%s/couriers/meta-info/%d?startDate=%s&endDate=%s", os.Getenv("host"), courierID, startDate, endDate)
	response, errResp := http.Get(url)
	if errResp != nil {
		return 0, 0, nil, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return 0, 0, response, nil
	}

	body, _ := io.ReadAll(response.Body)

	type output struct {
		Courier
		Rating   int32 `json:"rating"`
		Earnings int32 `json:"earnings"`
	}

	var o output
	errJs := json.Unmarshal(body, &o)
	if errJs != nil {
		return 0, 0, nil, errJs
	}
	return o.Rating, o.Earnings, response, nil

}

func TestCourierController(t *testing.T) {
	var couriers []Courier
	var count = 10
	for i := 0; i < count; i++ {
		var cour Courier
		cour.Regions = append(cour.Regions, 2, 3, 4, 5)
		cour.WorkingHours = append(cour.WorkingHours, "12:00-15:00")
		cour.CourierType = "FOOT"
		couriers = append(couriers, cour)
	}

	out, resp, err := PostCouriers(couriers)
	require.NoError(t, err, "err post couriers")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, count, len(out), "err with len couriers in post")
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 100)
		courier, response, err := GetCourier(out[i].CourierID)
		require.NoError(t, err, "err get courier")
		require.Equal(t, http.StatusOK, response.StatusCode, "HTTP status code")
		require.Equal(t, *courier, out[i])
	}

	out, resp, err = GetCouriers(10, 5)
	require.NoError(t, err, "err post couriers")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, 5, len(out), "err with len couriers in get")

	out, resp, err = GetCouriers(10, 0)
	require.NoError(t, err, "err post couriers")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, 10, len(out), "err with len couriers in get")

	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 200)
		responseDel, errDel := DeleteCourier(out[i].CourierID)
		require.NoError(t, errDel, "err del courier")
		require.Equal(t, http.StatusOK, responseDel.StatusCode, "HTTP status code")
		_, responseGet, errGet := GetCourier(out[i].CourierID)
		require.NoError(t, errGet, "err get courier")
		require.Equal(t, http.StatusNotFound, responseGet.StatusCode, "HTTP status code")
	}
}

func TestCourierControllerErrors(t *testing.T) {
	_, response, err := GetCourier(9)
	require.NoError(t, err, "err get courier")
	require.Equal(t, http.StatusNotFound, response.StatusCode, "not found error")

	response, err = GetAnything(fmt.Sprintf("%s/couriers/ggggg", os.Getenv("host")))
	require.NoError(t, err, "err get courier")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")

	response, err = GetAnything(fmt.Sprintf("%s/couriers?limit=abra&offset=tre", os.Getenv("host")))
	require.NoError(t, err, "err get courier")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

	data := []byte(`
		{
			"num":6.13,
			"strs":["a","b"]
		}
	`)
	response, err = PostAnything(fmt.Sprintf("%s/couriers", os.Getenv("host")), data)
	require.NoError(t, err, "err post courier")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

}

func TestOrderController(t *testing.T) {
	var orders []Order
	var count = 10
	for i := 0; i < count; i++ {
		var ord Order
		ord.Cost = 2
		ord.CompletedTime = ""
		ord.Regions = 2
		ord.Weight = 12
		ord.DeliveryHours = append(ord.DeliveryHours, "12:00-15:00", "45:00-46:00")
		orders = append(orders, ord)
	}
	out, resp, err := PostOrders(orders)
	require.NoError(t, err, "err post orders")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, count, len(out), "err with len orders in post")

	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 100)
		order, response, err := GetOrder(out[i].OrderID)
		require.NoError(t, err, "err get order")
		require.Equal(t, http.StatusOK, response.StatusCode, "HTTP status code")
		require.Equal(t, *order, out[i])
	}

	out, resp, err = GetOrders(10, 5)
	require.NoError(t, err, "err post orders")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, 5, len(out), "err with len orders in get")

	out, resp, err = GetOrders(10, 0)
	require.NoError(t, err, "err post couriers")
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, 10, len(out), "err with len orders in get")

	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 200)
		responseDel, errDel := DeleteOrder(out[i].OrderID)
		require.NoError(t, errDel, "err del order")
		require.Equal(t, http.StatusOK, responseDel.StatusCode, "HTTP status code")
		_, responseGet, errGet := GetOrder(out[i].OrderID)
		require.NoError(t, errGet, "err get order")
		require.Equal(t, http.StatusNotFound, responseGet.StatusCode, "HTTP status code")
	}

}

func TestOrderControllerErrors(t *testing.T) {
	_, response, err := GetOrder(9)
	require.NoError(t, err, "err get order")
	require.Equal(t, http.StatusNotFound, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

	response, err = GetAnything(fmt.Sprintf("%s/orders/ggggg", os.Getenv("host")))
	require.NoError(t, err, "err get order")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

	response, err = GetAnything(fmt.Sprintf("%s/orders?limit=abra&offset=tre", os.Getenv("host")))
	require.NoError(t, err, "err get order")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

	data := []byte(`
		{
			"num":6.13,
			"strs":["a","b"]
		}
	`)
	response, err = PostAnything(fmt.Sprintf("%s/orders", os.Getenv("host")), data)
	require.NoError(t, err, "err post order")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "not found error")
	time.Sleep(time.Millisecond * 100)

}

func TestRatingCouriers(t *testing.T) {
	//complete + rating

	var courier1 Courier
	var courier2 Courier

	courier1.CourierType = "BIKE"
	courier2.CourierType = "AUTO"

	couriers := []Courier{courier1, courier2}

	couriersPost, response, err := PostCouriers(couriers)
	require.NoError(t, err, "error post couriers")
	require.Equal(t, http.StatusOK, response.StatusCode, "status code post couriers")

	var order1 Order
	var order2 Order
	var order3 Order

	order1.Cost = 250
	order2.Cost = 500
	order3.Cost = 750

	orders := []Order{order1, order2, order3}
	ordersPost, response, err := PostOrders(orders)
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error post orders")
	require.Equal(t, http.StatusOK, response.StatusCode, "status code post orders")

	var info1 CompleteInfo
	var info2 CompleteInfo
	var info3 CompleteInfo

	info1.CourierId = couriersPost[0].CourierID
	info1.OrderId = ordersPost[0].OrderID
	info1.CompleteTime = "2023-05-15T16:13:56Z"

	info2.CourierId = couriersPost[0].CourierID
	info2.OrderId = ordersPost[1].OrderID
	info2.CompleteTime = "2023-05-15T17:13:56Z"

	info3.CourierId = couriersPost[0].CourierID
	info3.OrderId = ordersPost[2].OrderID
	info3.CompleteTime = "2023-05-15T18:13:56Z"

	info := []CompleteInfo{info1, info2, info3}

	_, response, err = CompleteOrders(info)
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error complete orders")
	require.Equal(t, http.StatusOK, response.StatusCode, "status code complete orders")

	var badInfo1 CompleteInfo
	badInfo1.CourierId = 1
	badInfo1.OrderId = 2
	badInfo1.CompleteTime = "2023-05-15Z"
	req := []CompleteInfo{badInfo1}

	_, response, err = CompleteOrders(req)
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error complete orders")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "status code complete orders")

	var badInfo2 CompleteInfo
	badInfo2.CourierId = couriersPost[1].CourierID
	badInfo2.OrderId = ordersPost[0].OrderID
	badInfo2.CompleteTime = "2023-05-15T18:13:56Z"
	req1 := []CompleteInfo{badInfo2}

	_, response, err = CompleteOrders(req1)
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error complete orders")
	require.Equal(t, http.StatusBadRequest, response.StatusCode, "status code complete orders")

	_, earning, response, err := CalculateRating(couriersPost[0].CourierID, "2023-05-15T16:13:56Z", "2023-05-15T18:13:56Z")
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error rating")
	require.Equal(t, http.StatusOK, response.StatusCode, "status code ")
	require.Equal(t, 3*(order1.Cost+order2.Cost+order3.Cost), earning, "earning error")

	rating, earning, response, err := CalculateRating(couriersPost[0].CourierID, "2023-05-15T16:13:56Z", "2023-05-15T17:13:56Z")
	time.Sleep(time.Millisecond * 100)
	require.NoError(t, err, "error rating")
	require.Equal(t, http.StatusOK, response.StatusCode, "status code ")
	require.Equal(t, 3*(order1.Cost+order2.Cost), earning, "earning error")
	require.Equal(t, int32(4), rating, "rating error")
}
