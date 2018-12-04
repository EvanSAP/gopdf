package commerce

import "encoding/json"

//
//[GIN] 2018/12/04 - 15:15:18 | 200 |     389.871µs |       127.0.0.1 | POST     /event/product.addtocart
//Event Posted: product.addtocart
//{"baseSiteId":"electronics","cartId":"00001022","productId":"5913124","userId":"hao.xu01@sap.com"}
//Event Posted: product.addtocart
//{"baseSiteId":"electronics","cartId":"00001022","productId":"5913124","userId":"hao.xu01@sap.com"}
//[GIN] 2018/12/04 - 15:15:18 | 200 |    2.226652ms |       127.0.0.1 | POST     /event/product.addtocart
//Event Posted: order.created
//{"orderCode":"00001023"}
//[GIN] 2018/12/04 - 15:15:32 | 200 |     201.093µs |       127.0.0.1 | POST     /event/order.created
//Event Posted: product.addtocart
//{"baseSiteId":"electronics","cartId":"00001024","productId":"5913124","userId":"hao.xu01@sap.com"}
//[GIN] 2018/12/04 - 15:17:19 | 200 |     207.381µs |       127.0.0.1 | POST     /event/product.addtocart
//Event Posted: product.addtocart
//{"baseSiteId":"electronics","cartId":"00001024","productId":"5913124","userId":"hao.xu01@sap.com"}
//[GIN] 2018/12/04 - 15:17:19 | 200 |    2.345401ms |       127.0.0.1 | POST     /event/product.addtocart
//Event Posted: order.created
//{"orderCode":"00001025"}
//[GIN] 2018/12/04 - 15:17:33 | 200 |     376.391µs |       127.0.0.1 | POST     /event/order.created
//[GIN] 2018/12/04 - 15:39:30 | 200 |     463.032µs |       10.52.1.1 | GET      /
//➜  gopdf git:(develop) ✗

//{"orderCode":"00001025"}
type OrderCreatedEvent struct {
	OrderCode string `json:"orderCode"`
}

func UnmarshalOrderCreatedEvent(b []byte) OrderCreatedEvent {
	value := OrderCreatedEvent{}
	json.Unmarshal(b, &value)
	return value
}

//{"baseSiteId":"electronics","cartId":"00001024","productId":"5913124","userId":"hao.xu01@sap.com"}
type ProductAddToCartEvent struct {
	BaseSiteId string `json:"baseSiteId"`
	CartId string `json:"cartId"`
	ProductId string `json:"productId"`
	UserId string `json:"userId"`
}


type CommerceApi struct {

}