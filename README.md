This is the modified implementation of Needelman-Wunsch algorithm
[Link](http://en.wikipedia.org/wiki/Needleman-Wunsch_algorithm)
that can be found [here](https://github.com/aebruno/nwalgo). 

## Installation 
```golang
go get -u github.com/MichaelLeachim/wunsch
```

## Reasons behind the implementation

There are three reasons behind this implentation:

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

### Item interface

It must implement the Id method. Item Id is an equivalent of a `rune` in a `string` item. 

```go
type Item interface {
	Id() string
}
```

Example of String implementation

```go

type StringItem string

func (i StringItem) Id() string {
	return string(i)
}

```

### Alignment process

```go 
TODO: write up on usage 
```





