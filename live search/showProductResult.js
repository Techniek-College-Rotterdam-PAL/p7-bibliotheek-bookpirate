$(document).ready(function(){
    function load_data(query) {
        $.ajax({
            url:            "../src/liveSearch.php",
            type:           "POST",
            data:           {query:query},
            success:        function(data)
            {
                $('#invoice-result').html('');
                $('#invoice-result').html(data);
            }
        });
    }
    $('#search').keyup(function(){
        var search = $(this).val();
        if (search != '') {
            load_data(search);
        } else {
            load_data();
        }
    });
});