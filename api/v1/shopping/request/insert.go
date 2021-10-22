package request

type InsertShoppingCart struct {
	UserId      string `json:"userid"`
	Description string `json:"description"`
}
