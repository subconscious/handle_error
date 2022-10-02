package handle_error

import "log"

func Handle_error(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
