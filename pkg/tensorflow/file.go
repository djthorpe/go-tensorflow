package tensorflow

///////////////////////////////////////////////////////////////////////////////
// TYPES

type File struct {
	*Tensor
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

/*
func NewFromFile(ctx *Scope, path string) File {
	scope := NewScope(ctx)
	file := op.ReadFile(scope.subScope("ReadFile"), op.Const(scope.subScope("filename"), path))
	tensor := NewTensor(ctx, Output(file))
	return File{tensor}
}

*/
