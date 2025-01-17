#
# Build libsqlitefunctions.so
#
# https://hub.docker.com/_/gcc?tab=tags
#
FROM docker.io/library/gcc:11@sha256:7e68daa92728bbb415cd66d3825357d67db65032baa36e3c631c80dad461bab2 AS gccbuild
WORKDIR /work
COPY ./build/extension-functions.c .
RUN gcc -fPIC -shared -O3 extension-functions.c -o libsqlitefunctions.so -lm

#
# Run datasette.
#
# https://hub.docker.com/r/datasetteproject/datasette/tags
#
FROM docker.io/datasetteproject/datasette:0.61.1@sha256:cca7d450d320f8f99ffbbbc4ec382c745f1a5a9fe3626f0ab781866f8f12a3d4
WORKDIR /work

# Install dependencies.
RUN apt-get -qq update \
 && apt-get -qq --no-install-recommends install \
      brotli \
      curl \
      gzip \
      varnish \
 && rm -rf /var/lib/apt/lists/*
# https://github.com/curiousleo/datasette-leaflet-geojson/archive/1e402abeb77192e0b8d51504b46055f1e1b4cf4d.tar.gz
RUN pip install \
      datasette-block-robots \
      datasette-cluster-map \
      datasette-graphql \
      datasette-vega \
 && true
# https://github.com/DarthSim/hivemind/releases
RUN curl -sSL https://github.com/DarthSim/hivemind/releases/download/v1.1.0/hivemind-v1.1.0-linux-amd64.gz \
    | gunzip - >/usr/bin/hivemind \
 && chmod +x /usr/bin/hivemind

COPY --from=gccbuild /work/libsqlitefunctions.so .
# The database files are renamed to .sqlite3 on purpose.
#
# The only way I can find in the documentation to load settings from a file is to use "Configuration
# directory mode", i.e. passing a directory to `datasette`.
#
# https://docs.datasette.io/en/0.60.1/settings.html#config-dir
#
# However, we also want users to be able to download the SQLite files. This requires us to pass
# `--immutable <dbfile>` to `datasette`.
#
# If we use `--immutable <dbfile>` in combination with "Configuration directory mode", Datasette
# will pick up any *.db files in the given directory and show them. If any of those SQLite files
# were passed in with `--immutable`, they will be shown twice: once as an immutable database, once
# as a mutable database.
#
# To avoid Datasette listing the databases twice, we give the SQLite files a file ending other than
# "db" so they are not picked up and listed as databases as a result of using "Configuration
# directory mode".
COPY ./build/Marktstammdatenregister.db.br ./Marktstammdatenregister.sqlite3.br
COPY ./build/Procfile .
COPY ./build/datasette .
COPY ./build/default.vcl .
COPY ./build/metadata.yaml .
COPY ./build/settings.json .
COPY ./build/varnish .

EXPOSE 8080

CMD hivemind
