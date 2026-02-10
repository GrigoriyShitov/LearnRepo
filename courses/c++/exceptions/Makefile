CC = g++
SRC = src/*.cpp
INCLUDE = -Iinclude
OUTPUT = build/app.exe
FLAGS = -Wextra

all: compile run
	

compile:
	$(CC) $(SRC) $(INCLUDE) -o $(OUTPUT) $(FLAGS)

run:
	./$(OUTPUT)

clean:
	rm build/*.exe