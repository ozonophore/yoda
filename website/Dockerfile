FROM node as build

WORKDIR /app

COPY ./public ./public
COPY ./src ./src
COPY ./package.json ./
COPY ./package-lock.json ./
COPY ./tsconfig.json ./

RUN npm install
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html
# COPY ./nginx/default.prod.conf /etc/nginx/conf.d/default.conf

VOLUME /var/log/nginx/log

EXPOSE 80