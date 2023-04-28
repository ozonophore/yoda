FROM node as build

COPY ./openapi ./openapi

WORKDIR /app

COPY website/public website/public
COPY website/src website/src
COPY website/package.json website/
COPY website/tsconfig.json website/

RUN npm install
RUN npm run generate
RUN npm run build

FROM arm64v8/nginx as prod_arm

LABEL org.opencontainers.image.documentation="https://github.com/ozonophore/yoda/blob/master/website/README.md"

COPY --from=build /app/website/build /usr/share/nginx/html
#COPY website/nginx/default.prod.conf /etc/nginx/conf.d/default.conf
COPY website/nginx/default.prod.conf /etc/nginx/templates/default.conf.template

VOLUME /var/log/nginx/log

EXPOSE 80

FROM nginx as prod

LABEL org.opencontainers.image.documentation="https://github.com/ozonophore/yoda/blob/master/website/README.md"

COPY --from=build /app/website/build /usr/share/nginx/html
#COPY website/nginx/default.prod.conf /etc/nginx/conf.d/default.conf
COPY website/nginx/default.prod.conf /etc/nginx/templates/default.conf.template

VOLUME /var/log/nginx/log

EXPOSE 80