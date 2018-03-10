# api-graphl
My journey with golang and graphql

## graphl-1
URL:`http://localhost:7171/graphql?query={who{id,name}}`

## graphql-2
URL :`http://localhost:1234/graphql?query={checkuser(id:1,name:%22ali%22){id,name}}`

## graphql-3
Get id and title
`curl -XPOST -H 'Content-Type:application/graphql'  -d '{ post(id:1) { id, title} }' http://localhost:3000/graphql`

Add Comment

`curl -XPOST -H 'Content-Type:application/graphql'  -d '{ post(id:1) { id, title, comment{name}} }' http://localhost:3000/graphql`
