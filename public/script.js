var $ = jQuery;

var hourPrice = 200;


$(document).ready(function () {
  bindSlider();
  formChanged();
  show();
});

function bindSlider() {
  $("#range").on("input",formChanged);
  $(".checkbox-option").change(formChanged);
}

function formChanged() {
  var val = $("#range").val();

  var price =  190 + (val * hourPrice);
  $("#value").text("Přibližná cena: " + price + " Kč");
}


function show() {
  if (window.location.href.indexOf("submitted") > 0) {
    $(".submited").show();
  }
}
