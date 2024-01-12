package server

import (
	"encoding/json"
	"go-discover-server/core"
	"go-discover-server/log"
	"net/http"
)

func RunHttpServer(addr string) {
	http.HandleFunc("/api/devices", handleGetAllDevices)

	var err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic("could not start http server: " + err.Error())
	}
}

func handleGetAllDevices(w http.ResponseWriter, r *http.Request) {
	var err error
	var devices = core.GetAllDevice()
	var success = SuccessMessage{
		Code: 1,
		Data: devices,
	}

	w.Header().Set("Content-Type", "application/json")

	var jsonData []byte
	jsonData, err = json.Marshal(success)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(createError(-1, "could not get devices"))
		log.E("could not parse json data: %s", err.Error())
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.E("could not send json data: %s", err.Error())
	}
}
