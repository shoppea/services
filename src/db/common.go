package db

func DBInsert(obj interface{}) (err error){
	err = SharedConnection().Create(obj).Error
	if err != nil {
		return
	}
	return
}
