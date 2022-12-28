package gdextension

import (
	"reflect"

	. "github.com/godot-go/godot-go/pkg/gdextensionffi"
)

// // TypeName represents a the name of a type in Godot
// type TypeName string

// func (n TypeName) compare(other TypeName) int {
// 	if n < other {
// 		return -1
// 	} else if n == other {
// 		return 0
// 	}
// 	return 1
// }

// // MethodName represents a method name in Godot.
// type MethodName string

// // SignalName represents a signal name in Godot.
// type SignalName string

// // PropertyName represents a property name in Godot.
// type PropertyName string

// // ConstantName represents a constant name in Godot.
// type ConstantName string

// PropertySetGet holds metadata of the getter and setting functions of a Godot property.
type PropertySetGet struct {
	Index   int
	Setter  string
	Getter  string
	_setptr *MethodBind
	_getptr *MethodBind
	Type    GDExtensionVariantType
}

type ClassInfo struct {
	Name                      string
	NameAsStringNamePtr       GDExtensionConstStringNamePtr
	ParentName                string
	ParentNameAsStringNamePtr GDExtensionConstStringNamePtr
	Level                     GDExtensionInitializationLevel
	MethodMap                 map[string]*MethodBind
	SignalNameMap             map[string]struct{}
	VirtualMethodMap          map[string]GDExtensionClassCallVirtual
	PropertyNameMap           map[string]struct{}
	ConstantNameMap           map[string]struct{}
	ParentPtr                 *ClassInfo
	ClassType                 reflect.Type
	InheritType               reflect.Type
}

func NewClassInfo(name, parentName string, level GDExtensionInitializationLevel, classType, inheritType reflect.Type, parentPtr *ClassInfo) *ClassInfo {
	return &ClassInfo{
		Name:                name,
		NameAsStringNamePtr: NewStringNameWithLatin1Chars(name).AsGDExtensionStringNamePtr(),
		ParentName:          parentName,
		Level:               level,
		MethodMap:           map[string]*MethodBind{},
		SignalNameMap:       map[string]struct{}{},
		VirtualMethodMap:    map[string]GDExtensionClassCallVirtual{},
		PropertyNameMap:     map[string]struct{}{},
		ConstantNameMap:     map[string]struct{}{},
		ParentPtr:           parentPtr,
		ClassType:           classType,
		InheritType:         inheritType,
	}
}

var (
	classdbCurrentLevel GDExtensionInitializationLevel = GDEXTENSION_INITIALIZATION_CORE
)
