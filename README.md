# Kitchen Server
The Kitchen part. 

# Docker build and run:
```shell
docker build ./ -t kitchen_image
```
```shell
docker run -d --rm -p 7500:7500 --name kitchen_container kitchen_image go run main
```
```shell
docker stop kitchen_container
```
```shell
docker run -d --rm -p 7500:7500 --name kitchen_container kitchen_image go run main
```
```shell
docker build ./ -t kitchen_image
```
To remove Docker image and container:
```shell
docker stop kitchen_container
```
```shell
docker rm kitchen_container
```
```shell
docker rmi kitchen_image
```

# View in browser addresses:
To check if the kitchen server is running:
localhost:8000/
```
