#/*
# * Copyright (C) 2020 Intel Corporation
# * SPDX-License-Identifier: BSD-3-Clause
# */

FROM centos:8

RUN yum -y install curl
RUN useradd -u 1001 --home /opt/isecl-k8s-extensions  --shell /bin/false isecl-controller
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
RUN chmod +x kubectl
RUN mkdir -p /var/log/isecl-k8s-extensions
COPY isecl-k8s-controller /isecl-k8s-controller

ENTRYPOINT ["/isecl-k8s-controller"]
