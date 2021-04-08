package Parse

type Applyable interface {
	Apply(string) (bool, []ParseTree)
}
