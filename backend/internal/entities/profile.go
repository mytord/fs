package entities

type Profile struct {
	Id        int
	Email     string `validate:"required,email"`
	Password  string `validate:"required,max=1000"`
	FirstName string `validate:"required,max=255"`
	LastName  string `validate:"required,max=255"`
	City      string `validate:"required,max=255"`
	Age       int    `validate:"required,gte=0,lte=130"`
	Interests string `validate:"max=2000"`
}
