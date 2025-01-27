package codepath

import "fmt"

// Package refers to the name of the package
const Package = "@lolopinto/ent"

// ActionPackage refers to the name of the action package
const ActionPackage = Package + "/action"

// AuthPackage refers to the name of the auth package where ent-auth stuff is
const AuthPackage = Package + "/auth"

// SchemaPackage refers to the name of the schema package
const SchemaPackage = Package + "/schema"

// GraphQLPackage refers to the name of the graphql package
const GraphQLPackage = Package + "/graphql"

func GetFilePathForInternalFile() string {
	return fmt.Sprintf("src/ent/internal.ts")
}

func GetFilePathForEntIndexFile() string {
	return fmt.Sprintf("src/ent/index.ts")
}

func GetInternalImportPath() string {
	return "src/ent/internal"
}

func GetExternalImportPath() string {
	return "src/ent/"
}

func GetFilePathForInternalGQLFile() string {
	return fmt.Sprintf("src/graphql/resolvers/internal.ts")
}

func GetFilePathForExternalGQLFile() string {
	return fmt.Sprintf("src/graphql/resolvers/index.ts")
}

func GetImportPathForInternalGQLFile() string {
	return fmt.Sprintf("src/graphql/resolvers/internal")
}

func GetImportPathForExternalGQLFile() string {
	return fmt.Sprintf("src/graphql/resolvers/")
}
