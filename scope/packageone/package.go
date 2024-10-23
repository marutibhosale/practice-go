package packagetwo

var privateVar = "private variable"                       // start with lower case cant export outside package
var PublicVar = "public variables access outside package" // can import outside as start with Caps

func notExported() { // cant import outside package

}

func Exported() { // can export outside of packge

}
