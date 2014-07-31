from debian

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get -y install golang git

RUN mkdir -p /var/www/staticd
RUN echo "OK" > /var/www/staticd/README.txt

RUN go get github.com/j4/staticd
CMD staticd.go -port 8080 -path /var/www/staticd

EXPOSE 8080