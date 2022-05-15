<?php

namespace App\Models;

/**
 * Article
 *
 * @property integer $id id
 * @property integer $user_id 用户id
 * @property string $content 正文
 * @property string $create_time 创建时间
 * @property string $update_time 更新时间
 */
class Article {

    private $id = null;

    private $user_id = null;

    private $content = null;

    private $create_time = null;

    private $update_time = null;

    public function __construct($object = null) {
        $this->cast($object);
    }

    public function cast($object) {
        if (is_array($object) || is_object($object)) {
            foreach ($object as $key => $value) {
                $this->$key = $value;
            }
        }
    }

    public function __get($key) {
        return $this->$key;
    }

    public function __set($key, $value) {
        if (property_exists($this, $key)) {
            $this->$key = $value;
            return true;
        }
        return false;
    }
}

// TEST
$article = new Article([
    'id' => 1,
    'user_id' => 2,
    'content' => 'hello world',
    'create_time' => '2022-05-01 12:00:00',
    'update_time' => '2022-05-01 12:00:00',
    'faked' => true
]);
var_dump($article);
$article->user_id = 5;
$article->category_id = 6;
var_dump($article);
var_dump($article->user_id);
// var_dump($article->faked);
var_dump($article->content);