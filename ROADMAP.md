# Roadmap
### Prototyping and questions to Answer Early
 - [ ] use TypeScript to define the data or make up a simple syntax that could easily be written to from TypeScript or Go
 - [ ] 

### First Release v0.1
The first release or minimum viable features to have this be usable are:

 - [ ] Data and Routes are described via a TypeScript file or similar looking yet simple syntax. 
 - [ ] Read Data Descriptions
   - [ ] How many properties
   - [ ] Their names
   - [ ] Their types  
 - [ ] Read List of Routes/Endpoints
   - [ ] HTTP Method
   - [ ] Data that goes to it
   - [ ] Data that comes from it
   - [ ] endpoint path
 - [ ] Create a Static HTML file from the data and endpoints displaying them
   - [ ] Say what routes exist
   - [ ] Say what data goes to that route
   - [ ] Say what data is returned from that route (assuming no errors)
      
### Second Release Goals
After the bare minimum, to make this more useful and enticing, Targete the following goals

 - [ ] Descritpions for each property within a `data` object
 - [ ] Group endpoints/routes together
 - [ ] CSS on HTML file to make it pretty
 - [ ] Allow comments in the file
 - [ ] API Level properties, things that apply to the server or all endpoints
   - [ ] Base URL
   - [ ] Auth Type
   - [ ] application/json vs some binary format over gRPC

### Phase Three
 - [ ] Read From Multiple files
 - [ ] Describe Errors

### Long-term Goals
 - [ ] Webserver to serve up the documentation
 - [ ] Data constraints, "must be greater than 0"
 - [ ] Given data structures, automatically generate example data

```rust
struct Person = {
  age: u8,
  name: String,
};
// =>  { "age": 72, "name": "bob" }
```
 - [ ] produce PDF output
