VERSION=${1:-v0.0.54}
PLATFORM=${2:-linux/amd64}

git clone --depth 1 --branch "$VERSION" https://github.com/replicase/pgcapture

cd pgcapture

echo "building rueian/pgcapture:$VERSION in $PLATFORM platform"

PLATFORM="$PLATFORM" ./dockerbuild.sh build

cd ../
rm -rf pgcapture
