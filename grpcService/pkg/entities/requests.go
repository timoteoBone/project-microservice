package entities

type Status struct {
	Message string
}

type CreateUserRequest struct {
	Name string
	Pass string
	Age  uint32
}

type CreateUserResponse struct {
	Status Status
	UserId string
}

type GetUserRequest struct {
	UserID string
}

type GetUserResponse struct {
	Name string
	Id   string
	Age  uint32
}
