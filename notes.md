# Notas

Se debe declarar un package y una función main.

## Export

Para exportar simplemente ponemos la primera letra en Mayúsculas.

## Enums en Go

De forma predeterminada no existen pero se pueden crear asi

```go
const (
    Variable = "test"
)
```

## Funciones

Las funciones pueden devolver varios valores

## Manejo de errores

En go no tenemos Try catch hay que gestionar los errores manualmente.

Para lanzar un nuevo error

```go
return "", errors.New("Error")
```

## Tipado

En funciones puedes encadenar los tipos es decir declaras dos variables del mismo tipo.

```go
name, channel string
```

## For

En go solo tenemos este ciclo
Y en go podemos dejar los valores que no queremos guardar con _