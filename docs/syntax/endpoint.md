# Endpoint
An endpoint is one single REST endpoint that can be hit. 
It is limited to one REST method. 
To have one endpoint with multiple methods, see Group. 

## Properties
```typescript
type endpoint = {
  method: HTTPMethod;
  response: Data;
  body?: Data;
  path: String;
  description?: String;
  error?: Data;
  params?: Data;
}
```
### Method
The REST Method, one of `GET`, `POST`, `PUT`, `OPTIONS`, `DELETE`or `PATCH`.
### Response
The structure of the data that is expected to be returned. See [data](./data.md) for more details.
### Body
Optionally descibe the data that is expected for this payload. For `GET` requests, this would be omitted
### Error
Optionally provide what a payload would look like if an error is returned outside of the standard status code and message.
### Path
The path this endpoint will serve.
### Params
If path params are inclduded, use a data type object to describe the expected data types for each path parameter.
### Description
> ** Note **
> perhaps change to use /// to denote a description above the object like Rust docs
Optionally Describe this endpoint with a string body. 
## Examples

```
endpoint GetBrand = {
  method: "POST",
  response: Brand,
  body: CreateBodyPayload,
  path: "/brands/:id", 
  description: "Pull info on a single Brand"
}
```
