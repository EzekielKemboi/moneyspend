build:
	go build -o moneyspend

run:
	go build -o moneyspend
	./moneyspend

version:
	./moneyspend -v

clean:
	rm ./moneyspend
