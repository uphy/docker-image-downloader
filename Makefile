build:
	docker build -t uphy/image-downloader .
	docker run -it --rm -v "`pwd`:/data" uphy/image-downloader hello-world:latest