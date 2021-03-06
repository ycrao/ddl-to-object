<?php

namespace App\Models;

/**
 * AdminRole admin_roles
 *
 * generated by ddl-to-object <https://github.com/ycrao/ddl-to-object>
 *  
 * @property integer $id id 
 * @property string $name name 
 * @property string $slug slug 
 * @property mixed|string $created_at created_at 
 * @property mixed|string $updated_at updated_at  
 * @author unknown
 */
class AdminRole {

 
    /**
     * @var integer - id
     */
    private $id = null;

    /**
     * @var string - name
     */
    private $name = null;

    /**
     * @var string - slug
     */
    private $slug = null;

    /**
     * @var mixed|string - created_at
     */
    private $created_at = null;

    /**
     * @var mixed|string - updated_at
     */
    private $updated_at = null;


    /**
     * construct from object or array
     *
     * @param object|array $object
     * @return void
     */
    public function __construct($object = null) {
        if (is_array($object) || is_object($object)) {
            foreach ($object as $key => $value) {
                $this->$key = $value;
            }
        }
    }


    /**
     * magic get
     *
     * @param property name $key
     * @return mixed
     */
    public function __get($key) {
        return $this->$key;
    }

    /**
     * magic set
     *
     * @param property name $key
     * @param property value $value
     * @return boolean
     */
    public function __set($key, $value) {
        if (property_exists($this, $key)) {
            $this->$key = $value;
            return true;
        }
        return false;
    }
}