package main

type Person interface {
	GetName() string
}

type Person2 interface {
	GetName() string
}

type PersonImp struct {
	Name string
}

func (p PersonImp) GetName() string {
	return p.Name
}

func InterfacePerson(p Person) Person {
	p.GetName()

	return p
}

func InterfacePerson2(p Person2) Person2 {
	p.GetName()

	return p
}

func GenericPerson[T Person](p T) T {
	p.GetName()

	return p
}

func main() {
	var p Person = PersonImp{Name: "Gaston"}

	InterfacePerson(p)

	InterfacePerson2(p) // runtime.assertI2I

	//func assertI2I(inter *interfacetype, tab *itab) *itab {
	//	if tab == nil {
	//	// explicit conversions require non-nil interface value.
	//	panic(&TypeAssertionError{nil, nil, &inter.typ, ""})
	//}
	//	if tab.inter == inter {
	//	return tab
	//}
	//	return getitab(inter, tab._type, false)
	//}

	var p1 Person = PersonImp{Name: "Gaston"}

	GenericPerson[Person](p1) // equivalent to InterfacePerson2(p)
}
