#!/usr/bin/env bash
run_mysql_and_service_and_migration() {
    run_mysql_and_service_without_migration
    run_migration_migrate
}

run_mysql_and_service_without_migration() {
    echo "Start Running Mysql and Service"
    docker-compose up -d
    echo "Mysql and Service Started"
}

run_migration_up() {
    echo "### Starting migration UP ###"
    cd migrations
    ls
    cd ../
    read -p "Version : " version
    docker run -v "$PWD"/migrations:/migrations --network ina-adm-regions_net migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(mysql:3306)/test_db" up $version
    echo "Finished migration version : "$version
}

run_migration_down() {
    echo "### Starting migration DOWN ###"
    cd migrations
    ls
    cd ../
    read -p "Version : " version
    docker run -v "$PWD"/migrations:/migrations --network ina-adm-regions_net migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(mysql:3306)/test_db" down $version
    echo "Finished Rollback migration version : "$version
}

run_download_dependencies() {
    go mod init ${PWD}
    go mod download
}

echo "\n===============Indonesia Administrative Regions===============\n"

echo "1. Run All (Mysql, Service without Migration)"
echo "2. Run Migration UP"
echo "3. Run Migration Down"
echo "4. Run All (Mysql, Service with Migration)"
echo "5. Download dependencies (create go.mod and go.sum)"

read -p "Pick options : " option

if [[ ${option} == 1 ]]
then
    run_download_dependencies
    run_mysql_and_service_without_migration
elif [[ ${option} == 2 ]]
then
    run_migration_up
elif [[ ${option} == 3 ]]
then
    run_migration_down
elif [[ ${option} == 4 ]]
then
    run_download_dependencies
    run_mysql_and_service_without_migration
    run_migration_up
elif [[ ${option} == 5 ]]
then
    run_download_dependencies
else
    run_download_dependencies
    run_mysql_and_service_without_migration
    run_migration_up
fi