<?php
$arr = [5, 2, 3, 1, 4];

function quick_sort($arr)
{
    $len =  count($arr);
    if ($len <= 1){
        return $arr;
    }

    $mid = $left = $right = [];
    $base = $arr[0];

    foreach ($arr as $val)
    {
        if ($val > $base) {
            $right[] = $val;
        } elseif ($val < $base) {
            $left[] = $val;
        } else {
            $mid[] = $val;
        }
    }

    $left = quick_sort($left);
    $right = quick_sort($right);

    return array_merge($left, $mid, $right);
}

$res = quick_sort($arr);
print_r($res);