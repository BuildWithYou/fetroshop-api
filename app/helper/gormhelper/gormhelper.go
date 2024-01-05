package gormhelper

func IsNotFound(err error) bool {
	return err.Error() == "record not found"
}
