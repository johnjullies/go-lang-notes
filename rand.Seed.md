https://stackoverflow.com/questions/45268998/how-do-i-seed-random-number-generator-on-official-tour-of-go

By default rand.Intn uses the [globalRand.Intn](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L277). It's created internally, refer [here](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L236). So when you set via rand.Seed

`rand.Seed(time.Now().UTC().UnixNano())`

Then globalRand uses the new seed value.
