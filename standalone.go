package main

func OneFunction() string {
	return "standalone: OneFunction"
}

func Declare() string {
	return "standalone declares: " + OneFunction()
}
