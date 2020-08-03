


function httpGet(theUrl)
{

    const xhr = new XMLHttpRequest();
    xhr.open('GET', '/jugadores');
    xhr.send();
    xhr.onload = function(e) {
    if (this.status == 200) {
        var newjson = JSON.parse(xhr.responseText)
        let counter  = 0
        newjson.forEach(element => {
            document.getElementById("place").innerHTML += `
            Datos de la api:
            ----------------- <br>
            <br>
            Numero de dato: ${counter} ${element.trofeos.personales}<br>`
            counter++
        });
    }
    };

}

console.log(httpGet("/jugadores"))