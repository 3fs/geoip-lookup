function init_map() {
    gps = document.getElementById('lats').innerHTML.split('|');
    var myOptions = {
        zoom: 14,
        center: new google.maps.LatLng(gps[0], gps[1]),
        mapTypeId: google.maps.MapTypeId.ROADMAP
    };
    map = new google.maps.Map(document.getElementById("gmap_canvas"), myOptions);
    marker = new google.maps.Marker({
        map: map,
        position: new google.maps.LatLng(gps[0], gps[1])
    });
    infowindow = new google.maps.InfoWindow({
        content: "<span style='height:auto !important; display:block; white-space:nowrap; overflow:hidden !important;'><strong style='font-weight:400;'>IP location</strong></span>"
    });
    google.maps.event.addListener(marker, "click", function() {
        infowindow.open(map, marker);
    });
    infowindow.open(map, marker);
}

google.maps.event.addDomListener(window, "load", init_map);
