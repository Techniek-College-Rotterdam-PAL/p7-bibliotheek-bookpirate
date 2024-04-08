<?php
require_once '../config/connect.php';

$return = '';
if (isset($_POST["query"])) {
    $search = mysqli_real_escape_string($con, $_POST["query"]);
    $query = "SELECT * FROM products
    WHERE id LIKE  '%" . $search . "%'
    OR product_name LIKE '%" . $search . "%'
    ";
} else {
    $query = "SELECT * FROM products";
}

$result = mysqli_query($con, $query);
if (mysqli_num_rows($result) > 0) {
    while ($row = mysqli_fetch_array($result)) {
        $return .= '
        <tr onclick="item(' . $row["id"] . ')">
        <td>' . $row["id"] . '</td>
        <td>' . $row["product_name"] . '</td>
        <td>' . $row["product_description"] . '</td>
        <td>' . $row["product_quantity"] . '</td>
        <td>' . $row["product_price"] . '</td>
        </tr>';
    }
    echo $return;
} else {
    echo '<tr><td colspan="5" style="text-align: center;">No results Found.</td></tr>';
}
