var URL = "https://xmeme07.herokuapp.com/";
// var URL= "http://localhost:8080/"

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

  fetch(URL + "memes", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      swal.fire("Congrats," + name, "You meme is added to Xmeme", "success");
      loadMeme();
    });
}

export function loadMeme() {
  var requestOptions = {
    method: "GET",
    redirect: "follow",
  };

  fetch(URL + "memes", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      var cards = "";

      for (
        var i = 0;
        result.result != null && i < result.result.length;
        i += 1
      ) {
        var ele = result.result[i];
        var temp = `<div class="m-4 shadow-lg max-w-lg mx-auto">
          <div class="p-2 mt-0.5 bg-gray-50 flex flex-col justify-around rounded-t-lg">
            <div class="flex justify-between">
              <div class="">
                <div class="mx-2 flex justify-between text-gray-500 text-md uppercase font-semibold
                    text-md uppercase font-semibold">
                  <h4 class="text-gray-500 overflow-x-hidden"> ${ele.time} -</h4>
                  <h4 class="ml-1 text-gray-500 overflow-x-hidden"> ${ele.name}</h4>
                </div>
                <div class="mx-2 flex justify-between text-gray-700">
                  <p class="text-lg font-bold overflow-x-hidden">${ele.caption}</p>
                </div>
              </div>
              <div class="flex my-auto ">
                  <svg onclick="editMeme(${ele.id})" class=" cursor-pointer h-8 w-8 border-2 text-white border-green-700 bg-green-700 rounded-full p-1 m-1 hover:bg-green-900"
                    xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                  </svg>
                  <svg onclick="delMeme(${ele.id})" class=" cursor-pointer h-8 w-8 border-2 text-white border-red-500 bg-red-500 rounded-full p-1 m-1 hover:bg-red-700"
                    xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                      d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
              </div>
            </div>
          </div>
          <img class="max-h-96 w-full rounded-b-lg mx-auto" src="${ele.url}" \>
        </div>`;
        cards +=temp;
      }

      document.getElementById("cards").innerHTML = cards;
    });
}
