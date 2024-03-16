## graphi-engine
PostgreSQL -> GraphQL の間を補完するHasuraみたいなものをGolangで作ってみる

## GraphQLで生成したいクエリ
例えば次のような関係で表現されるデータベースがあるとする
```
- user
    id
    name

- comment
    id
    author_id
    content

user 1:N comment
```

このとき、GraphQLのSchemaは次のようになっていてほしい
```
type User {
    id: ID!
    name: String!
    comments: [Comment!]!
}

type Comment {
    id: ID!
    author_id: ID!
    author: User!
    content: String!
}

type query_root {
    user(where: user_bool_exp! offset: Int! limit: Int! order_by: [user_order_by!]): [user!]!
    comment(where: comment_bool_exp! offset: Int! limit: Int! order_by: [comment_order_by!]): [comment!]!
}

type user_bool_exp {
    id: Int_comparison_exp
    name: String_comparison_exp
    comments: comment_bool_exp
    _and: [user_bool_exp!]
    _or: [user_bool_exp!]
    _not: user_bool_exp
}

type user_order_by {
    id: order_by
    name: order_by
    comments: comment_order_by
}

type comment_bool_exp {
    id: Int_comparison_exp
    author_id: Int_comparison_exp
    content: String_comparison_exp
    _and: [comment_bool_exp!]
    _or: [comment_bool_exp!]
    _not: comment_bool_exp
}

type comment_order_by {
    id: order_by
    author_id: order_by
    content: order_by
}
```

これを生成するためのgqlparserのASTは

```go
package test

import "testing"
import (
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestGraphQLAstParser(t *testing.T) {
    baseType := ""
    
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
        	// フィールドがならぶ
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
            // フィールドがならぶ
    	},
    })
}
```

こんな感じで生成できるようになるはず