var url= "https://xmeme07.herokuapp.com/"
// var url= "http://localhost:8080/"

export function addMeme(e) {
  e.preventDefault();

  var name = document.getElementById("name").value;
  var caption = document.getElementById("caption").value;
  var url = document.getElementById("url").value;

  var myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    name: name,
    url: url,
    caption: caption,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch(url+"memes", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      swal.fire("Hurray!", "id=" + result.id.toString(), "success");
      loadMeme();
    })
    .catch((error) => console.log("error", error));
}

export function loadMeme() {
  var requestOptions = {
    method: "GET",
    redirect: "follow",
  };

  fetch(url+"memes", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      var cards = "";
      for (var i = 0; i < result.result.length; i += 1) {
        var ele = result.result[i];
        var temp = `<div class="m-4 shadow-lg max-w-lg mx-auto">
                        <div class="p-2 mt-0.5 bg-gray-50 flex flex-col justify-around rounded-t-lg">
                            <div class="mx-2 flex justify-between text-gray-500 text-md uppercase font-semibold
                                text-md uppercase font-semibold">
                                <h4 class="text-gray-500"> ${ele.name}</h4>
                                <h4 class="text-gray-500"> ${ele.time} </h4>
                            </div>
                            <p class="mx-2 text-gray-700 font-bold">${ele.caption}</p>
                        </div>
                        <img class="max-h-96 w-full rounded-b-lg mx-auto" src="${ele.url}" \>
                    </div>`;
        cards += temp;
      }

      document.getElementById("cards").innerHTML = cards;
    })
    .catch((error) => console.log("error", error));
}
