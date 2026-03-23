FROM ubuntu:22.04

SHELL ["/bin/bash", "-c"]

RUN apt-get update && \
    apt-get install -y openjdk-8-jdk wget unzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN wget -O kotlin.zip https://github.com/JetBrains/kotlin/releases/download/v1.9.24/kotlin-compiler-1.9.24.zip && \
    unzip kotlin.zip -d /opt && \
    rm kotlin.zip && \
    chmod +x /opt/kotlinc/bin/kotlinc

RUN wget -O gradle.zip https://services.gradle.org/distributions/gradle-8.14-bin.zip && \
    unzip gradle.zip -d /opt && \
    rm gradle.zip

ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH="/opt/kotlinc/bin:/opt/gradle-8.14/bin:${PATH}"

WORKDIR /app
COPY . /app

CMD ["/bin/bash"]