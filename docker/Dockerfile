#FROM mysql:8.0.26
#FROM golang:1.16 AS base
FROM ubuntu:20.04

RUN apt-get update
#RUN apt-get install \
#    install mysql-server 
 

# Set the working directory
WORKDIR /

##COPY . .


COPY /bank-api /usr/bin


# Remove all soure code
RUN cd / && \
    rm -rf /mockup

# return to root
WORKDIR /

CMD ["sh","-c","/usr/bin/bank-api"]


