package entity
// generated by ddl-to-object <https://github.com/ycrao/ddl-to-object>

import (
	"database/sql"
)

// QrtzSimpropTrigger QRTZ_SIMPROP_TRIGGERS
type QrtzSimpropTrigger struct {  
	SchedName string `json:"sched_name" db:"SCHED_NAME"`  // SCHED_NAME 
	TriggerName string `json:"trigger_name" db:"TRIGGER_NAME"`  // TRIGGER_NAME 
	TriggerGroup string `json:"trigger_group" db:"TRIGGER_GROUP"`  // TRIGGER_GROUP 
	StrProp1 sql.NullString `json:"str_prop_1" db:"STR_PROP_1"`  // STR_PROP_1 
	StrProp2 sql.NullString `json:"str_prop_2" db:"STR_PROP_2"`  // STR_PROP_2 
	StrProp3 sql.NullString `json:"str_prop_3" db:"STR_PROP_3"`  // STR_PROP_3 
	IntProp1 int32 `json:"int_prop_1" db:"INT_PROP_1"`  // INT_PROP_1 
	IntProp2 int32 `json:"int_prop_2" db:"INT_PROP_2"`  // INT_PROP_2 
	LongProp1 int64 `json:"long_prop_1" db:"LONG_PROP_1"`  // LONG_PROP_1 
	LongProp2 int64 `json:"long_prop_2" db:"LONG_PROP_2"`  // LONG_PROP_2 
	DecProp1 float64 `json:"dec_prop_1" db:"DEC_PROP_1"`  // DEC_PROP_1 
	DecProp2 float64 `json:"dec_prop_2" db:"DEC_PROP_2"`  // DEC_PROP_2 
	BoolProp1 sql.NullString `json:"bool_prop_1" db:"BOOL_PROP_1"`  // BOOL_PROP_1 
	BoolProp2 sql.NullString `json:"bool_prop_2" db:"BOOL_PROP_2"`  // BOOL_PROP_2  
}