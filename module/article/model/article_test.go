package articlemodel

import "testing"

func TestValidate(t *testing.T) {
	dataTest := ArticleCreate{
		Title:       "",
		Description: "",
	}

	err := dataTest.Validate()

	if err != nil {
		t.Errorf("Test Valdiate() failed, expected ErrTitleCannotBeEmpty, but got %v", err)
	}

	dataTest.Title = "Sapa Lao Cai"
	err = dataTest.Validate()
	if err != nil {
		t.Errorf("Test Valdiate() failed, expected ErrTitleCannotBeEmpty, but got %v", err)
	}

	t.Log("Test Validate() passed")
}
