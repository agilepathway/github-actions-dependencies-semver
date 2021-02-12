#!/bin/bash

GAUGE_VERSION=v1.1.7
TRUNCATED_GAUGE_VERSION="${GAUGE_VERSION:1}"

wget https://github.com/getgauge/gauge/releases/download/$GAUGE_VERSION/gauge-$TRUNCATED_GAUGE_VERSION-linux.x86_64.zip
unzip -o gauge-$TRUNCATED_GAUGE_VERSION-linux.x86_64.zip -d /usr/local/bin
su codespace -c "gauge install html-report"
# su codespace -c "gauge install python"
# add more plugins here as required
