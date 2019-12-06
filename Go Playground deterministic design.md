The environment in which the programs in the Playground are executed is deterministic.

So the [Go Playground](https://play.golang.org/) by design does not allow to create truly pseudo-random outputs.

This is done intentionally for purpose of caching results to minimize CPU/memory usage for consequent runs. So the engine can evaluate your program just once and serve the same cached output every time when you or anyone else run it again.

For the same purpose the start time is locked to a constant.

Read more: https://blog.golang.org/playground
