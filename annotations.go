package core

// Annotation is a concept in v2fly. This struct is only for documentation. It is not used anywhere.
// Annotations begin with "v2fly:" in comment, as metadata of functions or types.
type Annotation struct {
	// API is for types or functions that can be used in other libs. Possible values are:
	//
	// * v2fly:api:beta for types or functions that are ready for use, but maybe changed in the future.
	// * v2fly:api:stable for types or functions with guarantee of backward compatibility.
	// * v2fly:api:deprecated for types or functions that should not be used anymore.
	//
	// Types or functions without api annotation should not be used externally.
	API string
}
