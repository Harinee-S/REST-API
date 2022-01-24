package modules

type users struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"num"`
	PromoCode   string `json:"promo_code"`
	Reference   string `json:"refer"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []users `json:"data"`
	Message string  `json:"message"`
}
