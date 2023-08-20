VERSION=${1:-v0.0.52}
PLATFORM=${2:-linux/amd64}

git clone --depth 1 --branch "$VERSION" https://github.com/rueian/pgcapture

cd pgcapture

echo "building rueian/pgcapture:$VERSION in $PLATFORM platform"

PLATFORM="$PLATFORM" ./dockerbuild.sh build

cd ../
rm -rf pgcapture
