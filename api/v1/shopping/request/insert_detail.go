package request

type DetailItemInShopCart struct {
	ID        string `json:"id"`
	ProductId string `json:"productid"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
	UserId    string `json:"userid"`
}
