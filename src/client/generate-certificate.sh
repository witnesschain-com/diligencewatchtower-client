#!/bin/sh

SUDO="sudo"

$SUDO 2>/dev/null

if [ "$?" = "127" ] # command sudo NOT found
then
	SUDO=""
fi

host_name=`cat config.json  | grep \"host_name\" | cut -f4 -d'"'`

echo "$host_name"

if [ "$host_name" != "" ]
then
	echo "Trying to get TLS certificate for '$host_name' ..."
 
	$SUDO certbot certonly --force-renewal -n --register-unsafely-without-email --agree-tos --standalone -d $host_name
else
	echo "===> WARNING: No 'host_name' found in 'config.json'"
fi
