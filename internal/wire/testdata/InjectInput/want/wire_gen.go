// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar(
	foo Foo,
) FooBar {
	bar := provideBar()
	fooBar := provideFooBar(
		foo,
		bar,
	)
	return fooBar
}
