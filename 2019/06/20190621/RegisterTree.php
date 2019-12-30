<?php

class RegisterTree {
    private $container = [];

    public function set ($key, $value)
    {
        if (isset($this->container[$key])) {
            die("当前实例已存在");
        } else {
            $this->container[$key] = $value;
        }
    }

    public function get($key)
    {
        if(isset($this->container[$key])) {
            return $this->container[$key];
        } else {
            die("当前实例不存在");
        }
    }

    public function unset($key)
    {
        if (isset($this->container[$key])) {
            unset($this->container[$key]);
        } else {
            die("当前实例不存在");
        }
    }
}

$registerTree = new RegisterTree();
$registerTree->set('k1', "v1");
echo $registerTree->get('k1') . PHP_EOL;

$registerTree->unset('k1');
echo $registerTree->get('k1') . PHP_EOL;

