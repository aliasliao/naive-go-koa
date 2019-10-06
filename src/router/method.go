package router

type Method int

const (
	GET Method = iota
	POST
	PUT
	PATCH
	DELETE
	ANY
)

func (m Method) String() string {
	return []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"ANY",
	}[m]
}

func StringToMethod(str string) Method {
	switch str {
	case GET.String():
		return GET
	case POST.String():
		return POST
	case PUT.String():
		return PUT
	case PATCH.String():
		return PATCH
	case DELETE.String():
		return DELETE
	default:
		return ANY
	}
}
