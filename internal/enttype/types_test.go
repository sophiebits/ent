package enttype_test

import (
	"fmt"
	"go/types"
	"strconv"
	"strings"
	"testing"

	"github.com/lolopinto/ent/internal/enttype"
	"github.com/lolopinto/ent/internal/schemaparser"
	"github.com/stretchr/testify/assert"
)

type expType struct {
	db                  string
	graphql             string
	graphqlImports      []enttype.FileImport
	graphqlPanics       bool
	goTypePanics        bool
	castToMethod        string
	zeroValue           interface{}
	nullableType        enttype.Type
	nonNullableType     enttype.Type
	defaultGQLFieldName string
	elemGraphql         string
	errorType           bool
	contextType         bool
	goType              string
	tsType              string
}

func TestStringType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() string {
	return ""
	}`)

	assert.IsType(t, &enttype.StringType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Text()",
		graphql: "String!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLString"),
		},
		zeroValue:    strconv.Quote(""),
		castToMethod: "cast.ToString",
		goType:       "string",
		nullableType: &enttype.NullableStringType{},
		tsType:       "string",
	}, ret)
}

func TestEnumishType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

		import "github.com/lolopinto/ent/ent"

func f() ent.NodeType {
	return ent.NodeType("user")
}`)

	assert.IsType(t, &enttype.StringType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Text()",
		graphql: "String!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLString"),
		},
		zeroValue: strconv.Quote(""),
		// this probably doesn't work correctly in practice because strong types broken?
		// would need castToMethod plus special enum cast
		castToMethod: "cast.ToString",
		goType:       "ent.NodeType",
		nullableType: &enttype.NullableStringType{},
		tsType:       "string",
	}, ret)
}

func TestNullableStringType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() *string {
	return nil
	}`)

	assert.IsType(t, &enttype.NullableStringType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Text()",
		graphql: "String",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLString"),
		},
		zeroValue:       strconv.Quote(""),
		castToMethod:    "cast.ToNullableString",
		goType:          "*string",
		nonNullableType: &enttype.StringType{},
		tsType:          "string | null",
	}, ret)
}

func TestBoolType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() bool {
	return true
	}`)

	assert.IsType(t, &enttype.BoolType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Boolean()",
		graphql: "Boolean!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLBoolean"),
		},
		zeroValue:    "false",
		castToMethod: "cast.ToBool",
		goType:       "bool",
		nullableType: &enttype.NullableBoolType{},
		tsType:       "boolean",
	}, ret)
}

func TestNullableBoolType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() *bool {
	return nil
	}`)

	assert.IsType(t, &enttype.NullableBoolType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Boolean()",
		graphql: "Boolean",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLBoolean"),
		},
		zeroValue:       "false",
		castToMethod:    "cast.ToNullableBool",
		nonNullableType: &enttype.BoolType{},
		goType:          "*bool",
		tsType:          "boolean | null",
	}, ret)
}

func TestIDType(t *testing.T) {
	testType(t, expType{
		db:      "postgresql.UUID()",
		graphql: "ID!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLID"),
		},
		zeroValue:    "",
		castToMethod: "cast.ToUUIDString",
		nullableType: &enttype.NullableIDType{},
		tsType:       "ID",
	}, returnType{
		entType: &enttype.IDType{},
	})
}

func TestNullableIDType(t *testing.T) {
	testType(t, expType{
		db:      "postgresql.UUID()",
		graphql: "ID",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLID"),
		},
		zeroValue:       "",
		castToMethod:    "cast.ToNullableUUIDString",
		nonNullableType: &enttype.IDType{},
		tsType:          "ID | null",
	}, returnType{
		entType: &enttype.NullableIDType{},
	})
}

func TestIntegerType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() int {
	return 1
	}`)

	assert.IsType(t, &enttype.IntegerType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Integer()",
		graphql: "Int!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLInt"),
		},
		zeroValue:    "0",
		castToMethod: "cast.ToInt",
		goType:       "int",
		nullableType: &enttype.NullableIntegerType{},
		tsType:       "number",
	}, ret)
}

func TestNullableIntegerType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() *int {
	return nil
	}`)

	assert.IsType(t, &enttype.NullableIntegerType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Integer()",
		graphql: "Int",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLInt"),
		},
		zeroValue:       "0",
		castToMethod:    "cast.ToNullableInt",
		nonNullableType: &enttype.IntegerType{},
		goType:          "*int",
		tsType:          "number | null",
	}, ret)
}

func TestFloat64Type(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() float64 {
	return 1.0
	}`)
	testFloatType(t, ret, "float64")
}

func TestFloat32Type(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() float32 {
	return 1.0
	}`)
	testFloatType(t, ret, "float32")
}

func testFloatType(t *testing.T, ret returnType, goType string) {
	assert.IsType(t, &enttype.FloatType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Float()",
		graphql: "Float!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			enttype.NewGQLFileImport("GraphQLFloat"),
		},
		zeroValue:    "0.0",
		castToMethod: "cast.ToFloat",
		goType:       goType,
		nullableType: &enttype.NullableFloatType{},
		tsType:       "number",
	}, ret)
}

func TestNullableFloat64Type(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() *float64 {
	return nil
	}`)
	testNullableFloatType(t, ret, "*float64")
}

func TestNullableFloat32Type(t *testing.T) {
	ret := getTestReturnType(t, `package main 

func f() *float32 {
	return nil
	}`)
	testNullableFloatType(t, ret, "*float32")
}

func testNullableFloatType(t *testing.T, ret returnType, goType string) {
	assert.IsType(t, &enttype.NullableFloatType{}, ret.entType)
	testType(t, expType{
		db:      "sa.Float()",
		graphql: "Float",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLFloat"),
		},
		zeroValue:       "0.0",
		castToMethod:    "cast.ToNullableFloat",
		goType:          goType,
		nonNullableType: &enttype.FloatType{},
		tsType:          "number | null",
	}, ret)
}

func TestTimestampType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

	import "time"

func f() time.Time {
	return time.Time{}
	}`)

	assert.IsType(t, &enttype.TimestampType{}, ret.entType)
	testType(t, expType{
		db:      "sa.TIMESTAMP()",
		graphql: "Time!",
		graphqlImports: []enttype.FileImport{
			enttype.NewGQLFileImport("GraphQLNonNull"),
			{
				Type:       "GraphQLTime",
				ImportType: enttype.EntGraphQL,
			},
		},
		zeroValue:           "time.Time{}",
		castToMethod:        "cast.ToTime",
		nullableType:        &enttype.NullableTimestampType{},
		defaultGQLFieldName: "time",
		goType:              "time.Time",
		tsType:              "Date",
	}, ret)
}

func TestNullableTimestampType(t *testing.T) {
	ret := getTestReturnType(t, `package main 

	import "time"

func f() *time.Time {
	return nil
	}`)

	assert.IsType(t, &enttype.NullableTimestampType{}, ret.entType)
	testType(t, expType{
		db:      "sa.TIMESTAMP()",
		graphql: "Time",
		graphqlImports: []enttype.FileImport{
			{
				Type:       "GraphQLTime",
				ImportType: enttype.EntGraphQL,
			},
		},
		zeroValue:           "time.Time{}",
		castToMethod:        "cast.ToNullableTime",
		nonNullableType:     &enttype.TimestampType{},
		defaultGQLFieldName: "time",
		goType:              "*time.Time",
		tsType:              "Date | null",
	}, ret)
}

type testCase struct {
	code string
	exp  expType
	fn   func(ret *returnType, exp *expType)
}

func TestNamedType(t *testing.T) {
	defaultFn := func(ret *returnType, exp *expType) {
		namedType := ret.entType.(*enttype.NamedType)
		exp.nullableType = enttype.NewNamedType(
			namedType.GetActualType(),
			true,
			false,
		)
		exp.nonNullableType = enttype.NewNamedType(
			namedType.GetActualType(),
			false,
			true,
		)
	}

	testTestCases(
		t,
		&enttype.NamedType{},
		map[string]testCase{
			"context": {
				`package main
	
			import "context"
	
			func f() context.Context {
				return context.TODO()
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "Context!",
					defaultGQLFieldName: "context",
					contextType:         true,
					goType:              "context.Context",
					zeroValue:           "nil",
					castToMethod:        "cast.UnmarshallJSON",
				},
				defaultFn,
			},
			"error": {
				`package main
	
			func f() error {
				return nil
			}`,
				expType{
					db:            "sa.Text()",
					graphqlPanics: true,
					errorType:     true,
					goType:        "error",
					zeroValue:     "nil",
					castToMethod:  "cast.UnmarshallJSON",
				},
				defaultFn,
			},
			"models.User": {
				`package main
	
				import "github.com/lolopinto/ent/internal/test_schema/models"
	
				func f() models.User {
					// we have no reason to do this but test that default behavior does the right thing
				return models.User{}
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "User!",
					defaultGQLFieldName: "user",
					goType:              "models.User",
					zeroValue:           "models.User{}",
					castToMethod:        "cast.UnmarshallJSON",
				},
				defaultFn,
			},
		},
	)
}

func TestPointerType(t *testing.T) {
	defaultFn := func(ret *returnType, exp *expType) {
		pointerType := ret.entType.(*enttype.PointerType)
		exp.nullableType = enttype.NewPointerType(
			pointerType.GetActualType(),
			true,
			false,
		)
		exp.nonNullableType = enttype.NewPointerType(
			pointerType.GetActualType(),
			false,
			true,
		)
	}

	testTestCases(
		t, &enttype.PointerType{},
		map[string]testCase{
			"models.User": {
				`package main

	import "github.com/lolopinto/ent/internal/test_schema/models"

	func f() *models.User {
	return nil
}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "User",
					defaultGQLFieldName: "user",
					goType:              "*models.User",
					zeroValue:           "nil",
					castToMethod:        "cast.UnmarshallJSON",
				},
				defaultFn,
			},
			"stringSlice": {
				`package main
	
				func f() *[]string {
					return &[]string{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[String!]",
					goType:       "*[]string",
					zeroValue:    "nil",
					castToMethod: "cast.UnmarshallJSON",
				},
				defaultFn,
			},
			"stringPointerSlice": {
				`package main
	
				func f() *[]*string {
					return &[]*string{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[String]",
					goType:       "*[]*string",
					zeroValue:    "nil",
					castToMethod: "cast.UnmarshallJSON",
				},
				defaultFn,
			},
			"models.UserPointer": {
				`package main
	
				import "github.com/lolopinto/ent/internal/test_schema/models"
	
				func f() *[]*models.User {
				return nil
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "[User]",
					defaultGQLFieldName: "users",
					goType:              "*[]*models.User",
					zeroValue:           "nil",
					castToMethod:        "cast.UnmarshallJSON",
				},
				defaultFn,
			},
		},
	)
}

func TestSliceType(t *testing.T) {
	testTestCases(
		t,
		&enttype.SliceType{},
		map[string]testCase{
			"string": {
				`package main
	
				func f() []string {
					return []string{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[String!]!",
					elemGraphql:  "String!",
					zeroValue:    "nil",
					goType:       "[]string",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"stringPointer": {
				`package main
	
				func f() []*string {
					return []*string{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[String]!",
					elemGraphql:  "String",
					zeroValue:    "nil",
					goType:       "[]*string",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"bool": {
				`package main
	
				func f() []bool {
					return []bool{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Boolean!]!",
					elemGraphql:  "Boolean!",
					zeroValue:    "nil",
					goType:       "[]bool",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"boolPointer": {
				`package main
	
				func f() []*bool {
					return []*bool{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Boolean]!",
					elemGraphql:  "Boolean",
					zeroValue:    "nil",
					goType:       "[]*bool",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"int": {
				`package main
	
				func f() []int {
					return []int{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Int!]!",
					elemGraphql:  "Int!",
					zeroValue:    "nil",
					goType:       "[]int",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"intPointer": {
				`package main
	
				func f() []*int {
					return []*int{}
			}`,
				expType{
					db: "sa.Text()",

					graphql:      "[Int]!",
					elemGraphql:  "Int",
					zeroValue:    "nil",
					goType:       "[]*int",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"float64": {
				`package main
	
				func f() []float64 {
					return []float64{}
			}`,
				expType{
					db:           "sa.Text()",
					zeroValue:    "nil",
					graphql:      "[Float!]!",
					elemGraphql:  "Float!",
					goType:       "[]float64",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"float64Pointer": {
				`package main
	
				func f() []*float64 {
					return []*float64{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Float]!",
					elemGraphql:  "Float",
					zeroValue:    "nil",
					goType:       "[]*float64",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"float32": {
				`package main
	
				func f() []float32 {
					return []float32{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Float!]!",
					elemGraphql:  "Float!",
					zeroValue:    "nil",
					goType:       "[]float32",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"float32Pointer": {
				`package main
	
				func f() []*float32 {
					return []*float32{}
			}`,
				expType{
					db:           "sa.Text()",
					graphql:      "[Float]!",
					elemGraphql:  "Float",
					zeroValue:    "nil",
					goType:       "[]*float32",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"time": {
				`package main
	
				import "time"
	
				func f() []time.Time {
					return []time.Time{}
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "[Time!]!",
					defaultGQLFieldName: "times",
					elemGraphql:         "Time!",
					zeroValue:           "nil",
					goType:              "[]time.Time",
					castToMethod:        "cast.UnmarshallJSON",
				},
				nil,
			},
			"timePointer": {
				`package main
	
				import "time"
	
				func f() []*time.Time {
					return []*time.Time{}
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "[Time]!",
					defaultGQLFieldName: "times",
					elemGraphql:         "Time",
					zeroValue:           "nil",
					goType:              "[]*time.Time",
					castToMethod:        "cast.UnmarshallJSON",
				},
				nil,
			},
			"models.User": {
				`package main
	
				import "github.com/lolopinto/ent/internal/test_schema/models"
	
				func f() []models.User {
					return []models.User{}
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "[User!]!",
					defaultGQLFieldName: "users",
					elemGraphql:         "User!",
					zeroValue:           "nil",
					goType:              "[]models.User",
					castToMethod:        "cast.UnmarshallJSON",
				},
				nil,
			},
			"models.UserPointer": {
				`package main
	
				import "github.com/lolopinto/ent/internal/test_schema/models"
	
				func f() []*models.User {
					return []*models.User{}
			}`,
				expType{
					db:                  "sa.Text()",
					graphql:             "[User]!",
					defaultGQLFieldName: "users",
					elemGraphql:         "User",
					zeroValue:           "nil",
					goType:              "[]*models.User",
					castToMethod:        "cast.UnmarshallJSON",
				},
				nil,
			},
		},
	)
}

func TestArrayType(t *testing.T) {
	// I assume no need to test every single case ala slices
	// if we run into issues, revisit.
	testTestCases(
		t,
		&enttype.ArrayType{},
		map[string]testCase{
			"string": {
				`package main

	func f() [2]string {
		return [2]string{}
	}`,
				expType{
					graphql:      "[String!]!",
					elemGraphql:  "String!",
					db:           "sa.Text()",
					zeroValue:    "nil",
					goType:       "[2]string",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"stringPointer": {
				`package main
	
				func f() [2]*string {
					return [2]*string{}
				}`,
				expType{
					graphql:      "[String]!",
					elemGraphql:  "String",
					db:           "sa.Text()",
					zeroValue:    "nil",
					goType:       "[2]*string",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
		},
	)
}

func TestMapType(t *testing.T) {
	// I assume no need to test every single case ala slices
	// if we run into issues, revisit.
	testTestCases(
		t,
		&enttype.MapType{},
		map[string]testCase{
			"string": {
				`package main

	func f() map[string]string {
		return map[string]string{}
	}`,
				expType{
					graphql:      "Map",
					db:           "sa.Text()",
					zeroValue:    "nil",
					goType:       "map[string]string",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
			"stringPointer": {
				`package main
	
				func f() map[string]*bool {
					return map[string]*bool{}
				}`,
				expType{
					graphql:      "Map",
					db:           "sa.Text()",
					zeroValue:    "nil",
					goType:       "map[string]*bool",
					castToMethod: "cast.UnmarshallJSON",
				},
				nil,
			},
		},
	)
}

// use testTypeDirectly for new time types?
//or testhelper.GetCodeWithSchema
func TestEnumType(t *testing.T) {
	testTypeDirectly(t,
		map[string]*typeTestCase{
			"nullable": {
				&enttype.NullableEnumType{
					Type:        "AccountStatus",
					GraphQLType: "AccountStatus",
					Values: []string{
						"NOT_VERIFIED",
						"VERIFIED",
						"DEACTIVATED",
						"DISABLED",
					},
				},
				expType{
					db:      "sa.Text()",
					graphql: "AccountStatus",
					graphqlImports: []enttype.FileImport{
						{
							Type:       "AccountStatus",
							ImportType: enttype.Enum,
						},
					},
					tsType:       "AccountStatus | null",
					goTypePanics: true,
					nonNullableType: &enttype.EnumType{
						Type:        "AccountStatus",
						GraphQLType: "AccountStatus",
						Values: []string{
							"NOT_VERIFIED",
							"VERIFIED",
							"DEACTIVATED",
							"DISABLED",
						},
					},
				},
				nil,
			},
			"not nullable": {
				&enttype.EnumType{
					Type:        "AccountStatus",
					GraphQLType: "AccountStatus",
					Values: []string{
						"NOT_VERIFIED",
						"VERIFIED",
						"DEACTIVATED",
						"DISABLED",
					},
				},
				expType{
					db:      "sa.Text()",
					graphql: "AccountStatus!",
					graphqlImports: []enttype.FileImport{
						enttype.NewGQLFileImport("GraphQLNonNull"),
						{
							Type:       "AccountStatus",
							ImportType: enttype.Enum,
						},
					},
					tsType:       "AccountStatus",
					goTypePanics: true,
					nullableType: &enttype.NullableEnumType{
						Type:        "AccountStatus",
						GraphQLType: "AccountStatus",
						Values: []string{
							"NOT_VERIFIED",
							"VERIFIED",
							"DEACTIVATED",
							"DISABLED",
						},
					},
				},
				nil,
			},
			"not nullable db enum": {
				&enttype.EnumType{
					Type:        "AccountStatus",
					GraphQLType: "AccountStatus",
					Values: []string{
						"NOT_VERIFIED",
						"VERIFIED",
						"DEACTIVATED",
						"DISABLED",
					},
					EnumDBType: true,
				},
				expType{
					db: fmt.Sprintf("postgresql.ENUM(%s, %s, %s, %s, name=%s)",
						strconv.Quote("NOT_VERIFIED"),
						strconv.Quote("VERIFIED"),
						strconv.Quote("DEACTIVATED"),
						strconv.Quote("DISABLED"),
						strconv.Quote("account_status"),
					),
					graphql: "AccountStatus!",
					graphqlImports: []enttype.FileImport{
						enttype.NewGQLFileImport("GraphQLNonNull"),
						{
							Type:       "AccountStatus",
							ImportType: enttype.Enum,
						},
					},
					tsType:       "AccountStatus",
					goTypePanics: true,
					nullableType: &enttype.NullableEnumType{
						Type:        "AccountStatus",
						GraphQLType: "AccountStatus",
						Values: []string{
							"NOT_VERIFIED",
							"VERIFIED",
							"DEACTIVATED",
							"DISABLED",
						},
					},
				},
				nil,
			},
			"nullable db enum": {
				&enttype.NullableEnumType{
					Type:        "AccountStatus",
					GraphQLType: "AccountStatus",
					Values: []string{
						"NOT_VERIFIED",
						"VERIFIED",
						"DEACTIVATED",
						"DISABLED",
					},
					EnumDBType: true,
				},
				expType{
					db: fmt.Sprintf("postgresql.ENUM(%s, %s, %s, %s, name=%s)",
						strconv.Quote("NOT_VERIFIED"),
						strconv.Quote("VERIFIED"),
						strconv.Quote("DEACTIVATED"),
						strconv.Quote("DISABLED"),
						strconv.Quote("account_status"),
					),
					graphql: "AccountStatus",
					graphqlImports: []enttype.FileImport{
						{
							Type:       "AccountStatus",
							ImportType: enttype.Enum,
						},
					},
					tsType:       "AccountStatus | null",
					goTypePanics: true,
					nonNullableType: &enttype.EnumType{
						Type:        "AccountStatus",
						GraphQLType: "AccountStatus",
						Values: []string{
							"NOT_VERIFIED",
							"VERIFIED",
							"DEACTIVATED",
							"DISABLED",
						},
					},
				},
				nil,
			},
		},
	)
}

func TestTimestamptzType(t *testing.T) {
	testTypeDirectly(t, map[string]*typeTestCase{
		"nullable": {
			&enttype.NullableTimestamptzType{},
			expType{
				db:      "sa.TIMESTAMP(timezone=True)",
				graphql: "Time",
				graphqlImports: []enttype.FileImport{
					{
						Type:       "GraphQLTime",
						ImportType: enttype.EntGraphQL,
					},
				},
				tsType:              "Date | null",
				nonNullableType:     &enttype.TimestamptzType{},
				castToMethod:        "cast.ToNullableTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
		"not nullable": {
			&enttype.TimestamptzType{},
			expType{
				db:      "sa.TIMESTAMP(timezone=True)",
				graphql: "Time!",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLNonNull"),
					{
						Type:       "GraphQLTime",
						ImportType: enttype.EntGraphQL,
					},
				},
				tsType:              "Date",
				nullableType:        &enttype.NullableTimestamptzType{},
				castToMethod:        "cast.ToTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
	})
}

func TestTimeType(t *testing.T) {
	testTypeDirectly(t, map[string]*typeTestCase{
		"nullable": {
			&enttype.NullableTimeType{},
			expType{
				db:      "sa.Time()",
				graphql: "String",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLString"),
				},
				tsType:              "Date | null",
				nonNullableType:     &enttype.TimeType{},
				castToMethod:        "cast.ToNullableTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
		"not nullable": {
			&enttype.TimeType{},
			expType{
				db:      "sa.Time()",
				graphql: "String!",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLNonNull"),
					enttype.NewGQLFileImport("GraphQLString"),
				},
				tsType:              "Date",
				nullableType:        &enttype.NullableTimeType{},
				castToMethod:        "cast.ToTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
	})
}

func TestTimetzType(t *testing.T) {
	testTypeDirectly(t, map[string]*typeTestCase{
		"nullable": {
			&enttype.NullableTimetzType{},
			expType{
				db:      "sa.Time(timezone=True)",
				graphql: "String",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLString"),
				},
				tsType:              "Date | null",
				nonNullableType:     &enttype.TimetzType{},
				castToMethod:        "cast.ToNullableTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
		"not nullable": {
			&enttype.TimetzType{},
			expType{
				db:      "sa.Time(timezone=True)",
				graphql: "String!",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLNonNull"),
					enttype.NewGQLFileImport("GraphQLString"),
				},
				tsType:              "Date",
				nullableType:        &enttype.NullableTimetzType{},
				castToMethod:        "cast.ToTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
	})
}

func TestDateType(t *testing.T) {
	testTypeDirectly(t, map[string]*typeTestCase{
		"nullable": {
			&enttype.NullableDateType{},
			expType{
				db:      "sa.Date()",
				graphql: "Time",
				graphqlImports: []enttype.FileImport{
					{
						Type:       "GraphQLTime",
						ImportType: enttype.EntGraphQL,
					},
				},
				tsType:              "Date | null",
				nonNullableType:     &enttype.DateType{},
				castToMethod:        "cast.ToNullableTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
		"not nullable": {
			&enttype.DateType{},
			expType{
				db:      "sa.Date()",
				graphql: "Time!",
				graphqlImports: []enttype.FileImport{
					enttype.NewGQLFileImport("GraphQLNonNull"),
					{
						Type:       "GraphQLTime",
						ImportType: enttype.EntGraphQL,
					},
				},
				tsType:              "Date",
				nullableType:        &enttype.NullableDateType{},
				castToMethod:        "cast.ToTime",
				defaultGQLFieldName: "time",
				zeroValue:           "time.Time{}",
			},
			nil,
		},
	})
}

func testTypeDirectly(t *testing.T, testCases map[string]*typeTestCase) {
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {

			ret := returnType{
				entType: tt.typ,
			}
			if tt.fn != nil {
				tt.fn(&ret, &tt.exp)
			}
			testType(t, tt.exp, ret)
		})
	}
}

// when testing the type directly e.g. typescript...
type typeTestCase struct {
	typ enttype.Type
	exp expType
	fn  func(ret *returnType, exp *expType)
}

type returnType struct {
	entType enttype.Type
	goType  types.Type
}

func getTestReturnType(t *testing.T, code string) returnType {
	pkg, fn, err := schemaparser.FindFunction(code, "main", "f")
	assert.Nil(t, err)
	assert.NotNil(t, fn)
	assert.NotNil(t, pkg)

	assert.NotNil(t, fn.Type.Results)
	results := fn.Type.Results
	assert.Len(t, results.List, 1)

	goType := pkg.TypesInfo.TypeOf(results.List[0].Type)
	return returnType{
		goType:  goType,
		entType: enttype.GetType(goType),
	}
}

func testTestCases(t *testing.T, expType enttype.Type, testCases map[string]testCase) {
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			ret := getTestReturnType(t, tt.code)
			assert.IsType(t, expType, ret.entType)

			if tt.fn != nil {
				tt.fn(&ret, &tt.exp)
			}
			testType(t, tt.exp, ret)
		})
	}
}

func testType(t *testing.T, exp expType, ret returnType) {
	typ := ret.entType
	if exp.graphqlPanics {
		assert.Panics(t, func() { typ.GetGraphQLType() })
	} else {
		assert.Equal(t, exp.graphql, typ.GetGraphQLType())
		gqlType, ok := typ.(enttype.TSGraphQLType)
		if ok {
			assert.Equal(t, exp.graphqlImports, gqlType.GetTSGraphQLImports())
		} else {
			// not a gqlType. this should be 0
			assert.Len(t, exp.graphqlImports, 0)
		}
	}

	entType, ok := typ.(enttype.EntType)
	if ok {
		assert.Equal(t, exp.db, entType.GetDBType())
		if exp.goTypePanics {
			assert.Panics(t, func() { entType.GetCastToMethod() })
			assert.Panics(t, func() { entType.GetZeroValue() })
		} else {
			assert.Equal(t, exp.castToMethod, entType.GetCastToMethod())
			assert.Equal(t, exp.zeroValue, entType.GetZeroValue())

		}
	}

	nullableType, ok := typ.(enttype.NullableType)
	if ok {
		nullType := nullableType.GetNullableType()

		assert.Equal(t, exp.nullableType, nullType)
		if exp.graphqlPanics {
			assert.Panics(t, func() { nullType.GetGraphQLType() })
		} else {
			assert.False(t, strings.HasSuffix(nullType.GetGraphQLType(), "!"))
		}

		if ret.goType != nil {
			// GetNullableType should return expected nullType
			assert.Equal(t, nullType, enttype.GetNullableType(ret.goType, true))
		}
	} else {
		if ret.goType != nil {
			// should return self if not nullable
			assert.Equal(t, typ, enttype.GetNullableType(ret.goType, false))
		}
	}

	nonNullableType, ok := typ.(enttype.NonNullableType)
	if ok {
		nonNullType := nonNullableType.GetNonNullableType()

		assert.Equal(t, exp.nonNullableType, nonNullType)
		if exp.graphqlPanics {
			assert.Panics(t, func() { nonNullType.GetGraphQLType() })
		} else {
			assert.True(t, strings.HasSuffix(nonNullType.GetGraphQLType(), "!"))
		}

		if ret.goType != nil {
			// GetNonNullableType should return expected nonNullType
			assert.Equal(t, nonNullType, enttype.GetNonNullableType(ret.goType, true))
		}
	} else {
		if ret.goType != nil {
			// should return self if not non-nullable
			assert.Equal(t, typ, enttype.GetNonNullableType(ret.goType, false))
		}
	}

	defaultFieldNameType, ok := typ.(enttype.DefaulFieldNameType)
	if ok {
		assert.Equal(t, exp.defaultGQLFieldName, defaultFieldNameType.DefaultGraphQLFieldName())
	}

	listType, ok := typ.(enttype.ListType)
	if ok {
		assert.Equal(t, exp.elemGraphql, listType.GetElemGraphQLType())
	}

	if ret.goType != nil {
		assert.Equal(t, exp.goType, enttype.GetGoType(ret.goType))
	}

	tsType, ok := typ.(enttype.TSType)
	if ok {
		assert.Equal(t, exp.tsType, tsType.GetTSType())
	} else {
		assert.Equal(t, "", exp.tsType)
	}

	assert.Equal(t, exp.errorType, enttype.IsErrorType(typ))
	assert.Equal(t, exp.contextType, enttype.IsContextType(typ))
}
