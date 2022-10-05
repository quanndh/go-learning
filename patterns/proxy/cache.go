package proxy

import "fmt"

type Cache struct {
	application *Application
	cached      map[string]string
}

func InitCache() *Cache {
	return &Cache{
		application: &Application{},
		cached:      make(map[string]string),
	}
}

func (c Cache) HandleRequest(url string, method string) (int, string) {
	cachedKey := fmt.Sprintf("%v-%v", url, method)

	if method == "get" {
		cachedData, ok := c.cached[cachedKey]

		if ok {
			if method == "get" {
				return 200, fmt.Sprintf("%v (returned from cache)", cachedData)
			}
		}
	}

	code, msg := c.application.HandleRequest(url, method)

	if method == "get" {
		c.cached[cachedKey] = msg
	}

	return code, fmt.Sprintf("%v (returned from db)", msg)
}
