FROM ubuntu:18.04

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    ca-certificates \
    xterm \
    tigervnc-common \
    tigervnc-standalone-server \
    tigervnc-viewer \
    xdotool \
    curl \
    git \
    xauth \
    python-numpy \
    libglib2.0-bin \
    gvfs-bin \
    libgconf2-4 \
    libgtk2.0-0 \
    libgtk-3-0 \
    libcanberra-gtk3-module \
    libx11-xcb-dev \
    libgconf2-dev \
    libasound2 \
    libxtst-dev \
    libxss1 \
    libnotify4 \
    libnss3 \
    xdg-utils && \
    rm -rf /var/lib/apt/lists/*

RUN curl -L "https://downloads.sourceforge.net/project/turbovnc/2.2/turbovnc_2.2_amd64.deb?r=https%3A%2F%2Fsourceforge.net%2Fprojects%2Fturbovnc%2Ffiles%2F2.2%2Fturbovnc_2.2_amd64.deb%2Fdownload&ts=1544470485" -o ~/turbovnc.deb
RUN dpkg -i ~/turbovnc.deb

RUN curl -L https://github.com/felixrieseberg/windows95/releases/download/v1.3.0/windows95-linux_1.3.0_amd64.deb -o ~/windows95.deb
RUN dpkg -i ~/windows95.deb

RUN git clone --depth=1 https://github.com/kanaka/noVNC.git /var/novnc && ln -s /var/novnc/vnc_lite.html /var/novnc/index.html
RUN git clone --depth=1 https://github.com/kanaka/websockify /var/novnc/utils/websockify


# Create and configure the VNC user
ARG VNCPASS
ENV VNCPASS ${VNCPASS:-secret}

RUN useradd remote --create-home --shell /bin/bash --user-group --groups adm,sudo && \
    echo "remote:$VNCPASS" | chpasswd

EXPOSE 5900

COPY main.sh /

ENV VNC_PWD=secret

ENTRYPOINT ["/main.sh"]
CMD ["default"]


