package use_cases

type createCategoryUseCase struct {
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (uc *createCategoryUseCase) Execute(name string) {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: save category to database
	log.Println(category)
	return nil
}