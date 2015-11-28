var $ = jQuery;

var hourPrice = 200;


$(document).ready(function () {
  bindSlider();
  formChanged();
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
