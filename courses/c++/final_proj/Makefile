CC = g++
SRC = src/*.cpp
INCLUDE = -Iinclude -Ilib
OUTPUT = build/app.exe
FLAGS = -Wall -Wextra -lcurl

all: compile run
	

compile:
	$(CC) $(SRC) $(INCLUDE) -o $(OUTPUT) $(FLAGS)

run:
	./$(OUTPUT)

clean:
	rm build/