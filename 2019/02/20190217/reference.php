<?php

// 变量的引用
$a = 1;
$b = &$a;
echo $a . PHP_EOL;
echo $b . PHP_EOL;
$b = 2;
echo $a . PHP_EOL;
echo $b . PHP_EOL;


// 函数的引用传递
// 一般情况
function add($var)
{
    $var += 1;
    echo $var . PHP_EOL;
}

function add2(&$var)
{
    $var += 1;
    echo $var . PHP_EOL;
}

$a = 1;
add($a);
echo $a . PHP_EOL;

add2($a);
echo $a . PHP_EOL;

// 特殊情况 使用 call_user_func_array 来调用函数时
function add3(&$arr)
{
    $arr[] = 3;
    var_dump($arr);
}

$arr = [0, 1, 2];
call_user_func_array('add3', array(&$arr));
var_dump($arr);


// 函数的引用返回
function &test()
{
    static $b = 0;
    $b += 1;
    echo $b . PHP_EOL;
    return $b;
}

$a = test();
$a = 5;
$a = test();

// b 与 a 指向同一个位置
$a = &test();
$a = 5;
$a = test();

// 用于对象
class TestForObject
{
    private $data = 'test';

    public function & get()
    {
        return $this->data;
    }

    public function output()
    {
        echo $this->data . PHP_EOL;
    }
}

$test = new TestForObject();
$data = &$test->get();
$test->output();

$data = 'for';
$test->output();

$data = 'object';
$test->output();


// 对象的引用
class TestForNewObject
{
    public $data = '123';
}

$oa = new TestForNewObject();
echo $oa->data . PHP_EOL;
$ob = $oa;
echo $ob->data . PHP_EOL;

$oa->data = '234';
echo $oa->data;


$a = 1;
$b = &$a;
unset($a);
//unset($b);
echo $b . PHP_EOL;

