package Sessions

import ("proj/DB"
	"fmt"
	"encoding/json")

func StructToJson(model DB.DBMarshall, username string) (string , error)  {
	fmt.Println("model " , model)
	temp , err := model.DbtoStruct(username)
	fmt.Println(temp)
	if err != nil{
		return "" , err} // better error handling here. should never happen
	bt , err  := json.Marshal(temp)
	if err != nil{
		return "" , err}
	return string(bt) , nil

}

