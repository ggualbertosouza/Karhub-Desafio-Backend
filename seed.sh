#!/usr/bin/env bash
set -e

BASE_URL="${APP_BASE_URL:-http://localhost:3001}"

echo "Populating beer styles at: $BASE_URL"
echo

create_beerstyle () {
  local name="$1"
  local min="$2"
  local max="$3"

  echo "➡️  Creating: $name ($min°C to $max°C)"

  curl -s -o /dev/null -w "%{http_code}\n" \
    -X POST "$BASE_URL/beerstyles" \
    -H "Content-Type: application/json" \
    -d "{
      \"name\": \"$name\",
      \"minTemp\": $min,
      \"maxTemp\": $max
    }"
}

create_beerstyle "Weissbier" -1 3
create_beerstyle "Pilsens" -2 4
create_beerstyle "Weizenbier" -4 6
create_beerstyle "Red ale" -5 5
create_beerstyle "India pale ale" -6 7
create_beerstyle "IPA" -7 10
create_beerstyle "Dunkel" -8 2
create_beerstyle "Imperial Stouts" -10 13
create_beerstyle "Brown ale" 0 14

echo
echo "Beer styles populated successfully"
