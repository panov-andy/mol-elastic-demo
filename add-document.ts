#!/usr/local/bin/deno --allow-net

let host = "http://localhost:9200";

const data = await fetch(host);
const json = await data.json();
console.log(JSON.stringify(json));

let test_data = {
    answer: 42
};

fetch(`${host}/my-catalog/_doc/1`, {
    method: "PUT",
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(test_data),
});

//curl localhost:9200/my-catalog/_search?size=10 | jq