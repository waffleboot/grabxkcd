all:
	go build -o xkcd main/main.go

call:
	~/go/bin/go-callvis -nostd -focus github.com/waffleboot/grabxkcd/grabxkcd -file=call -format=png github.com/waffleboot/grabxkcd/main
