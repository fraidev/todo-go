# todo-go
Simple Crud made with Go. Using go-pg and chi.



### Local

```sh
cp sample.local.env .env.local          # Create .env file
docker-compose up -d      	        # Run Dependencies if needed
./scripts/migrate up      	        # Apply database migrations
```

After that, you can choose:
* Config your IDE to read `.env.local` when running or debugging.
* Just load `.env.local` and execute `go run .`.

### Docker

```sh
cp sample.docker.env .env       	# Create .env file
docker-compose up -d      	        # Run Dependencies
./scripts/migrate up      	        # Apply database migrations
./scripts/build      		        # Build service image
./scripts/run       		        # Run Service
```