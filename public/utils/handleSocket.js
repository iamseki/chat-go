var socket = null;

$('#btnMsn').click(function () {

  socket.send(JSON.stringify({ "Message": "P@y Attention!" }))
});

setTimeout(() => $('#usermsg').focus(), 800)

$(function () {
  var msgBox = $("#chatbox textarea")
  var messages = $("#messages")


  $(function () {
    $("#usermsg").keypress(function (e) {
      if (e.which == 13) {
        //submit form via ajax, this is not JS but server side scripting so not showing here
        socket.send(JSON.stringify({ "Message": $(this).val() }))

        $(this).val("");

        e.preventDefault();
        $("#chatbox").submit()
      }
    });
  });

  $("#chatbox").submit(function () {
    if (!msgBox.val()) return false;
    if (!socket) {
      alert("Error: There is no socket connection.")
      return false
    }
    socket.send(JSON.stringify({ "Message": msgBox.val() }))
    msgBox.val("")
    return false
  })
  if (!window["WebSocket"]) {
    alert("Error: your browser does not support web sockets.")
  }
  else {
    const host = location.host
    socket = new WebSocket(`ws://${host}/room`);
    socket.onclose = function () {
      alert("Conection has been closed.")
    }
    socket.onmessage = function (e) {
      var msg = JSON.parse(e.data);

      if (msg.Message == "P@y Attention!") {
        $('#chatbox').css({
          'animation-name': 'msn',
          'animation-duration': '0.1s',
          'animation-iteration-count': '5'
        })
        setTimeout(() => $('#chatbox').css({
          'animation-name': ''
        }), 800)
      }
      else {
        messages.append(
          $("<li>").append(
            $("<div>").append(
              $("<img>").css({
                height: 52,
                width: 52,
                verticalAlign: "middle"
              }).attr("src", msg.AvatarURL),
              $("<strong>").text(msg.Name)
            ),
            $("<span>").text(msg.Message)
          )
        );

        $("#chatbox").animate({ scrollTop: $('ul#messages').height() }, 'slow');
      }
    }
  }
})

var getUrlParameter = function getUrlParameter(sParam) {
  var sPageURL = window.location.search.substring(1),
    sURLVariables = sPageURL.split('&'),
    sParameterName,
    i;

  for (i = 0; i < sURLVariables.length; i++) {
    sParameterName = sURLVariables[i].split('=');

    if (sParameterName[0] === sParam) {
      return sParameterName[1] === undefined ? true : decodeURIComponent(sParameterName[1]);
    }
  }
};