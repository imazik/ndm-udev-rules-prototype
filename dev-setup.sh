#!/bin/bash
RED='\033[0;31m'
NC='\033[0m'

echo "Copy custom rules file inside udev rules path."
sudo cp custom.rules /etc/udev/rules.d/custom.rules
sudo chmod 777 /etc/udev/rules.d/custom.rules

echo -e "Going inside ${RED}$PWD/ndm-udev-util${NC}."
cd ndm-udev-util
echo "Building ndm-udev-util binary."
go build
echo -e "Coping ${RED}$PWD/ndm-udev-util${NC} binary to ${RED}/bin/ndm-udev-util${NC}."
sudo cp ndm-udev-util /bin/ndm-udev-util
echo -e "Deleting ${RED}$PWD/ndm-udev-util${NC} binary."
rm -f ndm-udev-util
cd ..
echo -e "Going back ${RED}$PWD${NC}."

echo "Going inside ndm-rest-api-module folder."
echo -e "Going inside ${RED}$PWD/ndm-rest-api-module${NC}."
cd ndm-rest-api-module
echo "Building ndm-rest-api-module binary."
go build
echo -e "Coping ${RED}$PWD/ndm-rest-api-module${NC} binary to ${RED}/bin/ndm-rest-api-module${NC}."
sudo cp ndm-rest-api-module /bin/ndm-rest-api-module
echo -e "Deleting ${RED}$PWD/ndm-rest-api-module${NC} binary."
rm -f ndm-rest-api-module
cd ..
echo -e "Going back ${RED}$PWD${NC}."

echo "Reloading udevadm."
sudo udevadm control --reload-rules
