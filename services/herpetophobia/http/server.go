package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"net/http"
	"snake/dao"
	"snake/objects"
	"strings"
	"text/template"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getUuid() string {
	return uuid.NewString()
}

func playContent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/playGame/")
		if id == "" {
			errorResp(w, 400, errors.New("Parameter 'id' is missing in url."))
		}
		tmpl, _ := template.ParseFiles("./static/websock.html")
		tmpl.Execute(w, id)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var _map objects.Map
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&_map)
		if err != nil {
			errorResp(w, 500, err)
			return
		}
		gameId := getUuid()
		dao.SaveMap(objects.Level{
			Id:        gameId,
			Secret:    _map.Secret,
			Counter:   0,
			Init:      _map.Init,
			Flag:      _map.Flag,
			CreatedAt: time.Now(),
		})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		resp := make(map[string]string)
		resp["msg"] = "Created"
		resp["id"] = gameId
		jsonResp, _ := json.Marshal(resp)
		_, _ = w.Write(jsonResp)
		return
	}
	errorResp(w, 405, errors.New("method not allowed"))
}

func gameList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		limOff := make(map[string]int64)
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&limOff)
		if err != nil {
			errorResp(w, 500, err)
			return
		}
		listIds := dao.ListId(limOff["limit"], limOff["offset"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		jsonResp, _ := json.Marshal(listIds)
		_, _ = w.Write(jsonResp)
		return
	}
	errorResp(w, 405, errors.New("method not allowed"))
}

func play(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errorResp(w, 500, err)
		return
	}
	msg := make(map[string]string)
	_ = conn.ReadJSON(&msg)
	if msg["id"] != "" {
		gameConn := NewGameConn(conn, msg["id"])
		go gameConn.Play()
		return
	}
	errorResp(w, 401, errors.New("can't find id"))
	conn.Close()
}

func errorResp(w http.ResponseWriter, code int, err error) {
	resp := make(map[string]string)
	resp["msg"] = err.Error()
	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func StartServ() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/playGame/", playContent)

	http.HandleFunc("/create", create)
	http.HandleFunc("/gameList", gameList)
	http.HandleFunc("/play", play)
	http.ListenAndServe(":8080", nil)
}
