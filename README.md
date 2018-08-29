This is the modified implementation of [Needelman-Wunsch](http://en.wikipedia.org/wiki/Needleman-Wunsch_algorithm)  algorithm.
Original you can find [here](https://github.com/aebruno/nwalgo). 

## Installation 
```golang
go get -u github.com/MichaelLeachim/wunsch
```

## Reasons behind the implementation

There are three reasons behind this implementation:

* To allow generic (not only strings) alignment between collections
* To allow alignment of more than two sequences
* To learn more about alignment algorithms

In short, this library allows to do the following:
```golang

AlignManyString([]string{"Hello", "Hallo", "Hillo", "Hwello", "lloha", "habar"}) => 
[]string{
"H---ello-----",
"H--a-llo-----",
"H-i--llo-----",
"-----lloha---",
"--------habar",
"Hw--ello-----"}
```

### By applying these steps

* Choosing the longest sequence
* Making fingerprint by conflating all items together
* Making resulting alignment by comparing fingerprint with each sequence

## Usage 

As the library is small, all usage is obvious from the test cases, but here are 
some code snippets that might prove useful

### Item struct

The `Id` field is for comparing items between each other. 
The Index field is for reference to the external data structure. 

For example, if we wanted to use `string` as our external data structure, 
the `Id` field would store the numeric representation of a `rune`. 
And the index field would store the position of that `rune` in an 
original `string`. 


```go
type Item struct {
	Id    int
	Index int
}
```

### Alignment process

```go 
TODO: write up on usage 
```





