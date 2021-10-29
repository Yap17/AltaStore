package request

type DetailItemInShopCart struct {
	ProductId string `json:"productid"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
	UserId    string `json:"userid"`
}
