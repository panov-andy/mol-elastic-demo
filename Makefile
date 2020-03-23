


start-elastic:
	docker run --rm -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.1
	wget --no-verbose --retry-connrefused -O- http://localhost:9200
	@echo "Elasticsearch started..."

stop-elastic:
	docker stop elasticsearch