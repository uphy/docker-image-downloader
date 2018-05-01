# Docker Image Downloader

Docker Image Downloader downloads the Docker images with `docker in docker`.

## Usage

Download `hello-world:latest`.

```
$ docker run -it --rm --privileged -v "`pwd`:/data" uphy/image-downloader -out image.tar hello-world:latest
```

Load the image to Docker.

```
$ docker image load --input image.tar
```
