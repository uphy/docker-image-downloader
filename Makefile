build:
	docker build -t uphy/image-downloader .
	docker run -it --rm --privileged -v "`pwd`:/data" uphy/image-downloader -out /data/image.tar hello-world hello-world