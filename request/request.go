package request;

type Request struct {
	hash               string;
	url                string;
	responseStatusCode int;
}

func (r *Request) SetHash(hash string) {
	r.hash = hash;
}

func (r *Request) GetHash() string {
	return r.hash;
}

func (r *Request) SetUrl(url string) {
	r.url = url;
}

func (r Request) GetUrl() string {
	return r.url;
}

func (r *Request) SetResponseStatusCode(code int) {
	r.responseStatusCode = code;
}

func (r Request) GetResponseStatusCode() int {
	return r.responseStatusCode;
}
