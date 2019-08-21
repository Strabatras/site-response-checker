package request;

type Request struct {
	finished           bool;
	hash               string;
	url                string;
	responseStatusCode int;
}

func (r *Request) SetFinished() {
	r.finished = true;
}
func (r Request) GetFinished() bool {
	return (r.finished == true);
}

func (r *Request) SetHash(hash string) {
	r.hash = hash;
}

func (r Request) GetHash() string {
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
