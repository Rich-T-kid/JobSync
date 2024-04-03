package Sessions

import ("proj/DB"
	"fmt"
	"encoding/base64"
    "encoding/json"
)
func StructToJson(model DB.DBMarshall, username string) (string , error)  {
	temp , err := model.DbtoStruct(username)
	fmt.Println(temp)
	if err != nil{
		return "" , err} // better error handling here. should never happen
	bt , err  := json.Marshal(temp)
	if err != nil{
		return "" , err}
	return string(bt) , nil

}


func EncodeJSONToBase64(jsonData string) string {
    // Convert string JSON data to byte slice
    jsonDataBytes := []byte(jsonData)

    // Encode JSON data to base64
    encodedData := base64.StdEncoding.EncodeToString(jsonDataBytes)

    return encodedData
}

// DecodeBase64ToJSON decodes base64 data to JSON
func DecodeBase64ToJSON(encodedData string) string {
    // Decode base64 data
    decodedData, _ := base64.StdEncoding.DecodeString(encodedData)
    

    // Convert byte slice to string
    jsonData := string(decodedData)

    return jsonData 
}
