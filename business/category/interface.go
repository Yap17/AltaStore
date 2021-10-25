package category

type Service interface {
	// Mengambil semua data kategori
	GetAllCategory() (*[]Category, error)

	// Mengambil data kategori berdasarkan kode kategori
	FindCategoryById(id string) (*Category, error)

	// Menambahan kategori baru
	InsertCategory(category *CategorySpec) error

	// Memperbarui produk kategori
	UpdateCategory(id string, category *CategorySpec) error

	// Menghapus produk kategori
	DeleteCategory(id string, userid string) error
}

type Repository interface {
	// Mengambil semua data kategori
	GetAllCategory() (*[]Category, error)

	// Mengambil data kategoru bedasarkan kode kategori
	FindCategoryById(id string) (*Category, error)

	// Menambahkan kategori baru
	InsertCategory(category Category) error

	// Memperbarui kategori
	UpdateCategory(id string, category Category) error

	// Menghapus produk kategori
	DeleteCategory(id string, userid string) error
}
