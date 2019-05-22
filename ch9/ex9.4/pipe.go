package main

func pipe(stage int) (in, out chan int) {
	out = make(chan int)
	first := out

	for i := 0; i < stage; i++ {
		in = out
		out = make(chan int)

		go func(in, out chan int) {
			for v := range in {
				out <- v
			}

			close(out)
		}(in, out)
	}

	return first, out

}
