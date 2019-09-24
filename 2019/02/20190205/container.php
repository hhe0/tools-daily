<?php

class Container
{
    private $realtion = [];

    public function __set($k, $v)
    {
        $this->realtion[$k] = $v;
    }

    public function __get($k)
    {
        return $this->realtion[$k];
    }
}

class A
{
    private $container;

    public function __construct(Container $container)
    {
        $this->container = $container;
    }

    public function doSomething()
    {
        $b = $this->container->b;
        echo $b . PHP_EOL;
    }
}

class B
{
    private $name = 'hhe';

    public function __toString()
    {
        return $this->name;
    }
}

$container = new Container();
$b = new B();
$container->b = $b;

$a = new A($container);
$a->doSomething();


