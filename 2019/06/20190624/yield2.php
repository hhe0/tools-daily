<?php

class Mysql {
    private static $conn = NULL;

    private function __construct()
    {
    }

    public static function getInstance()
    {
        if (self::$conn) {
            return self::$conn;
        } else {
            try {
                $conn = new PDO("mysql:host=127.0.0.1;", 'root', 'root12345');
                echo "连接成功" . PHP_EOL;
            } catch(PDOException $e) {
                echo $e->getMessage();
            }

            return $conn;
        }
    }
}

$t1 = microtime(true);

$conn = Mysql::getInstance();

foreach (get($conn) as $id => $row) {
    $row_id = $row['id'];
    $sql = <<<EOF
update `test`.`yield_test` set count = 5 where id = $row_id;
EOF;
    $conn->exec($sql);
}

$t2 = microtime(true);
echo '耗时'.round($t2-$t1,3).'秒' . PHP_EOL;
echo 'Now memory_get_usage: ' . memory_get_usage() . 'b' . PHP_EOL;


function get($conn) {
    $sql_select = <<<EOF
select id,count from `test`.`yield_test`;
EOF;
    $res = $conn->query($sql_select);
    while($row = $res->fetch()) {
        yield $row;
    }
}






