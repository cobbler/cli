# vim: ft=dockerfile
# hadolint ignore=DL3006
FROM opensuse/tumbleweed

# Create work directory
WORKDIR /workspace
VOLUME /workspace

# Install required dependencies
RUN zypper -n --gpg-auto-import-keys refresh && \
    zypper -n install \
        rpm-build \
        go \
    	bash-completion \
		zsh \
    	fish \
        git \
        fdupes

# Set up RPM build directories
RUN mkdir -p ~/rpmbuild/{BUILD,RPMS,SOURCES,SPECS,SRPMS}

ENTRYPOINT ["make", "build-rpm"]
