// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFromSet() Foo {
	foo := provideFoo()
	return foo
}

func injectFromNestedSet() Foo {
	foo := provideFoo()
	return foo
}

func injectFromSetWithDuplicateBindings() Foo {
	foo := provideFoo()
	return foo
}

func injectDuplicateInterface() Bar {
	bar := provideBar()
	return bar
}
