src - source location as Latitude and Longitude (13.388860,52.517037)
dst - destination as Latitude and Longitude (13.397634,52.529407), can be multiple

endpoints
  - local 
	http://localhost:8080/routes?src=xx.xxxx,yy.yyyy&dst=hh.hhhh,ll.llll
	example: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
  - heroku
	http://localhost:8080/routes?src=xx.xxxx,yy.yyyy&dst=hh.hhhh,ll.llll
	example: https://floating-dusk-65973.herokuapp.com/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219

response
{
  "source": "13.388860,52.517037",
  "routes": [
    {
      "destination": "13.397634,52.529407",
      "duration": 251.6,
      "distance": 1884.9
    },
    {
      "destination": "13.428555,52.523219",
      "duration": 394.1,
      "distance": 3841.7
    }
  ],
  "err": ""
}

