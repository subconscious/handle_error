package handle_error

import "log"

func handle_error(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
