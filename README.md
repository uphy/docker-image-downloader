# Docker Image Downloader

Docker Image Downloader downloads the Docker images without `docker pull` command.

## Usage

Download `hello-world:latest`.

```
$ docker run -it --rm -v "`pwd`:/data" uphy/image-downloader hello-world:latest
```

Load the image to Docker.

```
$ docker image load --input hello-world\:latest.tar
```
