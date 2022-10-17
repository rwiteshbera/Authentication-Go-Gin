Installation
- Create a `.env` file and add the secret keys. 
```bash
PORT = 
MONGODB_URI = 
SECRET_KEY = 
```
You can generate the SECRET_KEY from [https://www.javainuse.com/jwtgenerator](https://www.javainuse.com/jwtgenerator)

- go mod tidy ensures that the go. mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies, if there are some not used dependencies go mod tidy will remove those from go.
```bash
go mod tidy
```

- `air` command will build the project and start the server. Refer [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
```bash
air
```

#### Workflow
- **POST** : `http://localhost:5050/user/signup`
```json
{
    "first_name" : "",
    "last_name" : "",
    "email" : "",
    "password":""
}
```

- **POST** : `http://localhost:5050/user/login`
```json
{
    "email":"",
    "password":""
}
```
- Get specific user data by ID (Provide the `token` in request header)
- **GET** : `http://localhost:5050/users/<USER_ID>`

- Get all users' data (Provide the `token` in request header)
- **GET** : `http://localhost:5050/users`