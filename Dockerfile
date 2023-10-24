FROM node

WORKDIR /app

RUN npm install -g create-react-app
RUN npm install -g openapi

COPY website /app

ENTRYPOINT ["bash"]