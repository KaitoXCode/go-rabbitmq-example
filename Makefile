all: 
	@echo "No specific target selected. Options available: ['rmq-docker-up', 'rmq-docker-down']"

rmq-docker-up:
	@docker compose up --build -d

recon-docker-down:
	@docker compose down

producer-one-msg:
	@go run ./producers/pone/. $(RKEY)

producer-two-msg:
	@go run ./producers/ptwo/. $(RKEY)

consumer-one:
	@go run ./consumers/cone/.

consumer-two:
	@go run ./consumers/ctwo/.

consumer-three:
	@go run ./consumers/cthree/.
