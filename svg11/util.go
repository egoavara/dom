package svg11

type FuncStyle struct {
	name  string
	args  [][]string

}

func NewFuncStyle(name string) *FuncStyle {
	return &FuncStyle{
		name:name,
	}
}
func (s *FuncStyle) List(str string) {
	s.args = append(s.args,  []string{str})
}
func (s *FuncStyle) Map(k, v string) {
	s.args = append(s.args, []string{k, v})
}
func (s *FuncStyle) Build() string {
	intext := ""
	for _, v := range s.args {
		switch len(v){
		case 1:
			intext += v[0] + ", "
		case 2:
			intext += v[0] + " : " + v[1] + ", "
		}

	}
	if len(intext) >= 2 {
		intext = intext[:len(intext)-2]
	}
	return s.name + "(" + intext + ")"
}

type ListStyle struct {
	args []string
}

func NewListStyle() *ListStyle {
	return &ListStyle{}
}
func (s *ListStyle) Append(str string) {
	s.args = append(s.args, str)
}
func (s *ListStyle) SkipMark() {
	s.args = append(s.args, "...")
}
func (s *ListStyle) Length() int {
	return len(s.args)
}
func (s *ListStyle) Build() string {
	intext := ""
	for _, v := range s.args {
		intext += v + ", "
	}
	if len(intext) >= 2 {
		intext = intext[:len(intext)-2]
	}
	return "[" + intext + "]"
}

type TupleStyle struct {
	args []string
}

func NewTupleStyle() *TupleStyle {
	return &TupleStyle{}
}
func (s *TupleStyle) Append(str string) {
	s.args = append(s.args, str)
}
func (s *TupleStyle) Build() string {
	intext := ""
	for _, v := range s.args {
		intext += v + ", "
	}
	if len(intext) >= 2 {
		intext = intext[:len(intext)-2]
	}
	return "(" + intext + ")"
}