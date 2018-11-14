# go-api-here

Go: API that handles requests for isochrones from HERE Maps API.  The returned JSON is ready for use in LeafletJS.com

- This was written for using in a different project and the functionality is narrow in scope.
- The API returns JSON in text format of only the GeoJSON portion of what HERE returns.

__*Deployment:*__ *http://zotact1.ddns.net:8001/v1/here-isochrone/{lng}/{lat}/{time}/{appid}/{appcode}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in seconds
- __*appid*__ => HERE app_id
- __*appcode*__ => HERE app_code