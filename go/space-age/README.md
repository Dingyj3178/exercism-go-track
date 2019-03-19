# Space Age

Given an age in seconds, calculate how old someone would be on:

   - Earth: orbital period 365.25 Earth days, or 31557600 seconds
   - Mercury: orbital period 0.2408467 Earth years
   - Venus: orbital period 0.61519726 Earth years
   - Mars: orbital period 1.8808158 Earth years
   - Jupiter: orbital period 11.862615 Earth years
   - Saturn: orbital period 29.447498 Earth years
   - Uranus: orbital period 84.016846 Earth years
   - Neptune: orbital period 164.79132 Earth years

So if you were told someone were 1,000,000,000 seconds old, you should
be able to say that they're 31.69 Earth-years old.

If you're wondering why Pluto didn't make the cut, go watch [this
youtube video](http://www.youtube.com/watch?v=Z_2gbGXzFbs).

## No Stub

This may be the first Go track exercise you encounter without a stub: a
pre-existing `space_age.go` file for your solution. You may not see stubs in
the future and should begin to get comfortable with creating your own Go files
for your solutions.

One way to figure out what the function signature(s) you would need is to look
at the corresponding \*\_test.go file. It will show you what the package level
functions(s) should be that the test will use to verify the solution.


## Running the tests

To run the tests run the command `go test` from within the exercise directory.

If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem`
flags:

    go test -v --bench . --benchmem

Keep in mind that each reviewer will run benchmarks on a different machine, with
different specs, so the results from these benchmark tests may vary.

## Further information

For more detailed information about the Go track, including how to get help if
you're having trouble, please visit the exercism.io [Go language page](http://exercism.io/languages/go/resources).

## Source

Partially inspired by Chapter 1 in Chris Pine's online Learn to Program tutorial. [http://pine.fm/LearnToProgram/?Chapter=01](http://pine.fm/LearnToProgram/?Chapter=01)

## Submitting Incomplete Solutions
It's possible to submit an incomplete solution so you can see how others have completed the exercise.

## Note 

自己写的时候只想到了用Case
```golang
func Age(s float64, p Planet) float64 {
	sp := s / 31557600
	switch p {
	case "Earth":

	case "Mercury":
		sp = sp / 0.2408467

	case "Venus":
		sp = sp / 0.61519726

	case "Mars":
		sp = sp / 1.8808158

	case "Jupiter":
		sp = sp / 11.862615

	case "Saturn":
		sp = sp / 29.447498

	case "Uranus":
		sp = sp / 84.016846

	case "Neptune":
		sp = sp / 164.79132
	}
	return sp
}
```

但是这里也可以用MAP：
```golang
var secondsInOneEarthYear = 31557600.0
var orbitalPeriodsInEarthYears = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}
type Planet string

func Age(seconds float64, planet Planet) float64 {

	secondsInAYearOnPlanet := secondsInOneEarthYear * orbitalPeriodsInEarthYears[planet]

	years := seconds / secondsInAYearOnPlanet

	return years

}
```