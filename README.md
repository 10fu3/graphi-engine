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