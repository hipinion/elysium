package elysium

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	RESULTS_PER_PAGE = 25
)

type QueryData struct {
	limit  int
	offset int
	page   int
}

func NewQueryData() QueryData {
	qd := QueryData{}
	return qd
}

func (q *QueryData) Generate(opts url.Values) {
	page := opts.Get("page")
	fmt.Println(page)
	if page != "" {
		fmt.Println("eh")
		qp, err := strconv.ParseInt(page, 10, 32)
		if err != nil {
			fmt.Println(err)
			q.page = 1
			q.offset = 0
			q.limit = RESULTS_PER_PAGE
		} else {
			q.page = int(qp)
			q.offset = RESULTS_PER_PAGE * (q.page - 1)
			q.limit = RESULTS_PER_PAGE
		}
	} else {
		q.page = 1
		q.offset = 0
		q.limit = RESULTS_PER_PAGE
	}
}

func Paginate() {

}
