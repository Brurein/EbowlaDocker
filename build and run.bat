docker build -t ebowla .

docker run -v .:/3bowla/host --rm ebowla mimikatz.exe

docker image list

docker load --input ebowla.img --quiet

docker image list

docker rmi -f $(docker images -aq)