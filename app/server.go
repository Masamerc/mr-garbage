package app

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Masamerc/mr-garbage/garbage"
	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RequestBody struct {
	Day string `json:"day"`
}

func Serve(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/", returnMessage).Methods("GET")
	r.HandleFunc("/day", BroadcastGarbageInfo).Methods("POST")
	r.HandleFunc("/callback", basicReply)

	log.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func returnMessage(w http.ResponseWriter, r *http.Request) {
	resp := Response{Status: "ok", Message: "ok"}
	respJson, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
}

func BroadcastGarbageInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON request body", http.StatusBadRequest)
		return
	}

	bot := GetBot()
	schedule := garbage.GetScheduleFromRawSchedule()

	if requestBody.Day == "Week" {
		Broadcast(bot, garbage.GetWeeklySchedule(schedule))
	} else {
		garbages := schedule[requestBody.Day]
		for _, garbage := range garbages {
			Broadcast(bot, "REMINDER!\n"+garbage.FormatMessage(true))
		}
	}
}

func basicReply(w http.ResponseWriter, r *http.Request) {
	bot := GetBot()

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	schedule := garbage.GetScheduleFromRawSchedule()

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				var replyText string

				if strings.Contains(strings.ToLower(message.Text), "week") {
					replyText = garbage.GetWeeklySchedule(schedule)
				} else {
					replyText = garbage.GetGarbageInfoFromUserMessage(message.Text)
				}

				if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyText)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
