<?php

class TryTest
{
    public static function test()
    {
        try {
            $i = 0;
//            throw new Exception('exception happens');
            return 1;
        } catch (Exception $e) {
            echo $e->getMessage() . PHP_EOL;
            return 2;
        } finally {
            // 如果finally中存在return，则返回该值；否则返回对应的try/catch块中的返回值
            return 3;
        }
    }
}

echo TryTest::test();
