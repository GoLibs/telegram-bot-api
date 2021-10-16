package go_telegram_bot_api

import (
	"encoding/json"
	"io/ioutil"
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

func (w *webhook) Listen() (updatesChannel chan Update) {
	if w.path == "" {
		w.path = "/"
	}
	updatesChannel = make(chan Update)
	go w.listen(updatesChannel)
	return
}

func (w *webhook) listen(updatesChannel chan Update) (err error) {
	handler := http.NewServeMux()
	handler.HandleFunc(w.path, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		var (
			update  Update
			data    []byte
			dataErr error
		)
		defer request.Body.Close()
		data, dataErr = ioutil.ReadAll(request.Body)
		if dataErr != nil {
			updatesChannel <- newUpdateError(err)
			return
		}
		dataErr = json.Unmarshal(data, &update)
		if dataErr != nil {
			updatesChannel <- newUpdateError(err)
			return
		}
		update.raw = data
		updatesChannel <- update
	})
	err = http.ListenAndServe(w.url, handler)
	if err != nil {
		updatesChannel <- newUpdateError(err)
	}
	return
}
