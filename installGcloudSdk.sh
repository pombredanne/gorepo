#!/bin/bash
gcloudpkg=google-cloud-sdk-166.0.0-linux-x86_64.tar.gz

cd ${HOME}

if [ -d google-cloud-sdk ]
then
	echo "gcloud dir already exists"
	exit
fi

read -r -p "Install $gcloudpkg [y/N] " resp
if [[ ! "$resp" =~ ^([yY][eE][sS]|[yY])+$ ]]
then
	exit
fi
echo "Starting"

if [ ! -e /tmp/$gcloudpkg ]
then
	wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/$gcloudpkg -O /tmp/$gcloudpkg
fi
tar -zxf /tmp/$gcloudpkg

./google-cloud-sdk/install.sh

echo -e "\nTo compleate the setup, open a new shell and run\ngcloud init"
