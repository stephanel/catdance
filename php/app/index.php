<?php

include('stores.php');
include('view.class.php');

$i = rand(0, sizeof($images)-1);
$url = $images[$i];

$view = new View('index.template');
$view->apply('url', $url);
$view->echo();

?>