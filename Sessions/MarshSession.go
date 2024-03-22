package Sessions

import ("proj/DB"
	"encoding/json")

func StructToJson(model DB.DBMarshall, username string) (string , error)  {
	temp , err := model.DbtoStruct(username)
	if err != nil{
		return "" , err} // better error handling here. should never happen
	bt , err  := json.Marshal(temp)
	if err != nil{
		return "" , err}
	return string(bt) , nil

}

