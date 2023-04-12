FROM node as build

COPY ./openapi ./openapi

WORKDIR /app

COPY website/public ./public
COPY website/src ./src
COPY website/package.json ./
COPY website/tsconfig.json ./

RUN npm install
RUN npm run generate
RUN npm run build

FROM arm64v8/nginx

LABEL org.opencontainers.image.documentation="https://github.com/ozonophore/yoda/blob/master/website/README.md"

COPY --from=build /app/build /usr/share/nginx/html
#COPY website/nginx/default.prod.conf /etc/nginx/conf.d/default.conf
COPY website/nginx/default.prod.conf /etc/nginx/templates/default.conf.template

VOLUME /var/log/nginx/log

EXPOSE 80