FROM nginx:1.15.2-alpine
COPY ./nginx-cfg/nginx.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT ["nginx","-g","daemon off;"]