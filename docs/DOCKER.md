
### Try out in docker image on you machine

```bash
docker run -p 8000:8000 -e PORT=8000 boofny/golive-docker
```

## Run the image and try these routes
### Routes 
- [/ping](http://localhost:8000/ping) <-- GET
- [/user/{id}](http://localhost:8000/user) <-- Url params
- [/posting](http://localhost:8000/posting) <-- POST
  - ```bash
    curl -X POST http://localhost:8000/posting    
    -H "Content-Type: application/json"    
    -d '{"name": "john", "email": "johnBram@gmail.com"}'
    ```
- [/v1/ping](http://localhost:8000/v1/ping) <-- For route grouping

### Get a more info on my docker repo
[Click here](https://hub.docker.com/repository/docker/boofny/golive-docker/general) 
