package main

import (
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/iancoleman/strcase"
	"github.com/lolopinto/ent/internal/action"
	intcodegen "github.com/lolopinto/ent/internal/codegen"
	 "github.com/lolopinto/ent/internal/schema"

	"github.com/pkg/errors"
)

// inspired by resolvergen from gqlgen
type entGraphQLResolverPlugin struct {
	nodes    schema.NodeMapInfo
	codePath *intcodegen.CodePath
}

var _ plugin.CodeGenerator = &entGraphQLResolverPlugin{}

func (p *entGraphQLResolverPlugin) Name() string {
	return "ent_graphql_resolver"
}

func (p *entGraphQLResolverPlugin) castToString(field *codegen.Field) bool {
	objName := field.Object.Name

	nodeData := p.nodes.GetNodeDataFromGraphQLName(objName)
	if nodeData == nil {
		return false
	}

	entField := nodeData.GetFieldByName(field.GoFieldName)
	//spew.Dump(field, entField)
	if entField == nil {
		return false
	}

	// todo need to make this generic enough in the future
	// castToSomething
	// then call a getCastToBlah code
	return entField.GetGraphQLTypeForField() == "String!"
}

func (p *entGraphQLResolverPlugin) loadObjectFromContext(field *codegen.Field) bool {
	if len(field.Args) != 1 {
		return false
	}
	firstArg := field.Args[0]

	// for now just assume always id. TODO?
	if firstArg.Name != "id" {
		return false
	}

	// field.Object.Name = Query
	// field.GoFieldName = Contact/User etc.
	nodeData := p.nodes.GetNodeDataFromGraphQLName(field.GoFieldName)
	return nodeData != nil
}

func (p *entGraphQLResolverPlugin) fieldEdge(field *codegen.Field) bool {
	objName := field.Object.Name

	nodeData := p.nodes.GetNodeDataFromGraphQLName(objName)
	if nodeData == nil {
		return false
	}

	edge := nodeData.GetFieldEdgeByName(field.GoFieldName)
	return edge != nil
}

func (p *entGraphQLResolverPlugin) pluralEdge(field *codegen.Field) bool {
	objName := field.Object.Name

	nodeData := p.nodes.GetNodeDataFromGraphQLName(objName)
	if nodeData == nil {
		return false
	}

	fkeyEdge := nodeData.GetForeignKeyEdgeByName(field.GoFieldName)
	if fkeyEdge != nil {
		return true
	}

	assocEdge := nodeData.GetAssociationEdgeByName(field.GoFieldName)
	return assocEdge != nil
}

func (p *entGraphQLResolverPlugin) mutation(field *codegen.Field) action.Action {
	if field.Object.Name != "Mutation" {
		return nil
	}
	//	spew.Dump(field)

	//	spew.Dump(field.TypeReference.Definition.Name)

	// Name -> userCreate, GoFieldName -> UserCreate
	//	spew.Dump(field.GoFieldName, field.GoReceiverName)
	//	spew.Dump(field)
	return p.nodes.GetActionFromGraphQLName(field.Name)
	//	spew.Dump(field.Name, action)
}

func (p *entGraphQLResolverPlugin) getActionPath(a action.Action) string {
	path, err := strconv.Unquote(p.codePath.PathToModels)
	if err != nil {
		panic(errors.Wrap(err, "could not unquote path"))
	}

	// TODO these names are broken. fix it
	// this is equivalent to nodeData.PackageName which is what we're using when
	// generating the file
	// need to make sure we support contact_date etc
	// to be consistent with writeActionFile
	actionPath := filepath.Join(path, strcase.ToSnake(a.GetNodeInfo().Node), "action")
	return actionPath
}

// ResolverBuild is the object passed to the template to generate the graphql code
type ResolverBuild struct {
	*codegen.Data

	PackageName  string
	ResolverType string
}

func (p *entGraphQLResolverPlugin) GenerateCode(data *codegen.Data) error {
	resolverBuild := &ResolverBuild{
		Data:         data,
		ResolverType: "Resolver",
	}

	return templates.Render(templates.Options{
		PackageName:     "graphql",
		Filename:        "graphql/resolver.go",
		Data:            resolverBuild,
		GeneratedHeader: true,
		Template:        readTemplateFile("ent_graphql_resolver.gotmpl"),
		Funcs: template.FuncMap{
			"castToString":          p.castToString,
			"loadObjectFromContext": p.loadObjectFromContext,
			"fieldEdge":             p.fieldEdge,
			"pluralEdge":            p.pluralEdge,
			"mutation":              p.mutation,
			"actionMethodName":      action.GetActionMethodName,
			"actionFields":          action.GetFields,
			"actionPath":            p.getActionPath,
		},
	})
}

func newGraphQLResolverPlugin(data *codegenData) plugin.Plugin {
	return &entGraphQLResolverPlugin{
		nodes:    data.allNodes,
		codePath: data.codePath,
	}
}
