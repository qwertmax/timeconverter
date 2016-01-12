Time Converter API

http://s.q-man.ru:3000/users

## User List

```bash 
GET /users
```

http://s.q-man.ru:3000/user/12

```json
[{
	id: 12,
	login: "123",
	password: "1",
	email: "StanleeLOD@gmail.com",
	cities: null
},{
	id: 13,
	login: "123",
	password: "1",
	email: "StanleeLOD@gmail.com",
	cities: null
}]
```

## Get User

```bash
GET /user/:id
```

http://s.q-man.ru:3000/user/12

```json
{
	id: 12,
	login: "123",
	password: "1",
	email: "StanleeLOD@gmail.com",
	cities: null
}
```

## Create User

POST /user

```json
{
	login: "test",
	password: "test",
	email: "email@example.com"
}
```

PUT /user/:id

```json
{
	login: "test",
	password: "test",
	email: "email@example.com"
}
```

$$ Remove User

DELETE /user/:id
