package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/formatter"
	"strings"
	"testing"
)
import (
	"github.com/vektah/gqlparser/v2/ast"
)

func TestGraphQLAstParser(t *testing.T) {
	source := &ast.Source{
		Input: string(`
scalar DateTime
scalar JSON

enum order_by {
  asc
  asc_nulls_first
  desc
  desc_nulls_first
}

input Boolean_comparison_exp {
  _eq: Boolean
  _gt: Boolean
  _gte: Boolean
  _in: [Boolean!]!
  _is_null: Boolean
  _lt: Boolean
  _lte: Boolean
  _neq: Boolean
  _nin: [Boolean!]!
}

input Int_comparison_exp {
  _eq: Int
  _gt: Int
  _gte: Int
  _in: [Int!]!
  _is_null: Boolean
  _lt: Int
  _lte: Int
  _neq: Int
  _nin: [Int!]!
}

input String_comparison_exp {
  _eq: String
  _gt: String
  _gte: String
  _ilike: String
  _in: [String!]!
  _iregex: String
  _is_null: Boolean
  _istartswith: String
  _like: String
  _lt: String
  _lte: String
  _neq: String
  _nilike: String
  _nin: [String!]!
  _niregex: String
  _nis_null: Boolean
  _nlike: String
  _nregex: String
  _nsimilar: String
  _regex: String
  _similar: String
}

input timestamp_comparison_exp {
	_eq: timestamp
	_gt: timestamp
	_gte: timestamp
	_in: [timestamp!]!
	_is_null: Boolean
	_lt: timestamp
	_lte: timestamp
	_neq: timestamp
	_nin: [timestamp!]!
}

input user_bool_exp {
  _and: [user_bool_exp!]
  _not: user_bool_exp
  _or: [user_bool_exp!]
  id: String_comparison_exp
  name: String_comparison_exp
}

input user_order_by {
  id: order_by
  name: order_by
}

type user {
  id: String!
  name: String!
  authorComments(
     where: comment_bool_exp
	 order_by: [comment_order_by!]
	 limit: Int
	 offset: Int
  ): [comment!]!
}

input comment_bool_exp {
	_and: [comment_bool_exp!]
	_not: comment_bool_exp
	_or: [comment_bool_exp!]
	author_id: String_comparison_exp
	content: String_comparison_exp
}

input comment_order_by {
	author_id: order_by
	content: order_by
	created_at: order_by
	id: order_by
	updated_at: order_by
}

type comment {
	id: String!
	author_id: String!
    content: String!
    created_at: DateTime!
	updated_at: DateTime!
}

type mutation_root {
  user(
     where: user_bool_exp
     order_by: [user_order_by!]
	 limit: Int
	 offset: Int
  ): [user!]!
}
`),
	}
	// パース
	schema, err := gqlparser.LoadSchema(source)

	if err != nil {
		panic(err)
	}

	// パース結果を出力
	jsonData, err := json.Marshal(schema)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func TestGenerateGraphQLAstParser(t *testing.T) {

	baseType := `
scalar timestamp
scalar JSON

enum order_by {
  asc
  asc_nulls_first
  desc
  desc_nulls_first
}

input Boolean_comparison_exp {
  _eq: Boolean
  _gt: Boolean
  _gte: Boolean
  _in: [Boolean!]!
  _is_null: Boolean
  _lt: Boolean
  _lte: Boolean
  _neq: Boolean
  _nin: [Boolean!]!
}

input Int_comparison_exp {
  _eq: Int
  _gt: Int
  _gte: Int
  _in: [Int!]!
  _is_null: Boolean
  _lt: Int
  _lte: Int
  _neq: Int
  _nin: [Int!]!
}

input String_comparison_exp {
  _eq: String
  _gt: String
  _gte: String
  _ilike: String
  _in: [String!]!
  _iregex: String
  _is_null: Boolean
  _istartswith: String
  _like: String
  _lt: String
  _lte: String
  _neq: String
  _nilike: String
  _nin: [String!]!
  _niregex: String
  _nis_null: Boolean
  _nlike: String
  _nregex: String
  _nsimilar: String
  _regex: String
  _similar: String
}

input timestamp_comparison_exp {
	_eq: timestamp
	_gt: timestamp
	_gte: timestamp
	_in: [timestamp!]!
	_is_null: Boolean
	_lt: timestamp
	_lte: timestamp
	_neq: timestamp
	_nin: [timestamp!]!
}
`

	source := &ast.Source{
		Input: baseType,
	}
	schema, err := gqlparser.LoadSchema(source)

	if err != nil {
		panic(err)
	}

	schema.AddTypes(&ast.Definition{
		Kind: "OBJECT",
		Name: "comment",
		Fields: []*ast.FieldDefinition{
			{
				Name: "id",
				Type: &ast.Type{
					NamedType: "String",
				},
			},
			{
				Name: "author_id",
				Type: &ast.Type{
					NamedType: "String",
				},
			},
			{
				Name: "content",
				Type: &ast.Type{
					NamedType: "String",
				},
			},
			{
				Name: "created_at",
				Type: &ast.Type{
					NamedType: "timestamp",
				},
			},
		},
	}, &ast.Definition{
		Kind: "INPUT_OBJECT",
		Name: "comment_bool_exp",
		Fields: ast.FieldList{
			{
				Name: "author_id",
				Type: &ast.Type{
					NamedType: "String_comparison_exp",
				},
			},
			{
				Name: "content",
				Type: &ast.Type{
					NamedType: "String_comparison_exp",
				},
			},
			{
				Name: "created_at",
				Type: &ast.Type{
					NamedType: "timestamp_comparison_exp",
				},
			},
			{
				Name: "_and",
				Type: &ast.Type{
					NamedType: "[comment_bool_exp!]",
				},
			},
			{
				Name: "_not",
				Type: &ast.Type{
					NamedType: "comment_bool_exp",
				},
			},
			{
				Name: "_or",
				Type: &ast.Type{
					NamedType: "[comment_bool_exp]",
				},
			},
		},
	}, &ast.Definition{
		Kind: "INPUT_OBJECT",
		Name: "comment_order_by",
		Fields: ast.FieldList{
			{
				Name: "author_id",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
			{
				Name: "content",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
			{
				Name: "created_at",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
			{
				Name: "id",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
		},
	}, &ast.Definition{
		Kind: "OBJECT",
		Name: "user",
		Fields: []*ast.FieldDefinition{
			{
				Name: "id",
				Type: &ast.Type{
					NamedType: "String",
				},
			},
			{
				Name: "name",
				Type: &ast.Type{
					NamedType: "String",
				},
			},
			{
				Name: "authorComments",
				Type: &ast.Type{
					NamedType: "[comment!]",
					NonNull:   true,
				},
				Arguments: ast.ArgumentDefinitionList{
					{
						Name: "where",
						Type: &ast.Type{
							NamedType: "comment_bool_exp",
						},
					},
				},
			},
		},
	}, &ast.Definition{
		Kind: "INPUT_OBJECT",
		Name: "user_bool_exp",
		Fields: ast.FieldList{
			{
				Name: "id",
				Type: &ast.Type{
					NamedType: "String_comparison_exp",
				},
			},
			{
				Name: "name",
				Type: &ast.Type{
					NamedType: "String_comparison_exp",
				},
			},
			{
				Name: "_and",
				Type: &ast.Type{
					NamedType: "[user_bool_exp!]",
				},
			},
			{
				Name: "_not",
				Type: &ast.Type{
					NamedType: "user_bool_exp",
				},
			},
			{
				Name: "_or",
				Type: &ast.Type{
					NamedType: "[user_bool_exp!]",
				},
			},
		},
	}, &ast.Definition{
		Kind: "INPUT_OBJECT",
		Name: "user_order_by",
		Fields: ast.FieldList{
			{
				Name: "id",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
			{
				Name: "name",
				Type: &ast.Type{
					NamedType: "order_by",
				},
			},
			{
				Name: "authorComments",
				Type: &ast.Type{
					NamedType: "[comment_order_by!]",
				},
			},
		},
	}, &ast.Definition{
		Kind: "OBJECT",
		Name: "mutation_root",
		Fields: []*ast.FieldDefinition{
			{
				Name: "user",
				Type: &ast.Type{
					NamedType: "[user!]",
					NonNull:   true,
				},
				Arguments: ast.ArgumentDefinitionList{
					{
						Name: "where",
						Type: &ast.Type{
							NamedType: "user_bool_exp",
						},
					},
					{
						Name: "order_by",
						Type: &ast.Type{
							NamedType: "user_order_by",
						},
					},
					{
						Name: "limit",
						Type: &ast.Type{
							NamedType: "Int",
						},
					},
					{
						Name: "offset",
						Type: &ast.Type{
							NamedType: "Int",
						},
					},
				},
			},
			{
				Name: "comment",
				Type: &ast.Type{
					NamedType: "[comment!]",
					NonNull:   true,
				},
				Arguments: ast.ArgumentDefinitionList{
					{
						Name: "where",
						Type: &ast.Type{
							NamedType: "comment_bool_exp",
						},
					},
					{
						Name: "order_by",
						Type: &ast.Type{
							NamedType: "comment_order_by",
						},
					},
					{
						Name: "limit",
						Type: &ast.Type{
							NamedType: "Int",
						},
					},
				},
			},
		},
	})

	// map to slice
	var types []*ast.Definition
	for _, v := range schema.Types {
		types = append(types, v)
	}

	var buf bytes.Buffer
	astFormatter := formatter.NewFormatter(&buf)
	astFormatter.FormatSchema(schema)

	var a = strings.Split(buf.String(), "\n")

	for i, v := range a {
		fmt.Println(i+1, v)
	}

	fmt.Println()

	_schema, err := gqlparser.LoadSchema(&ast.Source{
		Input: buf.String(),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(_schema)
}
