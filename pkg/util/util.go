package util

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"unsafe"
)

func GetLocalIp() (string, error) {
	err := errors.New("can not find the client ip address")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", err
}

func GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macAddrs
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func CallReflect(any any, name string, args ...any) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	if v := reflect.ValueOf(any).MethodByName(name); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}
}

func SetUnExportedStructField(ptr any, field string, newValue any) error {
	v := reflect.ValueOf(ptr).Elem().FieldByName(field)
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	nv := reflect.ValueOf(newValue)
	if v.Kind() != nv.Kind() {
		return fmt.Errorf("expected kind :%s, get kind: %s", v.Kind(), nv.Kind())
	}
	v.Set(nv)
	return nil
}
