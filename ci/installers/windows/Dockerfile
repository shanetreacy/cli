FROM debian:wheezy
RUN apt-get update
#Install GCC
RUN apt-get -y install build-essential

#Install Wine
RUN dpkg --add-architecture i386
RUN apt-get install --no-install-recommends --assume-yes wine

#Install Unzip
RUN apt-get install unzip

#Install OpenSSL
RUN apt-get -y install openssl
RUN apt-get -y install libssl-dev
RUN apt-get -y install libcurl4-gnutls-dev
RUN apt-get -y install automake autoconf

#Install OSSLSignCode
ADD http://downloads.sourceforge.net/project/osslsigncode/osslsigncode/osslsigncode-1.7.1.tar.gz?r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fosslsigncode%2F&ts=1474579286&use_mirror=heanet /osslsigncode-1.7.1.tar.gz
RUN mkdir osslsigncode
RUN tar -xvzf /osslsigncode-1.7.1.tar.gz -C /osslsigncode
WORKDIR /osslsigncode/osslsigncode-1.7.1
RUN export LDFLAGS=-L/usr/local/opt/openssl/lib
RUN export CPPFLAGS=-I/usr/local/opt/openssl/include
RUN export PKG_CONFIG_PATH=/usr/local/opt/openssl/lib/pkgconfig

RUN ./configure
RUN make
RUN make install

#Download Innosetup 5
WORKDIR /
ADD https://s3.amazonaws.com/ci-cli-assets/Inno+Setup+5.zip /
RUN unzip ./Inno+Setup+5.zip

#Download the ISCC script
ADD https://gist.githubusercontent.com/derekstavis/8288379/raw/678c72b66e14dbedff4a5b7786a5235aea18e983/iscc /
