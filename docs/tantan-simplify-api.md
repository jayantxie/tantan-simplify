## API
- List all users
```$xslt
GET /users

Example:
$curl -XGET "http://localhost:80/users"
[
  {
    "id": "21341231231",
    "name": "Bob" ,
    "type": "user"
}, {
    "id": "31231242322",
    "name": "Samantha" ,
    "type": "user"
} ]
```
- Create a user
```$xslt
POST /users

allowed fields:
  name = string
Example:
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:80/users"
{
  "id": "11231244213",
  "name": "Alice" ,
  "type": "user"
}
```
- List a users all relationships
```$xslt
GET /users/:user_id/relationships

Example:
$curl -XGET "http://localhost:80/users/11231244213/relationships"
[
  {
    "user_id": "222333444",
    "state": "liked" ,
    "type": "relationship"
}, {
    "user_id": "333222444",
    "state": "matched" ,
    "type": "relationship"
}, {
    "user_id": "444333222",
    "state": "disliked" ,
    "type": "relationship"
} ]
```
- Create/update relationship state to another user.
```$xslt
PUT /users/:user_id/relationships/:other_user_id

allowed fields:
   state = "liked"|"disliked"
If two users have "liked" each other, then the state of the relationship is "matched"
Example:
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/11231244213/relationships/21341231231"
{
  "user_id": "21341231231",
  "state": "liked" ,
  "type": "relationship"
}
$curl -XPUT -d '{"state":"liked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
  "user_id": "11231244213",
  "state": "matched" ,
  "type": "relationship"
}
$curl -XPUT -d '{"state":"disliked"}'
"http://localhost:80/users/21341231231/relationships/11231244213"
{
  "user_id": "11231244213",
  "state": "disliked" ,
  "type": "relationship"
}
```