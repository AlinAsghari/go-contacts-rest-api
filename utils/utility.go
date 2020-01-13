package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Message(status bool, message string) (map[string]interface{}) {
	// return map is a dictionary ( key=string , value= interface{} (means object))
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func FloatToString(input_num float64) string  {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
func IntToString(input_num int) string  {
	return strconv.Itoa(input_num)
}
func StringToInt(input string) ( int , error )  {
	return strconv.Atoi(input)
}

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
