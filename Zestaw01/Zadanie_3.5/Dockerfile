FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y openjdk-8-jdk wget unzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN wget -O kotlin.zip https://github.com/JetBrains/kotlin/releases/download/v1.9.24/kotlin-compiler-1.9.24.zip && \
    unzip kotlin.zip -d /opt && \
    rm kotlin.zip && \
    chmod +x /opt/kotlinc/bin/kotlinc

ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH="/opt/kotlinc/bin:${PATH}"

CMD ["/bin/bash", "-c", "java -version && kotlinc -version"]