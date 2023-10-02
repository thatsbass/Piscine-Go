curl -s "https://learn.zone01dakar.sn/assets/superhero/all.json" | jq '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
