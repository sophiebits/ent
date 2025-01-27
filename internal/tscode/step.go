package tscode

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/lolopinto/ent/internal/action"
	"github.com/lolopinto/ent/internal/codegen"
	"github.com/lolopinto/ent/internal/codepath"
	"github.com/lolopinto/ent/internal/edge"
	"github.com/lolopinto/ent/internal/file"
	"github.com/lolopinto/ent/internal/schema"
	"github.com/lolopinto/ent/internal/schema/enum"
	"github.com/lolopinto/ent/internal/syncerr"
	"github.com/lolopinto/ent/internal/tsimport"
	"github.com/lolopinto/ent/internal/util"
)

type Step struct {
	m        sync.Mutex
	nodeType []enum.Data
	edgeType []enum.Data
}

func (s *Step) Name() string {
	return "codegen"
}

var nodeType = regexp.MustCompile(`(\w+)Type`)

func (s *Step) ProcessData(data *codegen.Data) error {
	var wg sync.WaitGroup
	wg.Add(len(data.Schema.Nodes))
	var serr syncerr.Error

	for key := range data.Schema.Nodes {
		go func(key string) {
			defer wg.Done()

			info := data.Schema.Nodes[key]
			nodeData := info.NodeData

			if err := s.accumulateConsts(nodeData); err != nil {
				serr.Append(err)
				return
			}

			if !info.ShouldCodegen {
				return
			}

			if nodeData.PackageName == "" {
				serr.Append(fmt.Errorf("invalid node with no package"))
				return
			}

			if err := writeBaseModelFile(nodeData, data.CodePath); err != nil {
				serr.Append(err)
				return
			}
			if err := writeEntFile(nodeData, data.CodePath); err != nil {
				serr.Append(err)
				return
			}

			if len(nodeData.ActionInfo.Actions) == 0 {
				return
			}

			if err := writeBuilderFile(nodeData, data.CodePath); err != nil {
				serr.Append(err)
			}

			// write all the actions concurrently
			var actionsWg sync.WaitGroup
			actionsWg.Add(len(nodeData.ActionInfo.Actions))
			for idx := range nodeData.ActionInfo.Actions {
				go func(idx int) {
					defer actionsWg.Done()

					action := nodeData.ActionInfo.Actions[idx]
					if err := writeBaseActionFile(nodeData, data.CodePath, action); err != nil {
						serr.Append(err)
					}

					if err := writeActionFile(nodeData, data.CodePath, action); err != nil {
						serr.Append(err)
					}

				}(idx)
			}
			actionsWg.Wait()

			// write base edge file for all the edges and then eventually one per edge...
			if !nodeData.EdgeInfo.HasConnectionEdges() {
				return
			}

			if err := writeBaseQueryFile(data.Schema, nodeData, data.CodePath); err != nil {
				serr.Append(err)
			}

			var edgesWg sync.WaitGroup
			edgesWg.Add(len(nodeData.EdgeInfo.Associations))

			for idx := range nodeData.EdgeInfo.Associations {
				go func(idx int) {
					defer edgesWg.Done()

					edge := nodeData.EdgeInfo.Associations[idx]

					if err := writeAssocEdgeQueryFile(data.Schema, nodeData, edge, data.CodePath); err != nil {
						serr.Append(err)
					}
				}(idx)
			}

			// edges with IndexLoaderFactory
			edges := nodeData.EdgeInfo.GetEdgesForIndexLoader()
			edgesWg.Add(len(edges))
			for idx := range edges {
				go func(idx int) {
					defer edgesWg.Done()

					edge := edges[idx]

					if err := writeCustomEdgeQueryFile(data.Schema, nodeData, edge, data.CodePath); err != nil {
						serr.Append(err)
					}
				}(idx)
			}
			edgesWg.Wait()
		}(key)
	}

	wg.Add(len(data.Schema.Enums))
	for key := range data.Schema.Enums {
		go func(key string) {
			defer wg.Done()

			info := data.Schema.Enums[key]

			// only lookup table enums get their own files
			if !info.LookupTableEnum() {
				return
			}

			serr.Append(writeEnumFile(info, data.CodePath))
		}(key)
	}
	wg.Wait()
	if err := serr.Err(); err != nil {
		return err
	}
	// sort data so that the enum is stable
	sort.Slice(s.nodeType, func(i, j int) bool {
		return s.nodeType[i].Name < s.nodeType[j].Name
	})
	sort.Slice(s.edgeType, func(i, j int) bool {
		return s.edgeType[i].Name < s.edgeType[j].Name
	})
	funcs := []func() error{
		func() error {
			return writeConstFile(s.nodeType, s.edgeType)
		},
		func() error {
			return writeInternalEntFile(data.Schema, data.CodePath)
		},
		func() error {
			return writeEntIndexFile()
		},
		func() error {
			return writeLoadAnyFile(s.nodeType, data.CodePath)
		},
	}

	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Step) addNodeType(name, value, comment string, nodeData *schema.NodeData) {
	s.m.Lock()
	defer s.m.Unlock()
	s.nodeType = append(s.nodeType, enum.Data{
		Name:    name,
		Value:   value,
		Comment: comment,
	})
}

func (s *Step) addEdgeType(name, value, comment string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.edgeType = append(s.edgeType, enum.Data{
		Name:    name,
		Value:   value,
		Comment: comment,
	})
}

// take what exists for go and convert it to typescript format
// should probably fix this at some point upstream
func (s *Step) accumulateConsts(nodeData *schema.NodeData) error {
	for key, group := range nodeData.ConstantGroups {
		if key != "ent.NodeType" && key != "ent.EdgeType" {
			continue
		}
		constType := strings.Split(key, ".")[1]

		for _, constant := range group.Constants {
			switch constType {
			case "NodeType":
				match := nodeType.FindStringSubmatch(constant.ConstName)
				if len(match) != 2 {
					return fmt.Errorf("%s is not a valid node type", constant.ConstName)
				}
				comment := strings.ReplaceAll(constant.Comment, constant.ConstName, match[1])

				s.addNodeType(match[1], constant.ConstValue, comment, nodeData)
				break

			case "EdgeType":
				constName, err := edge.TsEdgeConst(constant.ConstName)
				if err != nil {
					return err
				}
				comment := strings.ReplaceAll(constant.Comment, constant.ConstName, constName)

				s.addEdgeType(constName, constant.ConstValue, comment)
				break
			}
		}
	}
	return nil
}

var _ codegen.Step = &Step{}

// todo standardize this? same as in internal/code
type nodeTemplateCodePath struct {
	NodeData *schema.NodeData
	CodePath *codegen.CodePath
	Package  *codegen.ImportPackage
	Imports  []schema.ImportPath
}

func getFilePathForBaseModelFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/generated/%s_base.ts", nodeData.PackageName)
}

func getFilePathForModelFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/%s.ts", nodeData.PackageName)
}

func getFilePathForEnumFile(info *schema.EnumInfo) string {
	return fmt.Sprintf("src/ent/generated/%s.ts", strcase.ToSnake(info.Enum.Name))
}

func getFilePathForBaseQueryFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/generated/%s_query_base.ts", nodeData.PackageName)
}

func getFilePathForAssocEdgeQueryFile(nodeData *schema.NodeData, e *edge.AssociationEdge) string {
	return fmt.Sprintf(
		"src/ent/%s/query/%s.ts",
		nodeData.PackageName,
		strcase.ToSnake(e.TsEdgeQueryName()),
	)
}

func getFilePathForCustomEdgeQueryFile(nodeData *schema.NodeData, e edge.ConnectionEdge) string {
	return fmt.Sprintf(
		"src/ent/%s/query/%s.ts",
		nodeData.PackageName,
		strcase.ToSnake(e.TsEdgeQueryName()),
	)
}

func getImportPathForAssocEdgeQueryFile(nodeData *schema.NodeData, e *edge.AssociationEdge) string {
	return fmt.Sprintf(
		"src/ent/%s/query/%s",
		nodeData.PackageName,
		strcase.ToSnake(e.TsEdgeQueryName()),
	)
}

func getImportPathForCustomEdgeQueryFile(nodeData *schema.NodeData, e edge.ConnectionEdge) string {
	return fmt.Sprintf(
		"src/ent/%s/query/%s",
		nodeData.PackageName,
		strcase.ToSnake(e.TsEdgeQueryName()),
	)
}

func getImportPathForEnumFile(info *schema.EnumInfo) string {
	return fmt.Sprintf("src/ent/generated/%s", strcase.ToSnake(info.Enum.Name))
}

func getImportPathForModelFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/%s", nodeData.PackageName)
}

func getImportPathForBaseModelFile(packageName string) string {
	return fmt.Sprintf("src/ent/generated/%s_base", packageName)
}

func getImportPathForBaseQueryFile(packageName string) string {
	return fmt.Sprintf("src/ent/generated/%s_query_base", packageName)
}

func getFilePathForConstFile() string {
	return fmt.Sprintf("src/ent/const.ts")
}

func getFilePathForLoadAnyFile() string {
	return fmt.Sprintf("src/ent/loadAny.ts")
}

func getFilePathForBuilderFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/%s/actions/%s_builder.ts", nodeData.PackageName, nodeData.PackageName)
}

func getImportPathForBuilderFile(nodeData *schema.NodeData) string {
	return fmt.Sprintf("src/ent/%s/actions/%s_builder", nodeData.PackageName, nodeData.PackageName)
}

func getFilePathForActionBaseFile(nodeData *schema.NodeData, a action.Action) string {
	fileName := strcase.ToSnake(a.GetActionName())
	return fmt.Sprintf("src/ent/%s/actions/generated/%s_base.ts", nodeData.PackageName, fileName)
}

func getImportPathForActionBaseFile(nodeData *schema.NodeData, a action.Action) string {
	fileName := strcase.ToSnake(a.GetActionName())
	return fmt.Sprintf("src/ent/%s/actions/generated/%s_base", nodeData.PackageName, fileName)
}

func getFilePathForActionFile(nodeData *schema.NodeData, a action.Action) string {
	fileName := strcase.ToSnake(a.GetActionName())
	return fmt.Sprintf("src/ent/%s/actions/%s.ts", nodeData.PackageName, fileName)
}

func writeBaseModelFile(nodeData *schema.NodeData, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: nodeTemplateCodePath{
			NodeData: nodeData,
			CodePath: codePathInfo,
			Package:  codePathInfo.GetImportPackage(),
		},
		CreateDirIfNeeded:  true,
		AbsPathToTemplate:  util.GetAbsolutePath("base.tmpl"),
		TemplateName:       "base.tmpl",
		OtherTemplateFiles: []string{util.GetAbsolutePath("../schema/enum/enum.tmpl")},
		PathToFile:         getFilePathForBaseModelFile(nodeData),
		FormatSource:       true,
		TsImports:          imps,
		FuncMap:            imps.FuncMap(),
	})
}

func writeEntFile(nodeData *schema.NodeData, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()
	return file.Write(&file.TemplatedBasedFileWriter{
		Data: nodeTemplateCodePath{
			NodeData: nodeData,
			CodePath: codePathInfo,
			Package:  codePathInfo.GetImportPackage(),
		},
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("ent.tmpl"),
		TemplateName:      "ent.tmpl",
		PathToFile:        getFilePathForModelFile(nodeData),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
		EditableCode:      true,
		// only write this file once.
		// TODO need a flag to overwrite this later.
	}, file.WriteOnce())
}

func writeEnumFile(enumInfo *schema.EnumInfo, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()
	return file.Write(&file.TemplatedBasedFileWriter{
		// enum file can be rendered on its own so just render it
		Data:              enumInfo.Enum,
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("../schema/enum/enum.tmpl"),
		TemplateName:      "enum.tmpl",
		PathToFile:        getFilePathForEnumFile(enumInfo),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
	})
}

func writeBaseQueryFile(s *schema.Schema, nodeData *schema.NodeData, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: struct {
			NodeData *schema.NodeData
			Schema   *schema.Schema
			Package  *codegen.ImportPackage
		}{
			Schema:   s,
			NodeData: nodeData,
			Package:  codePathInfo.GetImportPackage(),
		},
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("ent_query_base.tmpl"),
		TemplateName:      "ent_query_base.tmpl",
		PathToFile:        getFilePathForBaseQueryFile(nodeData),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
	})
}

func writeAssocEdgeQueryFile(s *schema.Schema, nodeData *schema.NodeData, e *edge.AssociationEdge, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: struct {
			Edge    *edge.AssociationEdge
			Package *codegen.ImportPackage
		}{
			Edge:    e,
			Package: codePathInfo.GetImportPackage(),
		},
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("assoc_ent_query.tmpl"),
		TemplateName:      "assoc_ent_query.tmpl",
		PathToFile:        getFilePathForAssocEdgeQueryFile(nodeData, e),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
		EditableCode:      true,
	}, file.WriteOnce())
}

func writeCustomEdgeQueryFile(s *schema.Schema, nodeData *schema.NodeData, e edge.ConnectionEdge, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: struct {
			Package         *codegen.ImportPackage
			TsEdgeQueryName string
		}{
			Package:         codePathInfo.GetImportPackage(),
			TsEdgeQueryName: e.TsEdgeQueryName(),
		},
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("custom_ent_query.tmpl"),
		TemplateName:      "custom_ent_query.tmpl",
		PathToFile:        getFilePathForCustomEdgeQueryFile(nodeData, e),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
		EditableCode:      true,
	}, file.WriteOnce())
}

func writeConstFile(nodeData []enum.Data, edgeData []enum.Data) error {
	// sort data so that the enum is stable
	sort.Slice(nodeData, func(i, j int) bool {
		return nodeData[i].Name < nodeData[j].Name
	})
	sort.Slice(edgeData, func(i, j int) bool {
		return edgeData[i].Name < edgeData[j].Name
	})

	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: struct {
			NodeType enum.Enum
			EdgeType enum.Enum
		}{
			enum.Enum{
				Name:   "NodeType",
				Values: nodeData,
			},
			enum.Enum{
				Name:   "EdgeType",
				Values: edgeData,
			},
		},
		AbsPathToTemplate: util.GetAbsolutePath("const.tmpl"),
		TemplateName:      "const.tmpl",
		OtherTemplateFiles: []string{
			util.GetAbsolutePath("../schema/enum/enum.tmpl"),
		},
		PathToFile:   getFilePathForConstFile(),
		FormatSource: true,
		TsImports:    imps,
		FuncMap:      imps.FuncMap(),
	})
}

func writeLoadAnyFile(nodeData []enum.Data, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: struct {
			NodeData []enum.Data
			Package  *codegen.ImportPackage
		}{
			nodeData,
			codePathInfo.GetImportPackage(),
		},
		AbsPathToTemplate: util.GetAbsolutePath("loadAny.tmpl"),
		TemplateName:      "loadAny.tmpl",
		PathToFile:        getFilePathForLoadAnyFile(),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
	})
}

func getSortedInternalEntFileLines(s *schema.Schema) []string {
	lines := []string{
		"src/ent/const",
	}

	append2 := func(list *[]string, str string) {
		*list = append(*list, str)
	}

	var baseFiles []string
	for _, info := range s.Nodes {
		append2(&baseFiles, getImportPathForBaseModelFile(info.NodeData.PackageName))
	}

	var entFiles []string
	for _, info := range s.Nodes {
		append2(&entFiles, getImportPathForModelFile(info.NodeData))
	}

	var enums []string
	for _, enum := range s.Enums {
		if enum.LookupTableEnum() {
			append2(&enums, getImportPathForEnumFile(enum))
		}
	}

	var baseQueryFiles []string
	var queryFiles []string
	for _, info := range s.Nodes {
		hasBaseQueryFile := false
		if len(info.NodeData.EdgeInfo.Associations) != 0 {
			hasBaseQueryFile = true
			for _, edge := range info.NodeData.EdgeInfo.Associations {
				append2(&queryFiles, getImportPathForAssocEdgeQueryFile(info.NodeData, edge))
			}
		}

		for _, edge := range info.NodeData.EdgeInfo.GetEdgesForIndexLoader() {
			hasBaseQueryFile = true
			append2(&queryFiles, getImportPathForCustomEdgeQueryFile(info.NodeData, edge))
		}

		if hasBaseQueryFile {
			append2(&baseQueryFiles, getImportPathForBaseQueryFile(info.NodeData.PackageName))
		}
	}

	// bucket each group, make sure it's sorted within each bucket so that it doesn't randomly change
	// and make sure we get the order we want
	list := [][]string{
		baseFiles,
		entFiles,
		enums,
		baseQueryFiles,
		queryFiles,
	}
	for _, l := range list {
		sort.Strings(l)
		lines = append(lines, l...)
	}
	return lines
}

func writeInternalEntFile(s *schema.Schema, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data:              getSortedInternalEntFileLines(s),
		AbsPathToTemplate: util.GetAbsolutePath("internal.tmpl"),
		TemplateName:      "internal.tmpl",
		PathToFile:        codepath.GetFilePathForInternalFile(),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
	})
}

func writeEntIndexFile() error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		AbsPathToTemplate: util.GetAbsolutePath("index.tmpl"),
		TemplateName:      "index.tmpl",
		PathToFile:        codepath.GetFilePathForEntIndexFile(),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           imps.FuncMap(),
	})
}

func writeBuilderFile(nodeData *schema.NodeData, codePathInfo *codegen.CodePath) error {
	imps := tsimport.NewImports()

	return file.Write(&file.TemplatedBasedFileWriter{
		Data: nodeTemplateCodePath{
			NodeData: nodeData,
			CodePath: codePathInfo,
			Package:  codePathInfo.GetImportPackage(),
			Imports:  nodeData.GetImportsForBaseFile(),
		},
		CreateDirIfNeeded: true,
		AbsPathToTemplate: util.GetAbsolutePath("builder.tmpl"),
		TemplateName:      "builder.tmpl",
		PathToFile:        getFilePathForBuilderFile(nodeData),
		FormatSource:      true,
		TsImports:         imps,
		FuncMap:           getBuilderFuncs(imps),
	})
}

func getBuilderFuncs(imps *tsimport.Imports) template.FuncMap {
	m := imps.FuncMap()
	m["edgeInfos"] = action.GetEdgesFromEdges

	return m
}
