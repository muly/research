
// write simple (with no external dependencies) web app in golang
file: hello-docker.go


'' write docker file
file: Dockerfile


// build docker
srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker build -t dockergo .
// run docker
srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker run --rm -p 8080:8080 dockergo


srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker build -t dockergo1 .
srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker run --rm -p 8080:8080 dockergo1


srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker build -t xyz .
srini@srini-ubuntu:~/gows/src/github.com/muly/research/golang/docker$ docker run --rm -p 8080:8080 xyz


