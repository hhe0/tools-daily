<?php

class Test
{
    public static $static_var = '11';
    public $var = '22';

    public function print_static_var()
    {
        echo self::$static_var . PHP_EOL;
        // 不能使用self输出非静态变量
//        echo self::$var;
        // 不能使用this输出静态变量
//        echo $this->static_var;
        echo $this->var . PHP_EOL;

        // self可以调用非静态方法
        self::hello();
        // this可以调用非静态方法
        $this->hello();
        // self可以调用静态方法
        self::hi();
        // this可以调用静态方法
        $this->hi();
    }

    public function hello()
    {
        echo 'hello' . PHP_EOL;
    }

    public static function hi()
    {
        echo 'hi' . PHP_EOL;
    }
}

$test = new Test();
$test->print_static_var();