### Requirements:
```shell script
GNU Make v4.0 or later;
docker v20.10.7 or later;
docker-compose v1.25 or later.
```
### Instructions:
1. Rename `.env.example` file to `.env`; 
2. Define your environments like database user and password or api port;
3. Run application;
4. It accepts a `POST` request with these body:
   ```shell script
   {
      "length": 32,             # Password Length
      "has_letter": true,       # If it has letter
      "has_number": true,       # If it has number
      "has_special_char": true  # If it has special char
    }
   ```
5. You must define an `API-KEY` in your header, this api-key you define in your `.env` file;<br />
OBS.: It could be any value, but you must define this value in reader of request with name `api-token`;<br />
6. After running the project, you could make a request at this URL: `http://localhost:{YOUR_API_PORT}/password-gen/`.<br />
OBS.: If you didn't have defined your api port, the project will run at port `8095` by default.<br />
OBS2.: If you're using OS Windows, you could run this command to start: `docker-compose up --build`.

### Run application/instalation:
```shell script
make build-docker
```

### Run application without changes:
```shell script
make run-docker
```
or
```shell script
make up
```
