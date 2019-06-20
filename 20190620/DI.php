<?php
interface player {
    public function playMovie();
    public function playMusic();
}

class Computer implements player {
    public function playMovie()
    {
        echo "Welcome to the world of movies by computer." . PHP_EOL;
    }

    public function playMusic()
    {
        echo "Welcome to the world of music by computer." . PHP_EOL;
    }
}

class Phone implements player {
    public function playMovie()
    {
        echo "Welcome to the world of movies by phone." . PHP_EOL;
    }

    public function playMusic()
    {
        echo "Welcome to the world of music by phone." . PHP_EOL;
    }
}

class People {
    private $computer;
    private $phone;

    public function __construct()
    {
        $this->computer = new Computer();
        $this->phone = new Phone();
    }

    public function watchMovieByComputer()
    {
        $this->computer->playMovie();
    }

    public function listenMusicByComputer()
    {
        $this->computer->playMusic();
    }

    public function watchMovieByPhone()
    {
        $this->phone->playMovie();
    }

    public function listenMusicByPhone()
    {
        $this->phone->playMusic();
    }
}

$people = new People();
$people->watchMovieByComputer();
$people->listenMusicByComputer();

$people->watchMovieByPhone();
$people->watchMovieByComputer();