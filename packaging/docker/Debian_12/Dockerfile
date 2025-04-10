# vim: ft=dockerfile
FROM docker.io/library/debian:12

ENV DEBIAN_FRONTEND=noninteractive

# Create work directory
WORKDIR /workspace
VOLUME /workspace

# hadolint ignore=DL3008
RUN apt-get update -qq && \
    apt-get install --no-install-recommends -qqy \
    build-essential  \
    debhelper  \
    devscripts \
    dh-python  \
    rename \
    git \
    bash-completion \
    golang-go && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Make directory for debs
RUN mkdir -p /usr/src/cobbler-cli/deb-build
VOLUME /usr/src/cobbler-cli/deb-build

ENTRYPOINT ["make", "build-deb"]
