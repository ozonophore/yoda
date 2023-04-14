FROM homebrew/brew

RUN brew install golang-migrate

WORKDIR /migrations

COPY ./changelog /migrations/changelog
COPY ./run.sh /migrations/run.sh

ARG A_DB_HOST="host.docker.internal"
ARG A_DB_PORT="5432"
ARG A_DB_USER="user"
ARG A_DB_PASSWORD="password"
ARG A_DB_NAME="db"

ENV DB_HOST=$A_DB_HOST
ENV DB_PORT=$A_DB_PORT
ENV DB_USER=$A_DB_USER
ENV DB_PASSWORD=$A_DB_PASSWORD
ENV DB_NAME=$A_DB_NAME

CMD ["sh", "/migrations/run.sh"]