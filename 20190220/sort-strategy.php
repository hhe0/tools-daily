<?php
interface SortStrategy
{
    public function sort($arr = []);
}

class BubbleSort implements SortStrategy
{
    public function sort($arr = [])
    {
        $len = count($arr);

        for ($i = 0; $i < $len; $i++)
        {
            for ($j = 0; $j < $len - $i - 1; $j++)
            {
                if ($arr[$j] > $arr[$j + 1])
                {
                    list($arr[$j + 1], $arr[$j]) = array($arr[$j], $arr[$j + 1]);
                }
            }
        }

        return $arr;
    }
}

class QuickSort implements SortStrategy
{
    public function sort($arr = [])
    {
        $len = count($arr);
        if ($len <= 1) {
            return $arr;
        }

        $left = $mid = $right = [];
        $base = $arr[0];

        foreach ($arr as $val)
        {
            if ($val > $base) {
                $right[] = $val;
            } elseif ($val < $base) {
                $left[] = $val;
            } else {
                $mid[] = $val;
            }
        }

        $left = self::sort($left);
        $right = self::sort($right);

        return array_merge($left, $mid, $right);
    }
}

class Sort
{
    public $sortStrategy = [];

    public function __construct(SortStrategy $sortStrategy)
    {
        $this->sortStrategy = $sortStrategy;
    }

    public function sort($arr = [])
    {
        return $this->sortStrategy->sort($arr);
    }
}


$arr = [2, 5, 1, 3, 4];
$sort = new Sort(new BubbleSort());
$res = $sort->sort($arr);
print_r($res);

print_r($arr);

$sort_2 = new Sort(new QuickSort());
$res_2 = $sort_2->sort($arr);
print_r($res_2);