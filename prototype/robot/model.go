package robot

// Copy an interface that implements the method Clone which returns a copy of the object
type Copy interface {
	Clone() Copy
}
