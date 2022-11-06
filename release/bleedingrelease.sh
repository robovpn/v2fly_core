#!/usr/bin/env bash

RELBODY="https://github.com/v2fly/v2fly-core/commit/${RELEASE_SHA}"
JSON_DATA=$(echo "{}" | jq -c ".tag_name=\"${RELEASE_TAG}\"")
JSON_DATA=$(echo ${JSON_DATA} | jq -c ".prerelease=${PRERELEASE}")
JSON_DATA=$(echo ${JSON_DATA} | jq -c ".body=\"${RELBODY}\"")
RELEASE_DATA=$(curl --data "${JSON_DATA}" -H "Authorization: token ${GITHUB_TOKEN}" -X POST https://api.github.com/repos/v2fly/V2FlyBleedingEdgeBinary/releases)
echo $RELEASE_DATA
RELEASE_ID=$(echo $RELEASE_DATA | jq ".id")

function uploadfile() {
  FILE=$1
  CTYPE=$(file -b --mime-type $FILE)

  sleep 1
  curl -H "Authorization: token ${GITHUB_TOKEN}" -H "Content-Type: ${CTYPE}" --data-binary @$FILE "https://uploads.github.com/repos/v2fly/V2FlyBleedingEdgeBinary/releases/${RELEASE_ID}/assets?name=$(basename $FILE)"
  sleep 1
}

function upload() {
  FILE=$1
  DGST=$1.dgst
  openssl dgst -md5 $FILE | sed 's/([^)]*)//g' >>$DGST
  openssl dgst -sha1 $FILE | sed 's/([^)]*)//g' >>$DGST
  openssl dgst -sha256 $FILE | sed 's/([^)]*)//g' >>$DGST
  openssl dgst -sha512 $FILE | sed 's/([^)]*)//g' >>$DGST
  uploadfile $FILE
  uploadfile $DGST
}

ART_ROOT=${WORKDIR}/bazel-bin/release

pushd ${ART_ROOT}
{
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen version ${RELEASE_TAG}
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen project "v2flyunstable"
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-macos-64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-windows-64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-windows-32.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-windows-arm32-v7a.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-32.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-arm64-v8a.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-arm32-v7a.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-arm32-v6.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-arm32-v5.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-mips64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-mips64le.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-mips32.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-mips32le.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-ppc64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-ppc64le.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-riscv64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-linux-s390x.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-freebsd-64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-freebsd-32.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-openbsd-64.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-openbsd-32.zip
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen file v2fly-dragonfly-64.zip
} >Release.unsigned.unsorted
  go run github.com/v2fly/V2BuildAssist/v2buildutil gen sort < Release.unsigned.unsorted > Release.unsigned

  {
    echo "Build Finished"
    echo "https://github.com/v2fly/V2FlyBleedingEdgeBinary/releases/tag/${RELEASE_TAG}"
  } > buildcomment

  go run github.com/v2fly/V2BuildAssist/v2buildutil post commit "${RELEASE_SHA}" < buildcomment
popd

upload ${ART_ROOT}/v2fly-macos-64.zip
upload ${ART_ROOT}/v2fly-windows-64.zip
upload ${ART_ROOT}/v2fly-windows-32.zip
upload ${ART_ROOT}/v2fly-windows-arm32-v7a.zip
upload ${ART_ROOT}/v2fly-linux-64.zip
upload ${ART_ROOT}/v2fly-linux-32.zip
upload ${ART_ROOT}/v2fly-linux-arm64-v8a.zip
upload ${ART_ROOT}/v2fly-linux-arm32-v7a.zip
upload ${ART_ROOT}/v2fly-linux-arm32-v6.zip
upload ${ART_ROOT}/v2fly-linux-arm32-v5.zip
upload ${ART_ROOT}/v2fly-linux-mips64.zip
upload ${ART_ROOT}/v2fly-linux-mips64le.zip
upload ${ART_ROOT}/v2fly-linux-mips32.zip
upload ${ART_ROOT}/v2fly-linux-mips32le.zip
upload ${ART_ROOT}/v2fly-linux-ppc64.zip
upload ${ART_ROOT}/v2fly-linux-ppc64le.zip
upload ${ART_ROOT}/v2fly-linux-riscv64.zip
upload ${ART_ROOT}/v2fly-linux-s390x.zip
upload ${ART_ROOT}/v2fly-freebsd-64.zip
upload ${ART_ROOT}/v2fly-freebsd-32.zip
upload ${ART_ROOT}/v2fly-openbsd-64.zip
upload ${ART_ROOT}/v2fly-openbsd-32.zip
upload ${ART_ROOT}/v2fly-dragonfly-64.zip
upload ${ART_ROOT}/Release.unsigned
