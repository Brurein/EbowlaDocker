FROM ubuntu:16.04

WORKDIR /pip

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install git curl python2.7 -y
RUN curl https://bootstrap.pypa.io/pip/2.7/get-pip.py --output get-pip.py
RUN python2.7 get-pip.py
RUN pip install configobj
RUN pip install pycryptodome
RUN pip install pyparsing

WORKDIR /go

RUN curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
RUN tar -xvf go1.6.linux-amd64.tar.gz
RUN mv go /usr/local
RUN export PATH=$PATH:/usr/local/go/bin
RUN apt-get install cmake build-essential g++-mingw-w64-x86-64 g++-mingw-w64-i686 gcc-mingw-w64-i686 gcc-mingw-w64-x86-64 -y
RUN apt-get install golang-go -y

WORKDIR /3bowla

RUN git clone https://github.com/Brurein/Ebowla.git

RUN cp -R /3bowla/Ebowla/MemoryModule /3bowla/

COPY entrypoint.sh .

RUN chmod +x /3bowla/entrypoint.sh

ENTRYPOINT /3bowla/entrypoint.sh