FROM mcr.microsoft.com/vscode/devcontainers/universal:linux

COPY scripts/* /tmp/scripts/

USER root

RUN /tmp/scripts/install-hadolint.sh \
    # && /tmp/scripts/install-gauge.sh \
    && rm -rf /tmp/scripts

USER codespace
