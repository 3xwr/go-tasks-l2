package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()

	for n := range ch {
		println(n)
	}
}

//Мы получим deadlock, т.к. range по channel'у будет ждать значения до его закрытия,
//которого не происходит в исходном коде.
