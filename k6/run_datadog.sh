#!/bin/bash
DOCKER_CONTENT_TRUST=1 \
docker run -d \
    --name datadog \
    -v /var/run/docker.sock:/var/run/docker.sock:ro \
    -v /proc/:/host/proc/:ro \
    -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
    -e DD_SITE="datadoghq.com" \
    -e DD_API_KEY=996960a8d30f0f1b8db78d7d88a54efd \
    -e DD_DOGSTATSD_NON_LOCAL_TRAFFIC=1 \
    -p 8125:8125/udp \
    datadog/agent:latest