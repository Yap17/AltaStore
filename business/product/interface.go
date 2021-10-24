package product

type Service interface {
	//GetAllProduct If data not found will return nil without error
	GetAllProductByParameter(id, isActive, categoryName, code, name string) (*[]Product, error)

	//GetAllProduct If data not found will return nil without error
	GetAllProduct() (*[]Product, error)

	//FindProductById If data not found will return nil without error
	FindProductById(id string) (*Product, error)

	//InsertProduct Insert new Product into storage
	InsertProduct(Product *InsertProductSpec) error

	//UpdateProduct if data not found will return error
	UpdateProduct(id string, Product *UpdateProductSpec) error

	//DeleteProduct if data not found will return error
	DeleteProduct(id string, adminId string) error
}

type Repository interface {
	//GetAllProduct If data not found will return nil without error
	GetAllProductByParameter(id, isActive, categoryName, code, name string) (*[]Product, error)

	//GetAllProduct If data not found will return nil without error
	GetAllProduct() (*[]Product, error)

	//FindProductById If data not found will return nil without error
	FindProductById(id string) (*Product, error)

	//InsertProduct Insert new Product into storage
	InsertProduct(Product Product) error

	//UpdateProduct if data not found will return error
	UpdateProduct(Product Product) error

	//DeleteProduct if data not found will return error
	DeleteProduct(Product Product) error
}
