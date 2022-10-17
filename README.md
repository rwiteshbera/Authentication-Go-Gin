Installation
- Create a `.env` file and add the secret key. 
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