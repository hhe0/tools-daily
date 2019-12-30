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
    private $player;

    public function __construct(Player $player)
    {
        $this->player = $player;
    }

    public function watchMovie()
    {
        $this->player->playMovie();
    }

    public function listenMusic()
    {
        $this->player->playMusic();
    }
}

$people = new People(new Computer());
$people->watchMovie();
$people->listenMusic();

$people2 = new People(new Phone());
$people2->watchMovie();
$people2->listenMusic();
