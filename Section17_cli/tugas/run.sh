#!/bin/bash

tanggal=$(date +"%a %b %d %T %Z %Y")

mkdir "reafa  $tanggal"

mkdir -p "reafa  $tanggal/about_me/personal"
mkdir -p "reafa  $tanggal/about_me/professional"

mkdir "reafa  $tanggal/my_friends"

mkdir "reafa  $tanggal/my_system_info"

echo "https://www.facebook.com/$1" > "reafa  $tanggal/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$2" > "reafa  $tanggal/about_me/professional/linkedin.txt"

curl "https://raw.githubusercontent.com/mrefawaladam/automate/main/friends.txt" > "reafa  $tanggal/my_friends/list_of_my_friends.txt"

echo "My username: $USER" > "reafa  $tanggal/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "reafa  $tanggal/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "reafa  $tanggal/my_system_info/internet_connection.txt"

tree "reafa  $tanggal"
