$(document).ready(function(){
    $('#data').on('submit', function(e){
       e.preventDefault()
       $.ajax({
        url: '/',
        type: 'post',
        data: $('#data').serialize(),
        success: function(data) {
                   alert(data)
                }
        });
    });
})


