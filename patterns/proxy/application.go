package proxy

type Application struct {
}

func (a Application) HandleRequest(url string, method string) (int, string) {
	if url == "/user" && method == "get" {
		return 200, "id: 1"
	}

	if url == "/user" && method == "post" {
		return 201, "id: 2"
	}

	return 404, "not found"
}
