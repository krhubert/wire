// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func inject(
	name BarName,
	opts *FooOptions,
) *Bar {
	foo := provideFoo(
		opts,
	)
	bar := provideBar(
		foo,
		name,
	)
	return bar
}

func injectBarService(
	name BarName,
	opts *FakeBarService,
) *FooBar {
	fooOptions := provideFooOptions()
	foo := provideFoo(
		fooOptions,
	)
	fooBar := &FooBar{
		BarService: opts,
		Foo:        foo,
	}
	return fooBar
}

func injectFooBarService(
	name BarName,
	opts *FooOptions,
	bar *FakeBarService,
) *FooBar {
	foo := provideFoo(
		opts,
	)
	fooBar := &FooBar{
		BarService: bar,
		Foo:        foo,
	}
	return fooBar
}

func injectNone(
	name BarName,
	foo Foo,
	bar *FakeBarService,
) *FooBar {
	fooBar := &FooBar{
		BarService: bar,
		Foo:        foo,
	}
	return fooBar
}
