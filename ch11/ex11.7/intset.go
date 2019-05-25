package intset

type IntSet interface {
	Has(x int) bool
	Add(x int)
	UnionWith(t IntSet)
	String() string
	Len() int
	Remove(x int)
	Clear()
}
