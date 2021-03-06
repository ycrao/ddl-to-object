<?php

namespace App\Models;

/**
 * QrtzSimpropTrigger QRTZ_SIMPROP_TRIGGERS
 *
 * generated by ddl-to-object <https://github.com/ycrao/ddl-to-object>
 *  
 * @property string $SCHED_NAME SCHED_NAME 
 * @property string $TRIGGER_NAME TRIGGER_NAME 
 * @property string $TRIGGER_GROUP TRIGGER_GROUP 
 * @property string $STR_PROP_1 STR_PROP_1 
 * @property string $STR_PROP_2 STR_PROP_2 
 * @property string $STR_PROP_3 STR_PROP_3 
 * @property integer $INT_PROP_1 INT_PROP_1 
 * @property integer $INT_PROP_2 INT_PROP_2 
 * @property integer $LONG_PROP_1 LONG_PROP_1 
 * @property integer $LONG_PROP_2 LONG_PROP_2 
 * @property numeric $DEC_PROP_1 DEC_PROP_1 
 * @property numeric $DEC_PROP_2 DEC_PROP_2 
 * @property string $BOOL_PROP_1 BOOL_PROP_1 
 * @property string $BOOL_PROP_2 BOOL_PROP_2  
 * @author unknown
 */
class QrtzSimpropTrigger {

 
    /**
     * @var string - SCHED_NAME
     */
    private $SCHED_NAME = null;

    /**
     * @var string - TRIGGER_NAME
     */
    private $TRIGGER_NAME = null;

    /**
     * @var string - TRIGGER_GROUP
     */
    private $TRIGGER_GROUP = null;

    /**
     * @var string - STR_PROP_1
     */
    private $STR_PROP_1 = null;

    /**
     * @var string - STR_PROP_2
     */
    private $STR_PROP_2 = null;

    /**
     * @var string - STR_PROP_3
     */
    private $STR_PROP_3 = null;

    /**
     * @var integer - INT_PROP_1
     */
    private $INT_PROP_1 = null;

    /**
     * @var integer - INT_PROP_2
     */
    private $INT_PROP_2 = null;

    /**
     * @var integer - LONG_PROP_1
     */
    private $LONG_PROP_1 = null;

    /**
     * @var integer - LONG_PROP_2
     */
    private $LONG_PROP_2 = null;

    /**
     * @var numeric - DEC_PROP_1
     */
    private $DEC_PROP_1 = null;

    /**
     * @var numeric - DEC_PROP_2
     */
    private $DEC_PROP_2 = null;

    /**
     * @var string - BOOL_PROP_1
     */
    private $BOOL_PROP_1 = null;

    /**
     * @var string - BOOL_PROP_2
     */
    private $BOOL_PROP_2 = null;


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