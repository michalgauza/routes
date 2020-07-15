package main

func getOSMRUrl(source string, destination string) string {
	url := "http://router.project-osrm.org/route/v1/driving/"
	url += source
	url += ";"
	url += destination
	url += "?overview=false"
	return url
}
