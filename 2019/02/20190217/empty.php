<?php
var_dump(isset($a));

$a = '';
var_dump(empty($a));
var_dump(isset($a));

$b = 0;
var_dump(empty($b));
var_dump(isset($b));
unset($b);
var_dump(isset($b));

// 必须全部存在才返回True
var_dump(isset($a, $b));

class TestForEmpty
{
//    private $data = '111';

    public function test()
    {
        echo 'hello' . PHP_EOL;
    }
}
$obj = new TestForEmpty();
// 空的属性也会返回False
var_dump(empty($obj));

