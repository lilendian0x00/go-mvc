# go-mvc
A simple CRUD Application written in Golang based on MVC Architecture

## Technologies

### Backend ###
Web Framework: [Fiber](https://github.com/gofiber/fiber)

Database: [Postgresql](https://www.postgresql.org/)

Database ORM (Object–relational mapping): [GORM](https://github.com/go-gorm/gorm)

Configuration solution: [Viper](https://github.com/spf13/viper) 

-------------------
### Front-End ###
Main Framework: [ReactJS](https://reactjs.org/)

-------------------

## Run
⚠️**Disclaimer:** I've already done `npm run build` for the frontend source code (`frontend`dir) under `views/build`

0. install postgresql based on your OS from [Postgresql Website](https://www.postgresql.org/download/)
1. make a database and override the `DB_SOURCE` ENV inside `app.env` file
2. install golang based on your OS from [Golang Website](https://go.dev/dl/)
3. `git clone https://github.com/lilendian0x00/go-mvc.git`
4. `cd go-mvc`
5. `go build .`
6. `./go-mvc`