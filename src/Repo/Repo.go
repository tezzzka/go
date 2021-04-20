package Repo
import (
	"db"
	
)
func CreateRecord() bool {
	//Json's way is here.
	db.JsonStructToMongoDB()
	return true
}
func GetAllItems() string {
	return db.GetAllRecords()
}
func GetById(id int) string {
	return db.GetItem(id)
}
func DelItem(id int) {
	db.DelById(id)
}