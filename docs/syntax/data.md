# Data
Most of what we do surrounds data and in strabo, you can define it similar to how you would in TypeScript or Go.
Data are objects that can be referenced later to indicate data that an endpoint expects to receive or return.

```
data MyDataObject = {
    PropertyA: u8
    PropertyB: bool
    PropertyC: date
    PropertyD: string
}
```

By default, all of the data is expected to be sent and received as `application/json`.
So the above would translate as 

```json
{
    "PropertyA": 7,
    "PropertyB": false,
    "PropertyC": "2025-09-03",
    "PropertyD": "Hello, Strabo"
}
```

## Nesting
Data structures may be used inside of other data structures

```
data Job = {
    title: string
    salary: f32
    description: string
}

data Employee = {
    name: string
    startDate: date
    job: Job
}
```

## Extending Types
Often we send and return similar data structures. 
Objects can be extended using the `&` sign.

```
data Pet = {
    name: string
    age: u8
}

data Cat = Pet & {
    color: string
}
```
