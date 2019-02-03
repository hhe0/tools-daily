<?php

class A
{
    public static function who()
    {
        echo __CLASS__;
    }

    public function test()
    {
        // A
        self::who();
        // B
        static::who();
        // A
        self::what();
        // B
        static::what();
    }

    public function what()
    {
        echo __CLASS__;
    }
}

class B extends A
{
    public static function who()
    {
        echo __CLASS__;
    }

    public function what()
    {
        echo __CLASS__;
    }
}

$B = new B();

$B->test();