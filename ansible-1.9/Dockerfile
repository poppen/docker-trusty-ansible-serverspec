FROM buildpack-deps:trusty

RUN apt-get update && apt-get install -y --no-install-recommends \
                libpython2.7-dev \
                python-pip \
                software-properties-common \
        && apt-add-repository -y ppa:brightbox/ruby-ng \
        && apt-get update && apt-get install -y --no-install-recommends \
                rake \
                ruby2.3 \
        && rm -rf /var/lib/apt/lists/*

RUN pip install ansible-lint
RUN pip install 'ansible<2.0'

RUN gem install serverspec
