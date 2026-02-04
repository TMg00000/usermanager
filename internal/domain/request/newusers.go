package request

type Users struct {
	Username string `json:"name" bson:"name" validate:"required,min=3,max=30"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
	Hash     string `json:"-"`
}
