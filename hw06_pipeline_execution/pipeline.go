package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	prev := in
	for _, fun := range stages {
		prev = fun(prev)

		next := make(Bi)

		go func(done In, in In, next Bi) {
			defer close(next)

			for element := range in {
				select {
				case _, ok := <-done:
					if !ok {
						return
					}
				default:
					next <- element
				}
			}
		}(done, prev, next)

		prev = next
	}

	return prev
}
