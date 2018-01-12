#!/usr/bin/env bash

main() {
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    local db_container_id=$(docker ps -aqf "name=clarakmprojectsgo_db_1")
    local path_to_sql_ddl_file=${SCRIPT_DIR}/seed-data.sql
    cat ${path_to_sql_ddl_file} | docker exec -i ${db_container_id} /usr/bin/mysql -uclara -ppassword clarakm-projects
}

main "$@"