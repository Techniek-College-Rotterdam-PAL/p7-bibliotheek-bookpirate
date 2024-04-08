var modal = $('#myModal');
var btn = $('#btn');

$(document).ready(function() {
    btn.on('click',function() {
        modal.show();
    });
});

$('body').bind('click', function(e) {
    if ($(e.target).hasClass("modal")) {
        modal.fadeOut();
    }
});

var number = 0;

function item(id) {
    $.ajax({
        type:       "POST",
        url:        "../src/addInvoiceItem.php",
        data:       {id: id},
        success:    function(data) {
            console.log(data);
            onSuccess(data);
        },
        error:      function(xhr, status, error) {
            console.error('error:', error);
        },
    });
    function onSuccess(data) {
        var results = $.parseJSON(data);

        $('#item_results').prepend(
            '<tr id="item-'+number+'">'
            + '<td><a href="#" onclick="removeItem('+number+')">Remove</a></td>'
            + '<td><input type="hidden" name="product[]" value="' + results['id'] + '"'
            + '<p>' + results['product_name'] + '</p></td>'
            + '<td><p>' + results['product_description'] + '</p></td>'
            + '<td><input type="number" name="qty[]" id="input-qty" value="1"></td>'
            + '<td><p>' + results['product_price'] + '</p></td>'
            + '</tr>'
        );
        modal.fadeOut();
        number++;
    }

}

function removeItem(number) {
    const item =  document.getElementById("item-" + number);
    item.remove();
}