CC = g++
SRC = src/*.cpp
INCLUDE = -Iinclude
OUTPUT = build/app.exe
FLAGS = -Wall -Wextra

all: compile run
	

compile:
	$(CC) $(SRC) $(INCLUDE) -o $(OUTPUT) $(FLAGS)

run:
	./$(OUTPUT)

clean:
	rm build/