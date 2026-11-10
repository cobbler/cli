# vim: ft=dockerfile
# hadolint ignore=DL3006
#
# One-off build target for openSUSE Leap 16.0 / SLES 16.0 RPMs, used for
# manual Cobbler 4.0.0 release validation. Mirrors
# packaging/docker/openSUSE_Tumbleweed/Dockerfile but targets Leap 16.0,
# which is the actual release-validation target distro.
FROM registry.opensuse.org/opensuse/leap:16.0

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
