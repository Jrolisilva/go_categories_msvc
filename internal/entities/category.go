package entities

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCategory(name string) (*Category, error ){
	category := &Category{
		Name: name,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	// business rules
	if err := category.isValid(); err != nil {
	return category, nil
}

func (c *Category) isValid() error {
	if len(c.Name) < 3 {
		return fmt.Errorf("name must be at least 3 characters long %d", len(c.Name))
	}
	return nil
}