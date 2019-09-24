<?php

class Singleton
{
    private static $instance;

    private function __construct()
    {
    }

    public static function getInstance()
    {
        if (is_null(self::$instance))
        {
            self::$instance = new self();
        }

        return self::$instance;
    }
}

var_dump(Singleton::getInstance());
var_dump(Singleton::getInstance());
var_dump(Singleton::getInstance());
