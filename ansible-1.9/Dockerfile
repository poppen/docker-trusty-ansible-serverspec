FROM buildpack-deps:trusty

RUN apt-get update && apt-get install -y --no-install-recommends \
                python2.7 \
                libpython2.7-dev \
                software-properties-common \
        && curl -fSL 'https://bootstrap.pypa.io/get-pip.py' | python2 \
        && apt-add-repository -y ppa:brightbox/ruby-ng \
        && apt-get update && apt-get install -y --no-install-recommends \
                rake \
                ruby2.3 \
        && rm -rf /var/lib/apt/lists/*

RUN pip install --no-cache-dir 'ansible<2.0'
RUN pip install --no-cache-dir ansible-lint

RUN gem install serverspec
