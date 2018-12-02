package interfaces

type Response struct{}

func NewResponse() Response {
	return Response{}
}

func (r *Response) Success(c Context, result interface{}) error {
	return c.JSON(200, result)
}

func (r *Response) ErrInternalServerError(c Context, result interface{}) error {
	return c.JSON(500, result)
}
