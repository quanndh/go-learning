package proxy

type Nginx struct {
	cacheProxy        *Cache
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func InitNginx() *Nginx {
	cacheProxy := InitCache()

	return &Nginx{
		cacheProxy:        cacheProxy,
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n Nginx) HandleRequest(url string, method string) (int, string) {
	allow := n.checkRateLimit(url)

	if !allow {
		return 403, "forbidden"
	}

	return n.cacheProxy.HandleRequest(url, method)
}

func (n Nginx) checkRateLimit(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1

	return true
}
