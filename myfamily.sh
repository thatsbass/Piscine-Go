
curl -s "https://learn.zone01dakar.sn/assets/superhero/all.json" | jq -c --argjson ID $HERO_ID '.[] | select(.id == $ID) | .connections.relatives'| sed 's/"//g'
