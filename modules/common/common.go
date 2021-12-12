package common

const (
	//IMAGE_CATEGORIES
	CharacterCategories int = 1
	WeaponCategories        = 2
	MaterialCategories      = 3
)

type CustomError struct {
}

func RenderImage(category int, id string) (string, error) {

	switch category {
	case CharacterCategories:
		return "Character", nil
	}

	return "", NewCusomErrorWrapper()
}

/*
* Custom Error function
 */

func (c CustomError) Error() string {
	return "No categires found"

}

func NewCusomErrorWrapper() error {
	return CustomError{}
}
