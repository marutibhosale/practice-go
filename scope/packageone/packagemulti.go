package packageone // in same folder all files  we need to give same package name to all files at top
func Test() {
	println(privateVar) // taken from file paackge.go in same folder private variables
	notExported()
	println(PublicVar) // we can use both public and private var and func in same package
	Exported()
}
