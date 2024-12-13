# looking for a fix run as root

mkdir /mnt/storage/

mkdir /mnt/storage/files

cd filehandler/

docker compose -f "compose.yml" up -d --build

cd file-host/

docker compose -f "docker-compose.yml" up -d --build

cd ..

cd ..

cp html/ /var/www/ -r

echo "Done!"
