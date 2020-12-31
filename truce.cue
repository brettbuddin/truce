apis: [...#API]

#API: {
	version: =~"^[0-9]+$" // API Major version
	transports: [...#Transport]
	functions: [...#Function]
	types: [...#Type]
}

#Transport: {
	type: string
	versions: [...string]
	prefix: string
}

#Function: {
	name: =~"^[A-Z][a-zA-Z]*$" // function names are capitalised
	arguments: [...#Field]
	return: #Field
	transports: [...#FunctionTransport]
}

#Type: {
	name: =~"^[A-Z][a-zA-Z]*$" // type names are capitalised
	fields: [...#Field]
}

#primitives: ["bool", "int", "float", "byte", "string"]
#slices: [ for x in #primitives {"[]\(x)"}]
#all: #primitives + #slices

#Field: {
	name: string
	type: or(#all) | =~"^[A-Z][a-zA-Z]*$" | =~"^[[]][A-Z][a-zA-Z]*?" // can be primitive, CustomType, []primitive or []CustomType. 
}

#FunctionTransport: {
	type:   "http"
	path:   string
	method: "GET" | "POST" | "PUT" | "PATCH" | "OPTIONS" | "DELETE" | "HEAD" | "CONNECT" | "TRACE"
	arguments: [...#Argument]
}

#Argument: {
	name:  string
	value: "$body" | =~"^[$]path[.].+$" | =~"^[^$]?.*$" // can be a constant or $body or $path.var
}