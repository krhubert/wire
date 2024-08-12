// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectedFoo() string {
	s := provideS()
	string2 := s.Foo
	return string2
}

func injectedFooBar() Out {
	s := provideS()
	string2 := s.Foo
	int2 := s.Bar
	out := Out{
		Foo: string2,
		Bar: int2,
	}
	return out
}

func injectedFooUnusedBar() Unused {
	s := provideS()
	string2 := s.Foo
	unused := Unused{
		Foo: string2,
	}
	return unused
}
