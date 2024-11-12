#! /bin/bash

# Requires xorriso (sudo apt-get install -y xorriso, sudo yum install xorriso -y, or sudo zypper install -y xorriso)
if [ -z "$1" ]
  then
    echo "No cobbler server url supplied"
fi

cobbler_commit=df356046f3cf27be62a61001b982d5983800cfd9 # 3.3.6 as of 2024-10-09
cobbler_branch=release33
iso_url=https://cdimage.ubuntu.com/ubuntu-legacy-server/releases/20.04/release/ubuntu-20.04.1-legacy-server-amd64.iso
iso_os=ubuntu
valid_iso_checksum=00a9d46306fbe9beb3581853a289490bc231c51f
iso_filename=$(echo ${iso_url##*/})
valid_extracted_iso_checksum=dd0b3148e1f071fb86aee4b0395fd63b
valid_git_checksum=6c9511b26946dd3f1f072b9f40eaeccf  # master as of 4/2/2022

[ -d "./testing/cobbler_source" ] && git_checksum=$(find ./testing/cobbler_source/ -type f -exec md5sum {} \; | sort -k 2 | md5sum | awk '{print $1}')
if [ -d "./testing/cobbler_source" ] && [ $git_checksum == $valid_git_checksum ]; then
  echo "Cobbler code already cloned and the correct version is checked out"
else
  rm -rf ./testing/cobbler_source
  git clone --shallow-since="2021-09-01" https://github.com/cobbler/cobbler.git -b $cobbler_branch testing/cobbler_source
  cd ./testing/cobbler_source
  printf "Changing to version of Cobbler being tested.\n\n"
  git checkout $cobbler_commit > /dev/null 2>&1
  rm -rf .git  # remove .git dir so the checksum is consistent
  cd -
fi

echo $(pwd)
if [ -f "$iso_filename" ] && [ $(sha1sum $iso_filename | awk '{print $1}') == "$valid_iso_checksum" ]; then
  echo "ISO already downloaded"
else
  rm $iso_filename
  wget $iso_url
fi

extracted_iso_checksum=$(find extracted_iso_image -type f -exec md5sum {} \; | sort -k 2 | md5sum | awk '{print $1}')
if [ -d "extracted_iso_image" ] && [ $extracted_iso_checksum == $valid_extracted_iso_checksum ]; then
   echo "ISO already extracted"
else
   xorriso -osirrox on -indev $iso_filename -extract / extracted_iso_image
fi

docker build -f ./testing/cobbler_source/docker/develop/develop.dockerfile -t cobbler-dev .
docker compose -f testing/compose.yml up -d

SERVER_URL=$1
printf "### Waiting for Cobbler to become available on ${SERVER_URL} \n\n"

attempt_counter=0
max_attempts=48

until $(curl --connect-timeout 1 --output /dev/null --silent ${SERVER_URL}); do
  if [ ${attempt_counter} -eq ${max_attempts} ];then
    echo "Max attempts reached"
    # Debug logs
    docker compose -f ./testing/compose.yml logs
    exit 1
  fi

  attempt_counter=$(($attempt_counter+1))
  sleep 5
done

# Sleep 10 seconds to let the "cobbler import" succeed
sleep 10

docker compose -f testing/compose.yml logs