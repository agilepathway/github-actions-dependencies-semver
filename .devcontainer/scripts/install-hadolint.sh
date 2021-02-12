#!/bin/bash

export HADOLINT_VERSION=v1.22.1

mkdir -p /tmp/hadolint
cd /tmp/hadolint || exit
wget https://github.com/hadolint/hadolint/releases/download/$HADOLINT_VERSION/hadolint-Linux-x86_64
mv hadolint-Linux-x86_64 /usr/local/bin/hadolint
chmod +x /usr/local/bin/hadolint
rm -rf /tmp/hadolint
