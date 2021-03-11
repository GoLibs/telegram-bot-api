package go_telegram_bot_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/GoLibs/telegram-bot-api/responses"
	"github.com/GoLibs/telegram-bot-api/structs"
	"github.com/go-resty/resty/v2"
)

func (tb *TelegramBot) prepareRequest(config Config, request *resty.Request) {
	if val, isFile := config.(file); isFile && len(val.medias()) != 0 {
		for _, info := range val.medias() {
			if info.Reader != nil {
				request.SetFileReader(info.Field, info.Name, info.Reader)
			} else if info.Path != "" {
				request.SetFile(info.Field, info.Path)
			}
		}
		request.SetFormData(tb.getFormData(config))
	} else {
		request.SetHeader("Content-Type", "application/json")
		body, _ := config.marshalJSON()
		request = request.SetBody(string(body))
	}
}

func (tb *TelegramBot) getFormData(config Config) (fd map[string]string) {
	j, _ := config.marshalJSON()
	var result map[string]interface{}
	_ = json.Unmarshal(j, &result)
	fd = map[string]string{}
	for k, v := range result {
		switch v.(type) {
		case string:
			fd[k] = v.(string)
		case float64:
			if k != "latitude" && k != "longitude" && k != "x_shift" && k != "y_shift" && k != "scale" {
				fd[k] = fmt.Sprintf("%d", int64(v.(float64)))
			} else {
				fd[k] = fmt.Sprintf("%f", v.(float64))
			}
		case bool:
			fd[k] = strconv.FormatBool(v.(bool))
		case map[string]interface{}:
			jm, _ := json.Marshal(v.(map[string]interface{}))
			fd[k] = string(jm)
		case []interface{}:
			jm, _ := json.Marshal(v.([]interface{}))
			fd[k] = string(jm)
		case [][]map[string]string:
			jm, _ := json.Marshal(v.([][]map[string]string))
			fd[k] = string(jm)
		default:
			fmt.Println(fmt.Sprintf("Unknown type: %T", v))
		}
	}
	return
}

func (tb TelegramBot) getMessageResponse(resp *resty.Response, config Config) (response *responses.Response, err error) {
	defer resp.RawBody().Close()
	j := json.NewDecoder(resp.RawBody())
	response = &responses.Response{}
	err = j.Decode(response)
	if err != nil {
		return
	}
	if !response.Ok {
		err = responses.Error{
			ErrorCode:   response.ErrorCode,
			Description: response.Description,
		}
		response = nil
		return
	}

	var responseVar = config.response()
	json.Unmarshal(response.Result, &responseVar)
	// fmt.Println(fmt.Sprintf("%T\n%T\n%s", response.Result, responseVar, string(response.Result)))
	switch responseVar.(type) {
	case *[]structs.Message:
		messages := responseVar.(*[]structs.Message)
		response.Messages = *messages
	case *[]structs.Update:
		updates := responseVar.(*[]structs.Update)
		response.Updates = *updates
	case *[]structs.ChatMember:
		updates := responseVar.(*[]structs.ChatMember)
		response.ChatMembers = *updates
	case *structs.ChatMember:
		response.ChatMember = responseVar.(*structs.ChatMember)
	case *structs.Message:
		response.Message = responseVar.(*structs.Message)
	case *structs.UserProfilePhotos:
		response.UserProfilePhotos = responseVar.(*structs.UserProfilePhotos)
	case *structs.File:
		response.File = responseVar.(*structs.File)
	case *structs.Chat:
		response.Chat = responseVar.(*structs.Chat)
	case *structs.StickerSet:
		response.StickerSet = responseVar.(*structs.StickerSet)
	case *structs.MessageId:
		response.MessageId = responseVar.(*structs.MessageId)
	case *bool:
		response.Bool = responseVar.(*bool)
	case *int64:
		response.Int = responseVar.(*int64)
	default:
		err = errors.New(fmt.Sprintf("unknown response result %s - %T - %+v", string(response.Result), responseVar, config))
	}
	return
}

func randomString() string {
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	return s[rand.Intn(len(s)-1)]
}
