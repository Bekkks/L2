package main

import (
	"fmt"
	"os"
)

//функция Foo возвращает ин-йс error
func Foo() error {
	var err *os.PathError = nil // здесь мы присваиваем err указатель на nil
	return err                  // возвращаем err
}

func main() {
	err := Foo()
	fmt.Println(err)        // вывод: <nil>
	fmt.Println(err == nil) //вывод: false
}

//почему так происходит ведь мы явно присвоили err значение nil
/* это не понять если не узнать как внутри устроен интерфейс в Го
он состоит из двух вещей:
	информация о типе данных - type (например: *os.PathError)
	и сами значения - data
В данном случае наш код:
	var err *os.PathError = nil
	внутри происходит так:
	type = *os.PathError
	data = nil
	то есть наш ин-йс можно представить так:
	error {
	type: *os.PathError
	data: nil

И по факту ин-йс уже не равен nil так как он хранит инфу о типе данных
В  Го ин-йс считается nil только если и тип и значение равно nil
Следовательно отсюда и вывод: значение ин-са = nil fmt.Println(err)
но он все равно не равен nil var err *os.PathError так как хранит инфу о типе
*/
