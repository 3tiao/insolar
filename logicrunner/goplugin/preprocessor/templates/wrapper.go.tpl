package main

import (
    {{- range $import, $i := .Imports }}
        {{$import}}
    {{- end }}
)

type ExtendableError struct{
    S string
}

func ( e *ExtendableError ) Error() string{
    return e.S
}

{{ range $method := .Methods }}
func INSMETHOD_{{ $method.Name }}(object []byte, data []byte) ([]byte, []byte, error) {
    ph := proxyctx.Current

    self := new({{ $.ContractType }})

    err := ph.Deserialize(object, self)
    if err != nil {
        e := &ExtendableError{ S: "[ Fake{{ $method.Name }} ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Data: " + err.Error() }
        return nil, nil, e
    }

    {{ $method.ArgumentsZeroList }}
    err = ph.Deserialize(data, &args)
    if err != nil {
        e := &ExtendableError{ S: "[ Fake{{ $method.Name }} ] ( INSMETHOD_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error() }
        return nil, nil, e
    }

{{ if $method.Results }}
    {{ $method.Results }} := self.{{ $method.Name }}( {{ $method.Arguments }} )
{{ else }}
    self.{{ $method.Name }}( {{ $method.Arguments }} )
{{ end }}

    state := []byte{}
    err = ph.Serialize(self, &state)
    if err != nil {
        return nil, nil, err
    }

{{ range $i := $method.ErrorInterfaceInRes }}
    ret{{ $i }} = ph.MakeErrorSerializable(ret{{ $i }})
{{ end }}

    ret := []byte{}
    err = ph.Serialize([]interface{} { {{ $method.Results }} }, &ret)

    return state, ret, err
}
{{ end }}


{{ range $f := .Functions }}
func INSCONSTRUCTOR_{{ $f.Name }}(data []byte) ([]byte, error) {
    ph := proxyctx.Current
    {{ $f.ArgumentsZeroList }}
    err := ph.Deserialize(data, &args)
    if err != nil {
        e := &ExtendableError{ S: "[ Fake{{ $f.Name }} ] ( INSCONSTRUCTOR_* ) ( Generated Method ) Can't deserialize args.Arguments: " + err.Error() }
        return nil, e
    }

    {{ $f.Results }} := {{ $f.Name }}( {{ $f.Arguments }} )
    if ret1 != nil {
        return nil, ret1
    }

    ret := []byte{}
    err = ph.Serialize(ret0, &ret)
    if err != nil {
        return nil, err
    }

    return ret, err
}
{{ end }}
