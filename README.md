# go-web-srv

    cd go-web-srv
    mkdir mysql-data portainer-data grafana-data
    docker swarm init --advertise-addr <ip>
    docker stack deploy -c docker-compose.yml servers
