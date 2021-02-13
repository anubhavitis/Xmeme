import { loadMeme, addMeme } from "./util.js";

window.onload = function () {
  loadMeme();
};

document.getElementById("form").onsubmit = function (e) {
  addMeme(e);
};
