# Naming Scheme

## Type References

1. environment - env
2. stack - stack

## Struct members

Struct type name reference outlined in, "Type References," followed by the camel-cased name of the member.

```Go
envStack := newStack()
```

## Call Endpoints

Short code reference: "/call/${callMethodHash or provided endpoint}"

Standard reference: "localhost:8080/call/${callMethodHash or provided endpoint}"

## Comments

Heading, body separated by, " - ."

```MD
// Lorem - lorem ipsum dolor sit amet
```