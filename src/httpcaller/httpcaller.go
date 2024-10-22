package httpcaller

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type HttpCaller struct {
	BaseUrl string
	Url     url.URL
	Params  url.Values
	query   string
}

func New(baseurl string) *HttpCaller {
	urlaux, _ := url.Parse(baseurl + "?")
	caller := &HttpCaller{
		BaseUrl: baseurl,
		Url:     *urlaux,
		Params:  url.Values{},
	}

	return caller
}

func (me *HttpCaller) SetQuery(query string) {
	me.query = query
}

func (me *HttpCaller) GetQuery() string {
	return me.query
}

func (me *HttpCaller) UpdateQuery() {
	me.Url.RawQuery = me.Params.Encode()
	me.SetQuery(me.Url.String())
}

func (me *HttpCaller) SendRequest(query string) {

	client := &http.Client{}
	request, err := http.NewRequest("GET", me.query, nil)
	if err != nil {
		log.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%q\n", string(body))

}
