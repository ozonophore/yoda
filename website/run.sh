docker run --name some-nginx -v build:/usr/share/nginx/html:ro -v nginx:/etc/nginx/conf.d:ro -d nginx
