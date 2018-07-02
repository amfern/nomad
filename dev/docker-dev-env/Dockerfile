# docker build -f dev/docker-dev-env/Dockerfile -t nomad-dev-env .
# docker run -ti -v /tmp:/tmp nomad-dev-env bash -c "cp /output/nomad /tmp/"
# adb root
# adb remount
# adb push /tmp/nomad /system/xbin/

# cd /storage/emulated/0
# nomad agent -config /storage/emulated/0/server.hcl
# nomad agent -config client1.hcl -network-interface=wlan0
# nomad agent -config client2.hcl -network-interface=wlan0

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
                              curl \
                              python \
                              apt-transport-https

RUN wget -q --output-document=android-ndk.zip https://dl.google.com/android/repository/android-ndk-r17-linux-x86_64.zip && \
	  unzip android-ndk.zip && \
	  rm -f android-ndk.zip && \
	  mv android-ndk-r17 android-ndk-linux
# --------------------- end android ndk --------------------------

# configure golang
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# install gomobile
# RUN go get golang.org/x/mobile/cmd/gomobile
# RUN gomobile init -ndk /opt/android-ndk-linux

WORKDIR /go/src/github.com/hashicorp/nomad

# -------------- setup toolchain -----------------
env ARCH arm
env TOOLCHAIN_DIR /opt/my-toolchain
RUN $ANDROID_NDK/build/tools/make_standalone_toolchain.py \
    --arch $ARCH \
    --api 28 \
    --stl=libc++ \
    --install-dir=$TOOLCHAIN_DIR

RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 GOARM=7 go install std

RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 GOARM=7 go install os/user


# Add the standalone toolchain to the search path.
ENV PATH $PATH:$TOOLCHAIN_DIR

COPY . .

# make dev-ui
# RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
# RUN apt-get install -y nodejs

# RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
# RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
# RUN apt-get update && \
#             apt-get install -y --no-install-recommends yarn

# RUN make dev-ui

RUN mkdir /output
RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 GOARM=7 go build -x -o /output/nomad
