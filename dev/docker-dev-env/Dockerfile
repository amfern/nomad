# docker build -f dev/docker-dev-env/Dockerfile -t nomad-dev-env .
# docker run -ti -v /tmp:/tmp nomad-dev-env bash -c "cp /output/nomad /tmp/"
# adb root
# adb remount
# adb push /tmp/nomad /system/xbin/

# mkdir /data/data/nomad
# cd /data/data/nomad
# vim /data/system/packages.list
# nomad 10113 0 /data/user/0/nomad default:targetSdkVersion=27 3003
# nomad agent -config /storage/emulated/0/server.hcl
# ./nomad agent -config client1.hcl -network-interface=wlan0
# ./nomad agent -config client2.hcl -network-interface=wlan0

# ./dist/nomad job run dev/docker-dev-env/example.nomad

# nomad agent -dev -data-dir=`pwd` -network-interface=wlan0

FROM golang:rc

RUN apt update -y && \
    apt install -y \
        build-essential \
        git \
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
                              apt-transport-https \
                              libtinfo5

RUN wget -q --output-document=android-ndk.zip https://dl.google.com/android/repository/android-ndk-r17-linux-x86_64.zip && \
	  unzip android-ndk.zip && \
	  rm -f android-ndk.zip && \
	  mv android-ndk-r17 android-ndk-linux
# --------------------- end android ndk --------------------------

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

RUN go get -v # COPY . .

# make dev-ui
# RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
# RUN apt-get install -y nodejs

# RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
# RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
# RUN apt-get update && \
#             apt-get install -y --no-install-recommends yarn

# RUN make dev-ui
# RUN go build -x -o /output/nomad

# RUN gomobile build -target=android
RUN mkdir /output
RUN CC=$TOOLCHAIN_DIR/bin/clang CXX=$TOOLCHAIN_DIR/bin/clang++ GOOS=android GOARCH=$ARCH CGO_ENABLED=1 GOARM=7 go build -o /output/nomad



# ############################## example output on debian ############################
# default via 10.0.0.138 dev eno1
# ip route
# 10.0.0.0/24 dev eno1 proto kernel scope link src 10.0.0.18 
# 172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1 linkdown 
# 192.168.100.0/24 dev virbr0 proto kernel scope link src 192.168.100.1

# ####################### example output on android ####################
# ip route
# 10.0.0.0/24 dev wlan0  proto kernel  scope link  src 10.0.0.32

# android doesn't have concept of default gateway
# in android the first interface in the list is the default
# need to add new file similar to https://github.com/hashicorp/go-sockaddr/blob/e12d9401a74f025fe672cd1a84b2081c773990d3/route_info_linux.go, route_info_android.go that will select the first interface on the list

