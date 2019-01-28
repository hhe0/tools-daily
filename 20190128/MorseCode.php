<?php

namespace chap1;

/**
 * 摩斯电码转换机
 */
class MorseCodeChanger
{
    private static $morseCodeMap = [
        'A' => '.-',
        'B' => '-...',
        'C' => '-.-.',
        'D' => '-..',
        'E' => '.',
        'F' => '..-.',
        'G' => '--.',
        'H' => '....',
        'I' => '..',
        'J' => '.---',
        'K' => '-.-',
        'L' => '.-..',
        'M' => '--',
        'N' => '-.',
        'O' => '---',
        'P' => '.--.',
        'Q' => '--.-',
        'R' => '.-.',
        'S' => '...',
        'T' => '-',
        'U' => '..-',
        'V' => '...-',
        'W' => '.--',
        'X' => '-..-',
        'Y' => '-.--',
        'Z' => '--..',
        // 空格符用来分隔单词
        ' ' => '-....-'
    ];

    public static function decode($morseStr)
    {
        $wordStr  = '';
        $morseArr = explode(' ', $morseStr);

        foreach ($morseArr as $eachMorse) {
            if ($char = self::findChar($eachMorse)) {
                $wordStr .= $char;
            }
        }

        return $wordStr;
    }

    public static function encode($str)
    {
        $wordMorseStr = '';
        $charArr      = str_split($str);

        $first = true;
        foreach ($charArr as $char) {
            if ($morseStr = self::findCode($char)) {
                if ($first) {
                    $wordMorseStr .= $morseStr;
                    $first = false;
                } else {
                    $wordMorseStr .= ' ' . $morseStr;
                }
            }
        }

        return $wordMorseStr;
    }

    private static function findCode($str)
    {
        $char = '';

        foreach (self::$morseCodeMap as $morseChar => $morseCode) {
            if (strtoupper($str) == $morseChar) {
                $char .= $morseCode;
            }
        }

        return $char;
    }

    private static function findChar($str)
    {
        $char = '';

        foreach (self::$morseCodeMap as $morseChar => $morseCode) {
            if ($str == $morseCode) {
                $char .= $morseChar;
            }
        }

        return $char;
    }
}

/**
 * 小孩之间使用摩斯电码进行交流
 */
class Kid
{
    private $receiver;
    private $receiveMsg = null;

    public function connect(Kid $kid)
    {
        $this->receiver = $kid;
    }

    public function send($msg)
    {
        $this->receiver->receiveMsg = $this->encode($msg);
    }

    public function receive()
    {
        echo $this->decode($this->receiveMsg);
    }

    // 发送消息之前必须先对信息进行编码
    private function encode($msg)
    {
        return MorseCodeChanger::encode($msg);
    }

    // 接收消息后必须对信息进行解码
    private function decode($msg)
    {
        return MorseCodeChanger::decode($msg);
    }
}

$xiaoMing = new Kid();
$xiaoHong = new Kid();

$xiaoMing->connect($xiaoHong);

// 发送的消息会被转换为摩斯电码，以摩斯电码的形式发送出去
$xiaoMing->send('HELLO WoRLD');
// 接收到的信息会都转换为大写
$xiaoHong->receive();