FROM ubuntu:14.04

RUN apt-get update --yes

RUN apt-get install --yes golang
#RUN apt-get install -y --yes nano 
#RUN apt-get install --yes wget 
#RUN apt-get install --yes dialog 
#RUN apt-get install --yes net-tools
#RUN apt-get install -y --yes nginx
#RUN rm -v /etc/nginx/nginx.conf
#ADD nginx.conf /etc/nginx/
#RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# Expose ports
#EXPOSE 80

# Set the default command to execute
# when creating a new container
#CMD service nginx start

WORKDIR /srv

CMD bash
