package main

func main() {
	c := make(chan int, 1)
	for done := false; !done; {
		select {
		default:
			print(1)
			done = true
		case <-c:
			print(2)
			c = nil
		case c <- 1:
			print(3)
		}
	}
}

/*
	Key points:

	For the 1st loop step, only the operation c <- 1 is non-blocking. So the last case branch is chosen.

	For the 2nd loop step, only the operation <-c is non-blocking. So the first case branch is chosen.

	For the 3rd loop step, both the channel operations are blocking. So the default branch is chosen.
*/
