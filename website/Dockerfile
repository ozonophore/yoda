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

FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html
COPY website/nginx/default.prod.conf /etc/nginx/conf.d/default.conf

VOLUME /var/log/nginx/log

EXPOSE 80