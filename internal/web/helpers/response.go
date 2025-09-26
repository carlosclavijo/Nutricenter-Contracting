package helpers

type Response[T any] struct {
	Success bool   `json:"success"`
	Length  int    `json:"length,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   *Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
