#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

passphrase="test1234"

# RED=`tput setaf 1`
# GREEN=`tput setaf 2`
# RESET=`tput sgr0`
GREEN=""
RED=""
RESET=""

random_string() {
  local __resultvar=$1
  local length=${2:-6} # Default is 6
  # Newer mac might have shasum instead of sha1sum
  if  command -v shasum &> /dev/null
    then
      eval $__resultvar="'$(date +%s.%N | shasum | fold -w ${length} | head -n 1)'"
    else
      eval $__resultvar="'$(date +%s.%N | sha1sum | fold -w ${length} | head -n 1)'"
  fi
}

check_response() {
  result=$1
  expected_string=$2
  if [[ $result != *$expected_string* ]]; then
    echo "${GREEN}ERROR:${RESET} command failed. The expected string: $expected_string not found in the result: $result"
    exit 1
  fi
}

check_response_and_report() {
  result=$1
  expected_string=$2
  check_response "$result" "$expected_string"
  echo "${GREEN}SUCCESS: ${RESET} Result contains expected substring: $expected_string"
}

response_does_not_contain() {
  result=$1
  unexpected_string=$2
  if [[ $result == *$unexpected_string* ]];then
    echo "ERROR: command failed. The unexpected string: $unexpected_string found in the result: $result"
    exit 1
  fi
  echo "${GREEN}SUCCESS: ${RESET}Result does not contain unexpected substring: $unexpected_string"
}

create_new_account(){
  local  __resultvar=$1
  random_string name
  eval $__resultvar="'$name'"

  roles=$2

  echo "Account name: $name"

  echo "Generate key for $name"
  echo $passphrase | dclcli keys add "$name"

  address=$(dclcli keys show $name -a)
  pubkey=$(dclcli keys show $name -p)

  echo "Jack proposes account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dclcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=$roles --from jack --yes)
  check_response "$result" "\"success\": true"
  echo "$result"

  echo "Alice approves account for \"$name\" with roles: \"$roles\""
  result=$(echo $passphrase | dclcli tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"success\": true"
  echo "$result"
}

create_new_vendor_account(){

  _name=$1
  _vid=$2

  echo $passphrase | dclcli keys add "$_name"
  address=$(dclcli keys show $_name -a)
  pubkey=$(dclcli keys show $_name -p)

  test_divider

  echo "Jack proposes account for \"$_name\" with Vendor role"
  result=$(echo $passphrase | dclcli tx auth propose-add-account --address="$address" --pubkey="$pubkey" --roles=Vendor --vid=$_vid --from jack --yes)
  check_response "$result" "\"success\": true"

  test_divider

  echo "Alice approves account for \"$_name\" with Vendor role"
  result=$(echo $passphrase | dclcli tx auth approve-add-account --address="$address" --from alice --yes)
  check_response "$result" "\"success\": true"

}

create_model_and_version() {
  _vid=$1
  _pid=$2
  _softwareVersion=$3
  _softwareVersionString=$4
  _user_address=$5
  result=$(echo 'test1234' | dclcli tx model add-model --vid=$_vid --pid=$_pid --deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=$_user_address --yes)
  check_response "$result" "\"success\": true"
  result=$(echo 'test1234' | dclcli tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=$_vid --pid=$_pid --softwareVersion=$_softwareVersion --softwareVersionString=$_softwareVersionString --from=$_user_address --yes)
  check_response "$result" "\"success\": true"
}

test_divider() {
  echo ""
  echo "--------------------------"
  echo ""
}

