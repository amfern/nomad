# docker build -f dev/docker-dev-env/Dockerfile -t nomad-dev-env .
# sudo docker run -ti -v `pwd`:/go/src/github.com/hashicorp/nomad nomad-dev-env bash

FROM debian:sid

RUN apt update -y && \
    apt install -y \
        build-essential \
        git \
        golang \
        liblxc1

# --------------------- start android ndk --------------------------
WORKDIR /opt

ENV ANDROID_NDK /opt/android-ndk-linux
ENV ANDROID_NDK_HOME /opt/android-ndk-linux

RUN apt-get update && apt-get install -y --no-install-recommends \
	                            unzip \
	                            wget \
                              python

RUN wget -q --output-document=android-ndk.zip https://dl.google.com/android/repository/android-ndk-r17-linux-x86_64.zip && \
	  unzip android-ndk.zip && \
	  rm -f android-ndk.zip && \
	  mv android-ndk-r17 android-ndk-linux
# --------------------- end android ndk --------------------------

# configure golang
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# install gomobile
RUN go get golang.org/x/mobile/cmd/gomobile
RUN gomobile init -ndk /opt/android-ndk-linux

WORKDIR /go/src/github.com/hashicorp/nomad

# -------------- setup toolchain -----------------
env ARCH arm
env TOOLCHAIN_DIR /opt/my-toolchain
RUN $ANDROID_NDK/build/tools/make_standalone_toolchain.py \
    --arch $ARCH \
    --api 26 \
    --stl=libc++ \
    --install-dir=$TOOLCHAIN_DIR

RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 go install std

# Add the standalone toolchain to the search path.
ENV PATH $PATH:$TOOLCHAIN_DIR

COPY . .
RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 go build -o anomad
