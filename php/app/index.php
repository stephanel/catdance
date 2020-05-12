<?php

include('stores.php');

$i = rand(0, sizeof($images)-1);
$url = $images[$i];

$html = file_get_contents('index.view.php');

$html = str_replace("{url}", $url, $html);

echo $html;

?>