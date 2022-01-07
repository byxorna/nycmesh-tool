#!/bin/bash
# used to fetch the auth token for use by controller
set -e
UISP_HOSTNAME="${UISP_HOSTNAME:-uisp.mesh}"
USERNAME="${USERNAME:-$(id)}"

echo "This script logs you in to UISP_HOSTNAME=$UISP_HOSTNAME, so you can add the token to ~/.nycmesh-tool.yaml as uisp.x-auth-token"

read -p "UISP Username ($USERNAME): " username
read -sp "UISP Password: " uisp_password

if [[ -n $username ]] ; then
  USERNAME="$username"
fi

curl -vk -X POST \
  -H  "accept: application/json" \
  -H  "Content-Type: application/json" \
  -d "{  \"username\": \"${USERNAME}\",  \"password\": \"$uisp_password\",  \"sessionTimeout\": 0}" \
 "https://${UISP_HOSTNAME}/nms/api/v2.1/user/login"

echo
echo "Your x-auth-token header should be printed above. Copy and paste that into ~/.nycmesh-tool.yaml as uisp.x-auth-token:"
