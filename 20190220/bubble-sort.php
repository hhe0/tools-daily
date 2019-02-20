<?php

$arr = [5, 2, 3, 1];
$len = count($arr);

for ($i = 0; $i < $len; $i++)
{
    for ($j = 0; $j < $len - $i - 1; $j++)
    {
        if ($arr[$j] < $arr[$j + 1])
        {
            list($arr[$j + 1], $arr[$j]) = [$arr[$j], $arr[$j + 1]];
        }
    }
}

print_r($arr);
