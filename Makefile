BINARY=engine
engine:
	go build -o ${BINARY} app/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t hotel_management .

run:
	docker-compose up --build -d

down:
	docker-compose down

stop:
	docker-compose pause