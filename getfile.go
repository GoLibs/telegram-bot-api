package go_telegram_bot_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoLibs/telegram-bot-api/structs"
)

type getFile struct {
	parent *TelegramBot
	fileId string
}

func (sv *getFile) marshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FileId string `json:"file_id"`
	}{
		FileId: sv.fileId,
	})
}

func (sv *getFile) response() interface{} {
	f := &structs.File{}
	f.Download = func(path string) error {
		fileAddress := fmt.Sprintf(`https://api.telegram.org/pngFile/bot%s/%s`, sv.parent.apiToken, f.FilePath)
		resp, err := sv.parent.client.SetDebug(true).R().Get(fileAddress)
		if err != nil {
			return err
		}
		defer resp.RawBody().Close()
		b, _ := ioutil.ReadAll(resp.RawBody())
		if path == "" {
			path = f.FilePath
			fPath := filepath.Base(path)
			os.Mkdir(strings.ReplaceAll(path, fPath, ""), os.ModePerm)
		}
		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
		return nil
	}
	return f
}

func (sv *getFile) method() string {
	return "POST"
}

func (sv *getFile) endpoint() string {
	return "getFile"
}

func (sv *getFile) SetFileId(fileId string) *getFile {
	sv.fileId = fileId
	return sv
}

/*func (sv *getFile) Send() (message *structs.Message, err error) {
	message, err = sv.parent.Send(sd)
	return
}
*/
