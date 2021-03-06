# go-api-here

Go: API that handles requests for isochrones from HERE Maps API.  The returned JSON is ready for use in LeafletJS.com

- This was written for using in a different project and the functionality is narrow in scope.
- The API returns JSON that make up the verticies of the isochrone (polygon).

__*Usage:*__ *http://myserver:8003/v1/here-isochrone/{lng}/{lat}/{time}/{appid}/{appcode}*

- __*lng*__ => longitude (decimal degrees)
- __*lat*__ => latitude (decimal degrees)
- __*time*__ => drive time polygon in seconds
- __*appid*__ => HERE app_id
- __*appcode*__ => HERE app_code


__*Example API Call & Return Value:*__

-   http://myserver:8003/v1/here-isochrone/-95.9671744/36.1332013/60/my_app_id/my_app_code
-   {"here":"[36.1340332,-95.9716702],[36.1340332,-95.9683228],[36.1342049,-95.9678078],[36.1347198,-95.9676361],[36.1360931,-95.9676361],[36.1366081,-95.9674644],[36.1366081,-95.9671211],[36.1355782,-95.9667778],[36.1352348,-95.9657478],[36.1348915,-95.9657478],[36.1345482,-95.9667778],[36.1342049,-95.9667778],[36.1340332,-95.9662628],[36.1340332,-95.9635162],[36.1338615,-95.9630013],[36.1333466,-95.9628296],[36.1328316,-95.9630013],[36.1326599,-95.9635162],[36.1326599,-95.9642029],[36.1324883,-95.9647179],[36.1314583,-95.9650612],[36.1312866,-95.9655762],[36.1312866,-95.9662628],[36.131115,-95.9667778],[36.1306,-95.9669495],[36.1299133,-95.9669495],[36.1293983,-95.9671211],[36.1292267,-95.9676361],[36.1292267,-95.9683228],[36.1293983,-95.9688377],[36.1299133,-95.9690094],[36.1306,-95.9690094],[36.131115,-95.9691811],[36.1318016,-95.9698677],[36.1321449,-95.9708977],[36.1324883,-95.971241],[36.1328316,-95.972271],[36.1333466,-95.9724426],[36.1338615,-95.9723568],[36.1340332,-95.9720993],[36.1340332,-95.9716702]"}

-   On error returns: {"here":"msg"}
-   msg = Here API call error message
