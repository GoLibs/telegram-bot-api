package tgbotapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type webhook struct {
	parent *TelegramBot
	url    string
	path   string
}

func (w *webhook) SetUrl(address string) *webhook {
	w.url = address
	return w
}

func (w *webhook) SetPath(pathUrl string) *webhook {
	w.path = pathUrl
	return w
}

func (w *webhook) Listen() error {
	if w.path == "" {
		w.path = "/"
	}

	return w.listen()
}

func (w *webhook) listen() (err error) {
	handler := http.NewServeMux()

	handler.HandleFunc(w.path, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)

		var (
			update  Update
			data    []byte
			dataErr error
		)

		defer request.Body.Close()

		data, dataErr = io.ReadAll(request.Body)
		if dataErr != nil {
			w.parent.updates <- newUpdateError(err)
			return
		}

		dataErr = json.Unmarshal(data, &update)
		if dataErr != nil {
			w.parent.updates <- newUpdateError(err)
			return
		}

		update.raw = data

		w.parent.updates <- update
	})

	return http.ListenAndServe(w.url, handler)
}
