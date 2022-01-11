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
	UserId int64
}

type GetUserRequest struct {
	UserID int64
}

type GetUserResponse struct {
	Name string
	Id   int64
	Age  uint32
}
