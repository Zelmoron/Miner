APP_NAME=Minesweeper
SRC=cmd/app/main.go
build:
	@echo "Build"
	go build -o $(APP_NAME) $(SRC)
run: 
	@echo "Server run"
	./$(APP_NAME)
	