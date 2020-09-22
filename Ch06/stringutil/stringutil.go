package stringutil

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

//for naked retrun, without taking care of the return order

func FullNameNakedreturn(f, l string) (full string, length int) {
	full = f + "" + l
	length = len(full)
	return
}
