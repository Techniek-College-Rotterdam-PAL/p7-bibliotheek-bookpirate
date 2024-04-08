$(document).ready(function(){
    function load_data(query) {
        $.ajax({
            url:            "http://127.0.0.1:8080/api/v1/search-books",
            type:           "POST",
            data:           {query:query},
            success:        function(data)
            {
                $('#liveSearch-result').html('');
                $('#liveSearch-result').html(data);
            }
        });
    }
    $('#liveSearch').keyup(function(){
        var search = $(this).val();
        if (search != '') {
            load_data(search);
        } else {
            load_data();
        }
    });
});