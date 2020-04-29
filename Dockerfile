# docker build -t jasonkofo/router:master  --build-arg SSH_KEY="`cat ~/.ssh/id_rsa`" .

FROM golang:1.13

ARG SSH_KEY

# Authorize SSH Host
RUN mkdir -p /root/.ssh && \
	chmod 0700 /root/.ssh && \
	ssh-keyscan github.com > /root/.ssh/known_hosts

# We need this key so that we can read our private IMQS git repos from github
RUN echo "$SSH_KEY" > /root/.ssh/id_rsa && \
	chmod 600 /root/.ssh/id_rsa

RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

RUN mkdir /build
WORKDIR /build

# Cache downloads
COPY go.mod go.sum /build/
RUN go mod download

# Compile
COPY . /build/
RUN go build main.go

EXPOSE 2001
