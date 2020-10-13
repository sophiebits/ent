package schema_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/lolopinto/ent/internal/schema"
	"github.com/lolopinto/ent/internal/schema/base"
	"github.com/lolopinto/ent/internal/schema/input"
	"github.com/lolopinto/ent/internal/schema/testhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrimaryKeyFieldConstraint(t *testing.T) {
	testConstraints(
		t,
		map[string]string{
			"user.ts": testhelper.GetCodeWithSchema(
				`import {Field, StringType, BaseEntSchema} from "{schema}";

				export default class User extends BaseEntSchema {
					fields: Field[] = [
						StringType({
							name: 'firstName',
						}),
						StringType({
							name: 'lastName',
						}),
					];
				}
			`,
			),
		},
		map[string]*schema.NodeData{
			"User": {
				Constraints: constraintsWithNodeConstraints("users"),
			},
		},
	)
}

func TestForeignKeyFieldConstraint(t *testing.T) {
	testConstraints(
		t,
		map[string]string{
			"user.ts": testhelper.GetCodeWithSchema(
				`import {Field, StringType, BaseEntSchema} from "{schema}";

				export default class User extends BaseEntSchema {
					fields: Field[] = [
						StringType({
							name: 'firstName',
						}),
						StringType({
							name: 'lastName',
						}),
					];
				}
			`,
			),
			"contact.ts": testhelper.GetCodeWithSchema(
				`import {Field, StringType, BaseEntSchema, UUIDType} from "{schema}";

				export default class Contact extends BaseEntSchema {
					fields: Field[] = [
						StringType({
							name: 'firstName',
						}),
						StringType({
							name: 'lastName',
						}),
						UUIDType({
							name: "ownerID",
							foreignKey: ["User", "ID"],
						}),
					];
				}
			`,
			),
		},
		map[string]*schema.NodeData{
			"User": {
				Constraints: constraintsWithNodeConstraints("users"),
			},
			"Contact": {
				Constraints: constraintsWithNodeConstraints(
					"contacts",
					&input.Constraint{
						Name:    "contacts_owner_id_fkey",
						Type:    input.ForeignKey,
						Columns: []string{"ownerID"},
						ForeignKey: &input.ForeignKeyInfo{
							TableName: "users",
							Columns:   []string{"id"},
							OnDelete:  "CASCADE",
						},
					}),
			},
		},
	)
}

func TestUniqueFieldConstraint(t *testing.T) {
	testConstraints(
		t,
		map[string]string{
			"user.ts": testhelper.GetCodeWithSchema(
				`import {Field, StringType, BaseEntSchema} from "{schema}";

				export default class User extends BaseEntSchema {
					fields: Field[] = [
						StringType({
							name: 'firstName',
						}),
						StringType({
							name: 'lastName',
						}),
						StringType({
							name: "emailAddress",
							unique: true,
						})
					];
				}
			`,
			),
		},
		map[string]*schema.NodeData{
			"User": {
				Constraints: constraintsWithNodeConstraints("users",
					&input.Constraint{
						Name:    "users_unique_email_address",
						Type:    input.Unique,
						Columns: []string{"emailAddress"},
					},
				),
			},
		},
	)
}

func TestConstraints(t *testing.T) {
	testCases := map[string]testCase{
		"multi-column-primary key": {
			code: map[string]string{
				"user_photo.ts": testhelper.GetCodeWithSchema(`
					import {Schema, Field, UUIDType, Constraint, ConstraintType} from "{schema}";

					export default class UserPhoto implements Schema {
						fields: Field[] = [
							UUIDType({
								name: 'UserID',
							}),
							UUIDType({
								name: 'PhotoID',
							}),
						];

						constraints: Constraint[] = [
							{
								name: "user_photos_pkey",
								type: ConstraintType.PrimaryKey,
								columns: ["UserID", "PhotoID"],
							},
						];
					}
				`),
			},
			expectedMap: map[string]*schema.NodeData{
				"UserPhoto": {
					Constraints: []*input.Constraint{
						{
							Name:    "user_photos_pkey",
							Type:    input.PrimaryKey,
							Columns: []string{"UserID", "PhotoID"},
						},
					},
				},
			},
		},
		"multi-column-unique key": {
			code: map[string]string{
				"user.ts": testhelper.GetCodeWithSchema(`
					import {Field, StringType, BaseEntSchema} from "{schema}";

					export default class User extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: 'firstName',
							}),
							StringType({
								name: 'lastName',
							}),
						];
					}
				`),
				"contact.ts": testhelper.GetCodeWithSchema(`
					import {BaseEntSchema, Field, UUIDType, StringType, Constraint, ConstraintType} from "{schema}";

					export default class Contact extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: "emailAddress",
							}),
							UUIDType({
								name: "userID",
								foreignKey: ["User", "ID"],
							}),
						];

						constraints: Constraint[] = [
							{
								name: "contacts_unique_email",
								type: ConstraintType.Unique,
								columns: ["emailAddress", "userID"],
							},
						];
					}
				`),
			},
			expectedMap: map[string]*schema.NodeData{
				"User": {
					Constraints: constraintsWithNodeConstraints("users"),
				},
				"Contact": {
					Constraints: constraintsWithNodeConstraints("contacts",
						&input.Constraint{
							Name:    "contacts_user_id_fkey",
							Type:    input.ForeignKey,
							Columns: []string{"userID"},
							ForeignKey: &input.ForeignKeyInfo{
								TableName: "users",
								Columns:   []string{"id"},
								OnDelete:  "CASCADE",
							},
						},
						&input.Constraint{
							Name:    "contacts_unique_email",
							Type:    input.Unique,
							Columns: []string{"emailAddress", "userID"},
						}),
				},
			},
		},
		"multi-column-foreign key": {
			code: map[string]string{
				"user.ts": testhelper.GetCodeWithSchema(`
					import {Field, StringType, BaseEntSchema} from "{schema}";

					export default class User extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: 'firstName',
							}),
							StringType({
								name: 'lastName',
							}),
							StringType({
								name: 'emailAddress',
								unique: true,
							}),
						];
					}
				`),
				"contact.ts": testhelper.GetCodeWithSchema(`
					import {BaseEntSchema, Field, UUIDType, StringType, Constraint, ConstraintType} from "{schema}";

					export default class Contact extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: "emailAddress",
							}),
							UUIDType({
								name: "userID",
							}),
						];

						constraints: Constraint[] = [
							{
								name: "contacts_user_fkey",
								type: ConstraintType.ForeignKey,
								columns: ["userID", "emailAddress"],
								fkey: {
									tableName: "users", 
									ondelete: "CASCADE",
									columns: ["ID", "emailAddress"],
								}
							},
						];
					}
				`),
			},
			expectedMap: map[string]*schema.NodeData{
				"User": {
					Constraints: constraintsWithNodeConstraints("users", &input.Constraint{
						Name:    "users_unique_email_address",
						Type:    input.Unique,
						Columns: []string{"emailAddress"},
					}),
				},
				"Contact": {
					Constraints: constraintsWithNodeConstraints("contacts",
						&input.Constraint{
							Name:    "contacts_user_fkey",
							Type:    input.ForeignKey,
							Columns: []string{"userID", "emailAddress"},
							ForeignKey: &input.ForeignKeyInfo{
								TableName: "users",
								Columns:   []string{"ID", "emailAddress"},
								OnDelete:  "CASCADE",
							},
						},
					),
				},
			},
		},
		"check constraint no columns": {
			code: map[string]string{
				"item.ts": testhelper.GetCodeWithSchema(`
					import {Field, FloatType, BaseEntSchema, Constraint, ConstraintType} from "{schema}";

					export default class Item extends BaseEntSchema {
						fields: Field[] = [
							FloatType({
								name: 'price',
							}),
						];

						constraints: Constraint[] = [
							{
								name: "item_positive_price",
								type: ConstraintType.Check,
								condition: 'price > 0',
								columns: [],
							},
						];
					}`),
			},
			expectedMap: map[string]*schema.NodeData{
				"Item": {
					Constraints: constraintsWithNodeConstraints("items", &input.Constraint{
						Name:      "item_positive_price",
						Type:      input.Check,
						Columns:   []string{},
						Condition: "price > 0",
					}),
				},
			},
		},
		"check constraint multiple columns": {
			code: map[string]string{
				"item.ts": testhelper.GetCodeWithSchema(`
					import {Field, FloatType, BaseEntSchema, Constraint, ConstraintType} from "{schema}";

					export default class Item extends BaseEntSchema {
						fields: Field[] = [
							FloatType({
								name: 'price',
							}),
							FloatType({
								name: 'discount_price',
							}),
						];

						constraints: Constraint[] = [
							{
								name: "item_positive_price",
								type: ConstraintType.Check,
								// TODO condition is required when type == Check
								condition: 'price > 0',
								// TODO need to test this later when we have mixed everything in since we may not
								// want this...
								columns: ['price'],
							},
							{
								name: "item_positive_discount_price",
								type: ConstraintType.Check,
								// TODO condition is required when type == Check
								condition: 'discount_price > 0',
								columns: ['discount_price'],
							},
							{
								name: "item_price_greater_than_discount",
								type: ConstraintType.Check,
								// TODO condition is required when type == Check
								condition: 'price > discount_price',
								columns: ['price', 'discount_price'],
							},
						];
					}`),
			},
			expectedMap: map[string]*schema.NodeData{
				"Item": {
					Constraints: constraintsWithNodeConstraints("items", &input.Constraint{
						Name:      "item_positive_price",
						Type:      input.Check,
						Columns:   []string{"price"},
						Condition: "price > 0",
					},
						&input.Constraint{
							Name:      "item_positive_discount_price",
							Type:      input.Check,
							Columns:   []string{"discount_price"},
							Condition: "discount_price > 0",
						},
						&input.Constraint{
							Name:      "item_price_greater_than_discount",
							Type:      input.Check,
							Columns:   []string{"price", "discount_price"},
							Condition: "price > discount_price",
						}),
				},
			},
		},
	}

	runTestCases(t, testCases)
}

func TestInvalidConstraints(t *testing.T) {
	testCases := map[string]testCase{
		"missing fkey field": {
			code: map[string]string{
				"user.ts": testhelper.GetCodeWithSchema(`
					import {Field, StringType, BaseEntSchema} from "{schema}";

					export default class User extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: 'firstName',
							}),
							StringType({
								name: 'lastName',
							}),
							StringType({
								name: 'emailAddress',
								unique: true,
							}),
						];
					}
				`),
				"contact.ts": testhelper.GetCodeWithSchema(`
					import {BaseEntSchema, Field, UUIDType, StringType, Constraint, ConstraintType} from "{schema}";

					export default class Contact extends BaseEntSchema {
						fields: Field[] = [
							StringType({
								name: "emailAddress",
							}),
							UUIDType({
								name: "userID",
							}),
						];

						constraints: Constraint[] = [
							{
								name: "contacts_user_fkey",
								type: ConstraintType.ForeignKey,
								columns: ["userID", "emailAddress"],
							},
						];
					}
				`),
			},
			expectedErr: fmt.Errorf("ForeignKey cannot be nil when type is ForeignKey"),
		},
		"missing condition check constraint": {
			code: map[string]string{
				"item.ts": testhelper.GetCodeWithSchema(`
					import {Field, FloatType, BaseEntSchema, Constraint, ConstraintType} from "{schema}";

					export default class Item extends BaseEntSchema {
						fields: Field[] = [
							FloatType({
								name: 'price',
							}),
						];

						constraints: Constraint[] = [
							{
								name: "item_positive_price",
								type: ConstraintType.Check,
								columns: [],
							},
						];
					}`),
			},
			expectedErr: fmt.Errorf("Condition is required when constraint type is Check"),
		},
	}
	runTestCases(t, testCases)
}

type testCase struct {
	code        map[string]string
	expectedMap map[string]*schema.NodeData
	expectedErr error
}

func runTestCases(t *testing.T, testCases map[string]testCase) {
	for key, tt := range testCases {
		t.Run(key, func(t *testing.T) {

			if tt.expectedErr != nil {
				assert.PanicsWithValue(t, tt.expectedErr.Error(), func() {
					testConstraints(t, tt.code, tt.expectedMap)
				})
			} else {
				testConstraints(t, tt.code, tt.expectedMap)
			}
		})
	}
}

func testConstraints(t *testing.T, code map[string]string, expectedMap map[string]*schema.NodeData) {
	absPath, err := filepath.Abs(".")
	require.NoError(t, err)

	schema := testhelper.ParseSchemaForTest(
		t,
		absPath,
		code,
		base.TypeScript,
	)

	for k, expNodeData := range expectedMap {
		info := schema.Nodes[k+"Config"]
		require.NotNil(t, info, "expected %s to exist in schema", k)
		nodeData := info.NodeData

		expConstraints := expNodeData.Constraints
		constraints := nodeData.Constraints

		assert.Len(t, constraints, len(expConstraints))

		for i, expConstraint := range expConstraints {
			constraint := constraints[i]

			assert.Equal(t, expConstraint.Name, constraint.Name)
			assert.Equal(t, expConstraint.Columns, constraint.Columns)
			assert.Equal(t, expConstraint.Type, constraint.Type)
			assert.Equal(t, expConstraint.Condition, constraint.Condition)

			if expConstraint.ForeignKey == nil {
				require.Nil(t, constraint.ForeignKey)
			} else {
				require.NotNil(t, constraint.ForeignKey)

				assert.Equal(t, expConstraint.ForeignKey.TableName, constraint.ForeignKey.TableName)
				assert.Equal(t, expConstraint.ForeignKey.OnDelete, constraint.ForeignKey.OnDelete)
				assert.Equal(t, expConstraint.ForeignKey.Columns, constraint.ForeignKey.Columns)
			}
		}
	}
}

func constraintsWithNodeConstraints(tableName string, constraints ...*input.Constraint) []*input.Constraint {
	return append([]*input.Constraint{
		{
			Name:    fmt.Sprintf("%s_id_pkey", tableName),
			Type:    input.PrimaryKey,
			Columns: []string{"ID"},
		},
	}, constraints...)
}