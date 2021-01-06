#!/bin/bash
if [ -z $1 ]; then 
	echo "Missining domain argument\nExample: crtshgrab domain.com"
else
	curl https://crt.sh/\?q\=%25.$1\&output\=json 2> /dev/null | jq -r ".[].name_value" | sort -u | grep -v "*" | grep -v "@"
fi

