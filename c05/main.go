package main

type customError struct {
	msg string
	num int
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	println(err)
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

//Программа выведет error.
//err хранит в себе значение формата customError, которое хранит в себе (указатель на структуру, значение)
//указатель на структуру != nil, получаем логичное err != nil.
