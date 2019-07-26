package request

type Url struct {
	url 				string	;
	responseStatusCode	int		;
}

func ( u *Url ) SetUrl( url string ) {
	u.url = url;
}

func ( u Url) GetUrl() string {
	return u.url;
}

func ( u *Url ) SetResponseStatusCode( code int ) {
	u.responseStatusCode = code;
}

func ( u Url) GetResponseStatusCode() int {
	return u.responseStatusCode;
}

