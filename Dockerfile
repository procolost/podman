FROM gitpod/workspace-full-vnc
RUN sudo apt-get update && \
    sudo apt-get install -y proxychains4 libgtk-3-dev  wireguard-tools && \
    sudo rm -rf /var/lib/apt/lists/* && sudo apt clean
