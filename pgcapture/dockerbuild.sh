VERSION=${1:-v0.0.40}
PLATFORM=${2:-linux/amd64}

git clone --depth 1 --branch "$VERSION" https://github.com/rueian/pgcapture

cd pgcapture || exit
sha=$(git rev-parse --short HEAD)
cd ../

echo "building rueian/pgcapture:$VERSION with git sha: $sha in $PLATFORM platform"

docker buildx build \
  --platform linux/amd64 \
  --output type=docker \
  --build-arg VERSION="$VERSION" \
  --build-arg SHA="$sha" \
  -t rueian/pgcapture:"$VERSION" pgcapture

rm -rf pgcapture
