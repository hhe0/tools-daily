<?php

function createRange($num)
{
    $data = [];

    for ($i = 0; $i < $num; $i++)
    {
        $data[] = time();
    }

    return $data;
}

function createRangeUsingYield($num)
{
    for ($i = 0; $i < $num; $i++)
    {
        // 使用yield优化性能
        yield time();
    }
}

$res = createRange(10);
foreach ($res as $val)
{
    sleep(1);
    echo $val . PHP_EOL;
}

$r = createRangeUsingYield(10);
foreach ($r as $val)
{
    sleep(1);
    echo $val . PHP_EOL;
}