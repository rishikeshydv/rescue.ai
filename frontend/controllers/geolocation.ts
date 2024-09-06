export default function getLngLat(){
    if (navigator && navigator.geolocation){
        navigator.geolocation.getCurrentPosition(
            function(position){
                const lng = position.coords.longitude;
                const lat = position.coords.latitude;
                return {lng:lng, lat:lat};
            }
        )
    }
}