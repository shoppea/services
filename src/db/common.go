package db


func DBInsert(obj interface{}) (err error){
	db := SharedConnection()
	err = db.Create(obj).Error
	if err != nil {
		return
	}
	return
}