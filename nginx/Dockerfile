FROM nginx:1.21-alpine

RUN mkdir /static/
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d